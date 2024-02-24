package p2p

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/p2p/pb"
)

type signatureDB struct {
	iWrappedDB
}

func (db *signatureDB) findLatestSignatureId(signer common.Address) (*string, error) {
	latests, err := db.findLatestSignaturesBySigner(signer, 1, 0)
	if err != nil {
		return nil, err
	} else if latests.len() == 0 {
		return nil, nil
	}
	id := latests.get(0).getID()
	return &id, nil
}

func (db *signatureDB) getSignatureExchangeResponse(signer common.Address, idAfter string, limit, offset int) (msg *pb.Stream, count int, err error) {
	// get latest signatures for each requested signer
	sigs, err := db.findSignatures(&idAfter, &signer, nil, nil, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	return sigs.signatureExchangeResponse(), sigs.len(), nil
}

func (db *signatureDB) hasSignature(id string, previousID *string) (bool, error) {
	sig, err := db.findSignatureByID(id)
	if errors.Is(err, database.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	} else if previousID == nil {
		return true, nil
	}
	return sig.getPreviousID() == *previousID, nil
}

func (db *signatureDB) getFindCommonSignatureRequest(
	signer common.Address,
	limit, offset int,
) (msg *pb.Stream, count int, fromID, toID string, err error) {
	// find local latest signatures (order by: id desc)
	sigs, err := db.findLatestSignaturesBySigner(signer, limit, offset)
	if err != nil {
		return nil, 0, "", "", err
	} else if sigs.len() == 0 {
		// reached the last
		return nil, 0, "", "", nil
	}
	return sigs.findCommonSignatureRequest()
}

func (db *signatureDB) getFindCommonSignatureResponse(
	req pb.ICommonSignatureRequest,
) (msg *pb.Stream, id string, found bool, err error) {
	sig, err := db.findSignatureByID(req.GetId())
	if err == nil {
		id = sig.getID()
	}
	found = err == nil && sig.getPreviousID() == req.GetPreviousId()
	return sig.findCommonSignatureResponse(), id, found, err
}
