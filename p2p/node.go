package p2p

import (
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
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/metrics"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	kaddht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	msgio "github.com/libp2p/go-msgio"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oasysgames/oasys-optimism-verifier/verifier"
	"github.com/oklog/ulid/v2"
	"google.golang.org/protobuf/proto"
)

const (
	pubsubTopic    = "/oasys-optimism-verifier/pubsub/1.0.0"
	streamProtocol = "/oasys-optimism-verifier/stream/1.0.0"
)

const (
	warnQueueLen = 30
)

var (
	eom = &pb.Stream{Body: &pb.Stream_Eom{Eom: nil}}

	errUnavailableStream = errors.New("unavailable stream")
	errSelfMessage       = errors.New("self message")
)

type Node struct {
	cfg             *config.P2P
	db              *database.Database
	h               host.Host
	dht             *kaddht.IpfsDHT
	bwm             *metrics.BandwidthCounter
	hubLayerChainID *big.Int
	ignoreSigners   map[common.Address]int

	topic *ps.Topic
	sub   *ps.Subscription
	log   log.Logger
}

func NewNode(
	cfg *config.P2P,
	db *database.Database,
	host host.Host,
	dht *kaddht.IpfsDHT,
	bwm *metrics.BandwidthCounter,
	hubLayerChainID uint64,
	ignoreSigners []common.Address,
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
		hubLayerChainID: new(big.Int).SetUint64(hubLayerChainID),
		ignoreSigners:   map[common.Address]int{},
		topic:           topic,
		sub:             sub,
		log:             log.New("worker", "p2p"),
	}
	worker.h.SetStreamHandler(streamProtocol, worker.handleStream)

	for _, addr := range ignoreSigners {
		worker.ignoreSigners[addr] = 1
	}

	return worker, nil
}

func (w *Node) Start(ctx context.Context) {
	defer w.topic.Close()
	defer w.sub.Cancel()

	wg := &sync.WaitGroup{}

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

	w.log.Info("Worker started", "id", w.PeerID(),
		"publish-interval", w.cfg.PublishInterval, "stream-timeout", w.cfg.StreamTimeout)
	wg.Wait()
	w.log.Info("Worker stopped")
}

func (w *Node) PeerID() peer.ID {
	return w.h.ID()
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
	type Job struct {
		from   peer.ID
		db     signatureDatabase
		remote pbSignature
	}

	wg := util.NewWorkerGroup(100) // each signer address
	running := &sync.Map{}         // stores IDs in process for each signer
	optimismDB := &optimismDatabase{db: w.db.Optimism}
	opstackDB := &opstackDatabase{db: w.db.OPStack}

	for {
		from, msg, err := subscribe(ctx, w.sub, w.h.ID())
		if errors.Is(err, context.Canceled) {
			// worker stopped
			return
		} else if errors.Is(err, errSelfMessage) {
			continue
		} else if err != nil {
			w.log.Error("Failed to subscribe", "peer", from, "err", err)
			continue
		}

		var jobs []*Job
		if t := msg.GetOptimismSignatureExchange(); t != nil {
			for _, remote := range t.Latests {
				jobs = append(jobs, &Job{from: from, db: optimismDB, remote: remote})
			}
		}
		if t := msg.GetOpstackSignatureExchange(); t != nil {
			for _, remote := range t.Latests {
				jobs = append(jobs, &Job{from: from, db: opstackDB, remote: remote})
			}
		}
		if len(jobs) == 0 {
			w.log.Warn("Unsupported pubsub message", "peer", from, "err", err)
			continue
		}

		for _, job := range jobs {
			wname := common.BytesToAddress(job.remote.GetSigner()).Hex()

			// skip if older than the ID being processed
			if proc, ok := running.Load(wname); ok &&
				strings.Compare(job.remote.GetId(), proc.(string)) < 1 {
				w.log.Debug("Skip pubsub",
					"peer", from, "signer", wname,
					"processed-id", proc, "remote-id", job.remote.GetId())
				continue
			}
			running.Store(wname, job.remote.GetId())

			// add new worker
			if !wg.Has(wname) {
				handler := func(ctx context.Context, rname string, data interface{}) {
					defer running.Delete(rname)

					if job, ok := data.(*Job); ok {
						st := time.Now()
						w.handleSignatureExchangeFromPubSub(ctx, job.db, job.from, job.remote)
						w.log.Debug("Worked pubsub",
							"peer", from, "signer", rname,
							"elapsed", time.Since(st), "remote-id", job.remote.GetId())
					}
				}
				wg.AddWorker(ctx, wname, handler)
			}

			wg.Enqueue(wname, job)

			qlen := len(wg.Queue(wname))
			w.log.Debug("Enqueue pubsub",
				"peer", from, "signer", wname,
				"remote-id", job.remote.GetId(), "queue-len", qlen)
			if qlen >= warnQueueLen {
				w.log.Warn("Long queue", "signer", wname, "queue-len", qlen)
			}
		}
	}
}

