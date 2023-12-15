package hublayer

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/stretchr/testify/suite"
)

type SubmitterTestSuite struct {
	SccTestSuite
}

func TestSubmitter(t *testing.T) {
	suite.Run(t, new(SubmitterTestSuite))
}

func (s *SubmitterTestSuite) TestWork() {
	var (
		indexes = 5

		signers    [20]common.Address
		signatures [][]*database.OptimismSignature
	)

	for i := range s.Range(0, len(signers)) {
		signers[i] = s.RandAddress()

		s.sm.Owners = append(s.sm.Owners, s.RandAddress())
		s.sm.Operators = append(s.sm.Operators, signers[i])
		s.sm.Stakes = append(
			s.sm.Stakes,
			new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10_000_000)),
		)
		s.sm.Candidates = append(s.sm.Candidates, true)
		s.sm.NewCursor = big.NewInt(0)
	}

	for i := range s.Range(0, indexes) {
		signatures = append(signatures, make([]*database.OptimismSignature, len(signers)))

		batchIndex := uint64(i)
		batchRoot := s.RandHash()
		batchSize := uint64(i)
		prevTotalElements := uint64(i + 1)
		extraData := []byte(fmt.Sprintf("%d", i))
		approved := i < indexes-1

		// create sample signatures
		for j := range s.Range(0, len(signers)) {
			sig, _ := s.db.Optimism.SaveSignature(
				nil, nil,
				signers[j],
				s.sccAddr,
				batchIndex,
				batchRoot,
				batchSize,
				prevTotalElements,
				extraData,
				approved,
				database.RandSignature(),
			)

			signatures[i][j] = sig
		}
		sort.Slice(signatures[i], func(x, y int) bool {
			a := signatures[i][x].Signer.Address.Hash().Big()
			b := signatures[i][y].Signer.Address.Hash().Big()
			return a.Cmp(b) == -1
		})

		// emit StateBatchAppended event to the test contract
		s.tscc.EmitStateBatchAppended(
			s.hub.TransactOpts(context.Background()),
			new(big.Int).SetUint64(batchIndex),
			batchRoot,
			new(big.Int).SetUint64(batchSize),
			new(big.Int).SetUint64(prevTotalElements),
			extraData)
		s.mining()
	}

	s.submitter.vs.Refresh(context.Background())

	for range s.Range(0, indexes/s.submitter.cfg.BatchSize+1) {
		go func() {
			time.Sleep(10 * time.Millisecond)
			s.hub.Commit()
		}()
		s.submitter.work(context.Background(), NewSccSubmitTask(s.hub, s.sccAddr, s.scc))
	}

	for i := range s.Range(0, indexes) {
		got, _ := s.sccv.AssertLogs(
			&bind.CallOpts{Context: context.Background()},
			big.NewInt(int64(i)),
		)

		s.Equal(s.sccAddr, got.StateCommitmentChain)

		s.Equal(uint64(i), got.BatchHeader.BatchIndex.Uint64())
		s.Equal(signatures[i][0].BatchRoot[:], got.BatchHeader.BatchRoot[:])
		s.Equal(uint64(i), got.BatchHeader.BatchSize.Uint64())
		s.Equal(uint64(i+1), got.BatchHeader.PrevTotalElements.Uint64())
		s.Equal([]byte(fmt.Sprintf("%d", i)), got.BatchHeader.ExtraData)

		s.Len(got.Signatures, len(signers)*65)
		for j, sig := range signatures[i] {
			start := j * 65
			end := start + 65
			s.Equal(sig.Signature[:], got.Signatures[start:end])
		}

		s.Equal(i < indexes-1, got.Approve)
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
