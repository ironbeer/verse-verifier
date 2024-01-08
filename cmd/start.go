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
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/ipc"
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

	startCmd.Flags().String(configFlag, "", "configuration file")
	startCmd.MarkFlagRequired(configFlag)
}

func runStartCmd(cmd *cobra.Command, args []string) {
	log.Info(fmt.Sprintf("Start %s", commandName), "version", version.SemVer())

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := mustNewServer(cmd)

	// start ipc server
	// Note: must start the IPC server before unlocking the wallet.
	s.mustStartIPC(ctx)

	// unlock walelts(wait forever)
	if s.ipc != nil {
		waitForUnlockWallets(ctx, s.conf, s.ks)
	}

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

	return s
}

func (s *server) mustStartIPC(ctx context.Context) {
	if !s.conf.IPC.Enable {
		return
	}

	ipc, err := ipc.NewIPCServer(commandName)
	if err != nil {
		log.Crit("Failed to start ipc server", "err", err)
	}

	s.ipc = ipc
	s.ipc.SetHandler(ipccmd.WalletUnlockCmd.NewHandler(s.ks))

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

	listen := strings.Split(s.conf.P2P.Listen, ":")
	host, dht, bwm, err := p2p.NewHost(ctx, listen[0], listen[1], p2pKey)
	if err != nil {
		log.Crit("Failed to setup libp2p", "err", err)
	}

	// etup peer discovery
	p2p.Bootstrap(ctx, host, dht)

	// connect to bootstrap peers
	bootnodes := p2p.ConvertPeers(s.conf.P2P.Bootnodes)
	if len(bootnodes) > 0 {
		ticker := util.NewTicker(time.Minute, 1)
		defer ticker.Stop()

		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					p2p.ConnectPeers(ctx, host, bootnodes)
				}
			}
		}()
	}

	// ignore self-signed signatures
	ignoreSigners := []common.Address{}
	if s.conf.Verifier.Enable {
		_, account := findWallet(s.conf, s.ks, s.conf.Verifier.Wallet)
		ignoreSigners = append(ignoreSigners, account.Address)
	}

	s.p2p, err = p2p.NewNode(&s.conf.P2P, s.db,
		host, dht, bwm, s.conf.HubLayer.ChainId, ignoreSigners)
	if err != nil {
		log.Crit("Failed to construct p2p node", "err", err)
	}

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
			case verse := <-sub.Next():
				s.discoveredVerses.Store(verse.ChainID, verse)

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

	sm, err := stakemanager.NewStakemanager(common.HexToAddress(StakeManagerAddress), s.hub)
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
			l2Client, err := ethutil.NewReadOnlyClient(verse.RPC)
			if err != nil {
				log.Error("Failed to construct verse-layer client", "err", err)
				return true
			}
			if scc_ != nil {
				s.verifier.AddWorker(verifierpkg.NewSccVerifyWorker(l2Client, sccAddr, scc_))
			}
			if l2oo_ != nil {
				s.verifier.AddWorker(verifierpkg.NewL2OOVerifyWorker(l2Client, l2ooAddr, l2oo_))
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

	if s.verifier != nil {
		// start verifier
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			s.verifier.Start(ctx)
		}()

		// start database optimizer
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()

			tick := util.NewTicker(s.conf.Verifier.OptimizeInterval, 1)
			defer tick.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-tick.C:
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

			for {
				select {
				case <-ctx.Done():
					return
				case sig := <-sub.Next():
					switch t := sig.(type) {
					case *database.OptimismSignature:
						s.p2p.PublishSignatures(ctx, []*database.OptimismSignature{t})
					case *database.OpstackSignature:
						// TODO
					}
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

	// start verse discovery
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.startVerseDiscovery(ctx)
	}()

	// start verse handler
	if s.verifier != nil || s.submitter != nil {
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

func waitForUnlockWallets(ctx context.Context, c *config.Config, ks *wallet.KeyStore) {
	wg := &sync.WaitGroup{}
	wg.Add(len(c.Wallets))

	for name, wallet := range c.Wallets {
		go func(name string, wallet config.Wallet) {
			defer wg.Done()

			_wallet, account, err := ks.FindWallet(common.HexToAddress(wallet.Address))
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

				if err := ks.Unlock(*account, strings.Trim(string(b), "\r\n\t ")); err != nil {
					log.Crit("Failed to unlock wallet using password file",
						"name", name, "address", wallet.Address, "err", err)
				}
			} else if ks.Unlock(*account, "") != nil {
				log.Info("Waiting for wallet unlock", "name", name, "address", wallet.Address)
				if err := ks.WaitForUnlock(ctx, _wallet); err != nil {
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
