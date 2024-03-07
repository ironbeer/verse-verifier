package database

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/oasysgames/oasys-optimism-verifier/contract/scc"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/suite"
)

func TestOptimismDatabase(t *testing.T) {
	suite.Run(t, new(OptimismDatabaseTestSuite))
}

type OptimismDatabaseTestSuite struct {
	DatabaseTestSuite

	db *OptimismDatabase
}

func (s *OptimismDatabaseTestSuite) SetupTest() {
	s.DatabaseTestSuite.SetupTest()
	s.db = s.DatabaseTestSuite.db.Optimism
}

func (s *OptimismDatabaseTestSuite) TestFindOrCreateSigner() {
	assert := func(got1, got2, got3 *Signer) {
		var count int
		s.rawdb.Table("signers").Select("COUNT(*)").Row().Scan(&count)
		s.Equal(3, count)

		s.Equal(uint64(1), got1.ID)
		s.Equal(uint64(2), got2.ID)
		s.Equal(uint64(3), got3.ID)
		s.Equal(s.ItoAddress(1), got1.Address)
		s.Equal(s.ItoAddress(2), got2.Address)
		s.Equal(s.ItoAddress(3), got3.Address)
	}

	addr1 := s.ItoAddress(1)
	addr2 := s.ItoAddress(2)
	addr3 := s.ItoAddress(3)

	got1, _ := s.db.FindOrCreateSigner(addr1)
	got2, _ := s.db.FindOrCreateSigner(addr2)
	got3, _ := s.db.FindOrCreateSigner(addr3)
	assert(got1, got2, got3)

	got1, _ = s.db.FindOrCreateSigner(addr1)
	got2, _ = s.db.FindOrCreateSigner(addr2)
	got3, _ = s.db.FindOrCreateSigner(addr3)
	assert(got1, got2, got3)
}

func (s *OptimismDatabaseTestSuite) TestFindSCCs() {
	var wants []*OptimismScc
	for range s.Range(0, 10) {
		want, _ := s.db.FindOrCreateSCC(s.RandAddress())
		wants = append(wants, want)
	}

	gots, _ := s.db.FindSCCs()
	s.Len(gots, len(wants))
	for i, want := range wants {
		s.Equal(want.Address, gots[i].Address)
	}
}

func (s *OptimismDatabaseTestSuite) TestFindOrCreateSCC() {
	assert := func(got0, got1, got2 *OptimismScc) {
		var count int
		s.rawdb.Table("optimism_sccs").Select("COUNT(*)").Row().Scan(&count)
		s.Equal(3, count)

		s.Equal(uint64(1), got0.ID)
		s.Equal(uint64(2), got1.ID)
		s.Equal(uint64(3), got2.ID)
		s.Equal(s.ItoAddress(1), got0.Address)
		s.Equal(s.ItoAddress(2), got1.Address)
		s.Equal(s.ItoAddress(3), got2.Address)
	}

	addr0 := s.ItoAddress(1)
	addr1 := s.ItoAddress(2)
	addr2 := s.ItoAddress(3)

	got0, _ := s.db.FindOrCreateSCC(addr0)
	got1, _ := s.db.FindOrCreateSCC(addr1)
	got2, _ := s.db.FindOrCreateSCC(addr2)
	assert(got0, got1, got2)

	got0, _ = s.db.FindOrCreateSCC(addr0)
	got1, _ = s.db.FindOrCreateSCC(addr1)
	got2, _ = s.db.FindOrCreateSCC(addr2)
	assert(got0, got1, got2)

	// test upsert
	s.db.SaveNextIndex(addr0, 10)
	got0, _ = s.db.FindOrCreateSCC(addr0)
	s.Equal(uint64(10), got0.NextIndex)
}

func (s *OptimismDatabaseTestSuite) TestFindState() {
	scc0 := s.createSCC()
	scc1 := s.createSCC()

	st0 := s.createState(scc0, 0)
	st1 := s.createState(scc1, 0)
	st2 := s.createState(scc1, 1)

	got0, _ := s.db.FindState(scc0.Address, 0)
	got1, _ := s.db.FindState(scc1.Address, 0)
	got2, _ := s.db.FindState(scc1.Address, 1)

	s.Equal(st0, got0)
	s.Equal(st1, got1)
	s.Equal(st2, got2)

	_, err := s.db.FindState(scc0.Address, 1)
	s.Error(err, ErrNotFound)

	_, err = s.db.FindState(scc1.Address, 2)
	s.Error(err, ErrNotFound)
}

