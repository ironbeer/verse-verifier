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
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
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

	wg := &sync.WaitGroup{}

	// load configuration file
	conf, err := loadConfig(cmd)
	if err != nil {
		log.Crit("Failed to load configuration file", "err", err)
	}

	// setup database
	if conf.Database.Path == "" {
		conf.Database.Path = conf.DatabasePath()
	}
	db, err := database.NewDatabase(&conf.Database)
	if err != nil {
		log.Crit("Failed to open database", "err", err)
	}

	// open geth keystore
	ks := wallet.NewKeyStore(conf.KeyStore)

	// start ipc server
	// note: start ipc server before unlocking wallet
	ipc := newIPC(conf, ks)
	if ipc != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ipc.Start(ctx)
		}()
	}

	// unlock walelts(wait forever)
	waitForUnlockWallets(ctx, conf, ks)

	// create hub-layer client
	hub, err := ethutil.NewReadOnlyClient(conf.HubLayer.RPC)
	if err != nil {
		log.Crit("Failed to create hub-layer client", "err", err)
	}

	// start block collector
	bkCollector := newBlockCollector(ctx, conf, db, hub)
	if bkCollector != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bkCollector.Start(ctx)
		}()
	}

	// start event collector
	evCollector := newEventCollector(ctx, conf, db, hub)
	if evCollector != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			evCollector.Start(ctx)
		}()
	}

	// Construct the L1 RPC Client
	wallet, account := findWallet(conf, ks, conf.Verifier.Wallet)
	l1Client, err := ethutil.NewWritableClient(
		new(big.Int).SetUint64(conf.HubLayer.ChainId),
		conf.HubLayer.RPC,
		wallet,
		account,
	)
	if err != nil {
		log.Crit("Failed to create hub-layer clinet", "err", err)
	}

	// construct state verifier
	verifier := newVerifier(conf, db, l1Client)

	//  start p2p
	p2p := newP2P(ctx, conf, db, verifier)
	wg.Add(1)
	go func() {
		defer wg.Done()
		p2p.Start(ctx)
	}()

	// start state verifier
	if verifier != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			startVerifier(ctx, conf, db, l1Client, verifier, p2p)
		}()
	}

	// start beacon worker
	if verifier != nil && conf.Beacon.Enable {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bw := beacon.NewBeaconWorker(
				&conf.Beacon,
				http.DefaultClient,
				beacon.Beacon{
					Signer:  verifier.SignerContext().Signer.String(),
					Version: version.SemVer(),
					PeerID:  p2p.PeerID().String(),
				},
			)
			bw.Start(ctx)
		}()
	}

	// start signature submitter
	submitter := newSubmitter(ctx, conf, ks, db, hub)
	if submitter != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			submitter.Start(ctx)
		}()
	}

	// start verse discovery worker
	wg.Add(1)
	go func() {
		defer wg.Done()
		startVerseDiscovery(ctx, conf, ks, l1Client, verifier, submitter)
	}()

	wg.Wait()
	log.Info("Stopped all workers")
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

func newIPC(c *config.Config, ks *wallet.KeyStore) *ipc.IPCServer {
	if !c.IPC.Enable {
		return nil
	}

	ipc, err := ipc.NewIPCServer(commandName)
	if err != nil {
		log.Crit("Failed to create ipc server", "err", err)
	}

	ipc.SetHandler(ipccmd.WalletUnlockCmd.NewHandler(ks))
	return ipc
}

func newP2P(
	ctx context.Context,
	c *config.Config,
	db *database.Database,
	verifier *verifierpkg.Verifier,
) *p2p.Node {
	// get p2p private key
	p2pKey, err := getOrCreateP2PKey(c.P2PKeyPath())
	if err != nil {
		log.Crit(err.Error())
	}

	// setup p2p node
	listens := strings.Split(c.P2P.Listen, ":")
	host, dht, bwm, err := p2p.NewHost(ctx, listens[0], listens[1], p2pKey)
	if err != nil {
		log.Crit(err.Error())
	}

	// connect to bootstrap peers and setup peer discovery
	p2p.Bootstrap(ctx, host, dht)
	bootstrapPeers := p2p.ConvertPeers(c.P2P.Bootnodes)
	if len(bootstrapPeers) > 0 {
		go func() {
			ticker := util.NewTicker(time.Minute, 1)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					p2p.ConnectPeers(ctx, host, bootstrapPeers)
				}
			}
		}()
	}

	// ignore self-signed signatures
	ignoreSigners := []common.Address{}
	if verifier != nil {
		ignoreSigners = append(ignoreSigners, verifier.SignerContext().Signer)
	}

	node, err := p2p.NewNode(&c.P2P, db, host, dht, bwm, c.HubLayer.ChainId, ignoreSigners)
	if err != nil {
		log.Crit("Failed to create p2p server", "err", err)
	}

	return node
}

func newBlockCollector(
	ctx context.Context,
	c *config.Config,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
) *collector.BlockCollector {
	if !c.Verifier.Enable {
		return nil
	}

	return collector.NewBlockCollector(&c.Verifier, db, hub)
}

func newEventCollector(
	ctx context.Context,
	c *config.Config,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
) *collector.EventCollector {
	if !c.Verifier.Enable {
		return nil
	}

	return collector.NewEventCollector(
		&c.Verifier, db, hub,
		common.HexToAddress(c.Wallets[c.Verifier.Wallet].Address),
	)
}

