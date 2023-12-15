package hublayer

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ValidatorStakings struct {
	sm    StakeManager
	cache *sync.Map
}

func NewValidatorStakings(sm StakeManager) *ValidatorStakings {
	return &ValidatorStakings{sm: sm, cache: &sync.Map{}}
}

func (vs *ValidatorStakings) Refresh(ctx context.Context) error {
	// total amount
	total, err := vs.sm.GetTotalStake(&bind.CallOpts{Context: ctx}, common.Big0)
	if err != nil {
		return err
	}
	vs.cache.Store(common.Address{}, total)

	// each validator
	cursor, howMany := big.NewInt(0), big.NewInt(250)
	for {
		result, err := vs.sm.GetValidators(&bind.CallOpts{Context: ctx}, common.Big0, cursor, howMany)
		if err != nil {
			return err
		} else if len(result.Owners) == 0 {
			break
		}

		for i, operator := range result.Operators {
			if operator != (common.Address{}) {
				vs.cache.Store(operator, result.Stakes[i])
			}
		}
		cursor = result.NewCursor
	}

	return nil
}

func (vs *ValidatorStakings) TotalStake() *big.Int {
	if val, ok := vs.cache.Load(common.Address{}); !ok {
		return big.NewInt(0)
	} else {
		return val.(*big.Int)
	}
}

func (vs ValidatorStakings) StakeBySigner(signer common.Address) *big.Int {
	if signer != (common.Address{}) {
		if val, ok := vs.cache.Load(signer); ok {
			return val.(*big.Int)
		}
	}
	return big.NewInt(0)
}

type StakeManager interface {
	GetTotalStake(callOpts *bind.CallOpts, epoch *big.Int) (*big.Int, error)

	GetValidators(callOpts *bind.CallOpts, epoch, cursol, howMany *big.Int) (struct {
		Owners     []common.Address
		Operators  []common.Address
		Stakes     []*big.Int
		Candidates []bool
		NewCursor  *big.Int
	}, error)
}
