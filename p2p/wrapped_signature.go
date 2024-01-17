package p2p

import (
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
)

type iWrappedSignature interface {
	getID() string
	getPreviousID() string
	findCommonSignatureResponse() *pb.Stream
}

type wrappedOptimismSignature database.OptimismSignature

func (s *wrappedOptimismSignature) getID() string {
	return s.ID
}

func (s *wrappedOptimismSignature) getPreviousID() string {
	return s.PreviousID
}

func (s *wrappedOptimismSignature) findCommonSignatureResponse() *pb.Stream {
	var found *pb.OptimismSignature
	if s != nil {
		found = s.protoSig()
	}
	return &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
		FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{Found: found},
	}}
}

func (s *wrappedOptimismSignature) protoSig() *pb.OptimismSignature {
	return &pb.OptimismSignature{
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
	}
}

type wrappedOpstackSignature database.OpstackSignature

func (s *wrappedOpstackSignature) getID() string {
	return s.ID
}

func (s *wrappedOpstackSignature) getPreviousID() string {
	return s.PreviousID
}

func (s *wrappedOpstackSignature) findCommonSignatureResponse() *pb.Stream {
	var found *pb.OpstackSignature
	if s != nil {
		found = s.protoSig()
	}
	return &pb.Stream{Body: &pb.Stream_FindCommonOpstackSignature{
		FindCommonOpstackSignature: &pb.FindCommonOpstackSignature{Found: found},
	}}
}

func (s *wrappedOpstackSignature) protoSig() *pb.OpstackSignature {
	return &pb.OpstackSignature{
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
	}
}

type wrappedSignatures interface {
	len() int
	get(index int) iWrappedSignature
	signatureExchangeResponse() *pb.Stream
	findCommonSignatureRequest() (msg *pb.Stream, from, to string, err error)
}

type wrappedOptimismSignatures []*database.OptimismSignature

func (sigs wrappedOptimismSignatures) ProtoSig() (pbSigs []*pb.OptimismSignature) {
	pbSigs = make([]*pb.OptimismSignature, len(sigs))
	for i, s := range sigs {
		pbSigs[i] = (*wrappedOptimismSignature)(s).protoSig()
	}
	return pbSigs
}

func (ss wrappedOptimismSignatures) len() int {
	return len(ss)
}

func (ss wrappedOptimismSignatures) get(index int) iWrappedSignature {
	return (*wrappedOptimismSignature)(ss[index])
}

func (ss wrappedOptimismSignatures) signatureExchangeResponse() *pb.Stream {
	return &pb.Stream{Body: &pb.Stream_OptimismSignatureExchange{
		OptimismSignatureExchange: &pb.OptimismSignatureExchange{
			Responses: ss.ProtoSig(),
		},
	}}
}

func (ss wrappedOptimismSignatures) findCommonSignatureRequest() (msg *pb.Stream, from, to string, err error) {
	locals := make([]*pb.FindCommonOptimismSignature_Local, len(ss))
	for i, sig := range ss {
		locals[i] = &pb.FindCommonOptimismSignature_Local{
			Id:         sig.ID,
			PreviousId: sig.PreviousID,
		}
	}

	return &pb.Stream{Body: &pb.Stream_FindCommonOptimismSignature{
		FindCommonOptimismSignature: &pb.FindCommonOptimismSignature{Locals: locals},
	}}, ss[0].ID, ss[len(ss)-1].ID, nil
}

type wrappedOpstackSignatures []*database.OpstackSignature

func (ss wrappedOpstackSignatures) ProtoSig() (pbSigs []*pb.OpstackSignature) {
	pbSigs = make([]*pb.OpstackSignature, len(ss))
	for i, s := range ss {
		pbSigs[i] = (*wrappedOpstackSignature)(s).protoSig()
	}
	return pbSigs
}

func (ss wrappedOpstackSignatures) len() int {
	return len(ss)
}

func (ss wrappedOpstackSignatures) get(index int) iWrappedSignature {
	return (*wrappedOpstackSignature)(ss[index])
}

func (ss wrappedOpstackSignatures) signatureExchangeResponse() *pb.Stream {
	return &pb.Stream{Body: &pb.Stream_OpstackSignatureExchange{
		OpstackSignatureExchange: &pb.OpstackSignatureExchange{
			Responses: ss.ProtoSig(),
		},
	}}
}

func (ss wrappedOpstackSignatures) findCommonSignatureRequest() (msg *pb.Stream, from, to string, err error) {
	locals := make([]*pb.FindCommonOpstackSignature_Local, len(ss))
	for i, sig := range ss {
		locals[i] = &pb.FindCommonOpstackSignature_Local{
			Id:         sig.ID,
			PreviousId: sig.PreviousID,
		}
	}

	return &pb.Stream{Body: &pb.Stream_FindCommonOpstackSignature{
		FindCommonOpstackSignature: &pb.FindCommonOpstackSignature{Locals: locals},
	}}, ss[0].ID, ss[len(ss)-1].ID, nil
}
