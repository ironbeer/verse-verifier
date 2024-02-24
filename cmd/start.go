package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/beacon"
	"github.com/oasysgames/oasys-optimism-verifier/cmd/ipccmd"
	"github.com/oasysgames/oasys-optimism-verifier/collector"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/contract/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/debug"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
	"github.com/oasysgames/oasys-optimism-verifier/metrics"
	"github.com/oasysgames/oasys-optimism-verifier/p2p"
	submitterpkg "github.com/oasysgames/oasys-optimism-verifier/submitter"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	verifierpkg "github.com/oasysgames/oasys-optimism-verifier/verifier"
	"github.com/oasysgames/oasys-optimism-verifier/version"
	"github.com/oasysgames/oasys-optimism-verifier/wallet"
	"github.com/spf13/cobra"
)

const (
	SccName             = "StateCommitmentChain"
	L2OOName            = "L2OutputOracle"
	StakeManagerAddress = "0x0000000000000000000000000000000000001001"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Verifier",
	Long:  "Start the Verifier",
	Run:   runStartCmd,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runStartCmd(cmd *cobra.Command, args []string) {
	log.Info(fmt.Sprintf("Start %s", commandName), "version", version.SemVer())

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := mustNewServer(cmd)

	// start metrics server
	s.mustStartMetrics(ctx)

	// start pprof server
	s.mustStartPprof(ctx)

	// start ipc server
	// Note: must start the IPC server before unlocking wallets.
	s.mustStartIPC(ctx)

	// unlock walelts(wait forever)
	s.waitForUnlockWallets(ctx)

	// start p2p
	// Note: must start the P2P before setup beacon worker.
	s.mustStartP2P(ctx)

	// setup workers
	s.setupCollector()
	s.mustSetupVerifier()
	s.mustSetupSubmitter()
	s.setupBeacon()

	// start workers
	s.start(ctx)
	s.wg.Wait()

	// stopped by signal
	log.Info("Stopped all workers")
}

type server struct {
	wg               sync.WaitGroup
	conf             *config.Config
	db               *database.Database
	ks               *wallet.KeyStore
	hub              ethutil.ReadOnlyClient
	smcache          *stakemanager.Cache
	ipc              *ipc.IPCServer
	p2p              *p2p.Node
	blockCollector   *collector.BlockCollector
	eventCollector   *collector.EventCollector
	verifier         *verifierpkg.Verifier
	submitter        *submitterpkg.Submitter
	bw               *beacon.BeaconWorker
	discoveredVerses sync.Map
	discoveredNotify chan struct{}
}

func mustNewServer(cmd *cobra.Command) *server {
	var err error

	s := &server{discoveredNotify: make(chan struct{}, 1)}

	// load configuration file
	if s.conf, err = loadConfig(cmd); err != nil {
		log.Crit("Failed to load configuration file", "err", err)
	}

	// setup database
	if s.conf.Database.Path == "" {
		s.conf.Database.Path = s.conf.DatabasePath()
	}
	if s.db, err = database.NewDatabase(&s.conf.Database); err != nil {
		log.Crit("Failed to open database", "err", err)
	}

	// open geth keystore
	s.ks = wallet.NewKeyStore(s.conf.KeyStore)

	// construct hub-layer client
	if s.hub, err = ethutil.NewReadOnlyClient(s.conf.HubLayer.RPC); err != nil {
		log.Crit("Failed to construct hub-layer client", "err", err)
	}

	// construct stakemanager cache
	sm, err := stakemanager.NewStakemanagerCaller(
		common.HexToAddress(StakeManagerAddress), s.hub)
	if err != nil {
		log.Crit("Failed to construct StakeManager", "err", err)
	}
	s.smcache = stakemanager.NewCache(sm)

	return s
}

func (s *server) mustStartMetrics(ctx context.Context) {
	if !s.conf.Metrics.Enable {
		return
	}

	s.wg.Add(1)
	metrics.Initialize(&s.conf.Metrics)

	go func() {
		defer s.wg.Done()
		if err := metrics.ListenAndServe(ctx); err != nil {
			log.Crit("Failed to start metrics server", "err", err)
		}
	}()
}

func (s *server) mustStartPprof(ctx context.Context) {
	if !s.conf.Debug.Pprof.Enable {
		return
	}

	s.wg.Add(1)

	ps := debug.NewPprofServer(&s.conf.Debug.Pprof)
	go func() {
		defer s.wg.Done()
		if err := ps.ListenAndServe(ctx); err != nil {
			log.Crit("Failed to start pprof server", "err", err)
		}
	}()
}

func (s *server) mustStartIPC(ctx context.Context) {
	if s.conf.IPC.Sockname == "" {
		log.Crit("IPC socket name is required")
	}

	ipc, err := ipc.NewIPCServer(s.conf.IPC.Sockname)
	if err != nil {
		log.Crit("Failed to start ipc server", "err", err)
	}

	s.ipc = ipc

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.ipc.Start(ctx)
	}()
}

