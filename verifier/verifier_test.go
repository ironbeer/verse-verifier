package verifier

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/contract/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/database"
	"github.com/oasysgames/oasys-optimism-verifier/ethutil"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/stretchr/testify/suite"

	"github.com/oasysgames/oasys-optimism-verifier/testhelper/backend"
)

type VerifierTestSuite struct {
	backend.BackendSuite

	verifier            *Verifier
	l2ToL1MessagePasser common.Address
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
	}, s.DB, s.Hub.SignerContext())

	s.verifier.AddWorker(NewSccVerifyWorker(s.Verse, s.SCCAddr, s.SCC))
	s.verifier.AddWorker(NewL2OOVerifyWorker(s.Verse, s.L2OOAddr, s.L2OO))

	s.l2ToL1MessagePasser = common.HexToAddress("0x4200000000000000000000000000000000000016")
}

func (s *VerifierTestSuite) TestVerifySCC() {
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

		batchRoot, err := calcMerkleRoot(stateRoots)
		s.Nil(err)

		wantSignature, err := NewSccMessage(
			s.Hub.ChainID(),
			s.SCCAddr,
			big.NewInt(int64(i)),
			batchRoot,
			tt.wantApproved).Signature(s.Hub.SignData)
		s.Nil(err)

		tt.batchRoot = batchRoot
		tt.wantSignature = wantSignature
	}

	// emit and collect `StateBatchAppended` events
	for i, tt := range cases {
		_, err := s.DB.Optimism.SaveState(&scc.SccStateBatchAppended{
			Raw:               types.Log{Address: s.SCCAddr},
			BatchIndex:        big.NewInt(int64(i)),
			BatchRoot:         tt.batchRoot,
			BatchSize:         big.NewInt(int64(batchSize)),
			PrevTotalElements: big.NewInt(int64(batchSize * i)),
			ExtraData:         []byte(fmt.Sprintf("test-%d", batchSize))})
		s.Nil(err)
	}

	// subscribe new signature
	published, _ := startAndWait(s.verifier, len(cases))

	// assert
	for i, tt := range cases {
		ui64 := uint64(i)

		got0, err := s.DB.Optimism.FindSignatures(nil, nil, nil, &ui64, 1, 0)
		got1 := published[i]
		s.Nil(err)

		s.IsType("string", got0[0].ID)
		s.IsType("string", got0[0].PreviousID)
		s.IsType("string", got1.ID)
		s.IsType("string", got1.PreviousID)
		s.Greater(got0[0].ID, got0[0].PreviousID)
		s.Greater(got1.ID, got1.PreviousID)

		s.Equal(s.Hub.Signer(), got0[0].Signer.Address)
		s.Equal(s.Hub.Signer(), got1.Signer.Address)

		s.Equal(s.SCCAddr, got0[0].OptimismScc.Address)
		s.Equal(s.SCCAddr, got1.OptimismScc.Address)

		s.Equal(ui64, got0[0].BatchIndex)
		s.Equal(ui64, got1.BatchIndex)

		s.Equal(tt.batchRoot[:], got0[0].BatchRoot[:])
		s.Equal(tt.batchRoot[:], got1.BatchRoot[:])

		s.Equal(uint64(batchSize), got0[0].BatchSize)
		s.Equal(uint64(batchSize), got1.BatchSize)

		s.Equal(uint64(batchSize*i), got0[0].PrevTotalElements)
		s.Equal(uint64(batchSize*i), got1.PrevTotalElements)

		s.Equal(fmt.Sprintf("test-%d", batchSize), string(got0[0].ExtraData))
		s.Equal(fmt.Sprintf("test-%d", batchSize), string(got1.ExtraData))

		s.Equal(tt.wantApproved, got0[0].Approved)
		s.Equal(tt.wantApproved, got1.Approved)

		s.Equal(tt.wantSignature[:], got0[0].Signature[:])
		s.Equal(tt.wantSignature[:], got1.Signature[:])
	}
}

