package submitter

import (
	"bytes"
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
	tl2oo "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/l2oo"
	tscc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
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
	s.StakeManager.NewCursor = big.NewInt(0)
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
		VerifierAddress:   s.SCCVAddr.String(),
		Multicall2Address: s.Mcall2Addr.String(),
		UseMulticall:      true,
	}, s.DB, s.StakeManager)
}

func (s *SubmitterTestSuite) TestSubmitSCC() {
	ctx := context.Background()
	indexes := s.Range(0, 5)
	nextIndex := 2
	signers := s.StakeManager.Operators

	// save dummy signatures
	events := make([]*tscc.SccStateBatchAppended, len(indexes))
	signatures := make([][]*database.OptimismSignature, len(indexes))
	for i := range indexes {
		events[i] = s.EmitStateBatchAppended(i)
		signatures[i] = make([]*database.OptimismSignature, len(signers))

		for j := range s.Range(0, len(signers)) {
			signatures[i][j], _ = s.DB.Optimism.SaveSignature(
				nil, nil,
				signers[j],
				s.SCCAddr,
				events[i].BatchIndex.Uint64(),
				events[i].BatchRoot,
				events[i].BatchSize.Uint64(),
				events[i].PrevTotalElements.Uint64(),
				events[i].ExtraData,
				i < len(indexes)-1,
				database.RandSignature(),
			)
		}

		sort.Sort(optimismSignatures(signatures[i]))
	}

	// set the `SCC.nextIndex`
	s.TSCC.SetNextIndex(s.Hub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Commit()

	// submitter do the work.
	s.submitter.stakemanager.Refresh(ctx)
	task, _ := NewTask(s.Hub, s.SCCAddr, s.SCC, s.SCCV, s.Mcall2)
	go s.submitter.work(ctx, task)
	time.Sleep(time.Millisecond * 25)
	s.Hub.Commit()

	// assert
	length, _ := s.TSCCV.SccAssertLogsLen(&bind.CallOpts{Context: ctx})
	s.Equal(uint64(3), length.Uint64())

	for i := range indexes {
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
			s.Equal(i < len(indexes)-1, got.Approve)

			s.Len(got.Signatures, len(signers)*65)
			for j, sig := range signatures[i] {
				start := j * 65
				end := start + 65
				s.Equal(sig.Signature[:], got.Signatures[start:end])
			}
		}
	}
}

func (s *SubmitterTestSuite) TestSubmitL2OO() {
	ctx := context.Background()
	indexes := s.Range(0, 5)
	nextIndex := 2
	signers := s.StakeManager.Operators

	// save dummy signatures
	events := make([]*tl2oo.L2ooOutputProposed, len(indexes))
	signatures := make([][]*database.OpstackSignature, len(indexes))
	for i := range indexes {
		events[i] = s.EmitOutputProposed(i)
		signatures[i] = make([]*database.OpstackSignature, len(signers))

		for j := range s.Range(0, len(signers)) {
			signatures[i][j], _ = s.DB.OPStack.SaveSignature(
				nil, nil,
				signers[j],
				s.L2OOAddr,
				events[i].L2OutputIndex.Uint64(),
				events[i].OutputRoot,
				events[i].L2BlockNumber.Uint64(),
				events[i].L1Timestamp.Uint64(),
				i < len(indexes)-1,
				database.RandSignature(),
			)
		}

		sort.Sort(opstackSignatures(signatures[i]))
	}

	// set the `L2OO.nextVerifyIndex`
	s.TL2OO.SetNextVerifyIndex(s.Hub.TransactOpts(ctx), big.NewInt(int64(nextIndex)))
	s.Hub.Commit()

	// submitter do the work.
	s.submitter.stakemanager.Refresh(ctx)
	task, _ := NewTask(s.Hub, s.L2OOAddr, s.L2OO, s.SCCV, s.Mcall2)
	go s.submitter.work(ctx, task)
	time.Sleep(time.Millisecond * 25)
	s.Hub.Commit()

	// assert
	length, _ := s.TSCCV.L2ooAssertLogsLen(&bind.CallOpts{Context: ctx})
	s.Equal(uint64(len(indexes)-nextIndex), length.Uint64())

	for i := range indexes {
		if i < nextIndex {
			got, err := s.TSCCV.L2ooAssertLogs(
				&bind.CallOpts{Context: ctx},
				big.NewInt(int64(i+nextIndex+1)))
			s.ErrorContains(err, "execution reverted")
			s.Equal(common.Address{}, got.L2OutputOracle)
		} else {
			got, err := s.TSCCV.L2ooAssertLogs(
				&bind.CallOpts{Context: ctx},
				big.NewInt(int64(i-nextIndex)))
			s.NoError(err)
			s.Equal(s.L2OOAddr, got.L2OutputOracle)
			s.Equal(events[i].L2OutputIndex.Uint64(), got.L2OutputIndex.Uint64())
			s.Equal(events[i].OutputRoot, got.L2Output.OutputRoot)
			s.Equal(events[i].L1Timestamp.Uint64(), got.L2Output.Timestamp.Uint64())
			s.Equal(events[i].L2BlockNumber.Uint64(), got.L2Output.L2BlockNumber.Uint64())
			s.Equal(i < len(indexes)-1, got.Approve)

			s.Len(got.Signatures, len(signers)*65)
			for j, sig := range signatures[i] {
				start := j * 65
				end := start + 65
				s.Equal(sig.Signature[:], got.Signatures[start:end])
			}
		}
	}
}

