package verifier

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/verse"
	"github.com/stretchr/testify/suite"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
)

type VerifierTestSuite struct {
	backend.BackendSuite

	verifier *Verifier
	task     verse.VerifiableVerse
}

func TestSccVerifier(t *testing.T) {
	suite.Run(t, new(VerifierTestSuite))
}

func (s *VerifierTestSuite) SetupTest() {
	s.BackendSuite.SetupTest()

	// setup verifier
	s.verifier = NewVerifier(&config.Verifier{
		Interval:            50 * time.Millisecond,
		Concurrency:         10,
		StateCollectLimit:   2,
		StateCollectTimeout: time.Second,
	}, s.DB, s.SignableHub.SignerContext())

	s.task = verse.NewOPLegacy(s.DB, s.Hub, s.SCCAddr).WithVerifiable(s.Verse)
	s.verifier.AddTask(s.task)
}

func (s *VerifierTestSuite) TestVerify() {
	type case_ struct {
		batchRoot     common.Hash
		wantApproved  bool
		wantSignature database.Signature
	}
	cases := []*case_{
		{wantApproved: true},
		{wantApproved: false},
	}

	batchSize := 10

	// send transactions to verse-layer
	for i, tt := range cases {
		stateRoots := make([][32]byte, batchSize)
		for j, h := range s.sendVerseTX(s.RandAddress(), batchSize) {
			if tt.wantApproved {
				stateRoots[j] = h.Root
			} else {
				stateRoots[j] = s.RandHash()
			}
		}

		batchRoot, err := verse.CalcMerkleRoot(stateRoots)
		s.Nil(err)

		wantSignature, err := ethutil.NewMessage(
			s.SignableHub.ChainID(),
			s.SCCAddr,
			big.NewInt(int64(i)),
			batchRoot,
			tt.wantApproved).Signature(s.SignableHub.SignData)
		s.Nil(err)

		tt.batchRoot = batchRoot
		tt.wantSignature = wantSignature
	}

	// save events
	for i, tt := range cases {
		_, err := s.task.EventDB().Save(
			s.task.RollupContract(),
			&scc.SccStateBatchAppended{
				BatchIndex:        big.NewInt(int64(i)),
				BatchRoot:         tt.batchRoot,
				BatchSize:         big.NewInt(int64(batchSize)),
				PrevTotalElements: big.NewInt(int64(batchSize * i)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", batchSize))})
		s.Nil(err)
	}

	// subscribe new signature
	published := s.startAndWait(len(cases))

	// assert
	for i, tt := range cases {
		ui64 := uint64(i)

		got0, err := s.DB.OPSignature.Find(nil, nil, nil, &ui64, 1, 0)
		got1 := published[i]
		s.Nil(err)

		s.IsType("string", got0[0].ID)
		s.IsType("string", got0[0].PreviousID)
		s.IsType("string", got1.ID)
		s.IsType("string", got1.PreviousID)
		s.Greater(got0[0].ID, got0[0].PreviousID)
		s.Greater(got1.ID, got1.PreviousID)

		s.Equal(s.SignableHub.Signer(), got0[0].Signer.Address)
		s.Equal(s.SignableHub.Signer(), got1.Signer.Address)

		s.Equal(s.SCCAddr, got0[0].Contract.Address)
		s.Equal(s.SCCAddr, got1.Contract.Address)

		s.Equal(ui64, got0[0].RollupIndex)
		s.Equal(ui64, got1.RollupIndex)

		s.Equal(tt.batchRoot[:], got0[0].RollupHash[:])
		s.Equal(tt.batchRoot[:], got1.RollupHash[:])

		s.Equal(tt.wantApproved, got0[0].Approved)
		s.Equal(tt.wantApproved, got1.Approved)

		s.Equal(tt.wantSignature[:], got0[0].Signature[:])
		s.Equal(tt.wantSignature[:], got1.Signature[:])
	}
}

func (s *VerifierTestSuite) TestDeleteInvalidNextIndexSignature() {
	batches := s.Range(0, 10)
	batchSize := 5
	invalidBatch := 6

	// send transactions to verse-layer
	merkleRoots := make([][32]byte, len(batches))
	for batchIdx := range batches {
		elements := make([][32]byte, batchSize)
		for i, header := range s.sendVerseTX(s.RandAddress(), batchSize) {
			elements[i] = header.Root
		}
		if merkleRoot, err := verse.CalcMerkleRoot(elements); s.NoError(err) {
			merkleRoots[batchIdx] = merkleRoot
		}
	}

	createds := make([]*database.OptimismSignature, len(batches))
	for batchIdx, merkleRoot := range merkleRoots {
		// save verify waiting state
		s.task.EventDB().Save(
			s.task.RollupContract(),
			&scc.SccStateBatchAppended{
				BatchIndex:        big.NewInt(int64(batchIdx)),
				BatchRoot:         merkleRoot,
				BatchSize:         big.NewInt(int64(batchSize)),
				PrevTotalElements: big.NewInt(int64(batchIdx * batchSize)),
				ExtraData:         []byte(fmt.Sprintf("test-%d", batchIdx))})

		// run verification
		published := s.startAndWait(1)
		s.Len(published, 1)
		s.Equal(merkleRoot[:], published[0].RollupHash[:])
		createds[batchIdx] = published[0]
	}

	// increment `nextIndex`
	for batchIdx := range s.Range(0, invalidBatch) {
		s.TSCC.EmitStateBatchVerified(
			s.SignableHub.TransactOpts(context.Background()),
			big.NewInt(int64(batchIdx)),
			merkleRoots[batchIdx],
		)
		s.Hub.Commit()
	}

	// run `deleteInvalidSignature`, but nothing happens
	published := s.startAndWait(1)
	s.Len(published, 0)

	signer := s.SignableHub.Signer()
	gots, _ := s.DB.OPSignature.Find(nil, &signer, &s.SCCAddr, nil, 100, 0)
	s.Equal(len(batches), len(gots))

	for batchIdx := range batches {
		// should not be re-created
		s.Equal(createds[batchIdx].ID, gots[batchIdx].ID)
		s.Equal(createds[batchIdx].Signature, gots[batchIdx].Signature)
	}

	// update to invalid signature
	s.DB.OPSignature.Save(
		&createds[invalidBatch].ID,
		&createds[invalidBatch].PreviousID,
		createds[invalidBatch].Signer.Address,
		createds[invalidBatch].Contract.Address,
		createds[invalidBatch].RollupIndex,
		createds[invalidBatch].RollupHash,
		createds[invalidBatch].Approved,
		database.RandSignature())

	// run `deleteInvalidSignature`
	published = s.startAndWait(len(batches) - invalidBatch)
	s.Len(published, len(batches)-invalidBatch)

	gots, _ = s.DB.OPSignature.Find(nil, &signer, &s.SCCAddr, nil, 100, 0)
	s.Equal(len(batches), len(gots))

	for batchIdx := range batches {
		if batchIdx < invalidBatch {
			s.Equal(createds[batchIdx].ID, gots[batchIdx].ID)
		} else {
			// should be re-created
			s.NotEqual(createds[batchIdx].ID, gots[batchIdx].ID)
		}
		s.Equal(createds[batchIdx].Signature, gots[batchIdx].Signature)
	}
}

func (s *VerifierTestSuite) sendVerseTX(to common.Address, count int) (headers []*types.Header) {
	for i := 0; i < count; i++ {
		signedTx, err := s.SignableVerse.SignTx(types.NewTransaction(
			uint64(i), to, common.Big1, uint64(21_000), big.NewInt(875_000_000), nil))
		s.Nil(err)

		s.Verse.SendTransaction(context.Background(), signedTx)
		if h, err := s.Verse.HeaderByHash(context.Background(), s.Verse.Commit()); s.NoError(err) {
			headers = append(headers, h)
		}
	}
	return headers
}

func (s *VerifierTestSuite) startAndWait(count int) (published []*database.OptimismSignature) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()

	sub := s.verifier.SubscribeNewSignature(ctx)
	defer sub.Cancel()

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case sig := <-sub.Next():
				published = append(published, sig)
				if len(published) == count {
					return
				}
			}
		}

	}()

	go s.verifier.Start(ctx)
	<-ctx.Done()

	return published
}