func (w *Node) handleStream(s network.Stream) {
	defer closeStream(s)

	peer := s.Conn().RemotePeer()
	optimismDB := &optimismDatabase{db: w.db.Optimism}
	opstackDB := &opstackDatabase{db: w.db.OPStack}

	for {
		m, err := readStreamWithTimeout(context.Background(), s, w.cfg.StreamTimeout)
		if errors.Is(err, errUnavailableStream) {
			w.log.Error("Failed to read stream message", "peer", peer, "err", err)
			return
		} else if err != nil {
			w.log.Error(err.Error(), "peer", peer)
			continue
		}

		var (
			db     signatureDatabase
			exReqs []pbSignatureRequest
			exRess []pbSignature
			fcReqs []pbCommonSignatureRequest
		)
		switch t := m.Body.(type) {
		case *pb.Stream_OptimismSignatureExchange:
			db = optimismDB
			// received signature exchange request
			for _, req := range t.OptimismSignatureExchange.Requests {
				exReqs = append(exReqs, req)
			}
			// received signature exchange response
			for _, res := range t.OptimismSignatureExchange.Responses {
				exRess = append(exRess, res)
			}
		case *pb.Stream_OpstackSignatureExchange:
			db = opstackDB
			// received signature exchange request
			for _, req := range t.OpstackSignatureExchange.Requests {
				exReqs = append(exReqs, req)
			}
			// received signature exchange response
			for _, res := range t.OpstackSignatureExchange.Responses {
				exRess = append(exRess, res)
			}
		case *pb.Stream_FindCommonOptimismSignature:
			db = optimismDB
			// received FindCommonSignature request
			for _, req := range t.FindCommonOptimismSignature.Locals {
				fcReqs = append(fcReqs, req)
			}
		case *pb.Stream_FindCommonOpstackSignature:
			db = opstackDB
			// received FindCommonSignature request
			for _, req := range t.FindCommonOpstackSignature.Locals {
				fcReqs = append(fcReqs, req)
			}
		case *pb.Stream_Eom:
			// received last message
			return
		default:
			w.log.Warn("Received an unknown message", "peer", peer)
			return
		}

		if len(exReqs) > 0 {
			w.handleSignatureExchangeRequests(db, s, exReqs)
		}
		if len(exRess) > 0 {
			w.handleSignatureExchangeResponses(db, s, exRess)
		}
		if len(fcReqs) > 0 {
			w.handleFindCommonSignatureRequests(db, s, fcReqs)
		}
	}
}

