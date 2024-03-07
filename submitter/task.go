package submitter

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/contract/multicall2"
	"github.com/oasysgames/oasys-optimism-verifier/contract/sccverifier"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
)

type SubmitTask struct {
	l1Client     ethutil.SignableClient
	verifier     *sccverifier.OasysRollupVerifier
	multicall    *multicall2.Multicall2
	contractAddr common.Address
	nextIndexFn  func(*bind.CallOpts) (*big.Int, error)
	txCreator    txCreator
}

type txCreator = func(
	opts *bind.TransactOpts,
	verifier *sccverifier.OasysRollupVerifier,
	contractAddr common.Address,
	rollupIndex uint64,
	approved bool,
	signatures [][]byte,
) (unsignedTx *types.Transaction, err error)
