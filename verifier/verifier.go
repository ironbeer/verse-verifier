package verifier

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
)

type verifyWorkerContext struct {
	cfg       *config.Verifier
	db        *database.Database
	signerCtx *ethutil.SignerContext
	topic     *util.Topic
	log       log.Logger
}

type VerifyWorker interface {
	id() string
	rpc() string
	work(*verifyWorkerContext, context.Context)
}

// Worker to verify the events of OasysStateCommitmentChain.
type Verifier struct {
	cfg       *config.Verifier
	db        *database.Database
	signerCtx *ethutil.SignerContext
	topic     *util.Topic
	workers   *sync.Map
	log       log.Logger
}

// Returns the new verifier.
func NewVerifier(cfg *config.Verifier, db *database.Database, signerCtx *ethutil.SignerContext) *Verifier {
	return &Verifier{
		cfg:       cfg,
		db:        db,
		signerCtx: signerCtx,
		topic:     util.NewTopic(),
		workers:   &sync.Map{},
		log:       log.New("worker", "verifier"),
	}
}

// Start verifier.
func (w *Verifier) Start(ctx context.Context) {
	w.log.Info(
		"Worker started",
		"signer", w.signerCtx.Signer,
		"interval", w.cfg.Interval,
		"state-collect-limit", w.cfg.StateCollectLimit,
		"concurrency", w.cfg.Concurrency,
	)

	wg := util.NewWorkerGroup(w.cfg.Concurrency)
	running := &sync.Map{}

	tick := time.NewTicker(w.cfg.Interval)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker stopped")
			return
		case <-tick.C:
			w.workers.Range(func(key, val interface{}) bool {
				id, ok := key.(string)
				if !ok {
					return true
				}
				worker, ok := val.(VerifyWorker)
				if !ok {
					return true
				}

				// deduplication
				if _, ok := running.Load(id); ok {
					return true
				}
				running.Store(id, 1)

				if !wg.Has(id) {
					handler := func(ctx context.Context, rname string, data interface{}) {
						defer running.Delete(rname)

						if worker, ok := data.(VerifyWorker); ok {
							worker.work(&verifyWorkerContext{
								cfg:       w.cfg,
								db:        w.db,
								signerCtx: w.signerCtx,
								topic:     w.topic,
								log:       w.log.New(),
							}, ctx)
						}
					}
					wg.AddWorker(ctx, id, handler)
				}

				wg.Enqueue(id, worker)
				return true
			})
		}
	}
}

func (w *Verifier) SignerContext() *ethutil.SignerContext {
	return w.signerCtx
}

// TODO
func (w *Verifier) HasWorker(id, rpc string) bool {
	val, ok := w.workers.Load(id)
	if !ok {
		return false
	}
	// If the L2 RPC is changed, replace the worker.
	return rpc == val.(VerifyWorker).rpc()
}

func (w *Verifier) AddWorker(worker VerifyWorker) {
	w.workers.Store(worker.id(), worker)
}

func (w *Verifier) RemoveWorker(id string) {
	w.workers.Delete(id)
}

func (s *Verifier) SubscribeNewSignature(ctx context.Context) *SignatureSubscription {
	ch := make(chan interface{})
	cancel := s.topic.Subscribe(ctx, func(ctx context.Context, data interface{}) {
		ch <- data
	})
	return &SignatureSubscription{Cancel: cancel, ch: ch}
}

type SignatureSubscription struct {
	Cancel context.CancelFunc
	ch     chan interface{}
}

func (s *SignatureSubscription) Next() <-chan interface{} {
	return s.ch
}