func (w *Node) handleSignatureExchangeFromPubSub(
	ctx context.Context,
	db signatureDatabase,
	sender peer.ID,
	remote pbSignature,
) {
	signer := common.BytesToAddress(remote.GetSigner())
	logctx := []interface{}{
		"peer", sender,
		"signer", signer,
		"remote-id", remote.GetId(),
		"remote-previous-id", remote.GetPreviousId(),
		// "index", remote.BatchIndex, // TODO
	}

	if ok, err := db.verifySignature(w.hubLayerChainID, remote); err != nil || !ok {
		w.log.Error("Invalid signature", append(logctx, "verify", ok, "err", err)...)
		return
	}
	if _, ok := w.ignoreSigners[signer]; ok {
		w.log.Info("Ignored signature", logctx...)
		return
	}

	localId, err := db.findLatestSignatureId(signer)
	if err != nil {
		w.log.Error("Failed to find the latest signature", append(logctx, "err", err)...)
		return
	} else if strings.Compare(localId, remote.GetId()) == 1 {
		// fully synchronized or less than local
		return
	}

	// open stream to peer
	s, err := w.h.NewStream(ctx, sender, streamProtocol)
	if err != nil {
		w.log.Error("Failed to open stream", "peer", sender, "err", err)
		return
	}
	defer closeStream(s)

	var idAfter string
	if localId == "" {
		w.log.Info("Request all signatures", logctx...)
	} else {
		if found, err := w.findCommonLatestSignature(db, s, remote); err == nil {
			remoteSigner := common.BytesToAddress(found.GetSigner())
			if remoteSigner != signer {
				w.log.Error("Signer does not match", append(logctx, "remote-signer", remoteSigner)...)
				return
			}

			idAfter = found.GetId()
			w.log.Info("Found common signature from peer",
				"signer", signer, "id", found.GetId(), "previous-id", found.GetPreviousId())
		} else {
			if localID, err := ulid.ParseStrict(localId); err == nil {
				// Prevent out-of-sync by specifying the ID of 1 second ago
				ms := localID.Time() - 1000
				idAfter = ulid.MustNew(ms, ulid.DefaultEntropy()).String()
				logctx = append(logctx, "local-id", localId, "created-after", time.UnixMilli(int64(ms)))
			} else {
				w.log.Error("Failed to parse ULID", "local-id", localId, "err", err)
				return
			}
		}

		w.log.Info("Request signatures", append(logctx, "id-after", idAfter)...)
	}

	// send request to peer
	m := db.getSignatureExchangeRequest(signer, idAfter)
	if err = writeStream(s, m); err != nil {
		w.log.Error("Failed to send signature request", "err", err)
		return
	}

	if err := writeStream(s, eom); err != nil {
		w.log.Error("Failed to send end-of-message", "err", err)
		return
	}

	// wait for signature exchange response
	w.handleStream(s)
}

func (w *Node) handleSignatureExchangeRequests(
	db signatureDatabase,
	s network.Stream,
	reqs []pbSignatureRequest,
) {
	for _, req := range reqs {
		signer := common.BytesToAddress(req.GetSigner())
		idAfter := req.GetIdAfter()
		logctx := []interface{}{"signer", signer, "id-after", idAfter}

		w.log.Info("Received signature request", logctx...)

		for limit, offset := 1000, 0; ; offset += limit {
			m, count, err := db.getSignatureExchangeResponse(signer, idAfter, limit, offset)
			if err != nil {
				w.log.Error("Failed to find requested signatures",
					append(logctx, "err", err)...)
				break
			} else if count == 0 {
				break
			}

			// send response to peer
			if err := writeStream(s, m); err != nil {
				w.log.Error("Failed to send signatures", append(logctx, "err", err)...)
				return
			}

			w.log.Info("Sent signatures", "len", count)
		}
	}
}

func (w *Node) handleSignatureExchangeResponses(
	db signatureDatabase,
	s network.Stream,
	ress []pbSignature,
) {
	for _, res := range ress {
		signer := common.BytesToAddress(res.GetSigner())
		id := res.GetId()
		previousId := res.GetPreviousId()

		// scc := common.BytesToAddress(res.Scc) // TODO
		logctx := []interface{}{"signer", signer, "id", id, "previous-id", previousId}

		if ok, err := db.verifySignature(w.hubLayerChainID, res); !ok || err != nil {
			w.log.Error("Invalid signature", append(logctx, "verify", ok, "err", err)...)
			return
		}
		if _, ok := w.ignoreSigners[signer]; ok {
			w.log.Info("Ignored", logctx...)
			return
		}

		// deduplication
		if has, _ := db.hasSignature(id, &previousId); has {
			continue
		}

		if previousId != "" {
			if has, err := db.hasSignature(res.GetPreviousId(), nil); err != nil {
				w.log.Error("Failed to find previous signature", append(logctx, "err", err)...)
				return
			} else if !has {
				w.log.Warn("Previous ID does not exist", logctx...)
				return
			}
		}

		if err := db.saveSignature(res); err != nil {
			w.log.Error("Failed to save signature", append(logctx, "err", err)...)
			return
		}

		w.log.Info("Received new signature", logctx...)
	}
}

