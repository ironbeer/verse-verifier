package hublayer

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

type ValidatorStakingsTestSuite struct {
	testhelper.Suite

	sm *stakeManagerMock
	vs *ValidatorStakings
}

func TestValidatorStakings(t *testing.T) {
	suite.Run(t, new(ValidatorStakingsTestSuite))
}

func (s *ValidatorStakingsTestSuite) SetupTest() {
	s.sm = &stakeManagerMock{}
	s.vs = NewValidatorStakings(s.sm)

	for i := range s.Range(0, 1000) {
		s.sm.Owners = append(s.sm.Owners, s.RandAddress())
		s.sm.Operators = append(s.sm.Operators, s.RandAddress())
		s.sm.Stakes = append(s.sm.Stakes, big.NewInt(int64(i)))
		s.sm.Candidates = append(s.sm.Candidates, true)
		s.sm.NewCursor = big.NewInt(0)
	}
}

func (s *ValidatorStakingsTestSuite) TestRefresh() {
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
