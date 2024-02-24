package p2p

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/routing"
	ps "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	meter "github.com/oasysgames/oasys-optimism-verifier/metrics"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oklog/ulid/v2"
	"golang.org/x/sync/semaphore"
	"golang.org/x/time/rate"
	"google.golang.org/protobuf/proto"
)

const (
	pubsubTopic    = "/oasys-optimism-verifier/pubsub/1.0.0"
	streamProtocol = "/oasys-optimism-verifier/stream/1.0.0"
)

var (
	eom             = &pb.Stream{Body: &pb.Stream_Eom{Eom: nil}}
	tenMillionEther = new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10_000_000))

	// miscellaneous messages
	misc_SIGRECEIVED = []byte("SIGNATURES_RECEIVED")
)

type Node struct {
	cfg             *config.P2P
	db              *database.Database
	h               host.Host
	dht             routing.Routing
	bwm             *metrics.BandwidthCounter
	hpHelper        HolePunchHelper
	hubLayerChainID *big.Int
	ignoreSigners   map[common.Address]int
	stakemanager    *stakemanager.Cache

	topic *ps.Topic
	sub   *ps.Subscription
	log   log.Logger

	outboundSem, inboundSem     *semaphore.Weighted
	outboundThrot, inboundThrot *rate.Limiter

	meterPubsubSubscribed,
	meterPubsubUnknownMsg,
	meterPubsubWorkers,
	meterStreamOpend,
	meterStreamHandled,
	meterStreamClosed,
	meterStreamWrites,
	meterStreamReads,
	meterStreamUnknownMsg,
	meterHolePunchSuccess,
	meterHolePunchErrs,
	meterStreamOpenErrs,
	meterStreamReadErrs,
	meterStreamWriteErrs meter.Counter

	meterPeers,
	meterTCPConnections,
	meterUDPConnections,
	meterRelayConnections,
	meterRelayHopStreams,
	meterRelayStopStreams,
	meterVerifierStreams,
	meterPubsubJobs meter.Gauge
}

func NewNode(
	cfg *config.P2P,
	db *database.Database,
	host host.Host,
	dht routing.Routing,
	bwm *metrics.BandwidthCounter,
	hpHelper HolePunchHelper,
	hubLayerChainID uint64,
	ignoreSigners []common.Address,
	stakemanager *stakemanager.Cache,
) (*Node, error) {
	_, topic, sub, err := setupPubSub(context.Background(), host, pubsubTopic)
	if err != nil {
		return nil, err
	}

	worker := &Node{
		cfg:             cfg,
		db:              db,
		h:               host,
		dht:             dht,
		bwm:             bwm,
		hpHelper:        hpHelper,
		hubLayerChainID: new(big.Int).SetUint64(hubLayerChainID),
		ignoreSigners:   map[common.Address]int{},
		stakemanager:    stakemanager,
		topic:           topic,
		sub:             sub,
		log:             log.New("worker", "p2p"),

		outboundSem: semaphore.NewWeighted(int64(cfg.OutboundLimits.Concurrency)),
		inboundSem:  semaphore.NewWeighted(int64(cfg.InboundLimits.Concurrency)),
		outboundThrot: rate.NewLimiter(
			rate.Limit(cfg.OutboundLimits.Throttling), cfg.OutboundLimits.Throttling),
		inboundThrot: rate.NewLimiter(
			rate.Limit(cfg.InboundLimits.Throttling), cfg.InboundLimits.Throttling),

		meterPubsubSubscribed: meter.GetOrRegisterCounter([]string{"p2p", "pubsub", "subscribed"}, ""),
		meterPubsubUnknownMsg: meter.GetOrRegisterCounter([]string{"p2p", "pubsub", "unknown", "messages"}, ""),
		meterPubsubWorkers:    meter.GetOrRegisterCounter([]string{"p2p", "pubsub", "workers"}, ""),
		meterPubsubJobs:       meter.GetOrRegisterGauge([]string{"p2p", "pubsub", "jobs"}, ""),
		meterStreamOpend:      meter.GetOrRegisterCounter([]string{"p2p", "stream", "opened"}, ""),
		meterStreamHandled:    meter.GetOrRegisterCounter([]string{"p2p", "stream", "handled"}, ""),
		meterStreamClosed:     meter.GetOrRegisterCounter([]string{"p2p", "stream", "closed"}, ""),
		meterStreamWrites:     meter.GetOrRegisterCounter([]string{"p2p", "stream", "writes"}, ""),
		meterStreamReads:      meter.GetOrRegisterCounter([]string{"p2p", "stream", "reads"}, ""),
		meterStreamUnknownMsg: meter.GetOrRegisterCounter([]string{"p2p", "stream", "unknown", "messages"}, ""),
		meterHolePunchSuccess: meter.GetOrRegisterCounter([]string{"p2p", "holepunch", "successes"}, ""),
		meterHolePunchErrs:    meter.GetOrRegisterCounter([]string{"p2p", "holepunch", "errors"}, ""),
		meterStreamOpenErrs:   meter.GetOrRegisterCounter([]string{"p2p", "stream", "open", "errors"}, ""),
		meterStreamReadErrs:   meter.GetOrRegisterCounter([]string{"p2p", "stream", "read", "errors"}, ""),
		meterStreamWriteErrs:  meter.GetOrRegisterCounter([]string{"p2p", "stream", "write", "errors"}, ""),
		meterPeers:            meter.GetOrRegisterGauge([]string{"p2p", "peers"}, ""),
		meterTCPConnections:   meter.GetOrRegisterGauge([]string{"p2p", "tcp", "connections"}, ""),
		meterUDPConnections:   meter.GetOrRegisterGauge([]string{"p2p", "udp", "connections"}, ""),
		meterRelayConnections: meter.GetOrRegisterGauge([]string{"p2p", "relay", "connections"}, ""),
		meterRelayHopStreams:  meter.GetOrRegisterGauge([]string{"p2p", "relayhop", "streams"}, ""),
		meterRelayStopStreams: meter.GetOrRegisterGauge([]string{"p2p", "relaystop", "streams"}, ""),
		meterVerifierStreams:  meter.GetOrRegisterGauge([]string{"p2p", "verifier", "streams"}, ""),
	}

	for _, addr := range ignoreSigners {
		worker.ignoreSigners[addr] = 1
	}

	return worker, nil
}