func (s *server) mustStartP2P(ctx context.Context) {
	// get p2p private key
	p2pKey, err := getOrCreateP2PKey(s.conf.P2PKeyPath())
	if err != nil {
		log.Crit("Failed to get(or create) p2p key", "err", err)
	}

	// construct libp2p host
	host, dht, bwm, hpHelper, err := p2p.NewHost(ctx, &s.conf.P2P, p2pKey)
	if err != nil {
		log.Crit("Failed to construct libp2p host", "err", err)
	}

	// ignore self-signed signatures
	ignoreSigners := []common.Address{}
	if s.conf.Verifier.Enable {
		_, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
		ignoreSigners = append(ignoreSigners, account.Address)
	}

	s.p2p, err = p2p.NewNode(&s.conf.P2P, s.db, host, dht, bwm,
		hpHelper, s.conf.HubLayer.ChainId, ignoreSigners, s.smcache)
	if err != nil {
		log.Crit("Failed to construct p2p node", "err", err)
	}

	s.ipc.SetHandler(ipccmd.PingCmd.NewHandler(ctx, s.p2p.Host(), s.p2p.HolePunchHelper()))
	s.ipc.SetHandler(ipccmd.StatusCmd.NewHandler(s.p2p.Host()))

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.p2p.Start(ctx)
	}()
}

func (s *server) startVerseDiscovery(ctx context.Context) {
	if s.conf.VerseLayer.Discovery.Endpoint == "" {
		return
	}

	disc := config.NewVerseDiscovery(
		http.DefaultClient,
		s.conf.VerseLayer.Discovery.Endpoint,
		s.conf.VerseLayer.Discovery.RefreshInterval)

	sub := disc.Subscribe(ctx)
	defer sub.Cancel()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case verses := <-sub.Next():
				for _, verse := range verses {
					s.discoveredVerses.Store(verse.ChainID, verse)
				}

				select {
				case <-ctx.Done():
					return
				case s.discoveredNotify <- struct{}{}:
				}
			}
		}
	}()

	time.Sleep(time.Second)
	disc.Start(ctx)
}

func (s *server) setupCollector() {
	if !s.conf.Verifier.Enable {
		return
	}

	_, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)

	s.blockCollector = collector.NewBlockCollector(&s.conf.Verifier, s.db, s.hub)
	s.eventCollector = collector.NewEventCollector(&s.conf.Verifier, s.db, s.hub, account.Address)
}

func (s *server) mustSetupVerifier() {
	if !s.conf.Verifier.Enable {
		return
	}

	wallet, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
	signer, err := ethutil.NewWritableClient(
		new(big.Int).SetUint64(s.conf.HubLayer.ChainId), s.conf.HubLayer.RPC, wallet, account)
	if err != nil {
		log.Crit("Failed to construct verifier", "err", err)
	}

	s.verifier = verifierpkg.NewVerifier(&s.conf.Verifier, s.db, signer.SignerContext())
}

func (s *server) mustSetupSubmitter() {
	if !s.conf.Submitter.Enable {
		return
	}

	sm, err := stakemanager.NewStakemanager(
		common.HexToAddress(StakeManagerAddress), s.hub)
	if err != nil {
		log.Crit("Failed to construct submitter", "err", err)
	}

	s.submitter = submitterpkg.NewSubmitter(&s.conf.Submitter, s.db, sm)
}

func (s *server) setupBeacon() {
	if !s.conf.Beacon.Enable || !s.conf.Verifier.Enable {
		return
	}

	_, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
	s.bw = beacon.NewBeaconWorker(
		&s.conf.Beacon,
		http.DefaultClient,
		beacon.Beacon{
			Signer:  account.Address.Hex(),
			Version: version.SemVer(),
			PeerID:  s.p2p.PeerID().String(),
		},
	)
}