type dummySignature struct {
	signer    common.Address
	signature database.Signature
	key       string
}

type dummySignatures []*dummySignature

func (sigs dummySignatures) Len() int                           { return len(sigs) }
func (sigs dummySignatures) Get(i int) interface{}              { return sigs[i] }
func (sigs dummySignatures) Signer(i int) common.Address        { return sigs[i].signer }
func (sigs dummySignatures) Signature(i int) database.Signature { return sigs[i].signature }
func (sigs dummySignatures) Key(i int) string                   { return sigs[i].key }

func (s *SubmitterTestSuite) TestGetTopStakingSignatures() {
	var (
		groups = []struct {
			key        string
			stake      int64
			signatures []*dummySignature
		}{
			{"group-0", 1000, make([]*dummySignature, 5)},
			{"group-1", 2000, make([]*dummySignature, 10)},
			{"group-2", 10000, make([]*dummySignature, 15)}, // want
			{"group-3", 3000, make([]*dummySignature, 20)},
		}
		want = groups[2]

		minStake        = big.NewInt(0)
		totalStake      = big.NewInt(0)
		allSignatures   = dummySignatures{}
		stakeBySigner   = map[common.Address]*big.Int{}
		stakeBySignerFn = func(addr common.Address) *big.Int {
			return stakeBySigner[addr]
		}
	)

	for _, group := range groups {
		totalStake.Add(totalStake, big.NewInt(group.stake))

		for i := range s.Range(0, len(group.signatures)) {
			signer := s.RandAddress()
			group.signatures[i] = &dummySignature{
				signer:    signer,
				signature: database.RandSignature(),
				key:       group.key,
			}
			allSignatures = append(allSignatures, group.signatures[i])
			stakeBySigner[signer] = big.NewInt(group.stake / int64(len(group.signatures)))
		}
	}
	sort.Slice(want.signatures, func(i, j int) bool {
		return bytes.Compare(want.signatures[i].signer[:], want.signatures[j].signer[:]) == -1
	})

	// ok
	got0, got1, _ := getTopStakingSignatures(allSignatures, minStake, totalStake, stakeBySignerFn)
	s.Len(got0, len(want.signatures))
	s.Equal(want.signatures[0].signer, allSignatures[got1].signer)
	for i, want := range want.signatures {
		s.Equal(want.signature.Bytes(), got0[i])
	}

	// error
	_, _, err := getTopStakingSignatures(allSignatures, minStake, big.NewInt(20000), stakeBySignerFn)
	s.ErrorContains(err, "stake amount shortage")
}