func (s *VerifierTestSuite) TestVerifyL2OO() {
	type case_ struct {
		outputRoot    [32]byte
		wantApproved  bool
		wantSignature database.Signature
	}
	cases := []*case_{
		{wantApproved: true},
		{wantApproved: false},
	}

	l1Timestamp := time.Now().Unix()
	l2Block, err := s.Verse.BlockNumber(context.Background())
	s.Nil(err)

	// send transactions to verse-layer
	for i, tt := range cases {
		head := s.sendVerseTX(s.l2ToL1MessagePasser, 1)[0]

		proof, err := s.Verse.GetProof(
			context.Background(), s.l2ToL1MessagePasser, []string{}, head.Number)
		s.Nil(err)

		var outputRoot [32]byte
		if tt.wantApproved {
			outputV0 := &ethutil.OpstackOutputV0{
				StateRoot:                head.Root,
				MessagePasserStorageRoot: proof.StorageHash,
				BlockHash:                head.Hash(),
			}
			outputRoot = outputV0.OutputRoot()
		} else {
			outputRoot = s.RandHash()
		}

		wantSignature, err := NewL2ooMessage(
			s.Hub.ChainID(),
			s.L2OOAddr,
			big.NewInt(int64(i)),
			outputRoot,
			big.NewInt(l1Timestamp+int64(i)),
			big.NewInt(int64(l2Block)+int64(i)+1),
			tt.wantApproved).Signature(s.Hub.SignData)
		s.Nil(err)

		tt.outputRoot = outputRoot
		tt.wantSignature = wantSignature
	}

	// emit and collect `OutputProposed` events
	for i, tt := range cases {
		_, err := s.DB.OPStack.SaveProposal(&l2oo.OasysL2OutputOracleOutputProposed{
			Raw:           types.Log{Address: s.L2OOAddr},
			OutputRoot:    tt.outputRoot,
			L2OutputIndex: big.NewInt(int64(i)),
			L2BlockNumber: big.NewInt(int64(l2Block) + int64(i) + 1),
			L1Timestamp:   big.NewInt(l1Timestamp + int64(i)),
		})
		s.Nil(err)
	}

	// subscribe new signature
	_, published := startAndWait(s.verifier, len(cases))

	// assert
	for i, tt := range cases {
		ui64 := uint64(i)

		got0, err := s.DB.OPStack.FindSignatures(nil, nil, nil, &ui64, 1, 0)
		got1 := published[i]
		s.Nil(err)

		s.IsType("string", got0[0].ID)
		s.IsType("string", got0[0].PreviousID)
		s.IsType("string", got1.ID)
		s.IsType("string", got1.PreviousID)
		s.Greater(got0[0].ID, got0[0].PreviousID)
		s.Greater(got1.ID, got1.PreviousID)

		s.Equal(s.Hub.Signer(), got0[0].Signer.Address)
		s.Equal(s.Hub.Signer(), got1.Signer.Address)

		s.Equal(s.L2OOAddr, got0[0].OpstackL2OutputOracle.Address)
		s.Equal(s.L2OOAddr, got1.OpstackL2OutputOracle.Address)

		s.Equal(ui64, got0[0].L2OutputIndex)
		s.Equal(ui64, got1.L2OutputIndex)

		s.Equal(tt.outputRoot[:], got0[0].OutputRoot[:])
		s.Equal(tt.outputRoot[:], got1.OutputRoot[:])

		s.Equal(l2Block+ui64+1, got0[0].L2BlockNumber)
		s.Equal(l2Block+ui64+1, got1.L2BlockNumber)

		s.Equal(uint64(l1Timestamp)+ui64, got0[0].L1Timestamp)
		s.Equal(uint64(l1Timestamp)+ui64, got1.L1Timestamp)

		s.Equal(tt.wantApproved, got0[0].Approved)
		s.Equal(tt.wantApproved, got1.Approved)

		s.Equal(tt.wantSignature[:], got0[0].Signature[:])
		s.Equal(tt.wantSignature[:], got1.Signature[:])
	}
}