func (w *Node) Start(ctx context.Context) {
	defer w.topic.Close()
	defer w.sub.Cancel()
	w.h.SetStreamHandler(streamProtocol, w.newStreamHandler(ctx))

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.meterLoop(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.publishLoop(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		w.subscribeLoop(ctx)
	}()

	w.showBootstrapLog()
	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *Node) PeerID() peer.ID                  { return w.h.ID() }
func (w *Node) Host() host.Host                  { return w.h }
func (w *Node) Routing() routing.Routing         { return w.dht }
func (w *Node) HolePunchHelper() HolePunchHelper { return w.hpHelper }

func (w *Node) meterLoop(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 15)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			nwstat := newNetworkStatus(w.h)
			w.meterTCPConnections.Set(float64(nwstat.connections.tcp))
			w.meterUDPConnections.Set(float64(nwstat.connections.udp))
			w.meterRelayConnections.Set(float64(nwstat.connections.relay))
			w.meterRelayHopStreams.Set(float64(nwstat.streams.hop))
			w.meterRelayStopStreams.Set(float64(nwstat.streams.stop))
			w.meterVerifierStreams.Set(float64(nwstat.streams.verifier))
			w.meterPeers.Set(float64(w.h.Peerstore().Peers().Len()))
		}
	}
}

func (w *Node) publishLoop(ctx context.Context) {
	ticker := time.NewTicker(w.cfg.PublishInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.publishLatestSignatures(ctx)
		}
	}
}

