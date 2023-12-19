package backend

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	tl2oo "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/l2oo"
	tmcall2 "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/multicall2"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/scc"
	tsccv "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/sccverifier"
)

type BackendSuite struct {
	testhelper.Suite

	DB    *database.Database
	Hub   *TestBackend
	Verse *TestBackend

	StakeManager *StakeManagerMock

	Mcall2     *tmcall2.Multicall2
	Mcall2Addr common.Address

	SCC     *scc.Scc
	TSCC    *tscc.Scc
	SCCAddr common.Address

	L2OO     *l2oo.OasysL2OutputOracle
	TL2OO    *tl2oo.L2oo
	L2OOAddr common.Address

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

	// deploy `OasysL2OutputOracle` contract
	b.L2OOAddr, _, b.TL2OO, _ = tl2oo.DeployL2oo(b.Hub.TransactOpts(ctx), b.Hub)
	b.L2OO, _ = l2oo.NewOasysL2OutputOracle(b.L2OOAddr, b.Hub)
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

func (b *BackendSuite) EmitOutputProposed(index int) *tl2oo.L2ooOutputProposed {
	event := &tl2oo.L2ooOutputProposed{
		OutputRoot:    b.RandHash(),
		L2OutputIndex: big.NewInt(int64(index)),
		L2BlockNumber: big.NewInt(int64(index * 10)),
		L1Timestamp:   big.NewInt(time.Now().Unix()),
	}
	_, err := b.TL2OO.EmitOutputProposed(b.Hub.TransactOpts(context.Background()),
		event.OutputRoot, event.L2OutputIndex, event.L2BlockNumber, event.L1Timestamp)
	b.Nil(err)
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
