package stakemanager

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	"github.com/stretchr/testify/suite"
)

type CacheTestSuite struct {
	testhelper.Suite

	sm *backend.StakeManagerMock
	vs *Cache
}

func TestNewCache(t *testing.T) {
	suite.Run(t, new(CacheTestSuite))
}

func (s *CacheTestSuite) SetupTest() {
	s.sm = &backend.StakeManagerMock{}
	s.vs = NewCache(s.sm)

	for i := range s.Range(0, 1000) {
		s.sm.Owners = append(s.sm.Owners, s.RandAddress())
		s.sm.Operators = append(s.sm.Operators, s.RandAddress())
		s.sm.Stakes = append(s.sm.Stakes, big.NewInt(int64(i)))
		s.sm.Candidates = append(s.sm.Candidates, true)
		s.sm.NewCursor = big.NewInt(0)
	}
}

func (s *CacheTestSuite) TestRefresh() {
	s.Equal(common.Big0, s.vs.TotalStake())
	for _, signer := range s.sm.Operators {
		s.Equal(common.Big0, s.vs.StakeBySigner(signer))
	}

	s.Nil(s.vs.Refresh(context.Background()))

	s.Equal(big.NewInt(499500), s.vs.TotalStake())
	for i, signer := range s.sm.Operators {
		s.Equal(s.sm.Stakes[i], s.vs.StakeBySigner(signer))
	}
}