func newVerifier(c *config.Config, db *database.Database, l1Client ethutil.WritableClient) *verifierpkg.Verifier {
	if !c.Verifier.Enable {
		return nil
	}
	return verifierpkg.NewVerifier(&c.Verifier, db, l1Client.SignerContext())
}

func newSubmitter(
	ctx context.Context,
	c *config.Config,
	ks *wallet.KeyStore,
	db *database.Database,
	hub ethutil.ReadOnlyClient,
) *submitterpkg.Submitter {
	if !c.Submitter.Enable {
		return nil
	}

	sm, err := stakemanager.NewStakemanager(common.HexToAddress(StakeManagerAddress), hub)
	if err != nil {
		log.Crit("Failed to create StakeManager", "err", err)
	}

	return submitterpkg.NewSubmitter(&c.Submitter, db, sm)
}

func startVerseDiscovery(
	ctx context.Context,
	c *config.Config,
	ks *wallet.KeyStore,
	l1Client ethutil.ReadOnlyClient,
	verifier *verifierpkg.Verifier,
	submitter *submitterpkg.Submitter,
) {
	if !c.Verifier.Enable && !c.Submitter.Enable {
		return
	}

	notify := make(chan struct{}, 1)
	verses := &sync.Map{}
	for _, v := range c.VerseLayer.Directs {
		verses.Store(v.ChainID, v)
	}
	notify <- struct{}{}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-notify:
				verses.Range(func(key, value any) bool {
					verse, ok := value.(*config.Verse)
					if !ok {
						return true
					}

					// get contract address and construct rollup contracts
					var (
						sccAddr, l2ooAddr common.Address
						scc_              *scc.Scc
						l2oo_             *l2oo.OasysL2OutputOracle
						err               error
					)
					if s, ok := verse.L1Contracts[SccName]; ok {
						sccAddr = common.HexToAddress(s)
						if scc_, err = scc.NewScc(sccAddr, l1Client); err != nil {
							log.Error("Failed to construct OasysStateCommitmentChain client", "err", err)
							return true
						}
					}
					if s, ok := verse.L1Contracts[L2OOName]; ok {
						l2ooAddr = common.HexToAddress(s)
						if l2oo_, err = l2oo.NewOasysL2OutputOracle(l2ooAddr, l1Client); err != nil {
							log.Error("Failed to construct OasysL2OutputOracle client", "err", err)
							return true
						}
					}

					// add verse to Verifier
					if c.Verifier.Enable {
						l2Client, err := ethutil.NewReadOnlyClient(verse.RPC)
						if err != nil {
							log.Error("Failed to create verse-layer client", "err", err)
							return true
						}

						if scc_ != nil {
							verifier.AddWorker(verifierpkg.NewSccVerifyWorker(l2Client, sccAddr, scc_))
						}
						if l2oo_ != nil {
							verifier.AddWorker(verifierpkg.NewL2OOVerifyWorker(l2Client, l2ooAddr, l2oo_))
						}
					}

					// add verse to Submitter
					if c.Submitter.Enable {
						for _, t := range c.Submitter.Targets {
							if t.ChainID != verse.ChainID {
								continue
							}

							wallet, account := findWallet(c, ks, t.Wallet)
							l1Client, err := ethutil.NewWritableClient(
								new(big.Int).SetUint64(c.HubLayer.ChainId),
								c.HubLayer.RPC,
								wallet,
								account,
							)
							if err != nil {
								log.Error("Failed to create hub-layer client", "err", err)
								return true
							}

							if scc_ != nil {
								submitter.AddTask(submitterpkg.NewSccSubmitTask(l1Client, sccAddr, scc_))
							}
							if l2oo_ != nil {
								submitter.AddTask(submitterpkg.NewL2OOSubmitTask(l1Client, l2ooAddr, l2oo_))
							}

							break
						}
					}

					return true
				})
			}
		}
	}()

	if c.VerseLayer.Discovery.Endpoint == "" {
		return
	}

	// start verse discovery
	discv := config.NewVerseDiscovery(
		http.DefaultClient,
		c.VerseLayer.Discovery.Endpoint,
		c.VerseLayer.Discovery.RefreshInterval,
	)

	go func() {
		sub := discv.Subscribe(ctx)
		defer sub.Cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case verse := <-sub.Next():
				verses.Store(verse.ChainID, verse)
				notify <- struct{}{}
			}
		}
	}()

	time.Sleep(1 * time.Second)
	discv.Start(ctx)
}

func startVerifier(
	ctx context.Context,
	c *config.Config,
	db *database.Database,
	l1Client ethutil.WritableClient,
	verifier *verifierpkg.Verifier,
	p2p *p2p.Node,
) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		verifier.Start(ctx)
	}()

	// optimize database every hour
	wg.Add(1)
	go func() {
		defer wg.Done()

		tick := util.NewTicker(c.Verifier.OptimizeInterval, 1)
		defer tick.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				db.Optimism.RepairPreviousID(l1Client.Signer())
			}
		}
	}()

	// publish new signature via p2p
	wg.Add(1)
	go func() {
		defer wg.Done()

		sub := verifier.SubscribeNewSignature(ctx)
		defer sub.Cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case sig := <-sub.Next():
				switch t := sig.(type) {
				case *database.OptimismSignature:
					p2p.PublishSignatures(ctx, []*database.OptimismSignature{t})
				case *database.OpstackSignature:
					// TODO
				}
			}
		}
	}()

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