func (w *Node) handleFindCommonSignatureRequests(
	db signatureDatabase,
	s network.Stream,
	reqs []pbCommonSignatureRequest,
) {
	var (
		res   *pb.Stream
		found bool
		err   error
	)
	for _, req := range reqs {
		res, found, err = db.getFindCommonSignatureResponse(req)
		if errors.Is(err, database.ErrNotFound) {
			continue
		}
		if err != nil {
			w.log.Error("Failed to find signature", "remote-id", req.GetId(), "err", err)
			return
		}
		if found {
			break
		}
	}

	if err := writeStream(s, res); err == nil {
		w.log.Info("Sent FindCommonOptimismSignature response", "found", found)
	} else {
		w.log.Error("Failed to send FindCommonOptimismSignature response", "err", err)
	}
}

// Find the latest signature of the same ID and PreviousID from peer
func (w *Node) findCommonLatestSignature(
	db signatureDatabase,
	s network.Stream,
	remote pbSignature,
) (pbSignature, error) {
	signer := common.BytesToAddress(remote.GetSigner())
	limit, offset := 100, 0
	for {
		logctx := []interface{}{"signer", signer}

		// find local latest signatures (order by: id desc)
		req, from, to, err := db.getFindCommonSignatureRequest(signer, limit, offset)
		if err != nil {
			w.log.Error("Failed to find latest signatures", append(logctx, "err", err)...)
			return nil, err
		}
		if req == nil {
			break
		}
		logctx = append(logctx, "from", from, "to", to)

		// send request
		if err = writeStream(s, req); err != nil {
			w.log.Error(
				"Failed to send FindCommonSignature request",
				append(logctx, "err", err)...)
			return nil, err
		}
		w.log.Info("Sent FindCommonSignature request", logctx...)

		// read response
		res, err := readStreamWithTimeout(context.Background(), s, time.Second*5)
		if errors.Is(err, context.DeadlineExceeded) {
			w.log.Warn("Timeout or peer does not support FindCommonSignature", logctx...)
			return nil, err
		} else if err != nil {
			w.log.Error("Failed to read stream message", append(logctx, "err", err)...)
			return nil, err
		}

		found, err := db.handleFindCommonSignatureResponse(res)
		if err != nil {
			w.log.Error(err.Error(), logctx...)
			return nil, err
		}
		if found != nil {
			// found!
			return found, nil
		}

		offset += limit
	}

	w.log.Warn("Common signature not found", "signer", signer)
	return nil, errors.New("not found")
}

func (w *Node) publishLatestSignatures(ctx context.Context) {
	optimisms, err := w.db.Optimism.FindLatestSignaturePerSigners()
	if err != nil {
		w.log.Error("Failed to find latest optimism signatures", "err", err)
	}
	opstacks, err := w.db.OPStack.FindLatestSignaturePerSigners()
	if err != nil {
		w.log.Error("Failed to find latest opstack signatures", "err", err)
	}
	w.PublishSignatures(ctx, optimisms, opstacks)
}

func (w *Node) PublishSignatures(
	ctx context.Context,
	optimisms optimismSignatures,
	opstacks opstackSignatures,
) {
	var msg pb.PubSub
	if len(optimisms) > 0 {
		msg.OptimismSignatureExchange = &pb.OptimismSignatureExchange{
			Latests: optimisms.toProtoBufSig(),
		}
	}
	if len(opstacks) > 0 {
		msg.OpstackSignatureExchange = &pb.OpstackSignatureExchange{
			Latests: opstacks.toProtoBufSig(),
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
	if err := msgio.NewWriter(s).WriteMsg(data); err != nil {
		return errUnavailableStream
	}

	return nil
}

// Read protobuf message from libp2p stream.
// Note: Will wait forever, should cancel.
func readStream(s io.Reader) (*pb.Stream, error) {
	data, err := msgio.NewReader(s).ReadMsg()
	if err != nil {
		return nil, errUnavailableStream
	}

	data, err = decompress(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress stream message: %w", err)
	}

	var m pb.Stream
	if err := proto.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stream message: %w", err)
	}

	return &m, nil
}

func readStreamWithTimeout(
	parent context.Context,
	s io.Reader,
	timeout time.Duration,
) (m *pb.Stream, err error) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	defer cancel()

	go func() {
		defer cancel()
		m, err = readStream(s)
	}()
	<-ctx.Done()

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return nil, context.DeadlineExceeded
	}
	return m, err
}

