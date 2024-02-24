package database

import (
	"github.com/ethereum/go-ethereum/common"
)

type Block struct {
	ID uint64 `gorm:"primarykey"`

	Number       uint64 `gorm:"uniqueIndex"`
	Hash         common.Hash
	LogCollected bool // Deprecated
}

type Signer struct {
	ID uint64 `gorm:"primarykey"`

	Address common.Address `gorm:"uniqueIndex"`
}

type OptimismScc struct {
	ID uint64 `gorm:"primarykey"`

	Address   common.Address `gorm:"uniqueIndex"`
	NextIndex uint64
}

type OptimismState struct {
	ID uint64 `gorm:"primarykey"`

	OptimismSccID uint64 `gorm:"uniqueIndex:optimism_state_idx0,priority:1"`
	OptimismScc   OptimismScc

	// Value of `StateBatchAppended` event.
	BatchIndex        uint64 `gorm:"uniqueIndex:optimism_state_idx0,priority:2;index:optimism_state_idx1"`
	BatchRoot         common.Hash
	BatchSize         uint64
	PrevTotalElements uint64
	ExtraData         []byte
}

type OptimismSignature struct {
	ID         string `gorm:"primarykey;index:optimism_signature_idx3,priority:2"`
	PreviousID string

	SignerID uint64 `gorm:"uniqueIndex:optimism_signature_idx1,priority:1;index:optimism_signature_idx3,priority:1"`
	Signer   Signer

	OptimismSccID uint64 `gorm:"uniqueIndex:optimism_signature_idx1,priority:2"`
	OptimismScc   OptimismScc

	BatchIndex        uint64 `gorm:"uniqueIndex:optimism_signature_idx1,priority:3;index:optimism_signature_idx2"`
	BatchRoot         common.Hash
	BatchSize         uint64
	PrevTotalElements uint64
	ExtraData         []byte

	Approved  bool
	Signature Signature
}

type OpstackL2OutputOracle struct {
	ID uint64 `gorm:"primarykey"`

	Address         common.Address `gorm:"uniqueIndex"`
	NextVerifyIndex uint64
}

type OpstackProposal struct {
	ID uint64 `gorm:"primarykey"`

	OpstackL2OutputOracleID uint64 `gorm:"uniqueIndex:opstack_proposal_idx0,priority:1"`
	OpstackL2OutputOracle   OpstackL2OutputOracle

	L2OutputIndex uint64 `gorm:"uniqueIndex:opstack_proposal_idx0,priority:2;index:opstack_proposal_idx1"`

	// Value of `Types.OutputProposal` struct.
	OutputRoot    common.Hash
	L2BlockNumber uint64
	L1Timestamp   uint64
}

type OpstackSignature struct {
	ID         string `gorm:"primarykey;index:opstack_signature_idx3,priority:2"`
	PreviousID string

	SignerID uint64 `gorm:"uniqueIndex:opstack_signature_idx1,priority:1;index:opstack_signature_idx3,priority:1"`
	Signer   Signer

	OpstackL2OutputOracleID uint64 `gorm:"uniqueIndex:opstack_signature_idx1,priority:2"`
	OpstackL2OutputOracle   OpstackL2OutputOracle

	L2OutputIndex uint64 `gorm:"uniqueIndex:opstack_signature_idx1,priority:3;index:opstack_signature_idx2"`

	// Value of `Types.OutputProposal` struct.
	OutputRoot    common.Hash
	L2BlockNumber uint64
	L1Timestamp   uint64

	Approved  bool
	Signature Signature
}

type Misc struct {
	ID    string `gorm:"primarykey"`
	Value []byte
}