func (s *OptimismDatabaseTestSuite) TestSaveNextIndex() {
	scc0 := s.createSCC()
	scc1 := s.createSCC()

	s.Equal(uint64(0), scc0.NextIndex)
	s.Equal(uint64(0), scc1.NextIndex)

	s.db.SaveNextIndex(scc0.Address, 5)
	s.db.SaveNextIndex(scc1.Address, 10)

	scc0, _ = s.db.FindOrCreateSCC(scc0.Address)
	scc1, _ = s.db.FindOrCreateSCC(scc1.Address)

	s.Equal(uint64(5), scc0.NextIndex)
	s.Equal(uint64(10), scc1.NextIndex)
}

func (s *OptimismDatabaseTestSuite) TestSaveState() {
	scc_ := s.createSCC()
	batchIndex := uint64(1)

	_, err := s.db.FindState(scc_.Address, batchIndex)
	s.Error(err, ErrNotFound)

	ev := &scc.SccStateBatchAppended{
		BatchIndex:        big.NewInt(int64(batchIndex)),
		BatchRoot:         s.RandHash(),
		BatchSize:         big.NewInt(12),
		PrevTotalElements: big.NewInt(3),
		ExtraData:         []byte("test"),
		Raw: types.Log{
			Address: scc_.Address,
		},
	}

	got, _ := s.db.SaveState(ev)
	s.Equal(uint64(1), got.ID)
	s.Equal(*scc_, got.OptimismScc)
	s.Equal(ev.BatchIndex.Uint64(), got.BatchIndex)
	s.Equal(ev.BatchRoot[:], got.BatchRoot[:])
	s.Equal(ev.BatchSize.Uint64(), got.BatchSize)
	s.Equal(ev.PrevTotalElements.Uint64(), got.PrevTotalElements)
	s.Equal(ev.ExtraData, got.ExtraData)

	found, _ := s.db.FindState(scc_.Address, batchIndex)
	s.Equal(got, found)
}

func (s OptimismDatabaseTestSuite) TestSaveSignature() {
	signer := s.createSigner()
	scc := s.createSCC()
	batchIndex := uint64(1)
	batchRoot := s.RandHash()
	batchSize := 10
	prevTotalElements := 20
	extraData := []byte("test")
	approved := true
	signature := RandSignature()

	assert := func(got *OptimismSignature) {
		ulid.MustParse(got.ID)
		s.Equal(*signer, got.Signer)
		s.Equal(*scc, got.OptimismScc)
		s.Equal(uint64(1), got.BatchIndex)
		s.Equal(batchRoot, got.BatchRoot)
		s.Equal(uint64(batchSize), got.BatchSize)
		s.Equal(uint64(prevTotalElements), got.PrevTotalElements)
		s.Equal(extraData, got.ExtraData)
		s.Equal(approved, got.Approved)
		s.Equal(signature, got.Signature)
	}

	got0, _ := s.db.SaveSignature(nil, nil,
		signer.Address, scc.Address, batchIndex, batchRoot,
		uint64(batchSize), uint64(prevTotalElements), extraData, approved, signature)
	assert(got0)
	s.Equal("", got0.PreviousID)

	// overwrite test
	batchRoot = s.RandHash()
	batchSize += 1
	prevTotalElements += 1
	extraData = []byte("TEST")
	approved = false
	signature = RandSignature()

	got0, _ = s.db.SaveSignature(nil, nil,
		signer.Address, scc.Address, batchIndex, batchRoot,
		uint64(batchSize), uint64(prevTotalElements), extraData, approved, signature)
	assert(got0)
	s.Equal("", got0.PreviousID)

	// add new
	got1, _ := s.db.SaveSignature(nil, nil,
		signer.Address, scc.Address, batchIndex+1, batchRoot,
		uint64(batchSize), uint64(prevTotalElements), extraData, approved, signature)
	s.Equal(got0.ID, got1.PreviousID)

	// other signer
	got2, _ := s.db.SaveSignature(nil, nil,
		s.RandAddress(), scc.Address, batchIndex+2, batchRoot,
		uint64(batchSize), uint64(prevTotalElements), extraData, approved, signature)
	s.Equal("", got2.PreviousID)

	// overtaking test
	id1, id2 := util.ULID(nil).String(), util.ULID(nil).String()
	_, err := s.db.SaveSignature(&id1, &id2,
		s.RandAddress(), scc.Address, batchIndex+3, batchRoot,
		uint64(batchSize), uint64(prevTotalElements), extraData, approved, signature)
	s.ErrorContains(err, "previous id is overtaking")
}

