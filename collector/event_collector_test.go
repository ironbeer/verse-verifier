package collector

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	tl2oo "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/l2oo"
	tc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contract/scc"
	"github.com/stretchr/testify/suite"
)

type EventCollectorTestSuite struct {
	backend.BackendSuite

	collector *EventCollector
}

func TestEventCollector(t *testing.T) {
	suite.Run(t, new(EventCollectorTestSuite))
}

func (s *EventCollectorTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	s.collector = NewEventCollector(&config.Verifier{
		Interval:         time.Millisecond,
		EventFilterLimit: 1000,
	}, s.DB, s.Hub, s.Hub.Signer())
}

func (s *EventCollectorTestSuite) TestHandleStateBatchAppendedEvent() {
	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.EmitStateBatchAppended(i))
	}

	// collect `StateBatchAppended` events
	s.collector.work(context.Background())

	// assert
	for i := range s.Range(0, 10) {
		got, _ := s.DB.Optimism.FindState(s.SCCAddr, uint64(i))
		s.Equal(s.SCCAddr, got.OptimismScc.Address)
		s.Equal(emits[i].BatchIndex.Uint64(), got.BatchIndex)
		s.Equal(emits[i].BatchRoot[:], got.BatchRoot[:])
		s.Equal(emits[i].BatchSize.Uint64(), got.BatchSize)
		s.Equal(emits[i].PrevTotalElements.Uint64(), got.PrevTotalElements)
		s.Equal(emits[i].ExtraData, got.ExtraData)
	}
}

func (s *EventCollectorTestSuite) TestHandleStateBatchDeletedEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.EmitStateBatchAppended(i))
	}

	// collect `StateBatchAppended` events
	s.collector.work(ctx)

	// create signature records
	var creates []*database.OptimismSignature
	for i := range s.Range(0, 10) {
		sig, _ := s.DB.Optimism.SaveSignature(
			nil, nil,
			s.Hub.Signer(),
			s.SCCAddr,
			emits[i].BatchIndex.Uint64(),
			emits[i].BatchRoot,
			0,
			0,
			[]byte(nil),
			true,
			database.RandSignature(),
		)
		creates = append(creates, sig)
	}

	// emit `StateBatchDeleted` event
	s.TSCC.EmitStateBatchDeleted(
		s.Hub.TransactOpts(ctx),
		emits[5].BatchIndex,
		emits[5].BatchRoot,
	)
	s.Mining()

	// collect `StateBatchDeleted` events
	s.collector.work(ctx)

	// assert
	for i := range s.Range(0, 10) {
		var want error
		if i >= 5 {
			want = database.ErrNotFound
		}
		_, err0 := s.DB.Optimism.FindState(s.SCCAddr, uint64(i))
		_, err1 := s.DB.Optimism.FindSignatureByID(creates[i].ID)
		s.Equal(want, err0)
		s.Equal(want, err1)
	}
}

func (s *EventCollectorTestSuite) TestHandleStateBatchVerifiedEvent() {
	// emit `EmitStateBatchVerified` events
	for index := range s.Range(0, 5) {
		s.TSCC.EmitStateBatchVerified(
			s.Hub.TransactOpts(context.Background()),
			big.NewInt(int64(index)),
			s.RandHash(),
		)
		s.Mining()
	}

	// collect `EmitStateBatchVerified` events
	s.collector.work(context.Background())

	// assert
	scc, _ := s.DB.Optimism.FindOrCreateSCC(s.SCCAddr)
	s.Equal(uint64(5), scc.NextIndex)
}

func (s *EventCollectorTestSuite) TestHandleOutputProposedEvent() {
	// emit `OutputProposed` events
	var events []*tl2oo.L2ooOutputProposed
	for i := range s.Range(0, 10) {
		events = append(events, s.EmitOutputProposed(i))
	}

	// collect `OutputProposed` events
	s.collector.work(context.Background())

	// assert
	for i := range s.Range(0, 10) {
		got, err := s.DB.OPStack.FindProposal(s.L2OOAddr, uint64(i))
		s.Nil(err)
		s.Equal(s.L2OOAddr, got.OpstackL2OutputOracle.Address)
		s.Equal(events[i].OutputRoot[:], got.OutputRoot[:])
		s.Equal(events[i].L2OutputIndex.Uint64(), got.L2OutputIndex)
		s.Equal(events[i].L2BlockNumber.Uint64(), got.L2BlockNumber)
		s.Equal(events[i].L1Timestamp.Uint64(), got.L1Timestamp)
	}
}

