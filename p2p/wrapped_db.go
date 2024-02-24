package p2p

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
	"github.com/oasysgames/oasys-optimism-verifier/verifier"
)

type iWrappedDB interface {
	// Methods that wrap around the original database.
	findSignatureByID(id string) (iWrappedSignature, error)
	findSignatures(
		idAfter *string,
		signer *common.Address,
		contract *common.Address,
		index *uint64,
		limit, offset int,
	) (wrappedSignatures, error)
	findLatestSignaturesBySigner(signer common.Address, limit, offset int) (wrappedSignatures, error)
	findLatestSignaturePerSigners() (wrappedSignatures, error)

	// Additional methods.
	getSignatureExchangeRequest(signer common.Address, idAfter string) *pb.Stream
	handleFindCommonSignatureResponse(res *pb.Stream) (sig pb.ISignature, found bool, err error)
	saveSignature(pbMsg interface{}) error
	verifySignature(hubLayerChainID *big.Int, pbMsg interface{}) error
}

type wrappedOptimismDB struct {
	*database.OptimismDatabase
}

func (w *wrappedOptimismDB) findSignatureByID(id string) (iWrappedSignature, error) {
	row, err := w.FindSignatureByID(id)
	return (*wrappedOptimismSignature)(row), err
}

func (w *wrappedOptimismDB) findSignatures(
	idAfter *string,
	signer *common.Address,
	contract *common.Address,
	index *uint64,
	limit, offset int,
) (wrappedSignatures, error) {
	rows, err := w.FindSignatures(idAfter, signer, contract, index, limit, offset)
	return wrappedOptimismSignatures(rows), err
}

func (w *wrappedOptimismDB) findLatestSignaturesBySigner(signer common.Address, limit, offset int) (wrappedSignatures, error) {
	rows, err := w.FindLatestSignaturesBySigner(signer, limit, offset)
	return wrappedOptimismSignatures(rows), err
}

func (w *wrappedOptimismDB) findLatestSignaturePerSigners() (wrappedSignatures, error) {
	rows, err := w.FindLatestSignaturePerSigners()
	return wrappedOptimismSignatures(rows), err
}

func (w *wrappedOptimismDB) getSignatureExchangeRequest(signer common.Address, idAfter string) *pb.Stream {
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

func (w *wrappedOptimismDB) handleFindCommonSignatureResponse(res *pb.Stream) (sig pb.ISignature, found bool, err error) {
	t := res.GetFindCommonOptimismSignature()
	if t == nil {
		return nil, false, errors.New("unexpected response")
	}
	return t.Found, t.Found != nil, nil
}

func (w *wrappedOptimismDB) saveSignature(pbMsg interface{}) error {
	t, ok := pbMsg.(*pb.OptimismSignature)
	if !ok {
		return errors.New("unknown protobuf message")
	}

	_, err := w.SaveSignature(
		&t.Id,
		&t.PreviousId,
		common.BytesToAddress(t.Signer),
		common.BytesToAddress(t.Scc),
		t.BatchIndex,
		common.BytesToHash(t.BatchRoot),
		t.BatchSize,
		t.PrevTotalElements,
		t.ExtraData,
		t.Approved,
		database.BytesSignature(t.Signature))
	return err
}

func (w *wrappedOptimismDB) verifySignature(hubLayerChainID *big.Int, pbMsg interface{}) error {
	t, ok := pbMsg.(*pb.OptimismSignature)
	if !ok {
		return errors.New("not an optimism signature")
	}

	signer := common.BytesToAddress(t.Signer)
	scc := common.BytesToAddress(t.Scc)
	batchIndex := new(big.Int).SetUint64(t.BatchIndex)
	batchRoot := common.BytesToHash(t.BatchRoot)

	msg := verifier.NewSccMessage(hubLayerChainID, scc, batchIndex, batchRoot, t.Approved)
	err := msg.VerifySigner(t.Signature, signer)

	// possibly an old signature with an approved bug
	if _, ok := err.(*verifier.SignerMismatchError); ok {
		msg = verifier.NewSCCMessageWithApprovedBug(
			hubLayerChainID, scc, batchIndex, batchRoot, t.Approved)
		err = msg.VerifySigner(t.Signature, signer)
	}

	return err
}

// Wrapped `database.OPStackDatabase`.
type wrappedOpstackDB struct {
	*database.OPStackDatabase
}

func (w *wrappedOpstackDB) findSignatureByID(id string) (iWrappedSignature, error) {
	row, err := w.FindSignatureByID(id)
	return (*wrappedOpstackSignature)(row), err
}

func (w *wrappedOpstackDB) findSignatures(
	idAfter *string,
	signer *common.Address,
	contract *common.Address,
	index *uint64,
	limit, offset int,
) (wrappedSignatures, error) {
	rows, err := w.FindSignatures(idAfter, signer, contract, index, limit, offset)
	return wrappedOpstackSignatures(rows), err
}

func (w *wrappedOpstackDB) findLatestSignaturesBySigner(signer common.Address, limit, offset int) (wrappedSignatures, error) {
	rows, err := w.FindLatestSignaturesBySigner(signer, limit, offset)
	return wrappedOpstackSignatures(rows), err
}

func (w *wrappedOpstackDB) findLatestSignaturePerSigners() (wrappedSignatures, error) {
	rows, err := w.FindLatestSignaturePerSigners()
	return wrappedOpstackSignatures(rows), err
}

func (w *wrappedOpstackDB) getSignatureExchangeRequest(signer common.Address, idAfter string) *pb.Stream {
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

func (w *wrappedOpstackDB) handleFindCommonSignatureResponse(res *pb.Stream) (sig pb.ISignature, found bool, err error) {
	t := res.GetFindCommonOpstackSignature()
	if t == nil {
		return nil, false, errors.New("unexpected response")
	}
	return t.Found, t.Found != nil, nil
}

func (w *wrappedOpstackDB) saveSignature(pbMsg interface{}) error {
	t, ok := pbMsg.(*pb.OpstackSignature)
	if !ok {
		return errors.New("unknown protobuf message")
	}

	_, err := w.SaveSignature(
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

func (w *wrappedOpstackDB) verifySignature(hubLayerChainID *big.Int, pbMsg interface{}) error {
	t, ok := pbMsg.(*pb.OpstackSignature)
	if !ok {
		return errors.New("not an opstack signature")
	}

	msg := verifier.NewL2ooMessage(
		hubLayerChainID,
		common.BytesToAddress(t.L2Oo),
		new(big.Int).SetUint64(t.L2OutputIndex),
		common.BytesToHash(t.OutputRoot),
		new(big.Int).SetUint64(t.L1Timestamp),
		new(big.Int).SetUint64(t.L2BlockNumber),
		t.Approved)
	return msg.VerifySigner(t.Signature, common.BytesToAddress(t.Signer))
}