func (s OptimismDatabaseTestSuite) TestFindSignatureByID() {
	signer0 := s.createSigner()
	signer1 := s.createSigner()

	scc0 := s.createSCC()
	scc1 := s.createSCC()

	sig0 := s.createSignature(signer0, scc0, 0)
	sig1 := s.createSignature(signer1, scc1, 0)
	sig2 := s.createSignature(signer1, scc1, 1)

	got0, _ := s.db.FindSignatureByID(sig0.ID)
	got1, _ := s.db.FindSignatureByID(sig1.ID)
	got2, _ := s.db.FindSignatureByID(sig2.ID)

	s.Equal(sig0, got0)
	s.Equal(sig1, got1)
	s.Equal(sig2, got2)

	_, err := s.db.FindSignatureByID("")
	s.Error(err, ErrNotFound)
}

func (s OptimismDatabaseTestSuite) TestFindLatestSignaturesBySigner() {
	signer0 := s.createSigner()
	signer1 := s.createSigner()
	scc := s.createSCC()

	var wants0, wants1 []*OptimismSignature
	for _, index := range s.Range(0, 5) {
		wants0 = append(wants0, s.createSignature(signer0, scc, index))
	}
	for _, index := range s.Range(0, 10) {
		wants1 = append(wants1, s.createSignature(signer1, scc, index))
	}

	gots0, _ := s.db.FindLatestSignaturesBySigner(signer0.Address, 2, 0)
	s.Len(gots0, 2)
	s.Equal(wants0[len(wants0)-1].ID, gots0[0].ID)
	s.Equal(wants0[len(wants0)-2].ID, gots0[1].ID)

	gots1, _ := s.db.FindLatestSignaturesBySigner(signer1.Address, 2, 0)
	s.Len(gots1, 2)
	s.Equal(wants1[len(wants1)-1].ID, gots1[0].ID)
	s.Equal(wants1[len(wants1)-2].ID, gots1[1].ID)

	gots0, _ = s.db.FindLatestSignaturesBySigner(signer0.Address, 10, 2)
	s.Len(gots0, 3)
	s.Equal(wants0[len(wants0)-3].ID, gots0[0].ID)
	s.Equal(wants0[len(wants0)-4].ID, gots0[1].ID)
	s.Equal(wants0[len(wants0)-5].ID, gots0[2].ID)

	gots1, _ = s.db.FindLatestSignaturesBySigner(signer1.Address, 10, 2)
	s.Len(gots1, 8)
	s.Equal(wants1[len(wants1)-3].ID, gots1[0].ID)
	s.Equal(wants1[len(wants1)-4].ID, gots1[1].ID)
	s.Equal(wants1[len(wants1)-5].ID, gots1[2].ID)
	s.Equal(wants1[len(wants1)-6].ID, gots1[3].ID)
	s.Equal(wants1[len(wants1)-7].ID, gots1[4].ID)
	s.Equal(wants1[len(wants1)-8].ID, gots1[5].ID)
	s.Equal(wants1[len(wants1)-9].ID, gots1[6].ID)
	s.Equal(wants1[len(wants1)-10].ID, gots1[7].ID)
}

func (s OptimismDatabaseTestSuite) TestFindLatestSignaturePerSigners() {
	signer0 := s.createSigner()
	signer1 := s.createSigner()
	signer2 := s.createSigner()
	sccs := []*OptimismScc{s.createSCC(), s.createSCC(), s.createSCC()}

	var want0, want1, want2 *OptimismSignature
	for i := range s.Shuffle(s.Range(0, len(sccs))) {
		for _, index := range s.Range(0, 5) {
			want0 = s.createSignature(signer0, sccs[i], index)
		}
		for _, index := range s.Range(0, 10) {
			want1 = s.createSignature(signer1, sccs[i], index)
		}
		for _, index := range s.Range(0, 15) {
			want2 = s.createSignature(signer2, sccs[i], index)
		}
	}

	gots, _ := s.db.FindLatestSignaturePerSigners()
	s.Len(gots, 3)
	s.Equal(want0, gots[0])
	s.Equal(want1, gots[1])
	s.Equal(want2, gots[2])
}