func (s *server) verseNotifyHandler(ctx context.Context) {
	s.discoveredVerses.Range(func(key, value any) bool {
		verse, ok := value.(*config.Verse)
		if !ok {
			return true
		}

		var err error

		// construct rollup contract
		var scc_ *scc.Scc
		sccAddr := common.HexToAddress(verse.L1Contracts[SccName])
		if sccAddr != (common.Address{}) {
			if scc_, err = scc.NewScc(sccAddr, s.hub); err != nil {
				log.Error("Failed to construct OasysStateCommitmentChain client", "err", err)
				return true
			}
		}
		var l2oo_ *l2oo.OasysL2OutputOracle
		l2ooAddr := common.HexToAddress(verse.L1Contracts[L2OOName])
		if l2ooAddr != (common.Address{}) {
			if l2oo_, err = l2oo.NewOasysL2OutputOracle(l2ooAddr, s.hub); err != nil {
				log.Error("Failed to construct OasysL2OutputOracle client", "err", err)
				return true
			}
		}
		if scc_ == nil && l2oo_ == nil {
			return true
		}

		// add verse to Verifier
		if s.verifier != nil {
			var (
				workers  []verifierpkg.VerifyWorker
				l2Client ethutil.ReadOnlyClient
			)
			if scc_ != nil {
				workers = append(workers, verifierpkg.NewSccVerifyWorker(l2Client, sccAddr, scc_))
			}
			if l2oo_ != nil {
				workers = append(workers, verifierpkg.NewL2OOVerifyWorker(l2Client, l2ooAddr, l2oo_))
			}
			if len(workers) > 0 {
				l2Client, err = ethutil.NewReadOnlyClient(verse.RPC)
				if err != nil {
					log.Error("Failed to construct verse-layer client", "err", err)
					return true
				}
				for _, worker := range workers {
					s.verifier.AddWorker(worker)
				}
			}
		}

		// add verse to Submitter
		if s.submitter != nil {
			for _, t := range s.conf.Submitter.Targets {
				if t.ChainID != verse.ChainID {
					continue
				}

				wallet, account := findWallet(s.conf, s.ks, t.Wallet)
				l1Client, err := ethutil.NewWritableClient(
					new(big.Int).SetUint64(s.conf.HubLayer.ChainId), s.conf.HubLayer.RPC, wallet, account)
				if err != nil {
					log.Error("Failed to construct hub-layer client", "err", err)
					return true
				}

				verifier, err := sccverifier.NewOasysRollupVerifier(
					common.HexToAddress(s.conf.Submitter.VerifierAddress), l1Client)
				if err != nil {
					log.Error("Failed to construct OasysRollupVerifier contract", "err", err)
					return true
				}

				var multicall *multicall2.Multicall2
				if s.conf.Submitter.UseMulticall {
					multicall, err = multicall2.NewMulticall2(
						common.HexToAddress(s.conf.Submitter.Multicall2Address), l1Client)
					if err != nil {
						log.Error("Failed to construct OasysRollupVerifier contract", "err", err)
						return true
					}
				}

				// TODO
				if scc_ != nil {
					task, err := submitterpkg.NewTask(
						l1Client, sccAddr, scc_, verifier, multicall)
					if err != nil {
						log.Error("Failed to add SCC contract submit task", "err", err)
					} else {
						s.submitter.AddTask(t.ChainID, task)
					}
				}
				if l2oo_ != nil {
					task, err := submitterpkg.NewTask(
						l1Client, l2ooAddr, l2oo_, verifier, multicall)
					if err != nil {
						log.Error("Failed to add L2OO contract submit task", "err", err)
					} else {
						s.submitter.AddTask(t.ChainID, task)
					}
				}

				break
			}
		}

		return true
	})
}

