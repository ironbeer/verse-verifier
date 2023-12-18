package collector

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
	tc "github.com/oasysgames/oasys-optimism-verifier/testhelper/contracts/scc"
	"github.com/stretchr/testify/suite"
)

type EventCollectorTestSuite struct {
	testhelper.Suite
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

func (s *EventCollectorTestSuite) TestProcessStateBatchAppendedEvent() {
	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.EmitStateBatchAppendedEvent(i))
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

func (s *EventCollectorTestSuite) TestProcessStateBatchDeletedEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.EmitStateBatchAppendedEvent(i))
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

func (s *EventCollectorTestSuite) TestProcessStateBatchVerifiedEvent() {
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

func (s *EventCollectorTestSuite) TestNoHandleOtherEvent() {
	ctx := context.Background()

	// emit `StateBatchAppended` and `Other` events
	for i := range s.Range(0, 10) {
		s.EmitStateBatchAppendedEvent(i)
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

func (s *EventCollectorTestSuite) TestHandleReorganization() {
	ctx := context.Background()

	// emit `StateBatchAppended` events
	var emits []*tc.SccStateBatchAppended
	for i := range s.Range(0, 10) {
		emits = append(emits, s.EmitStateBatchAppendedEvent(i))
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
	s.EmitStateBatchAppendedEvent(4)
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
