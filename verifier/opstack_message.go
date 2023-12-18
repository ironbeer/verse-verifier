package verifier

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type OpstackMessage struct {
	AbiPacked []byte
	Eip712Msg string
}

func NewOpstackMessage(
	hubChainID *big.Int,
	l2oo_ common.Address,
	l2OutputIndex *big.Int,
	outputRoot common.Hash,
	l1Timestamp *big.Int,
	l2BlockNumber *big.Int,
	approved bool,
) *OpstackMessage {
	_approved := []byte{0}
	if approved {
		_approved = []byte{1}
	}

	abiPacked := bytes.Join([][]byte{
		common.LeftPadBytes(hubChainID.Bytes(), 32),
		l2oo_[:],
		common.LeftPadBytes(l2OutputIndex.Bytes(), 32),
		bytes.Join([][]byte{
			outputRoot[:], // bytes32
			common.LeftPadBytes(l1Timestamp.Bytes(), 16),   // uint128
			common.LeftPadBytes(l2BlockNumber.Bytes(), 16), // uint128
		}, nil),
		_approved,
	}, nil)
	_, msg := accounts.TextAndHash(crypto.Keccak256(abiPacked))

	return &OpstackMessage{
		AbiPacked: abiPacked,
		Eip712Msg: msg,
	}
}

func (m *OpstackMessage) Signature(signDataFn ethutil.SignDataFn) ([65]byte, error) {
	var sig [65]byte
	signed, err := signDataFn([]byte(m.Eip712Msg))
	if err != nil {
		return sig, err
	}
	copy(sig[:], signed)

	// Transform V from 0/1 to 27/28 according to the yellow paper
	sig[crypto.RecoveryIDOffset] += 27
	return sig, nil
}

func (m *OpstackMessage) Ecrecover(signature []byte) (common.Address, error) {
	hash := crypto.Keccak256([]byte(m.Eip712Msg))
	return ethutil.Ecrecover(hash, signature)
}

func (m *OpstackMessage) VerifySigner(signature []byte, signer common.Address) (bool, error) {
	if recoverd, err := m.Ecrecover(signature); err != nil {
		return false, err
	} else {
		return bytes.Equal(recoverd.Bytes(), signer.Bytes()), nil
	}
}
