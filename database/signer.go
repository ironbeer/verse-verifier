package database

import (
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func findOrCreateSigner(db *gorm.DB, signer common.Address) (*Signer, error) {
	row := &Signer{Address: signer}
	tx := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "address"}},
	}).Create(row)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return row, nil
}
