package backend

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	tmcall2 "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/multicall2"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/scc"
	tsccv "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/sccverifier"
)

type BackendSuite struct {
	DB    *database.Database
	Hub   *TestBackend
	Verse *TestBackend

	StakeManager *StakeManagerMock

	Mcall2     *tmcall2.Multicall2
	Mcall2Addr common.Address

	SCC     *scc.Scc
	TSCC    *tscc.Scc
	SCCAddr common.Address

	SCCV     *tsccv.Sccverifier
	SCCVAddr common.Address
}

func (b *BackendSuite) SetupTest() {
	ctx := context.Background()
	b.DB, _ = database.NewDatabase(&config.Database{Path: ":memory:"})

	// setup test chain
	b.Hub = NewTestBackend()
	b.Verse = NewTestBackend()
	b.StakeManager = &StakeManagerMock{}

	// deploy `Multicall2` contract
	b.Mcall2Addr, _, b.Mcall2, _ = tmcall2.DeployMulticall2(b.Hub.TransactOpts(ctx), b.Hub)
	b.Hub.Mining()

	// deploy `StateCommitmentChain` contract
	b.SCCAddr, _, b.TSCC, _ = tscc.DeployScc(b.Hub.TransactOpts(ctx), b.Hub)
	b.SCC, _ = scc.NewScc(b.SCCAddr, b.Hub)
	b.Hub.Mining()

	// deploy `OasysStateCommitmentChainVerifier` contract
	b.SCCVAddr, _, b.SCCV, _ = tsccv.DeploySccverifier(b.Hub.TransactOpts(ctx), b.Hub)
	b.Hub.Mining()
}

func (b *BackendSuite) Mining() {
	b.Hub.Commit()
	header, _ := b.Hub.HeaderByNumber(context.Background(), nil)
	b.DB.Block.SaveNewBlock(header.Number.Uint64(), header.Hash())
}

func (b *BackendSuite) EmitStateBatchAppendedEvent(index int) *tscc.SccStateBatchAppended {
	i64 := int64(index)
	event := &tscc.SccStateBatchAppended{
		BatchIndex:        big.NewInt(i64),
		BatchRoot:         [32]byte(common.BigToHash(big.NewInt(i64))),
		BatchSize:         big.NewInt(10),
		PrevTotalElements: big.NewInt(i64 * 10),
		ExtraData:         []byte("extra data"),
	}
	b.TSCC.EmitStateBatchAppended(
		b.Hub.TransactOpts(context.Background()), event.BatchIndex,
		event.BatchRoot, event.BatchSize, event.PrevTotalElements, event.ExtraData)
	b.Mining()
	return event
}

type StakeManagerMock struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
	NewCursor  *big.Int
}

func (b *StakeManagerMock) GetTotalStake(
	callOpts *bind.CallOpts,
	epoch *big.Int,
) (*big.Int, error) {
	tot := new(big.Int)
	for _, stake := range b.Stakes {
		tot.Add(tot, stake)
	}
	return tot, nil
}

func (b *StakeManagerMock) GetValidators(
	callOpts *bind.CallOpts,
	epoch, cursol, howMany *big.Int,
) (struct {
	Owners     []common.Address
	Operators  []common.Address
	Stakes     []*big.Int
	Candidates []bool
	NewCursor  *big.Int
}, error) {
	length := big.NewInt(int64(len(b.Owners)))
	if new(big.Int).Add(cursol, howMany).Cmp(length) >= 0 {
		howMany = new(big.Int).Sub(length, cursol)
	}

	start := cursol.Uint64()
	end := start + howMany.Uint64()

	ret := struct {
		Owners     []common.Address
		Operators  []common.Address
		Stakes     []*big.Int
		Candidates []bool
		NewCursor  *big.Int
	}{
		Owners:     b.Owners[start:end],
		Operators:  b.Operators[start:end],
		Stakes:     b.Stakes[start:end],
		Candidates: b.Candidates[start:end],
		NewCursor:  new(big.Int).Add(cursol, howMany),
	}

	return ret, nil
}