func (s *OptimismDatabaseTestSuite) TestFindSignatures() {
	concat := func(slices ...[]*OptimismSignature) (cpy []*OptimismSignature) {
		for _, s := range slices {
			cpy = append(cpy, s...)
		}
		return cpy
	}
	assert := func(gots, wants []*OptimismSignature) {
		s.Len(gots, len(wants))
		for i, want := range wants {
			s.Equal(want.ID, gots[i].ID, i)
			s.Equal(want.SignerID, gots[i].SignerID, i)
			s.Equal(want.OptimismSccID, gots[i].OptimismSccID, i)
		}
	}

	signer0 := s.createSigner()
	signer1 := s.createSigner()
	scc0 := s.createSCC()
	scc1 := s.createSCC()

	var creates0, creates1, creates2, creates3 []*OptimismSignature
	for _, index := range s.Range(0, 5) {
		creates0 = append(creates0, s.createSignature(signer0, scc0, index))
	}
	for _, index := range s.Range(5, 10) {
		creates1 = append(creates1, s.createSignature(signer1, scc0, index))
	}
	for _, index := range s.Range(10, 15) {
		creates2 = append(creates2, s.createSignature(signer0, scc1, index))
	}
	for _, index := range s.Range(15, 20) {
		creates3 = append(creates3, s.createSignature(signer1, scc1, index))
	}

	index := uint64(5)
	cases := []struct {
		name          string
		idAfter       *string
		signer        *common.Address
		scc           *common.Address
		index         *uint64
		limit, offset int
		wants         []*OptimismSignature
	}{
		{
			"all",
			nil, nil, nil, nil, 100, 0,
			concat(creates0, creates1, creates2, creates3),
		},
		{
			"idAfter",
			&creates2[1].ID, nil, nil, nil, 100, 0,
			concat(creates2[1:], creates3),
		},
		{
			"signer0",
			nil, &signer0.Address, nil, nil, 100, 0,
			concat(creates0, creates2),
		},
		{
			"scc0",
			nil, nil, &scc0.Address, nil, 100, 0,
			concat(creates0, creates1),
		},
		{
			"index=5",
			nil, nil, nil, &index, 100, 0,
			creates1[0:1],
		},
		{
			"limit",
			nil, nil, nil, nil, 8, 0,
			concat(creates0, creates1[:3]),
		},
		{
			"offset",
			nil, nil, nil, nil, 100, 13,
			concat(creates2[3:], creates3),
		},
		{
			"over offset",
			nil, nil, nil, nil, 100, 100,
			[]*OptimismSignature{},
		},
	}
	for _, c := range cases {
		s.Run(c.name, func() {
			gots, _ := s.db.FindSignatures(
				c.idAfter, c.signer, c.scc, c.index, c.limit, c.offset)
			assert(gots, c.wants)
		})
	}
}

func (s *OptimismDatabaseTestSuite) TestFindVerificationWaitingStates() {
	assert := func(gots []*OptimismState, wants []*OptimismState) {
		s.Len(gots, len(wants))
		for i, got := range gots {
			s.Equal(wants[i].ID, got.ID)
		}
	}

	// Create dummy records
	count := 10
	signer := s.createSigner()
	scc0, scc1 := s.createSCC(), s.createSCC()
	wants0, wants1 := make([]*OptimismState, count), make([]*OptimismState, count)
	for _, batchIndex := range s.Shuffle(s.Range(0, count)) {
		wants0[batchIndex] = s.createState(scc0, batchIndex)
		wants1[batchIndex] = s.createState(scc1, batchIndex)
	}

	gots, _ := s.db.FindVerificationWaitingStates(signer.Address, scc0.Address, 0, 100)
	assert(gots, wants0)

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc1.Address, 0, 100)
	assert(gots, wants1)

	// when limit is set
	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc0.Address, 0, 2)
	assert(gots, wants0[:2])

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc1.Address, 0, 4)
	assert(gots, wants1[:4])

	// when `nextIndex` is set to query
	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc0.Address, 2, 100)
	assert(gots, wants0[2:])

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc1.Address, 3, 100)
	assert(gots, wants1[3:])

	// when `nextIndex` is set to scc
	s.db.SaveNextIndex(scc0.Address, 3)
	s.db.SaveNextIndex(scc1.Address, 4)

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc0.Address, 0, 100)
	assert(gots, wants0[3:])

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc1.Address, 0, 100)
	assert(gots, wants1[4:])

	// when several signatures exist
	merge := func(a, b []*OptimismState) (m []*OptimismState) {
		m = append(m, a...)
		m = append(m, b...)
		return m
	}
	s.createSignature(signer, scc0, 6)
	s.createSignature(signer, scc1, 8)

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc0.Address, 0, 100)
	assert(gots, merge(wants0[3:6], wants0[7:]))

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc1.Address, 0, 100)
	assert(gots, merge(wants1[4:8], wants1[9:]))

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc0.Address, 8, 100)
	assert(gots, wants0[8:])

	gots, _ = s.db.FindVerificationWaitingStates(signer.Address, scc1.Address, 9, 100)
	assert(gots, wants1[9:])
}