func (w *Node) subscribeLoop(ctx context.Context) {
	type job struct {
		ctx    context.Context
		cancel context.CancelFunc
		peer   peer.ID
		db     *signatureDB
		remote pb.ISignature
		logctx []any
	}
	optimismDB := &signatureDB{&wrappedOptimismDB{w.db.Optimism}}
	opstackDB := &signatureDB{&wrappedOpstackDB{w.db.OPStack}}

	// Storing workers and jobs.
	workers := util.NewWorkerGroup(100)
	procs := &sync.Map{}

	for {
		peer, msg, err := subscribe(ctx, w.sub, w.h.ID())
		if errors.Is(err, context.Canceled) {
			// worker stopped
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else if err != nil {
			w.log.Error("Failed to subscribe", "peer", peer, "err", err)
			continue
		}
		w.meterPubsubSubscribed.Incr()

		var jobs []*job
		if t := msg.GetOptimismSignatureExchange(); t != nil {
			for _, remote := range t.Latests {
				jobs = append(jobs, &job{peer: peer, db: optimismDB, remote: remote})
			}
		}
		if t := msg.GetOpstackSignatureExchange(); t != nil {
			for _, remote := range t.Latests {
				jobs = append(jobs, &job{peer: peer, db: opstackDB, remote: remote})
			}
		}
		if len(jobs) == 0 {
			w.log.Warn("Unsupported pubsub message", "peer", peer, "err", err)
			w.meterPubsubUnknownMsg.Incr()
			continue
		}

		for _, jobx := range jobs {
			signer := common.BytesToAddress(jobx.remote.GetSigner())
			if _, ok := w.ignoreSigners[signer]; ok {
				continue // self created signature
			} else if w.stakemanager.StakeBySigner(signer).Cmp(tenMillionEther) == -1 {
				continue // low stake amount
			}

			// add new worker
			wname := signer.Hex()
			if !workers.Has(wname) {
				workers.AddWorker(ctx, wname, func(_ context.Context, rname string, data interface{}) {
					job := data.(*job)
					defer job.cancel()

					procs.Store(rname, job)
					defer procs.Delete(rname)

					w.handleSignatureExchangeFromPubSub(job.ctx, job.db, job.peer, job.remote)
					w.meterPubsubJobs.Decr()
				})
				w.meterPubsubWorkers.Incr()
			}

			if data, ok := procs.Load(wname); ok {
				proc := data.(*job)
				if peer == proc.peer {
					continue
				}
				if strings.Compare(jobx.remote.GetId(), proc.remote.GetId()) < 1 {
					w.log.Debug("Skipped old signature",
						append(proc.logctx, "skipped-peer", peer, "skipped-id", jobx.remote.GetId())...)
					continue
				}

				w.log.Info("Worker canceled because newer signature were received",
					append(proc.logctx, "newer-peer", peer, "newer-id", jobx.remote.GetId())...)
				proc.cancel()
			}

			jobx.ctx, jobx.cancel = context.WithCancel(ctx)
			jobx.logctx = []any{"peer", peer, "signer", wname, "remote-id", jobx.remote.GetId()}

			workers.Enqueue(wname, jobx)
			w.meterPubsubJobs.Incr()
		}
	}
}

func (w *Node) newStreamHandler(ctx context.Context) network.StreamHandler {
	optimismDB := &signatureDB{&wrappedOptimismDB{w.db.Optimism}}
	opstackDB := &signatureDB{&wrappedOpstackDB{w.db.OPStack}}

	return func(s network.Stream) {
		defer w.closeStream(s)

		w.meterStreamHandled.Incr()

		peer := s.Conn().RemotePeer()
		for {
			m, err := w.readStream(s)
			if t, ok := err.(*ReadWriteError); ok {
				w.log.Debug("Failed to read stream message", "peer", peer, "err", t)
				return
			} else if err != nil {
				w.log.Debug(err.Error(), "peer", peer)
				continue
			}

			var (
				db               *signatureDB
				commonRequests   []pb.ICommonSignatureRequest
				exchangeRequests []pb.ISignatureRequest
			)
			switch t := m.Body.(type) {
			case *pb.Stream_FindCommonOptimismSignature:
				// received FindCommonOptimismSignature request
				db = optimismDB
				for _, req := range t.FindCommonOptimismSignature.Locals {
					commonRequests = append(commonRequests, req)
				}
				if len(commonRequests) > 0 {
					// TODO
					// w.log.Info("Received FindCommonOptimismSignature request",
					// 	"from", remotes[0].Id, "to", remotes[len(remotes)-1].Id)
				}
			case *pb.Stream_FindCommonOpstackSignature:
				// received FindCommonOpstackSignature request
				db = opstackDB
				for _, req := range t.FindCommonOpstackSignature.Locals {
					commonRequests = append(commonRequests, req)
				}
			case *pb.Stream_OptimismSignatureExchange:
				// received OptimismSignatureExchange request
				db = optimismDB
				for _, req := range t.OptimismSignatureExchange.Requests {
					exchangeRequests = append(exchangeRequests, req)
				}
			case *pb.Stream_OpstackSignatureExchange:
				// received OpstackSignatureExchange request
				db = opstackDB
				for _, req := range t.OpstackSignatureExchange.Requests {
					exchangeRequests = append(exchangeRequests, req)
				}
			case *pb.Stream_Eom:
				// received last message
				return
			default:
				w.log.Warn("Received an unknown message", "peer", peer)
				w.meterStreamUnknownMsg.Incr()
				return
			}

			var disconnect bool
			if len(commonRequests) > 0 {
				disconnect = w.handleFindCommonSignatureRequests(db, s, commonRequests)
			} else if len(exchangeRequests) > 0 {
				disconnect = w.handleSignatureExchangeRequests(ctx, db, s, exchangeRequests)
			}

			if disconnect {
				return
			}
		}
	}
}

func (w *Node) handleSignatureExchangeFromPubSub(
	ctx context.Context,
	db *signatureDB,
	sender peer.ID,
	remote pb.ISignature,
) {
	signer := common.BytesToAddress(remote.GetSigner())
	logctx := []interface{}{
		"peer", sender,
		"signer", signer,
		"remote-latest-id", remote.GetId(),
		"remote-latest-previous-id", remote.GetPreviousId(),
		// "remote-latest-index", remote.BatchIndex, // TODO
	}

	if err := verifyULID(remote.GetId()); err != nil {
		w.log.Error("Invalid signature id", append(logctx, "err", err)...)
		return
	} else if err := db.verifySignature(w.hubLayerChainID, remote); err != nil {
		w.log.Error("Invalid signature", append(logctx, "err", err)...)
		return
	}

	localID, err := db.findLatestSignatureId(signer)
	if err != nil {
		w.log.Error("Failed to find the latest signature", append(logctx, "err", err)...)
		return
	} else if localID != nil && strings.Compare(*localID, remote.GetId()) == 1 {
		// fully synchronized or less than local
		return
	}

	// open stream to peer
	var s network.Stream
	openStream := func() error {
		if ss, err := w.openStream(ctx, sender); err != nil {
			return err
		} else {
			s = ss
			return nil
		}
	}
	returned := make(chan any)
	defer func() { close(returned) }()
	go func() {
		select {
		case <-ctx.Done():
			// canceled because newer signature were received
		case <-returned:
		}
		if s != nil {
			w.closeStream(s)
		}
	}()

	var idAfter string
	if localID == nil {
		w.log.Info("Request all signatures", logctx...)
	} else {
		if openStream() != nil {
			return
		}
		logctx = append(logctx, "local-id", localID)
		if found, err := w.findCommonLatestSignature(ctx, db, s, signer); err == nil {
			if rSigner := common.BytesToAddress(found.GetSigner()); rSigner != signer {
				w.log.Error("Signer does not match", append(logctx, "remote-signer", rSigner)...)
				return
			}

			idAfter = found.GetId()
			w.log.Info("Found common signature from peer",
				append(logctx, "found-id", idAfter, "found-previous-id", found.GetPreviousId())...)
		} else if errors.Is(err, database.ErrNotFound) {
			if parsed, err := ulid.ParseStrict(*localID); err == nil {
				// Prevent out-of-sync by specifying the ID of 1 second ago
				ms := parsed.Time() - 1000
				idAfter = ulid.MustNew(ms, ulid.DefaultEntropy()).String()
				logctx = append(logctx, "created-after", time.UnixMilli(int64(ms)))
			} else {
				w.log.Error("Failed to parse ULID", append(logctx, "err", err)...)
				return
			}
		} else {
			return
		}

		w.log.Info("Request signatures", append(logctx, "id-after", idAfter)...)
	}

	// send request to peer
	if s == nil && openStream() != nil {
		return
	}
	if err = w.writeStream(s, db.getSignatureExchangeRequest(signer, idAfter)); err != nil {
		w.log.Error("Failed to send signature request", append(logctx, "err", err)...)
		return
	}

	w.handleSignatureExchangeResponses(db, s)
}

func (w *Node) handleSignatureExchangeRequests(
	ctx context.Context,
	db *signatureDB,
	s network.Stream,
	requests []pb.ISignatureRequest,
) (disconnect bool) {
	peerID := s.Conn().RemotePeer()
	logctx := []interface{}{"peer", peerID}

	// number of signatures finds from database
	limit := w.cfg.InboundLimits.Throttling / w.cfg.InboundLimits.Concurrency

	// sending time limit
	isTimeup, timePenalty := func() (func() bool, func()) {
		limit := time.Now().Add(w.cfg.InboundLimits.MaxSendTime)
		return func() bool {
				return time.Now().After(limit)
			}, func() {
				limit = limit.Add(-(w.cfg.InboundLimits.MaxSendTime / 3))
			}
	}()

	// By finely acquiring the semaphore, it prevents
	// other peers from being blocked for a long time.
	sem := util.NewReleaseGuardSemaphore(w.inboundSem)
	defer sem.ReleaseALL()

	for _, req := range requests {
		signer := common.BytesToAddress(req.GetSigner())
		if w.stakemanager.StakeBySigner(signer).Cmp(tenMillionEther) == -1 {
			continue // low stake amount
		}

		logctx := append(logctx, "signer", signer, "id-after", req.GetIdAfter())
		w.log.Info("Received signature request", logctx...)

		for offset := 0; ; offset += limit {
			if isTimeup() {
				w.log.Warn("Time up", logctx...)
				return true
			} else if err := sem.Acquire(ctx, 1); err != nil {
				w.log.Error("Failed to acquire inbound semaphore", append(logctx, "err", err)...)
				return true
			}

			// get latest signatures for each requested signer
			m, count, err := db.getSignatureExchangeResponse(signer, req.GetIdAfter(), limit, offset)
			sem.ReleaseALL()
			if err != nil {
				w.log.Error("Failed to find requested signatures", append(logctx, "err", err)...)
				break
			} else if count < 1 {
				break // reached the last
			}

			w.throttling(w.inboundThrot, count,
				"in", "handleOptimismSignatureExchangeRequest", "peer", peerID)

			// send response to peer
			if err := w.writeStream(s, m); err != nil {
				w.log.Error("Failed to send signatures", append(logctx, "err", err)...)
				return true
			}
			w.log.Info("Sent signatures", append(logctx, "sents", count)...)

			// wait for received notify
			if m, err = w.readStream(s); err == nil && bytes.Equal(m.GetMisc(), misc_SIGRECEIVED) {
				w.log.Info("Received notification of receipt", logctx...)
			} else {
				timePenalty()
			}
		}
	}

	return false
}

func (w *Node) handleSignatureExchangeResponses(db *signatureDB, s network.Stream) {
	peerID := s.Conn().RemotePeer()
	logctx := []interface{}{"peer", peerID}

	for {
		m, err := w.readStream(s)
		if err != nil {
			w.log.Debug("Failed to read stream message", append(logctx, "err", err)...)
			return
		}

		// TODO
		var responses []pb.ISignature
		switch t := m.Body.(type) {
		case *pb.Stream_OptimismSignatureExchange:
			for _, res := range t.OptimismSignatureExchange.Responses {
				responses = append(responses, res)
			}
		case *pb.Stream_OpstackSignatureExchange:
			for _, res := range t.OpstackSignatureExchange.Responses {
				responses = append(responses, res)
			}
		case *pb.Stream_Eom:
			return
		default:
			w.log.Warn("Received an unknown message", logctx...)
			w.meterStreamUnknownMsg.Incr()
			return
		}
		if len(responses) == 0 {
			return
		}

		for _, res := range responses {
			signer := common.BytesToAddress(res.GetSigner())
			id, previousId := res.GetId(), res.GetPreviousId()
			logctx := append(logctx, "signer", signer, "id", id, "previous-id", previousId)
			// TODO
			// "scc", scc,
			// "index", res.BatchIndex)

			var (
				contract common.Address
				index    uint64
			)
			switch t := res.(type) {
			case *pb.OptimismSignature:
				contract = common.BytesToAddress(t.Scc)
				index = t.BatchIndex
			case *pb.OpstackSignature:
				contract = common.BytesToAddress(t.L2Oo)
				index = t.L2OutputIndex
			default:
				w.log.Error("Unknown signature", logctx...)
				return
			}
			logctx = append(logctx, "contract", contract, "index", index)

			if err := verifyULID(id); err != nil {
				w.log.Error("Invalid signature id", append(logctx, "err", err)...)
				return
			} else if err := db.verifySignature(w.hubLayerChainID, res); err != nil {
				w.log.Error("Invalid signature", append(logctx, "err", err)...)
				return
			}
			if _, ignore := w.ignoreSigners[signer]; ignore {
				w.log.Info("Ignored", logctx...)
				return
			}

			// deduplication
			if has, _ := db.hasSignature(id, &previousId); has {
				continue
			}

			// check if local is newer
			if locals, err := db.findSignatures(nil, &signer, &contract, &index, 1, 0); err != nil {
				w.log.Error("Failed to find local signature", append(logctx, "err", err)...)
				return
			} else if locals.len() > 0 && strings.Compare(locals.get(0).getID(), res.GetId()) == 1 {
				continue
			}

			if previousId != "" {
				_, err := db.findSignatureByID(previousId)
				if errors.Is(err, database.ErrNotFound) {
					w.log.Warn("Previous ID does not exist", logctx...)
				} else if err != nil {
					w.log.Error("Failed to find previous signature", append(logctx, "err", err)...)
					return
				}
			}

			if err := db.saveSignature(res); err != nil {
				w.log.Error("Failed to save signature", append(logctx, "err", err)...)
				return
			}
			w.log.Info("Received new signature", logctx...)
		}

		// send received notify
		w.writeStream(s, &pb.Stream{Body: &pb.Stream_Misc{Misc: misc_SIGRECEIVED}})
	}
}

func (w *Node) handleFindCommonSignatureRequests(
	db *signatureDB,
	s network.Stream,
	requests []pb.ICommonSignatureRequest,
) (disconnect bool) {
	var (
		m     *pb.Stream
		id    string
		found bool
		err   error
	)
	for _, req := range requests {
		m, id, found, err = db.getFindCommonSignatureResponse(req)
		if errors.Is(err, database.ErrNotFound) {
			continue
		}
		if err != nil {
			w.log.Error("Failed to find signature", "remote-id", req.GetId(), "err", err)
			return true
		}
		if found {
			break
		}
	}

	if err := w.writeStream(s, m); err != nil {
		w.log.Error("Failed to send FindCommonOptimismSignature response", "err", err)
		return true
	}

	if found {
		w.log.Info("Sent FindCommonOptimismSignature response", "found-id", id)
	} else {
		w.log.Info("Sent FindCommonOptimismSignature response", "found-id", nil)
	}
	return false
}

// Find the latest signature of the same ID and PreviousID from peer
func (w *Node) findCommonLatestSignature(
	ctx context.Context,
	db *signatureDB,
	s network.Stream,
	signer common.Address,
) (pb.ISignature, error) {
	peerID := s.Conn().RemotePeer()
	logctx := []interface{}{"peer", peerID, "signer", signer}
	limit := w.cfg.OutboundLimits.Throttling / w.cfg.OutboundLimits.Concurrency

	sem := util.NewReleaseGuardSemaphore(w.outboundSem)
	defer sem.ReleaseALL()

	for offset := 0; ; offset += limit {
		if err := sem.Acquire(ctx, 1); err != nil {
			w.log.Error("Failed to acquire outbound semaphore", append(logctx, "err", err)...)
			return nil, err
		}

		// find local latest signatures (order by: id desc)
		m, count, from, to, err := db.getFindCommonSignatureRequest(signer, limit, offset)
		sem.ReleaseALL()
		if err != nil {
			w.log.Error("Failed to find latest signatures", append(logctx, "err", err)...)
			return nil, err
		}

		if count < 1 {
			break // reached the last
		}
		w.throttling(w.outboundThrot, count, "in", "findCommonLatestSignature", "peer", peerID)

		logctx := append(logctx, "from", from, "to", to)

		// send request
		if err = w.writeStream(s, m); err != nil {
			w.log.Error(
				"Failed to send FindCommonSignature request",
				append(logctx, "err", err)...)
			return nil, err
		}
		w.log.Info("Sent FindCommonSignature request", logctx...)

		// read response
		res, err := w.readStream(s)
		if err != nil {
			w.log.Error("Failed to read stream message", append(logctx, "err", err)...)
			return nil, err
		}

		sig, found, err := db.handleFindCommonSignatureResponse(res)
		if err != nil {
			w.log.Error(err.Error(), logctx...)
			return nil, err
		}
		if found {
			return sig, nil
		}
	}

	w.log.Warn("Common signature not found", logctx...)
	return nil, database.ErrNotFound
}

// TODO
func (w *Node) publishLatestSignatures(ctx context.Context) {
	var optimisms []*database.OptimismSignature
	if sigs, err := w.db.Optimism.FindLatestSignaturePerSigners(); err != nil {
		w.log.Error("Failed to find latest optimism signatures", "err", err)
	} else {
		for _, sig := range sigs {
			if w.stakemanager.StakeBySigner(sig.Signer.Address).Cmp(tenMillionEther) >= 0 {
				optimisms = append(optimisms, sig)
			}
		}
	}

	var opstacks []*database.OpstackSignature
	if sigs, err := w.db.OPStack.FindLatestSignaturePerSigners(); err != nil {
		w.log.Error("Failed to find latest opstack signatures", "err", err)
	} else {
		for _, sig := range sigs {
			if w.stakemanager.StakeBySigner(sig.Signer.Address).Cmp(tenMillionEther) >= 0 {
				opstacks = append(opstacks, sig)
			}
		}
	}

	w.PublishSignatures(ctx, optimisms, opstacks)
}

func (w *Node) PublishSignatures(
	ctx context.Context,
	optimisms wrappedOptimismSignatures,
	opstacks wrappedOpstackSignatures,
) {
	var msg pb.PubSub
	if len(optimisms) > 0 {
		msg.OptimismSignatureExchange = &pb.OptimismSignatureExchange{
			Latests: optimisms.ProtoSig(),
		}
	}
	if len(opstacks) > 0 {
		msg.OpstackSignatureExchange = &pb.OpstackSignatureExchange{
			Latests: opstacks.ProtoSig(),
		}
	}
	if msg.OptimismSignatureExchange == nil && msg.OpstackSignatureExchange == nil {
		return
	}

	if err := publish(ctx, w.topic, &msg); err != nil {
		w.log.Error("Failed to publish latest signatures", "err", err)
	} else {
		w.log.Info("Publish latest signatures",
			"len(optimisms)", len(optimisms), "len(opstacks)", len(opstacks))
	}
}

func (w *Node) openStream(ctx context.Context, peer peer.ID) (network.Stream, error) {
	// If holepunch is available, attempt a direct connection.
	if !HasDirectConnection(w.h, peer) && w.hpHelper.Available(w.h) {
		if err := <-w.hpHelper.HolePunch(ctx, w.h, peer, DefaultHolePunchTimeout); err != nil {
			if !errors.Is(err, ErrPeerNotSupportHolePunch) {
				w.meterHolePunchErrs.Incr()
			}
		} else {
			w.meterHolePunchSuccess.Incr()
		}
	}

	// Note: `WithUseTransient` is required to open a stream via circuit relay.
	s, err := w.h.NewStream(network.WithUseTransient(ctx, streamProtocol), peer, streamProtocol)
	if err != nil {
		w.log.Error("Failed to open stream", "peer", peer, "err", err)
		w.meterStreamOpenErrs.Incr()
		return nil, err
	}

	w.meterStreamOpend.Incr()
	return s, nil
}

func (w *Node) writeStream(s network.Stream, m *pb.Stream) error {
	if w.cfg.StreamTimeout > 0 {
		s.SetWriteDeadline(time.Now().Add(w.cfg.StreamTimeout))
		defer s.SetWriteDeadline(time.Time{})
	}

	err := writeStream(s, m)
	_, isRWErr := err.(*ReadWriteError)
	_, isEOM := m.Body.(*pb.Stream_Eom)
	if err == nil {
		w.meterStreamWrites.Incr()
	} else if isRWErr && !isEOM {
		w.meterStreamWriteErrs.Incr()
	}
	return err
}

func (w *Node) readStream(s network.Stream) (m *pb.Stream, err error) {
	if w.cfg.StreamTimeout > 0 {
		s.SetReadDeadline(time.Now().Add(w.cfg.StreamTimeout))
		defer s.SetReadDeadline(time.Time{})
	}

	m, err = readStream(s)
	if err == nil {
		w.meterStreamReads.Incr()
		return m, nil
	} else if _, ok := err.(*ReadWriteError); ok {
		w.meterStreamReadErrs.Incr()
	}
	return nil, err
}

func (w *Node) closeStream(s network.Stream) {
	closeStream(s)
	w.meterStreamClosed.Incr()
}

func (w *Node) showBootstrapLog() {
	listens := []string{}
	for _, ma := range w.h.Network().ListenAddresses() {
		listens = append(listens, ma.String())
	}
	w.log.Info("Listening on: " + strings.Join(listens, ","))
	w.log.Info("Appended announce addresses: " + strings.Join(w.cfg.AppendAnnounce, ","))
	w.log.Info("No announce addresses: " + strings.Join(w.cfg.NoAnnounce, ","))
	w.log.Info("Connection filter addresses: " + strings.Join(w.cfg.ConnectionFilter, ","))
	if w.cfg.Transports.TCP {
		w.log.Info("Enabled TCP transport")
	}
	if w.cfg.Transports.QUIC {
		w.log.Info("Enabled QUIC transport")
	}
	w.log.Info("Bootnodes: " + strings.Join(w.cfg.Bootnodes, ","))
	w.log.Info("Enabled NAT Travasal features",
		"upnp", w.cfg.NAT.UPnP, "autonat", w.cfg.NAT.AutoNAT, "holepunch", w.hpHelper.Enabled())
	if w.cfg.RelayService.Enable {
		w.log.Info("Enabled circuit relay service")
	}
	if w.cfg.RelayClient.Enable {
		w.log.Info("Enabled circuit relay client, relay nodes: " + strings.Join(w.cfg.RelayClient.RelayNodes, ","))
	}
	w.log.Info("Worker started", "id", w.h.ID(),
		"publish-interval", w.cfg.PublishInterval,
		"stream-timeout", w.cfg.StreamTimeout,
		"outbound-limits-concurrency", w.cfg.OutboundLimits.Concurrency,
		"outbound-limits-throttling", w.cfg.OutboundLimits.Throttling,
		"inbound-limits-concurrency", w.cfg.InboundLimits.Concurrency,
		"inbound-limits-maxsendtime", w.cfg.InboundLimits.MaxSendTime,
		"inbound-limits-throttling", w.cfg.InboundLimits.Throttling,
	)
}

func (w *Node) throttling(limiter *rate.Limiter, num int, logCtx ...any) {
	rsv := limiter.ReserveN(time.Now(), num)
	if !rsv.OK() {
		w.log.Error("num is greater than burst", logCtx...)
		return
	}

	sleep := rsv.Delay()
	if sleep > 0 {
		w.log.Warn("Throttling", append(logCtx, "sleep", sleep)...)
		time.Sleep(sleep)
	}
}

// Write protobuf message to libp2p stream.
func writeStream(s io.Writer, m *pb.Stream) error {
	data, err := proto.Marshal(m)
	if err != nil {
		return err
	}

	data, err = compress(data)
	if err != nil {
		return err
	}

	// Note: Intentionally not closing with `Close()` as it would also close the stream.
	if err := msgio.NewWriter(s).WriteMsg(data); err != nil {
		return &ReadWriteError{err}
	}

	return nil
}

// Read protobuf message from libp2p stream.
// Note: Will wait forever, should cancel.
func readStream(s io.Reader) (*pb.Stream, error) {
	reader := msgio.NewReader(s)
	msg, err := reader.ReadMsg()
	if err != nil {
		return nil, &ReadWriteError{err}
	}

	// Note: Forgetting to call `ReleaseMsg()` can result
	// in high memory consumption within libp2p/go-buffer-pool.
	defer reader.ReleaseMsg(msg)

	data, err := decompress(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress stream message: %w", err)
	}

	var m pb.Stream
	if err := proto.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stream message: %w", err)
	}

	return &m, nil
}

