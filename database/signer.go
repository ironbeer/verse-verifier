package database

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

type SignerDatabase struct {
	rawdb *gorm.DB
	db    *Database
}

func (db *SignerDatabase) FindOrCreateSigner(signer common.Address) (row *Signer, err error) {
	err = db.rawdb.Transaction(func(txdb *gorm.DB) error {
		tx := txdb.Where("address = ?", signer).First(&row)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			row.Address = signer
			return txdb.Create(&row).Error
		}
		return tx.Error
	})
	if err != nil {
		return nil, err
	}
	return row, nil
}