func (s *VerifierTestSuite) TestDeleteInvalidSCCSignature() {
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
		if merkleRoot, err := calcMerkleRoot(elements); s.NoError(err) {
			merkleRoots[batchIdx] = merkleRoot
		}
	}

	createds := make([]*database.OptimismSignature, len(batches))
	for batchIdx, merkleRoot := range merkleRoots {
		// save verify waiting state
		s.DB.Optimism.SaveState(&scc.SccStateBatchAppended{
			Raw:               types.Log{Address: s.SCCAddr},
			BatchIndex:        big.NewInt(int64(batchIdx)),
			BatchRoot:         merkleRoot,
			BatchSize:         big.NewInt(int64(batchSize)),
			PrevTotalElements: big.NewInt(int64(batchIdx * batchSize)),
			ExtraData:         []byte(fmt.Sprintf("test-%d", batchIdx))})

		// run verification
		published, _ := startAndWait(s.verifier, 1)
		s.Len(published, 1)
		s.Equal(merkleRoot[:], published[0].BatchRoot[:])
		createds[batchIdx] = published[0]
	}

	// increment `nextIndex`
	for batchIdx := range s.Range(0, invalidBatch) {
		s.TSCC.EmitStateBatchVerified(
			s.Hub.TransactOpts(context.Background()),
			big.NewInt(int64(batchIdx)),
			merkleRoots[batchIdx],
		)
		s.Hub.Commit()
	}

	// run `deleteInvalidSignature`, but nothing happens
	published, _ := startAndWait(s.verifier, 1)
	s.Len(published, 0)

	signer := s.Hub.Signer()
	gots, _ := s.DB.Optimism.FindSignatures(nil, &signer, &s.SCCAddr, nil, 100, 0)
	s.Equal(len(batches), len(gots))

	for batchIdx := range batches {
		// should not be re-created
		s.Equal(createds[batchIdx].ID, gots[batchIdx].ID)
		s.Equal(createds[batchIdx].Signature, gots[batchIdx].Signature)
	}

	// update to invalid signature
	s.DB.Optimism.SaveSignature(
		&createds[invalidBatch].ID,
		&createds[invalidBatch].PreviousID,
		createds[invalidBatch].Signer.Address,
		createds[invalidBatch].OptimismScc.Address,
		createds[invalidBatch].BatchIndex,
		createds[invalidBatch].BatchRoot,
		createds[invalidBatch].BatchSize,
		createds[invalidBatch].PrevTotalElements,
		createds[invalidBatch].ExtraData,
		createds[invalidBatch].Approved,
		database.RandSignature())

	// run `deleteInvalidSignature`
	published, _ = startAndWait(s.verifier, len(batches)-invalidBatch)
	s.Len(published, len(batches)-invalidBatch)

	gots, _ = s.DB.Optimism.FindSignatures(nil, &signer, &s.SCCAddr, nil, 100, 0)
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

