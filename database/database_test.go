package database

import (
	"math/rand"
	"time"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"gorm.io/gorm"
)

type DatabaseTestSuite struct {
	testhelper.Suite

	db    *Database
	rawdb *gorm.DB
}

func (s *DatabaseTestSuite) SetupTest() {
	// Setup database
	db, err := NewDatabase(&config.Database{Path: ":memory:"})
	if err != nil {
		panic(err)
	}
	s.db = db
	s.rawdb = db.db
}

func (s *DatabaseTestSuite) createSigner() *Signer {
	signer := &Signer{Address: s.RandAddress()}
	s.NoDBError(s.rawdb.Create(signer))
	return signer
}

func (s *DatabaseTestSuite) createSCC() *OptimismScc {
	scc := &OptimismScc{Address: s.RandAddress()}
	s.NoDBError(s.rawdb.Create(scc))
	return scc
}

func (s *DatabaseTestSuite) createState(scc *OptimismScc, index int) *OptimismState {
	state := &OptimismState{
		OptimismScc:       *scc,
		BatchIndex:        uint64(index),
		BatchRoot:         s.ItoHash(index),
		BatchSize:         uint64(rand.Intn(99)),
		PrevTotalElements: uint64(rand.Intn(99)),
		ExtraData:         s.RandBytes(),
	}
	s.NoDBError(s.rawdb.Create(state))
	return state
}

func (s *DatabaseTestSuite) createSignature(
	signer *Signer,
	scc *OptimismScc,
	index int,
) *OptimismSignature {
	sig := &OptimismSignature{
		ID:          util.ULID(nil).String(),
		PreviousID:  util.ULID(nil).String(),
		Signer:      *signer,
		OptimismScc: *scc,
		BatchIndex:  uint64(index),
		BatchRoot:   s.RandHash(),
		Signature:   RandSignature(),
	}
	s.NoDBError(s.rawdb.Create(sig))
	return sig
}

func (s *DatabaseTestSuite) createL2OO() *OpstackL2OutputOracle {
	l2oo := &OpstackL2OutputOracle{Address: s.RandAddress()}
	s.NoDBError(s.rawdb.Create(l2oo))
	return l2oo
}

func (s *DatabaseTestSuite) createProposal(l2oo *OpstackL2OutputOracle, l2OutputIndex int) *OpstackProposal {
	proposal := &OpstackProposal{
		OpstackL2OutputOracle: *l2oo,
		OutputRoot:            s.RandHash(),
		L2OutputIndex:         uint64(l2OutputIndex),
		L2BlockNumber:         uint64(rand.Intn(99)),
		L1Timestamp:           uint64(time.Now().Unix()),
	}
	s.NoDBError(s.rawdb.Create(proposal))
	return proposal
}