func (s *EventCollectorTestSuite) TestHandleOutputsDeletedEvent() {
	ctx := context.Background()

	// emit `OutputProposed` events
	var events []*tl2oo.L2ooOutputProposed
	for i := range s.Range(0, 10) {
		events = append(events, s.EmitOutputProposed(i))
	}

	// collect `OutputProposed` events
	s.collector.work(ctx)

	// create signature records
	var creates []*database.OpstackSignature
	for i := range s.Range(0, 10) {
		sig, err := s.DB.OPStack.SaveSignature(
			nil, nil,
			s.Hub.Signer(),
			s.L2OOAddr,
			events[i].L2OutputIndex.Uint64(),
			events[i].OutputRoot,
			events[i].L2BlockNumber.Uint64(),
			events[i].L1Timestamp.Uint64(),
			true,
			database.RandSignature(),
		)
		s.Nil(err)
		creates = append(creates, sig)
	}

	// emit `OutputsDeleted` event
	_, err := s.TL2OO.EmitOutputsDeleted(
		s.Hub.TransactOpts(ctx),
		new(big.Int).Add(events[len(events)-1].L2OutputIndex, common.Big1),
		events[5].L2OutputIndex,
	)
	s.Nil(err)
	s.Mining()

	// collect `OutputsDeleted` events
	s.collector.work(ctx)

	// assert
	for i := range s.Range(0, 10) {
		proposal, err0 := s.DB.OPStack.FindProposal(s.L2OOAddr, uint64(i))
		sig, err1 := s.DB.OPStack.FindSignatureByID(creates[i].ID)

		if i >= 5 {
			s.Nil(proposal)
			s.Nil(sig)
			s.Equal(database.ErrNotFound, err0)
			s.Equal(database.ErrNotFound, err1)
		} else {
			s.NotNil(proposal)
			s.NotNil(sig)
			s.Nil(err0)
			s.Nil(err1)
		}
	}
}

func (s *EventCollectorTestSuite) TestHandleOutputVerifiedEvent() {
	// emit `OutputVerified` events
	for index := range s.Range(0, 5) {
		s.TL2OO.EmitOutputVerified(
			s.Hub.TransactOpts(context.Background()),
			big.NewInt(int64(index)),
			s.RandHash(),
			big.NewInt(int64(index*10)),
		)
		s.Mining()
	}

	// collect `OutputVerified` events
	s.collector.work(context.Background())

	// assert
	l2oo, _ := s.DB.OPStack.FindOrCreateL2OO(s.L2OOAddr)
	s.Equal(uint64(5), l2oo.NextVerifyIndex)
}

func (s *EventCollectorTestSuite) TestNoHandleOtherEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` and `Other` events
	for i := range s.Range(0, 10) {
		s.EmitStateBatchAppended(i)
		s.TSCC.EmitOtherEvent(s.Hub.TransactOpts(ctx), big.NewInt(11))
		s.Mining()
	}

	// collect `StateBatchAppended` events
	s.collector.work(ctx)

	// assert
	for i := range s.Range(0, 20) {
		var want error
		if i >= 10 {
			want = database.ErrNotFound
		}
		_, err := s.DB.Optimism.FindState(s.SCCAddr, uint64(i))
		s.ErrorIs(err, want)
	}
}

func (s *EventCollectorTestSuite) TestHandleSCCReorganization() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.EmitStateBatchAppended(i))
	}

	// collect `StateBatchAppended` events
	s.collector.work(ctx)

	// create signature records
	var creates []*database.OptimismSignature
	for i := range s.Range(0, 10) {
		sig, _ := s.DB.Optimism.SaveSignature(
			nil, nil,
			s.Hub.Signer(),
			s.SCCAddr,
			emits[i].BatchIndex.Uint64(),
			emits[i].BatchRoot,
			emits[i].BatchSize.Uint64(),
			emits[i].PrevTotalElements.Uint64(),
			emits[i].ExtraData,
			true,
			database.RandSignature(),
		)
		creates = append(creates, sig)
	}

	// simulate chain reorganization
	s.EmitStateBatchAppended(4)
	s.collector.work(ctx)

	// assert
	for i := range s.Range(0, 10) {
		_, err := s.DB.Optimism.FindState(s.SCCAddr, uint64(i))
		if i < 5 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}

		_, err = s.DB.Optimism.FindSignatureByID(creates[i].ID)
		if i < 4 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}
	}
}

func (s *EventCollectorTestSuite) TestHandleL2OOReorganization() {
	ctx := context.Background()

	// emit `OutputProposed` events
	var events []*tl2oo.L2ooOutputProposed
	for i := range s.Range(0, 10) {
		events = append(events, s.EmitOutputProposed(i))
	}

	// collect `OutputProposed` events
	s.collector.work(ctx)

	// create signature records
	var creates []*database.OpstackSignature
	for i := range s.Range(0, 10) {
		sig, err := s.DB.OPStack.SaveSignature(
			nil, nil,
			s.Hub.Signer(),
			s.L2OOAddr,
			events[i].L2OutputIndex.Uint64(),
			events[i].OutputRoot,
			events[i].L2BlockNumber.Uint64(),
			events[i].L1Timestamp.Uint64(),
			true,
			database.RandSignature(),
		)
		s.Nil(err)
		creates = append(creates, sig)
	}

	// simulate chain reorganization
	s.EmitOutputProposed(4)
	s.collector.work(ctx)

	// assert
	for i := range s.Range(0, 10) {
		_, err := s.DB.OPStack.FindProposal(s.L2OOAddr, uint64(i))
		if i < 5 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}

		_, err = s.DB.OPStack.FindSignatureByID(creates[i].ID)
		if i < 4 {
			s.NoError(err)
		} else {
			s.Error(err, database.ErrNotFound)
		}
	}
}
