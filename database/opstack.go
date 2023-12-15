package database

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-optimism-verifier/hublayer/contracts/l2oo"
	"github.com/oasysgames/oasys-optimism-verifier/util"
	"gorm.io/gorm"
)

type OPStackDatabase struct {
	db *gorm.DB
}

func (db *OPStackDatabase) FindL2OOs() ([]*OpstackL2OutputOracle, error) {
	var rows []*OpstackL2OutputOracle
	tx := db.db.Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

func (db *OPStackDatabase) FindOrCreateL2OO(l2oo_ common.Address) (*OpstackL2OutputOracle, error) {
	var row OpstackL2OutputOracle
	tx := db.db.Where("address = ?", l2oo_).First(&row)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		row.Address = l2oo_
		row.NextVerifyIndex = 0

		tx = db.db.Create(&row)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if tx.Error != nil {
		return nil, tx.Error
	}

	return &row, nil
}

func (db *OPStackDatabase) FindProposal(
	l2oo_ common.Address,
	l2OutputIndex uint64,
) (*OpstackProposal, error) {
	sub, err := db.l2ooIdSub(l2oo_)
	if err != nil {
		return nil, err
	}

	var row OpstackProposal
	tx := db.db.
		Joins("OpstackL2OutputOracle").
		Where("optimism_states.opstack_l2_output_oracle_id = (?)", sub).
		Where("optimism_states.l2_output_index = ?", l2OutputIndex).
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

// Return events waiting verification(order by BatchIndex).
func (db *OPStackDatabase) FindVerificationWaitingProposals(
	signer common.Address,
	l2ooAddr common.Address,
	nextVerifyIndex uint64,
	limit int,
) ([]*OpstackProposal, error) {
	signerSub, err := db.signerIdSub(signer)
	if err != nil {
		return nil, err
	}

	l2oo_, err := db.FindOrCreateL2OO(l2ooAddr)
	if err != nil {
		return nil, err
	}

	if l2oo_.NextVerifyIndex > nextVerifyIndex {
		nextVerifyIndex = l2oo_.NextVerifyIndex
	}

	sub := db.db.Model(&OpstackSignature{}).
		Select("l2_output_index").
		Where("opstack_l2_output_oracle_id = ? AND signer_id = (?)", l2oo_.ID, signerSub).
		Where("l2_output_index >= ?", nextVerifyIndex)
	if sub.Error != nil {
		return nil, sub.Error
	}

	var rows []*OpstackProposal
	tx := db.db.
		Joins("OpstackL2OutputOracle").
		Where("opstack_l2_output_oracle_id = ? AND l2_output_index >= ?", l2oo_.ID, nextVerifyIndex).
		Where("l2_output_index NOT IN (?)", sub).
		Order("l2_output_index ASC").
		Limit(limit).
		Find(&rows)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return rows, nil
}

func (db *OPStackDatabase) SaveNextVerifyIndex(l2ooAddr common.Address, nextVerifyIndex uint64) error {
	l2oo_, err := db.FindOrCreateL2OO(l2ooAddr)
	if err != nil {
		return err
	}

	l2oo_.NextVerifyIndex = nextVerifyIndex
	return db.db.Save(l2oo_).Error
}

// Save new state batch appended event to database.
func (db *OPStackDatabase) SaveProposal(e *l2oo.OasysL2OutputOracleOutputProposed) (*OpstackProposal, error) {
	row := &OpstackProposal{
		L2OutputIndex: e.L2OutputIndex.Uint64(),
		OutputRoot:    e.OutputRoot,
		L2BlockNumber: e.L2BlockNumber.Uint64(),
		L1Timestamp:   e.L1Timestamp.Uint64(),
	}

	err := db.db.Transaction(func(s *gorm.DB) error {
		l2oo_, err := newDB(s).OPStack.FindOrCreateL2OO(e.Raw.Address)
		if err != nil {
			return err
		}

		row.OpstackL2OutputOracle = *l2oo_
		return s.Create(row).Error
	})
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (db *OPStackDatabase) SaveSignature(
	id, previousID *string,
	signer common.Address,
	l2ooAddr common.Address,
	l2OutputIndex uint64,
	outputRoot common.Hash,
	l2BlockNumber uint64,
	l1Timestamp uint64,
	approved bool,
	signature Signature,
) (*OpstackSignature, error) {
	_signer, err := findOrCreateSigner(db.db, signer)
	if err != nil {
		return nil, err
	}

	l2oo_, err := db.FindOrCreateL2OO(l2ooAddr)
	if err != nil {
		return nil, err
	}

	values := map[string]interface{}{
		"signer_id":                   _signer.ID,
		"opstack_l2_output_oracle_id": l2oo_.ID,
		"l2_output_index":             l2OutputIndex,
		"output_root":                 outputRoot,
		"l2_block_number":             l2BlockNumber,
		"l1_timestamp":                l1Timestamp,
		"approved":                    approved,
		"signature":                   signature,
	}

	if previousID != nil {
		values["previous_id"] = *previousID
	} else {
		values["previous_id"] = gorm.Expr(`(SELECT IFNULL(
			(SELECT MAX(t.id) FROM opstack_signatures AS t WHERE t.signer_id = ?),
			""
		))`, _signer.ID)
	}

	var created OpstackSignature
	err = db.db.Transaction(func(s *gorm.DB) error {
		// Delete the same batch index signature as it may be recreated for reasons such as chain reorganization.
		if tx := s.Model(&OpstackSignature{}).
			Where("signer_id = ? AND opstack_l2_output_oracle_id = ?", _signer.ID, l2oo_.ID).
			// WARNING: Do not condition on signature comparison as this will result in a UNIQUE constraint error.
			Where("l2_output_index = ?", l2OutputIndex).
			Delete(&OpstackSignature{}); tx.Error != nil {
			return tx.Error
		}

		if id != nil {
			values["id"] = *id
		} else {
			values["id"] = util.ULID(nil).String()
		}

		if tx := s.Model(&OpstackSignature{}).Create(values); tx.Error != nil {
			return tx.Error
		}

		if tx := s.
			Joins("Signer").
			Joins("OpstackL2OutputOracle").
			First(&created, "opstack_signatures.id = ?", values["id"]); tx.Error != nil {
			return tx.Error
		}

		if strings.Compare(created.ID, created.PreviousID) <= 0 {
			return errors.New("previous id is overtaking")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (db *OPStackDatabase) FindLatestSignaturePerSigners() ([]*OpstackSignature, error) {
	var maxIds []string
	tx := db.db.Model(&OpstackSignature{}).
		Select("MAX(id)").
		Group("signer_id").
		Find(&maxIds)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var rows []*OpstackSignature
	tx = db.db.
		Joins("Signer").
		Joins("OpstackL2OutputOracle").
		Where("opstack_signatures.id IN (?)", maxIds).
		Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

func (db *OPStackDatabase) FindLatestSignaturesBySigner(
	signer common.Address,
	limit, offset int,
) ([]*OpstackSignature, error) {
	sub, err := db.signerIdSub(signer)
	if err != nil {
		return nil, err
	}

	var rows []*OpstackSignature
	tx := db.db.
		Joins("Signer").
		Joins("OpstackL2OutputOracle").
		Where("opstack_signatures.signer_id = (?)", sub).
		Order("opstack_signatures.id DESC").
		Limit(limit).
		Offset(offset).
		Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

func (db *OPStackDatabase) FindSignatureByID(id string) (*OpstackSignature, error) {
	var row OpstackSignature
	tx := db.db.
		Joins("Signer").
		Joins("OpstackL2OutputOracle").
		Where("opstack_signatures.id = ?", id).
		First(&row)

	if err := errconv(tx.Error); err != nil {
		return nil, err
	}
	return &row, nil
}

func (db *OPStackDatabase) FindSignatures(
	idAfter *string,
	signer *common.Address,
	l2oo_ *common.Address,
	l2OutputIndex *uint64,
	limit, offset int,
) ([]*OpstackSignature, error) {
	tx := db.db.
		Joins("Signer").
		Joins("OpstackL2OutputOracle").
		Order("opstack_signatures.id").
		Limit(limit).
		Offset(offset)

	if idAfter != nil {
		tx = tx.Where("opstack_signatures.id >= ?", *idAfter)
	}
	if signer != nil {
		if sub, err := db.signerIdSub(*signer); err != nil {
			return nil, err
		} else {
			tx = tx.Where("opstack_signatures.signer_id = (?)", sub)
		}
	}
	if l2oo_ != nil {
		if sub, err := db.l2ooIdSub(*l2oo_); err != nil {
			return nil, err
		} else {
			tx = tx.Where("opstack_signatures.opstack_l2_output_oracle_id = (?)", sub)
		}
	}
	if l2OutputIndex != nil {
		tx = tx.Where("opstack_signatures.l2_output_index = ?", *l2OutputIndex)
	}

	var rows []*OpstackSignature
	tx = tx.Find(&rows)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return rows, nil
}

// Delete states after the specified l2OutputIndex.
func (db *OPStackDatabase) DeleteProposals(l2oo_ common.Address, l2OutputIndex uint64) (int64, error) {
	var affected int64
	err := db.db.Transaction(func(s *gorm.DB) error {
		var ids []uint64
		tx := s.
			Model(&OpstackProposal{}).
			Joins("OpstackL2OutputOracle").
			Where("OpstackL2OutputOracle.address = ? AND l2_output_index >= ?", l2oo_, l2OutputIndex).
			Pluck("opstack_proposals.id", &ids)
		if tx.Error != nil {
			return tx.Error
		}

		tx = s.Where("id IN ?", ids).Delete(&OpstackProposal{})
		if tx.Error != nil {
			return tx.Error
		}

		affected = tx.RowsAffected
		return nil
	})
	if err != nil {
		return -1, err
	}

	return affected, nil
}

// Delete signatures after the specified l2OutputIndex.
func (db *OPStackDatabase) DeleteSignatures(
	signer common.Address,
	l2oo_ common.Address,
	l2OutputIndex uint64,
) (int64, error) {
	var affected int64
	err := db.db.Transaction(func(s *gorm.DB) error {
		var ids []string
		tx := s.
			Model(&OpstackSignature{}).
			Joins("Signer").
			Joins("OpstackL2OutputOracle").
			Where("Signer.address = ? AND OpstackL2OutputOracle.address = ?", signer, l2oo_).
			Where("opstack_signatures.l2_output_index >= ?", l2OutputIndex).
			Pluck("opstack_signatures.id", &ids)
		if tx.Error != nil {
			return tx.Error
		}

		tx = s.Where("id IN ?", ids).Delete(&OpstackSignature{})
		if tx.Error != nil {
			return tx.Error
		}

		affected = tx.RowsAffected
		return nil
	})
	if err != nil {
		return -1, err
	}

	return affected, nil
}

func (db *OPStackDatabase) signerIdSub(signer common.Address) (*gorm.DB, error) {
	sub := db.db.Model(&Signer{}).
		Select("id").
		Where("address = ?", signer)
	if sub.Error != nil {
		return nil, sub.Error
	}
	return sub, nil
}

func (db *OPStackDatabase) l2ooIdSub(l2oo common.Address) (*gorm.DB, error) {
	sub := db.db.Model(&OpstackL2OutputOracle{}).
		Select("id").
		Where("address = ?", l2oo)
	if sub.Error != nil {
		return nil, sub.Error
	}
	return sub, nil
}