func (s *server) start(ctx context.Context) {
	// start block collector
	if s.blockCollector != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.blockCollector.Start(ctx)
		}()
	}

	// start event collector
	if s.eventCollector != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.eventCollector.Start(ctx)
		}()
	}

	// start cache updater
	s.smcache.Refresh(ctx) // first time synchronous
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.smcache.RefreshLoop(ctx, time.Hour)
	}()

	// start verifier
	if s.verifier != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.verifier.Start(ctx)
		}()

		// start database optimizer
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()

			// optimize database every hour
			tick := util.NewTicker(s.conf.Verifier.OptimizeInterval, 1)
			defer tick.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-tick.C:
					// TODO
					s.db.Optimism.RepairPreviousID(s.verifier.SignerContext().Signer)
				}
			}
		}()

		// publish new signature via p2p
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()

			sub := s.verifier.SubscribeNewSignature(ctx)
			defer sub.Cancel()

			debounce := time.NewTicker(time.Second * 5)
			defer debounce.Stop()

			var optimisms, opstacks sync.Map
			for {
				select {
				case <-ctx.Done():
					return
				case sig := <-sub.Next():
					switch t := sig.(type) {
					case *database.OptimismSignature:
						optimisms.Store(t.Signer.Address, t)
					case *database.OpstackSignature:
						opstacks.Store(t.Signer.Address, t)
					}
				case <-debounce.C:
					// TODO
					var (
						pubOptimisms []*database.OptimismSignature
						pubOpstacks  []*database.OpstackSignature
					)
					optimisms.Range(func(_, value any) bool {
						pubOptimisms = append(pubOptimisms, value.(*database.OptimismSignature))
						return true
					})
					opstacks.Range(func(_, value any) bool {
						pubOpstacks = append(pubOpstacks, value.(*database.OpstackSignature))
						return true
					})
					optimisms, opstacks = sync.Map{}, sync.Map{}
					s.p2p.PublishSignatures(ctx, pubOptimisms, pubOpstacks)
				}
			}
		}()
	}

	// start submitter
	if s.submitter != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.submitter.Start(ctx)
		}()
	}

	// start verse discovery handler
	if s.verifier != nil || s.submitter != nil {
		// read verses in the configuration file
		for _, verse := range s.conf.VerseLayer.Directs {
			s.discoveredVerses.Store(verse.ChainID, verse)
		}
		go func() {
			s.discoveredNotify <- struct{}{}
		}()

		s.wg.Add(1)
		go func() {
			defer s.wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case <-s.discoveredNotify:
					s.verseNotifyHandler(ctx)
				}
			}
		}()
	}

	// start dynamic verse discovery
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.startVerseDiscovery(ctx)
	}()

	// start beacon worker
	if s.bw != nil {
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.bw.Start(ctx)
		}()
	}
}

func getOrCreateP2PKey(filename string) (crypto.PrivKey, error) {
	data, err := ioutil.ReadFile(filename)

	if err == nil {
		dec, peerID, err := p2p.DecodePrivateKey(string(data))
		if err != nil {
			log.Error("Failed to decode p2p private key", "err", err)
			return nil, err
		}

		log.Info("Loaded p2p private key", "file", filename, "id", peerID)
		return dec, nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		log.Error("Failed to load p2p private key", "err", err)
		return nil, err
	}

	priv, _, peerID, err := p2p.GenerateKeyPair()
	if err != nil {
		log.Error("Failed to generate p2p private key", "err", err)
		return nil, err
	}

	enc, err := p2p.EncodePrivateKey(priv)
	if err != nil {
		log.Error("Failed to encode p2p private key", "err", err)
		return nil, err
	}

	err = ioutil.WriteFile(filename, []byte(enc), 0644)
	if err != nil {
		log.Error("Failed to write p2p private key", "err", err)
		return nil, err
	}

	log.Info("Generated and saved to p2p private key", "file", filename, "id", peerID)
	return priv, nil
}

func (s *server) waitForUnlockWallets(ctx context.Context) {
	s.ipc.SetHandler(ipccmd.WalletUnlockCmd.NewHandler(s.ks))

	var wg sync.WaitGroup
	wg.Add(len(s.conf.Wallets))

	for name, wallet := range s.conf.Wallets {
		go func(name string, wallet config.Wallet) {
			defer wg.Done()

			_wallet, account, err := s.ks.FindWallet(common.HexToAddress(wallet.Address))
			if err != nil {
				log.Crit("Failed to find a wallet",
					"name", name, "address", wallet.Address, "err", err)
			}

			if wallet.Password != "" {
				b, err := ioutil.ReadFile(wallet.Password)
				if err != nil {
					log.Crit(
						"Failed to read password file",
						"name", name, "address", wallet.Address, "err", err)
				}

				if err := s.ks.Unlock(*account, strings.Trim(string(b), "\r\n\t ")); err != nil {
					log.Crit("Failed to unlock wallet using password file",
						"name", name, "address", wallet.Address, "err", err)
				}
			} else if s.ks.Unlock(*account, "") != nil {
				log.Info("Waiting for wallet unlock", "name", name, "address", wallet.Address)
				if err := s.ks.WaitForUnlock(ctx, _wallet); err != nil {
					log.Crit("Wallet was not unlocked",
						"name", name, "address", wallet.Address, "err", err)
				}
			}

			log.Info("Wallet unlocked", "name", name, "address", wallet.Address)
		}(name, wallet)
	}

	wg.Wait()
}

func findWallet(
	c *config.Config,
	ks *wallet.KeyStore,
	name string,
) (accounts.Wallet, *accounts.Account) {
	wallet, account, err := ks.FindWallet(common.HexToAddress(c.Wallets[name].Address))
	if err != nil {
		log.Crit("Wallet not found", "name", name)
	}
	return wallet, account
}
