package database

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSignerDatabase(t *testing.T) {
	suite.Run(t, new(SignerDatabaseTestSuite))
}

type SignerDatabaseTestSuite struct {
	DatabaseTestSuite

	db *SignerDatabase
}

func (s *SignerDatabaseTestSuite) SetupTest() {
	s.DatabaseTestSuite.SetupTest()
	s.db = s.DatabaseTestSuite.db.Signer
}

func (s *SignerDatabaseTestSuite) TestFindOrCreateSigner() {
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
