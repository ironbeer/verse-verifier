package submitter

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/contract/stakemanager"
	"github.com/oasysgames/oasys-optimism-verifier/database"
)

type signatureIterator struct {
	db           *database.Database
	stakemanager *stakemanager.Cache
	contract     common.Address
	rollupIndex  uint64
}

func (si *signatureIterator) next() ([]*database.OptimismSignature, error) {
	rows, err := si.db.OPSignature.Find(nil, nil, &si.contract, &si.rollupIndex, 1000, 0)
	if err != nil {
		return nil, err
	}

	rows, err = filterSignatures(rows, minStake,
		si.stakemanager.TotalStake(), si.stakemanager.SignerStakes())
	if err != nil {
		return nil, err
	}

	si.rollupIndex++
	return rows, nil
}

func filterSignatures(
	rows []*database.OptimismSignature,
	minStake, totalStake *big.Int,
	signerStakes map[common.Address]*big.Int,
) (filterd []*database.OptimismSignature, err error) {
	// group by RollupHash and Approved
	type group struct {
		stake *big.Int
		rows  []*database.OptimismSignature
	}
	groups := map[string]*group{}

	for _, row := range rows {
		stake, ok := signerStakes[row.Signer.Address]
		if !ok || stake.Cmp(minStake) == -1 {
			continue
		}

		key := fmt.Sprintf("%s:%v", row.RollupHash, row.Approved)
		if _, ok := groups[key]; !ok {
			groups[key] = &group{stake: new(big.Int)}
		}

		groups[key].stake = new(big.Int).Add(groups[key].stake, stake)
		groups[key].rows = append(groups[key].rows, row)
	}
	if len(groups) == 0 {
		return nil, nil
	}

	// find the group key with the highest stake
	var highest *group
	for key := range groups {
		if highest == nil || groups[key].stake.Cmp(highest.stake) == 1 {
			highest = groups[key]
		}
	}

	// check over half
	required := new(big.Int).Mul(new(big.Int).Div(totalStake, big.NewInt(100)), big.NewInt(51))
	if highest.stake.Cmp(required) == -1 {
		return nil, &StakeAmountShortage{required, highest.stake}
	}

	sort.Sort(database.OptimismSignatures(highest.rows))
	return highest.rows, nil
}

type StakeAmountShortage struct {
	required, actual *big.Int
}

func (err *StakeAmountShortage) Error() string {
	return fmt.Sprintf("stake amount shortage, required=%s actual=%s",
		fromWei(err.required), fromWei(err.actual))
}