func (s *VerifierTestSuite) TestDeleteInvalidL2OOSignature() {
	invalidOutputIdx := 6

	l2Block, err := s.Verse.BlockNumber(context.Background())
	s.Nil(err)

	// send transactions to verse-layer
	outputRoots := make([][32]byte, 10)
	for i := range outputRoots {
		head := s.sendVerseTX(s.l2ToL1MessagePasser, 1)[0]

		proof, err := s.Verse.GetProof(
			context.Background(), s.l2ToL1MessagePasser, []string{}, head.Number)
		s.Nil(err)

		outputV0 := &ethutil.OpstackOutputV0{
			StateRoot:                head.Root,
			MessagePasserStorageRoot: proof.StorageHash,
			BlockHash:                head.Hash(),
		}
		outputRoots[i] = outputV0.OutputRoot()
	}

	createds := make([]*database.OpstackSignature, len(outputRoots))
	for outputIdx, outputRoot := range outputRoots {
		// save verify waiting proposal
		s.DB.OPStack.SaveProposal(&l2oo.OasysL2OutputOracleOutputProposed{
			Raw:           types.Log{Address: s.L2OOAddr},
			OutputRoot:    outputRoot,
			L2OutputIndex: big.NewInt(int64(outputIdx)),
			L2BlockNumber: big.NewInt(int64(l2Block) + int64(outputIdx) + 1),
			L1Timestamp:   big.NewInt(time.Now().Unix()),
		})

		// run verification
		_, published := startAndWait(s.verifier, 1)
		s.Len(published, 1)
		s.Equal(outputRoot[:], published[0].OutputRoot[:])
		createds[outputIdx] = published[0]
	}

	// increments `nextVerifyIndex`
	_, err = s.TL2OO.SetNextVerifyIndex(
		s.Hub.TransactOpts(context.Background()), big.NewInt(int64(invalidOutputIdx)))
	s.Nil(err)
	s.Hub.Commit()

	// run `deleteInvalidSignature`, but nothing happens
	_, published := startAndWait(s.verifier, 1)
	s.Len(published, 0)

	gots, _ := s.DB.OPStack.FindSignatures(
		nil, &s.verifier.signerCtx.Signer, &s.L2OOAddr, nil, 100, 0)
	s.Equal(len(outputRoots), len(gots))

	for outputIdx := range outputRoots {
		// should not be re-created
		s.Equal(createds[outputIdx].ID, gots[outputIdx].ID)
		s.Equal(createds[outputIdx].Signature, gots[outputIdx].Signature)
	}

	// update to invalid signature
	_, err = s.DB.OPStack.SaveSignature(
		&createds[invalidOutputIdx].ID,
		&createds[invalidOutputIdx].PreviousID,
		createds[invalidOutputIdx].Signer.Address,
		createds[invalidOutputIdx].OpstackL2OutputOracle.Address,
		createds[invalidOutputIdx].L2OutputIndex,
		createds[invalidOutputIdx].OutputRoot,
		createds[invalidOutputIdx].L2BlockNumber,
		createds[invalidOutputIdx].L1Timestamp,
		createds[invalidOutputIdx].Approved,
		database.RandSignature())
	s.Nil(err)

	// run `deleteInvalidSignature`
	_, published = startAndWait(s.verifier, len(outputRoots)-invalidOutputIdx)
	s.Len(published, len(outputRoots)-invalidOutputIdx)

	gots, _ = s.DB.OPStack.FindSignatures(nil, &s.verifier.signerCtx.Signer, &s.L2OOAddr, nil, 100, 0)
	s.Equal(len(outputRoots), len(gots))

	for outputIdx := range outputRoots {
		if outputIdx < invalidOutputIdx {
			s.Equal(createds[outputIdx].ID, gots[outputIdx].ID)
		} else {
			// should be re-created
			s.NotEqual(createds[outputIdx].ID, gots[outputIdx].ID)
		}
		s.Equal(createds[outputIdx].Signature, gots[outputIdx].Signature)
	}
}

