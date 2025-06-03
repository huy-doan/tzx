package persistence

import (
	"context"
	"errors"

	repository "github.com/makeshop-jp/master-console/batch/domain/repository/payout"
	model "github.com/makeshop-jp/master-console/internal/domain/model/payout_record"
	object "github.com/makeshop-jp/master-console/internal/domain/object/payout"
	convert "github.com/makeshop-jp/master-console/internal/infrastructure/persistence/payout_record/convert"
	dto "github.com/makeshop-jp/master-console/internal/infrastructure/persistence/payout_record/dto"
	"github.com/makeshop-jp/master-console/internal/pkg/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PayoutRecordRepositoryImpl struct {
	db *gorm.DB
}

func NewPayoutRecordPersistence(db *gorm.DB) repository.PayoutRecordRepository {
	return &PayoutRecordRepositoryImpl{
		db: db,
	}
}

func (r *PayoutRecordRepositoryImpl) GetListTransferingByPayoutID(ctx context.Context, payoutID int) ([]*model.PayoutRecord, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var payoutRecords []*dto.PayoutRecord
	var count int64
	query := db.
		Model(&dto.PayoutRecord{}).
		Where("payout_id = ?", payoutID).
		Where("transfer_status = ?", object.TransferStatusWaitingTransfer).
		Where("transfer_requested_at IS NOT NULL").
		Where("transfer_executed_at IS NULL")
	err = query.Find(&payoutRecords).Count(&count).Error

	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}

	return convert.ToPayoutRecordModels(payoutRecords), nil
}

func (r *PayoutRecordRepositoryImpl) LockPayoutRecordByID(ctx context.Context, id int) (*model.PayoutRecord, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var dto dto.PayoutRecord
	err = db.Model(&dto).
		Where("id = ?", id).
		Where("transfer_status = ?", object.TransferStatusWaitingTransfer).
		Where("transfer_requested_at IS NOT NULL").
		Where("transfer_executed_at IS NULL").
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&dto).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return convert.ToPayoutRecordModel(&dto), nil
}

// UpdateByID updates a payout record by its ID
func (r *PayoutRecordRepositoryImpl) UpdateByID(ctx context.Context, id int, payoutRecord *model.PayoutRecord) (*model.PayoutRecord, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	dtoPayoutRecord := convert.ToPayoutRecordDTO(payoutRecord)
	err = db.Model(&dto.PayoutRecord{}).
		Where("id = ?", id).
		Updates(dtoPayoutRecord).Error

	if err != nil {
		return nil, err
	}

	return convert.ToPayoutRecordModel(dtoPayoutRecord), nil
}
