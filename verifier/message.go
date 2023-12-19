package verifier

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type Message struct {
	AbiPacked []byte
	Eip712Msg string
}

func (m *Message) Signature(signDataFn ethutil.SignDataFn) ([65]byte, error) {
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

func (m *Message) Ecrecover(signature []byte) (common.Address, error) {
	hash := crypto.Keccak256([]byte(m.Eip712Msg))
	return ethutil.Ecrecover(hash, signature)
}

func (m *Message) VerifySigner(signature []byte, signer common.Address) (bool, error) {
	if recoverd, err := m.Ecrecover(signature); err != nil {
		return false, err
	} else {
		return bytes.Equal(recoverd.Bytes(), signer.Bytes()), nil
	}
}

func NewSccMessage(
	hubChainID *big.Int,
	scc common.Address,
	batchIndex *big.Int,
	batchRoot [32]byte,
	approved bool,
) *Message {
	// See: https://github.com/oasysgames/oasys-optimism/blob/5186190c3250121179064b70d8e2fbd2d0a03ce3/packages/contracts/contracts/oasys/L1/rollup/OasysStateCommitmentChainVerifier.sol#L111-L119
	abiPacked := bytes.Join([][]byte{
		padUint256(hubChainID),
		scc[:],
		padUint256(batchIndex),
		batchRoot[:],
		padBool(approved),
	}, nil)
	_, msg := accounts.TextAndHash(crypto.Keccak256(abiPacked))

	return &Message{AbiPacked: abiPacked, Eip712Msg: msg}
}

func NewL2ooMessage(
	hubChainID *big.Int,
	l2oo_ common.Address,
	l2OutputIndex *big.Int,
	outputRoot common.Hash,
	l1Timestamp *big.Int,
	l2BlockNumber *big.Int,
	approved bool,
) *Message {
	abiPacked := bytes.Join([][]byte{
		padUint256(hubChainID),
		l2oo_[:],
		padUint256(l2OutputIndex),
		bytes.Join([][]byte{
			outputRoot[:],
			padUint128(l1Timestamp),
			padUint128(l2BlockNumber),
		}, nil),
		padBool(approved),
	}, nil)
	_, msg := accounts.TextAndHash(crypto.Keccak256(abiPacked))

	return &Message{AbiPacked: abiPacked, Eip712Msg: msg}
}

func padUint128(val *big.Int) []byte {
	return common.LeftPadBytes(val.Bytes(), 16)
}

func padUint256(val *big.Int) []byte {
	return common.LeftPadBytes(val.Bytes(), 32)
}

func padBool(val bool) []byte {
	if val {
		return []byte{1}
	}
	return []byte{0}
}