// Send end-of-message and close libp2p stream.
func closeStream(s network.Stream) {
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

type optimismSignatures []*database.OptimismSignature

func (sigs optimismSignatures) toProtoBufSig() (pbSigs []*pb.OptimismSignature) {
	for _, s := range sigs {
		pbSigs = append(pbSigs, &pb.OptimismSignature{
			Id:                s.ID,
			PreviousId:        s.PreviousID,
			Signer:            s.Signer.Address[:],
			Scc:               s.OptimismScc.Address[:],
			BatchIndex:        s.BatchIndex,
			BatchRoot:         s.BatchRoot[:],
			BatchSize:         s.BatchSize,
			PrevTotalElements: s.PrevTotalElements,
			ExtraData:         s.ExtraData,
			Approved:          s.Approved,
			Signature:         s.Signature[:],
		})
	}
	return pbSigs
}

type opstackSignatures []*database.OpstackSignature

func (sigs opstackSignatures) toProtoBufSig() (pbSigs []*pb.OpstackSignature) {
	for _, s := range sigs {
		pbSigs = append(pbSigs, &pb.OpstackSignature{
			Id:            s.ID,
			PreviousId:    s.PreviousID,
			Signer:        s.Signer.Address[:],
			L2Oo:          s.OpstackL2OutputOracle.Address[:],
			L2OutputIndex: s.L2OutputIndex,
			OutputRoot:    s.OutputRoot[:],
			L2BlockNumber: s.L2BlockNumber,
			L1Timestamp:   s.L1Timestamp,
			Approved:      s.Approved,
			Signature:     s.Signature[:],
		})
	}
	return pbSigs
}

type pbSignature interface {
	GetId() string         // signature id
	GetPreviousId() string // previous signature id
	GetSigner() []byte     // signer address
}

type pbSignatureRequest interface {
	GetIdAfter() string // signature id
	GetSigner() []byte  // signer address
}

type pbCommonSignatureRequest interface {
	GetId() string         // signature id
	GetPreviousId() string // previous signature id
}

type signatureDatabase interface {
	findLatestSignatureId(signer common.Address) (id string, err error)
	getSignatureExchangeRequest(signer common.Address, idAfter string) *pb.Stream
	getSignatureExchangeResponse(signer common.Address, idAfter string, limit, offset int) (msg *pb.Stream, count int, err error)
	hasSignature(id string, previousID *string) (bool, error)
	getFindCommonSignatureRequest(signer common.Address, limit, offset int) (msg *pb.Stream, from, to string, err error)
	getFindCommonSignatureResponse(req pbCommonSignatureRequest) (msg *pb.Stream, found bool, err error)
	handleFindCommonSignatureResponse(res *pb.Stream) (pbSignature, error)
	saveSignature(pbMsg interface{}) error
	verifySignature(hubLayerChainID *big.Int, pbMsg interface{}) (bool, error)
}

type optimismDatabase struct {
	db *database.OptimismDatabase
}

func (w *optimismDatabase) findLatestSignatureId(signer common.Address) (id string, err error) {
	latests, err := w.db.FindLatestSignaturesBySigner(signer, 1, 0)
	if err != nil {
		return "", err
	} else if len(latests) == 0 {
		return "", nil
	}
	return latests[0].ID, nil
}

func (w *optimismDatabase) getSignatureExchangeRequest(signer common.Address, idAfter string) *pb.Stream {
	return &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
		OptimismSignatureExchange: &pb.OptimismSignatureExchange{
			Requests: []*pb.OptimismSignatureExchange_Request{
				{
					Signer:  signer[:],
					IdAfter: idAfter,
				},
			},
		},
	}}
}

func (w *optimismDatabase) getSignatureExchangeResponse(signer common.Address, idAfter string, limit, offset int) (msg *pb.Stream, count int, err error) {
	// get latest signatures for each requested signer
	sigs, err := w.db.FindSignatures(&idAfter, &signer, nil, nil, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
		OptimismSignatureExchange: &pb.OptimismSignatureExchange{
			Responses: optimismSignatures(sigs).toProtoBufSig(),
		},
	}}, len(sigs), nil
}

