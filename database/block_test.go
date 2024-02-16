package database

import (
	"errors"
	"sort"
	"testing"

	"gorm.io/gorm"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"github.com/oasysgames/oasys-optimism-verifier/testhelper"
	"github.com/stretchr/testify/suite"
)

func TestBlockDatabase(t *testing.T) {
	suite.Run(t, new(BlockDatabaseTestSuite))
}

type BlockDatabaseTestSuite struct {
	testhelper.Suite

	db      *BlockDatabase
	rawdb   *gorm.DB
	creates []*Block
}

func (s *BlockDatabaseTestSuite) SetupTest() {
	db, err := NewDatabase(&config.Database{Path: ":memory:"})
	if err != nil {
		panic(err)
	}
	s.db = db.Block
	s.rawdb = db.db
	s.creates = []*Block{}

	for _, number := range s.Shuffle(s.Range(0, 50)) {
		block := &Block{Number: uint64(number + 1), Hash: s.ItoHash(number + 1)}
		s.rawdb.Create(block)
		s.creates = append(s.creates, block)
	}
	sort.Slice(s.creates, func(i, j int) bool {
		return s.creates[i].Number < s.creates[j].Number
	})

	s.creates = append([]*Block{nil}, s.creates...) // padding
}

func (s *BlockDatabaseTestSuite) TestFind() {
	got, _ := s.db.Find(10)
	s.Equal(uint64(10), got.Number)
	s.Equal(s.ItoHash(10), got.Hash)
	s.Equal(false, got.LogCollected)
}

func (s *BlockDatabaseTestSuite) TestFindHighest() {
	got, _ := s.db.FindHighest()
	s.Equal(uint64(50), got.Number)
	s.Equal(s.ItoHash(50), got.Hash)
	s.Equal(false, got.LogCollected)
}

func (s *BlockDatabaseTestSuite) TestFindUncollecteds() {
	assertGots := func(gots []*Block, expNumbers []int) {
		s.Equal(len(expNumbers), len(gots))

		for _, expNumber := range expNumbers {
			got := gots[0]
			gots = gots[1:]

			s.Equal(uint64(expNumber), got.Number)
			s.Equal(s.ItoHash(expNumber), got.Hash)
			s.Equal(false, got.LogCollected)
		}
	}

	// limit = 10
	gots, _ := s.db.FindUncollecteds(10)
	assertGots(gots, s.Range(1, 10+1))

	// limit = 100
	gots, _ = s.db.FindUncollecteds(100)
	assertGots(gots, s.Range(1, 50+1))

	s.db.db.Transaction(func(txdb *gorm.DB) error {
		db := newDB(txdb)
		db.db.Model(&Block{}).
			Where("number <= 25").Update("log_collected", true)

		// limit = 10
		gots, _ = db.Block.FindUncollecteds(10)
		assertGots(gots, s.Range(26, 35+1))

		// limit = 100
		gots, _ = db.Block.FindUncollecteds(100)
		assertGots(gots, s.Range(26, 50+1))

		return errors.New("ROLLBACK")
	})

	s.db.db.Transaction(func(tx *gorm.DB) error {
		db := newDB(tx)
		db.Block.SaveCollected(25, s.creates[25].Hash)

		// limit = 10
		gots, _ = db.Block.FindUncollecteds(10)
		assertGots(gots, s.Range(26, 35+1))

		// limit = 100
		gots, _ = db.Block.FindUncollecteds(100)
		assertGots(gots, s.Range(26, 50+1))

		return errors.New("ROLLBACK")
	})
}

func (s *BlockDatabaseTestSuite) TestSave() {
	number := uint64(100)

	s.db.SaveNewBlock(number, s.ItoHash(int(number)))

	got, _ := s.db.Find(number)
	s.Equal(number, got.Number)
	s.Equal(s.ItoHash(int(number)), got.Hash)
	s.Equal(false, got.LogCollected)
}

func (s *BlockDatabaseTestSuite) TestSaveCollected() {
	s.NoError(s.db.SaveCollected(10, s.creates[10].Hash))
	collected, _ := findCollectedBlock(s.rawdb)
	s.Equal(uint64(10), collected)

	s.NoError(s.db.SaveCollected(20, s.creates[20].Hash))
	collected, _ = findCollectedBlock(s.rawdb)
	s.Equal(uint64(20), collected)

	err := s.db.SaveCollected(51, s.RandHash())
	s.ErrorContains(err, "failed to find the target block")

	err = s.db.SaveCollected(20, s.RandHash())
	s.ErrorContains(err, "this block was removed due to reorganization")
}

func (s *BlockDatabaseTestSuite) TestDeletes() {
	s.NoError(s.db.Deletes(26))

	_, err := s.db.Find(25)
	s.NoError(err)

	for n := uint64(26); n <= 50; n++ {
		_, err := s.db.Find(n)
		s.ErrorIs(err, ErrNotFound)
	}

	_, err = findCollectedBlock(s.rawdb)
	s.ErrorIs(err, gorm.ErrRecordNotFound)

	s.NoError(saveCollectedBlock(s.rawdb, 5))
	s.NoError(s.db.Deletes(uint64(7)))

	collected, _ := findCollectedBlock(s.rawdb)
	s.Equal(uint64(5), collected)

	s.NoError(s.db.Deletes(uint64(2)))

	collected, _ = findCollectedBlock(s.rawdb)
	s.Equal(uint64(1), collected)
}
