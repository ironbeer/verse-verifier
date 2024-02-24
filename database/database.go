package database

import (
	"errors"

	"github.com/oasysgames/oasys-optimism-verifier/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

var (
	ErrNotFound = errors.New("not found")

	models = []interface{}{
		&Block{},
		&Signer{},
		&OptimismScc{},
		&OptimismState{},
		&OptimismSignature{},
		&OpstackL2OutputOracle{},
		&OpstackProposal{},
		&OpstackSignature{},
		&Misc{},
	}
)

type Database struct {
	rawdb *gorm.DB

	Signer   *SignerDatabase
	Block    *BlockDatabase
	Optimism *OptimismDatabase
	OPStack  *OPStackDatabase
}

func NewDatabase(cfg *config.Database) (*Database, error) {
	config := &gorm.Config{Logger: &mylogger{
		LogLevel:            gormlog.Info,
		LongQueryTime:       cfg.LongQueryTime,
		MinExaminedRowLimit: cfg.MinExaminedRowLimit,
	}}
	db, err := gorm.Open(sqlite.Open(cfg.Path), config)
	if err != nil {
		return nil, err
	}

	// workaround for "database is locked" error
	if rawdb, err := db.DB(); err != nil {
		return nil, err
	} else {
		rawdb.SetMaxOpenConns(1)
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return nil, err
		}
	}

	return newDB(db), nil
}

func (db *Database) Transaction(fn func(*Database) error) error {
	return db.rawdb.Transaction(func(tx *gorm.DB) error {
		return fn(newDB(tx))
	})
}

func newDB(rawdb *gorm.DB) *Database {
	var db Database
	db = Database{
		rawdb:    rawdb,
		Signer:   &SignerDatabase{rawdb: rawdb, db: &db},
		Block:    &BlockDatabase{rawdb: rawdb, db: &db},
		Optimism: &OptimismDatabase{rawdb: rawdb, db: &db},
		OPStack:  &OPStackDatabase{rawdb: rawdb, db: &db},
	}
	return &db
}

func errconv(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return err
}