func (w *optimismDatabase) hasSignature(id string, previousID *string) (bool, error) {
	sig, err := w.db.FindSignatureByID(id)
	if errors.Is(err, database.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	} else if previousID == nil {
		return true, nil
	}
	return sig.PreviousID == *previousID, nil
}

func (w *optimismDatabase) getFindCommonSignatureRequest(
	signer common.Address,
	limit, offset int,
) (msg *pb.Stream, from, to string, err error) {
	// find local latest signatures (order by: id desc)
	sigs, err := w.db.FindLatestSignaturesBySigner(signer, limit, offset)
	if err != nil {
		return nil, "", "", err
	} else if len(sigs) == 0 {
		// reached the last
		return nil, "", "", nil
	}

	// construct protobuf message
	locals := make([]*pb.FindCommonOptimismSignature_Local, len(sigs))
	for i, sig := range sigs {
		locals[i] = &pb.FindCommonOptimismSignature_Local{
			Id:         sig.ID,
			PreviousId: sig.PreviousID,
		}
	}

	return &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
		FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{Locals: locals},
	}}, sigs[0].ID, sigs[len(sigs)-1].ID, nil
}

func (w *optimismDatabase) getFindCommonSignatureResponse(req pbCommonSignatureRequest) (msg *pb.Stream, found bool, err error) {
	var pbSig *pb.OptimismSignature
	sig, err := w.db.FindSignatureByID(req.GetId())
	if err == nil && sig.PreviousID == req.GetPreviousId() {
		pbSig = optimismSignatures{sig}.toProtoBufSig()[0]
		found = true
	}

	return &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
		FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{Found: pbSig},
	}}, found, err
}

func (w *optimismDatabase) handleFindCommonSignatureResponse(res *pb.Stream) (pbSignature, error) {
	t := res.GetFindCommonOptimismSignature()
	if t == nil {
		return nil, errors.New("unexpected response")
	} else if t.Found == nil {
		return nil, nil
	}
	return t.Found, nil
}

func (w *optimismDatabase) saveSignature(pbMsg interface{}) error {
	sig, ok := pbMsg.(*pb.OptimismSignature)
	if !ok {
		return errors.New("unknown protobuf message")
	}

	_, err := w.db.SaveSignature(
		&sig.Id,
		&sig.PreviousId,
		common.BytesToAddress(sig.Signer),
		common.BytesToAddress(sig.Scc),
		sig.BatchIndex,
		common.BytesToHash(sig.BatchRoot),
		sig.BatchSize,
		sig.PrevTotalElements,
		sig.ExtraData,
		sig.Approved,
		database.BytesSignature(sig.Signature))
	return err
}

func (w *optimismDatabase) verifySignature(hubLayerChainID *big.Int, pbMsg interface{}) (bool, error) {
	sig := pbMsg.(*pb.OptimismSignature)
	if sig == nil {
		return false, errors.New("not an optimism signature")
	}
	if err := verifyULID(sig.Id); err != nil {
		return false, err
	}

	msg := verifier.NewSccMessage(
		hubLayerChainID,
		common.BytesToAddress(sig.Scc),
		new(big.Int).SetUint64(sig.BatchIndex),
		common.BytesToHash(sig.BatchRoot),
		sig.Approved)
	return msg.VerifySigner(sig.Signature, common.BytesToAddress(sig.Signer))
}

type opstackDatabase struct {
	db *database.OPStackDatabase
}

func (w *opstackDatabase) findLatestSignatureId(signer common.Address) (id string, err error) {
	latests, err := w.db.FindLatestSignaturesBySigner(signer, 1, 0)
	if err != nil {
		return "", err
	} else if len(latests) == 0 {
		return "", nil
	}
	return latests[0].ID, nil
}

func (w *opstackDatabase) getSignatureExchangeRequest(signer common.Address, idAfter string) *pb.Stream {
	return &pb.Stream{Body: &pb.Stream_OpstackSignatureExchange{
		OpstackSignatureExchange: &pb.OpstackSignatureExchange{
			Requests: []*pb.OpstackSignatureExchange_Request{
				{
					Signer:  signer[:],
					IdAfter: idAfter,
				},
			},
		},
	}}
}

