package database

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SignerDatabase struct {
	db *gorm.DB
}

func (db *SignerDatabase) FindOrCreateSigner(signer common.Address) (*Signer, error) {
	row := &Signer{Address: signer}
	tx := db.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "address"}},
	}).Create(row)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return row, nil
}

func (db *SignerDatabase) signerIdSub(signer common.Address) (*gorm.DB, error) {
	sub := db.db.Model(&Signer{}).
		Select("id").
		Where("address = ?", signer)
	if sub.Error != nil {
		return nil, sub.Error
	}
	return sub, nil
}
