package submitter

import (
	"context"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"github.com/stretchr/testify/suite"
)

type SubmitterTestSuite struct {
	backend.BackendSuite

	submitter *Submitter
}

func TestSubmitter(t *testing.T) {
	suite.Run(t, new(SubmitterTestSuite))
}

func (s *SubmitterTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	// Setup `StakeManager` contract
	for range s.Range(0, 10) {
		s.StakeManager.Owners = append(s.StakeManager.Owners, s.RandAddress())
		s.StakeManager.Operators = append(s.StakeManager.Operators, s.RandAddress())
		s.StakeManager.Stakes = append(
			s.StakeManager.Stakes,
			new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10_000_000)),
		)
		s.StakeManager.Candidates = append(s.StakeManager.Candidates, true)
	}

	// Setup submitter
	s.submitter = NewSubmitter(&config.Submitter{
		Interval:          0,
		Concurrency:       0,
		Confirmations:     0,
		GasMultiplier:     1.0,
		BatchSize:         20,
		Multicall2Address: s.MulticallAddr.String(),
		UseMulticall:      true, // TODO
	}, s.DB, s.StakeManager)
}

func (s *SubmitterTestSuite) TestSubmit() {
	ctx := context.Background()
	batchIndexes := s.Range(0, 5)
	nextIndex := 2
	signers := s.StakeManager.Operators

	// save dummy signatures
	events := make([]*tscc.SccStateBatchAppended, len(batchIndexes))
	signatures := make([][]*database.OptimismSignature, len(batchIndexes))
	for i := range batchIndexes {
		_, events[i] = s.EmitStateBatchAppended(i)
		signatures[i] = make([]*database.OptimismSignature, len(signers))

		for j := range s.Range(0, len(signers)) {
			signatures[i][j], _ = s.DB.OPSignature.Save(
				nil, nil,
				signers[j],
				s.SCCAddr,
				events[i].BatchIndex.Uint64(),
				events[i].BatchRoot,
				i < len(batchIndexes)-1,
				database.RandSignature(),
			)
		}

		sort.Sort(database.OptimismSignatures(signatures[i]))
	}

	// set the `SCC.nextIndex`
	s.TSCC.SetNextIndex(s.SignableHub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Commit()

	// submitter do the work.
	s.submitter.stakemanager.Refresh(ctx)
	go s.submitter.work(ctx, verse.NewOPLegacy(
		s.DB, s.Hub, s.SCCAddr).WithTransactable(s.SignableHub, s.SCCVAddr))
	time.Sleep(time.Millisecond * 25)
	s.Hub.Commit()

	// assert
	length, _ := s.TSCCV.SccAssertLogsLen(&bind.CallOpts{Context: ctx})
	s.Equal(uint64(3), length.Uint64())

	for i := range batchIndexes {
		if i < nextIndex {
			got, err := s.TSCCV.SccAssertLogs(
				&bind.CallOpts{Context: ctx},
				big.NewInt(int64(i+nextIndex+1)))
			s.ErrorContains(err, "execution reverted")
			s.Equal(common.Address{}, got.StateCommitmentChain)
		} else {
			got, err := s.TSCCV.SccAssertLogs(
				&bind.CallOpts{Context: ctx},
				big.NewInt(int64(i-nextIndex)))
			s.NoError(err)
			s.Equal(s.SCCAddr, got.StateCommitmentChain)
			s.Equal(events[i].BatchIndex.Uint64(), got.BatchHeader.BatchIndex.Uint64())
			s.Equal(events[i].BatchRoot, got.BatchHeader.BatchRoot)
			s.Equal(events[i].BatchSize.Uint64(), got.BatchHeader.BatchSize.Uint64())
			s.Equal(events[i].PrevTotalElements.Uint64(), got.BatchHeader.PrevTotalElements.Uint64())
			s.Equal(events[i].ExtraData, got.BatchHeader.ExtraData)
			s.Equal(i < len(batchIndexes)-1, got.Approve)

			s.Len(got.Signatures, len(signers)*65)
			for j, sig := range signatures[i] {
				start := j * 65
				end := start + 65
				s.Equal(sig.Signature[:], got.Signatures[start:end])
			}
		}
	}
}