func (w *opstackDatabase) getSignatureExchangeResponse(signer common.Address, idAfter string, limit, offset int) (msg *pb.Stream, count int, err error) {
	// get latest signatures for each requested signer
	sigs, err := w.db.FindSignatures(&idAfter, &signer, nil, nil, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return &pb.Stream{Body: &pb.Stream_OpstackSignatureExchange{
		OpstackSignatureExchange: &pb.OpstackSignatureExchange{
			Responses: opstackSignatures(sigs).toProtoBufSig(),
		},
	}}, len(sigs), nil
}

func (w *opstackDatabase) hasSignature(id string, previousID *string) (bool, error) {
	sig, err := w.db.FindSignatureByID(id)
	if errors.Is(err, database.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	} else if previousID == nil {
		return true, nil
	}
	return sig.PreviousID == *previousID, nil
}

func (w *opstackDatabase) getFindCommonSignatureRequest(
	signer common.Address,
	limit, offset int,
) (msg *pb.Stream, from, to string, err error) {
	// find local latest signatures (order by: id desc)
	sigs, err := w.db.FindLatestSignaturesBySigner(signer, limit, offset)
	if err != nil {
		return nil, "", "", err
	} else if len(sigs) == 0 {
		// reached the last
		return nil, "", "", nil
	}

	// construct protobuf message
	locals := make([]*pb.FindCommonOpstackSignature_Local, len(sigs))
	for i, sig := range sigs {
		locals[i] = &pb.FindCommonOpstackSignature_Local{
			Id:         sig.ID,
			PreviousId: sig.PreviousID,
		}
	}

	return &pb.Stream{Body: &pb.Stream_FindCommonOpstackSignature{
		FindCommonOpstackSignature: &pb.FindCommonOpstackSignature{Locals: locals},
	}}, sigs[0].ID, sigs[len(sigs)-1].ID, nil
}

func (w *opstackDatabase) getFindCommonSignatureResponse(req pbCommonSignatureRequest) (msg *pb.Stream, found bool, err error) {
	var pbSig *pb.OpstackSignature
	sig, err := w.db.FindSignatureByID(req.GetId())
	if err == nil && sig.PreviousID == req.GetPreviousId() {
		pbSig = opstackSignatures{sig}.toProtoBufSig()[0]
		found = true
	}

	return &pb.Stream{Body: &pb.Stream_FindCommonOpstackSignature{
		FindCommonOpstackSignature: &pb.FindCommonOpstackSignature{Found: pbSig},
	}}, found, err
}

func (w *opstackDatabase) handleFindCommonSignatureResponse(res *pb.Stream) (pbSignature, error) {
	t := res.GetFindCommonOpstackSignature()
	if t == nil {
		return nil, errors.New("unexpected response")
	} else if t.Found == nil {
		return nil, nil
	}
	return t.Found, nil
}

func (w *opstackDatabase) saveSignature(pbMsg interface{}) error {
	t, ok := pbMsg.(*pb.OpstackSignature)
	if !ok {
		return errors.New("unknown protobuf message")
	}

	_, err := w.db.SaveSignature(
		&t.Id,
		&t.PreviousId,
		common.BytesToAddress(t.Signer),
		common.BytesToAddress(t.L2Oo),
		t.L2OutputIndex,
		common.BytesToHash(t.OutputRoot),
		t.L2BlockNumber,
		t.L1Timestamp,
		t.Approved,
		database.BytesSignature(t.Signature))
	return err
}

func (w *opstackDatabase) verifySignature(hubLayerChainID *big.Int, pbMsg interface{}) (bool, error) {
	sig := pbMsg.(*pb.OpstackSignature)
	if sig == nil {
		return false, errors.New("not an opstack signature")
	}
	if err := verifyULID(sig.Id); err != nil {
		return false, err
	}

	msg := verifier.NewL2ooMessage(
		hubLayerChainID,
		common.BytesToAddress(sig.L2Oo),
		new(big.Int).SetUint64(sig.L2OutputIndex),
		common.BytesToHash(sig.OutputRoot),
		new(big.Int).SetUint64(sig.L1Timestamp),
		new(big.Int).SetUint64(sig.L2BlockNumber),
		sig.Approved)
	return msg.VerifySigner(sig.Signature, common.BytesToAddress(sig.Signer))
}