func (s *OptimismDatabaseTestSuite) TestDeleteStates() {
	assert := func(scc *OptimismScc, want []int) {
		var gots []int
		s.rawdb.Model(&OptimismState{}).
			Where("optimism_scc_id = ?", scc.ID).
			Order("batch_index").
			Pluck("batch_index", &gots)
		s.Equal(want, gots)
	}

	scc0 := s.createSCC()
	scc1 := s.createSCC()
	for _, i := range s.Shuffle(s.Range(0, 10)) {
		s.createState(scc0, i)
		s.createState(scc1, i)
	}

	assert(scc0, s.Range(0, 10))
	assert(scc1, s.Range(0, 10))

	rows0, _ := s.db.DeleteStates(scc0.Address, 3)
	rows1, _ := s.db.DeleteStates(scc1.Address, 6)

	s.Equal(int64(7), rows0)
	s.Equal(int64(4), rows1)
	assert(scc0, s.Range(0, 3))
	assert(scc1, s.Range(0, 6))
}

func (s *OptimismDatabaseTestSuite) TestDeleteSignatures() {
	assert := func(signer *Signer, scc *OptimismScc, want []int) {
		var gots []int
		s.rawdb.Model(&OptimismSignature{}).
			Where("signer_id = ? AND optimism_scc_id = ?", signer.ID, scc.ID).
			Order("batch_index").
			Pluck("batch_index", &gots)
		s.Equal(want, gots)
	}

	signer0 := s.createSigner()
	signer1 := s.createSigner()
	scc0 := s.createSCC()
	scc1 := s.createSCC()
	for _, i := range s.Shuffle(s.Range(0, 10)) {
		s.createSignature(signer0, scc0, i)
		s.createSignature(signer1, scc1, i)
	}

	assert(signer0, scc0, s.Range(0, 10))
	assert(signer1, scc1, s.Range(0, 10))

	rows0, _ := s.db.DeleteSignatures(signer0.Address, scc0.Address, 3)
	rows1, _ := s.db.DeleteSignatures(signer1.Address, scc1.Address, 6)

	s.Equal(int64(7), rows0)
	s.Equal(int64(4), rows1)
	assert(signer0, scc0, s.Range(0, 3))
	assert(signer1, scc1, s.Range(0, 6))
}

func (s *OptimismDatabaseTestSuite) TestSequentialSignaturesFinder() {
	signer := s.createSigner()
	scc := s.createSCC()

	var sigtree [5][]*OptimismSignature
	for _, depth := range s.Range(0, len(sigtree)) {
		sigtree[depth] = []*OptimismSignature{}

		var prevID string
		for i := 0; i <= depth; i++ {
			if depth > 0 && i < len(sigtree[depth-1]) {
				prevID = sigtree[depth-1][i].ID
			}

			sig := s.createSignature(signer, scc, (depth*10)+i)
			sig.PreviousID = prevID
			s.NoDBError(s.db.db.Save(sig))

			sigtree[depth] = append(sigtree[depth], sig)

			// one signature at last depth
			if depth == len(sigtree)-1 {
				break
			}
		}
	}

	assert := func(gots, wants []*OptimismSignature) {
		s.Len(gots, len(wants))
		for i, want := range wants {
			s.Equal(want.ID, gots[i].ID, i)
			s.Equal(want.PreviousID, gots[i].PreviousID, i)
		}
	}

	finder := s.db.SequentialSignaturesFinder("")
	gots0, _ := finder()
	gots1, _ := finder()
	gots2, _ := finder()
	gots3, _ := finder()
	gots4, _ := finder()
	gots5, _ := finder()

	assert(gots0, sigtree[0])
	assert(gots1, sigtree[1])
	assert(gots2, sigtree[2])
	assert(gots3, sigtree[3])
	assert(gots4, sigtree[4])
	assert(gots5, []*OptimismSignature{})
}
