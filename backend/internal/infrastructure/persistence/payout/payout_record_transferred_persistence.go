package persistence

import (
	"context"
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
	convert "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/convert"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PayoutRecordTransferredPersistenceImpl struct {
	db *gorm.DB
}

func NewPayoutRecordTransferredPersistence(db *gorm.DB) repository.PayoutRecordTransferredRepository {
	return &PayoutRecordTransferredPersistenceImpl{
		db: db,
	}
}

func (r *PayoutRecordTransferredPersistenceImpl) GetPayoutRecordTransferredDates(ctx context.Context) ([]model.PayoutRecordDate, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var results []*dto.PayoutRecordDate

	err = db.
		Table("payout_record pr").
		Select("pr.sending_date").
		Joins("JOIN payout p ON pr.payout_id = p.id").
		Where("pr.transfer_status = ?", object.PayoutRecordStatusTransferSuccess).
		Where("p.transfer_type = ?", object.PayoutTransferTypeNormalTransfer).
		Group("pr.sending_date").
		Order("pr.sending_date ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return convert.ToSendingDateModels(results), nil
}

func (r *PayoutRecordTransferredPersistenceImpl) GetPayoutRecordApplyNos(ctx context.Context) ([]model.PayoutRecordApplyNo, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var results []*dto.PayoutRecordApplyNo

	err = db.
		Table("payout_record pr").
		Select("pr.aozora_transfer_apply_no").
		Joins("JOIN payout p ON pr.payout_id = p.id").
		Where("pr.transfer_status = ?", object.PayoutRecordStatusTransferSuccess).
		Where("p.transfer_type = ?", object.PayoutTransferTypeBulkTransfer).
		Group("pr.aozora_transfer_apply_no").
		Order("pr.sending_date ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return convert.ToApplyNoModels(results), nil
}

func (r *PayoutRecordTransferredPersistenceImpl) GetPayoutRecordsWithMerchantTransferredByApplyNo(ctx context.Context, applyNo string) ([]*model.PayoutRecordWithMerchantID, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}
	var results []*dto.PayoutRecordWithMerchantID
	err = db.
		Table("payout_record pr").
		Select("DISTINCT pr.*, FIRST_VALUE(tr.merchant_id) OVER (PARTITION BY pr.transaction_id ORDER BY tr.id ASC) as merchant_id").
		Joins("JOIN payout p ON pr.payout_id = p.id").
		Joins("LEFT JOIN transaction_record tr ON pr.transaction_id = tr.transaction_id").
		Where("pr.aozora_transfer_apply_no = ?", applyNo).
		Where("pr.transfer_status = ?", object.PayoutRecordStatusTransferSuccess).
		Where("p.transfer_type = ?", object.PayoutTransferTypeBulkTransfer).
		Order("pr.id ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return convert.ToPayoutRecordWithMerchantIDModels(results), nil
}
func (r *PayoutRecordTransferredPersistenceImpl) GetPayoutRecordsWithMerchantBySendingDate(ctx context.Context, sendingDate time.Time) ([]*model.PayoutRecordWithMerchantID, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var results []*dto.PayoutRecordWithMerchantID
	err = db.
		Table("payout_record pr").
		Select("DISTINCT pr.*, FIRST_VALUE(tr.merchant_id) OVER (PARTITION BY pr.transaction_id ORDER BY tr.id ASC) as merchant_id").
		Joins("JOIN payout p ON pr.payout_id = p.id").
		Joins("LEFT JOIN transaction_record tr ON pr.transaction_id = tr.transaction_id").
		Where("pr.sending_date = ?", sendingDate).
		Where("pr.transfer_status = ?", object.PayoutRecordStatusTransferSuccess).
		Where("p.transfer_type = ?", object.PayoutTransferTypeNormalTransfer).
		Order("pr.id ASC").
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	return convert.ToPayoutRecordWithMerchantIDModels(results), nil
}

func (r *PayoutRecordTransferredPersistenceImpl) BulkUpdateTransferStatusByIDs(ctx context.Context, payoutRecords []*model.PayoutRecord) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	if len(payoutRecords) == 0 {
		return nil
	}

	var updateColumns = []string{
		"transfer_status",
		"aozora_transfer_status",
	}

	dtoRecords := convert.ToPayoutRecordDTOs(payoutRecords)

	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).Create(&dtoRecords).Error
}