func (s *VerifierTestSuite) TestCalcMerkleRoot() {
	cases := []struct {
		name string
		spec func()
	}{
		{
			"no elements",
			func() {
				_, err := calcMerkleRoot([][32]byte{})
				s.ErrorContains(err, "must provide at least one leaf hash")
			},
		},
		{
			"single element",
			func() {
				elements := [][32]byte{
					util.BytesToBytes32(
						common.FromHex(
							"0x56570de287d73cd1cb6092bb8fdee6173974955fdef345ae579ee9f475ea7432",
						),
					),
				}
				got, _ := calcMerkleRoot(elements)
				s.Equal(elements[0], got)
			},
		},
		{
			"more than one element",
			func() {
				wants := []string{
					"0x57d772147cdf27f5f67d679f0f3a513f8b87622ce598a3cf0b048ab178ddfc6e",
					"0x820919791e2ec4aea2fb218a7a3a5a89d06ba469585c824b60f0174ec13e1603",
					"0xe39e9f65a0fcee19f9b8aceadb3bbdbf7697be66b0632644e168d01dc103ddd6",
					"0x11f470d712bb3a84f0b01cb7c73493ec7d06eda480f567c99b9a6dc773679a72",
					"0x0faa9fa71909342540cabef2fdf911cf053141144b21d089641940533679cdd9",
					"0x0050d8ac9e23f46daf8be33332d201588cba3cee5c6141715756dc4b2c960ada",
					"0xfc61b646f502f97300b88afe15feaf046f90c8456f658273657d8a55e7fc79df",
					"0xa4329a43ffc1bc6195e1bddda04930ed1db6486df03a56a8df9a60bb2d5469e0",
					"0x9a68f697fd78c779e436dec655825d263066c5fee23f961fd15e9d14327ded6b",
					"0x437dc148af6b33ba532cf6e8d8c0c74ab680439cbd03f9000f7434fb217611b7",
					"0xeade7c5f57e013547c7cec95eff59e44616ab9bdadb73420545f741e4097f9c1",
					"0x7acdf7918c5b5dc8acac506737231e143f2dc6b8734ec02b3d92676852fd4880",
					"0x9b9b9244ced25fff4077e6bca56882d106981a5d949394ad509bb0b11e04d23a",
					"0xbba15d82445e21878b48a0e4b19854c4e0e75a68e644bdfb8ace0fc965264431",
				}
				sizes := []int{2, 3, 7, 9, 13, 63, 64, 123, 128, 129, 255, 1021, 1023, 1024}
				for i := range wants {
					want := util.BytesToBytes32(common.FromHex(wants[i]))
					size := sizes[i]

					s.Run(fmt.Sprintf("size %d", size), func() {
						elements := make([][32]byte, size)
						for i := range elements {
							bhash := common.FromHex(hexutil.EncodeBig(big.NewInt(int64(i))))
							elements[i] = util.BytesToBytes32(crypto.Keccak256(bhash))
						}
						got, _ := calcMerkleRoot(s.fillDefaultHashes(elements))
						s.Equal(want, got)
					})
				}
			},
		},
		{
			"odd number of elements",
			func() {
				elements := [][32]byte{
					util.BytesToBytes32(crypto.Keccak256(common.FromHex("0x12"))),
					util.BytesToBytes32(crypto.Keccak256(common.FromHex("0x34"))),
					util.BytesToBytes32(crypto.Keccak256(common.FromHex("0x56"))),
				}
				_, err := calcMerkleRoot(elements)
				s.Equal(nil, err)
			},
		},
	}

	for _, tt := range cases {
		s.Run(tt.name, tt.spec)
	}
}

func (s *VerifierTestSuite) sendVerseTX(to common.Address, count int) (headers []*types.Header) {
	for i := 0; i < count; i++ {
		signedTx, err := s.Verse.SignTx(types.NewTransaction(
			uint64(i), to, common.Big1, uint64(21_000), big.NewInt(875_000_000), nil))
		s.Nil(err)

		s.Verse.SendTransaction(context.Background(), signedTx)
		if h, err := s.Verse.HeaderByHash(context.Background(), s.Verse.Commit()); s.NoError(err) {
			headers = append(headers, h)
		}
	}
	return headers
}

func (s *VerifierTestSuite) fillDefaultHashes(elements [][32]byte) [][32]byte {
	fillhash := util.BytesToBytes32(
		crypto.Keccak256(common.FromHex("0x" + strings.Repeat("00", 32))),
	)

	filled := [][32]byte{}
	for i := 0; float64(i) < math.Pow(2, math.Ceil(math.Log2(float64(len(elements))))); i++ {
		if i < len(elements) {
			filled = append(filled, elements[i])
		} else {
			filled = append(filled, fillhash)
		}
	}

	return filled
}

func startAndWait(
	verifier *Verifier,
	count int,
) (sccs []*database.OptimismSignature, l2oos []*database.OpstackSignature) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
	defer cancel()

	sub := verifier.SubscribeNewSignature(ctx)
	defer sub.Cancel()

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case sig := <-sub.Next():
				switch t := sig.(type) {
				case *database.OptimismSignature:
					sccs = append(sccs, t)
				case *database.OpstackSignature:
					l2oos = append(l2oos, t)
				default:
					panic(fmt.Errorf("Unknown signature: %v", sig))
				}

				if len(sccs)+len(l2oos) == count {
					return
				}
			}
		}

	}()

	go verifier.Start(ctx)
	<-ctx.Done()

	return sccs, l2oos
}