// Send end-of-message and close libp2p stream.
func closeStream(s network.Stream) {
	s.SetWriteDeadline(time.Now().Add(time.Second / 2))
	defer s.SetWriteDeadline(time.Time{})

	writeStream(s, eom)
	s.Close()
}

// Publish new message.
func publish(ctx context.Context, topic *ps.Topic, m *pb.PubSub) error {
	data, err := proto.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal pubsub message: %w", err)
	}

	if data, err = compress(data); err != nil {
		return fmt.Errorf("failed to compress pubsub message: %w", err)
	}
	if err := topic.Publish(ctx, data); err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

// Subscribe new message.
// Note: Will wait forever, should cancel.
func subscribe(
	ctx context.Context,
	sub *ps.Subscription,
	self peer.ID,
) (peer.ID, *pb.PubSub, error) {
	recv, err := sub.Next(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("failed to subscribe pubsub message: %w", err)
	}

	if recv.ReceivedFrom == self || recv.GetFrom() == self {
		return "", nil, errSelfMessage
	}

	data, err := decompress(recv.Data)
	if err != nil {
		return "", nil, fmt.Errorf("failed to decompress pubsub message: %w", err)
	}

	var m pb.PubSub
	if err = proto.Unmarshal(data, &m); err != nil {
		return "", nil, fmt.Errorf("failed to unmarshal pubsub message: %w", err)
	}

	return recv.GetFrom(), &m, nil
}

func verifyULID(id string) error {
	if ulid, err := ulid.ParseStrict(id); err != nil {
		return err
	} else if ulid.Time() > uint64(time.Now().UnixMilli()) {
		return fmt.Errorf("future ulid: %s, timestamp: %d", id, ulid.Time())
	}
	return nil
}
