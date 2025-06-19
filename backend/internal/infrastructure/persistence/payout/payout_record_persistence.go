package persistence

import (
	"context"
	"fmt"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/convert"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const BATCH_SIZE = 100

type PayoutRecordRepositoryImpl struct {
	db            *gorm.DB
	filterBuilder *persistence.GormFilterBuilder
}

func NewPayoutRecordRepository(db *gorm.DB) repository.PayoutRecordRepository {
	return &PayoutRecordRepositoryImpl{
		db:            db,
		filterBuilder: persistence.NewGormFilterBuilder(),
	}
}

// CountByPayoutID counts the number of payout records associated with a specific payout ID
func (r *PayoutRecordRepositoryImpl) CountByPayoutID(ctx context.Context, payoutID int) (int, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return 0, err
	}
	var count int64

	err = db.
		Table("payout_record").
		Where("payout_id = ?", payoutID).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// CountByPayoutIDs counts the number of payout records associated with multiple payout IDs
func (r *PayoutRecordRepositoryImpl) CountByPayoutIDs(ctx context.Context, payoutIDs []int) (map[int]int, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}
	counts := make(map[int]int)

	if len(payoutIDs) == 0 {
		return counts, nil
	}

	var results []struct {
		PayoutID int
		Count    int64
	}

	err = db.
		Table("payout_record").
		Select("payout_id, COUNT(*) AS count").
		Where("payout_id IN (?)", payoutIDs).
		Group("payout_id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		counts[result.PayoutID] = int(result.Count)
	}

	return counts, nil
}

// SumAmountByPayoutID calculates the total amount of payout records associated with a specific payout ID
func (r *PayoutRecordRepositoryImpl) SumAmountByPayoutID(ctx context.Context, payoutID int) (int64, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return 0, err
	}
	var totalAmount int64

	err = db.WithContext(ctx).
		Table("payout_record").
		Where("payout_id = ?", payoutID).
		Select("SUM(amount)").
		Scan(&totalAmount).Error

	if err != nil {
		return 0, err
	}

	return totalAmount, nil
}

// SumAmountByPayoutIDs calculates the total amount of payout records associated with multiple payout IDs
func (r *PayoutRecordRepositoryImpl) SumAmountByPayoutIDs(ctx context.Context, payoutIDs []int) (map[int]int64, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	totalAmounts := make(map[int]int64)

	if len(payoutIDs) == 0 {
		return totalAmounts, nil
	}

	var results []struct {
		PayoutID int
		Sum      int64
	}

	err = db.
		Table("payout_record").
		Select("payout_id, SUM(amount) AS sum").
		Where("payout_id IN (?)", payoutIDs).
		Group("payout_id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	for _, result := range results {
		totalAmounts[result.PayoutID] = result.Sum
	}

	return totalAmounts, nil
}

// UpdateTransferStatusByPayoutID updates the transfer status for all payout records of a given payout ID
func (r *PayoutRecordRepositoryImpl) UpdateTransferStatusByPayoutID(ctx context.Context, payoutID int, transferStatus object.PayoutRecordStatus) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	err = db.WithContext(ctx).
		Table("payout_record").
		Where("payout_id = ?", payoutID).
		Update("transfer_status", int(transferStatus)).Error

	if err != nil {
		return fmt.Errorf("failed to update transfer status for payout records: %w", err)
	}

	return nil
}

func (r *PayoutRecordRepositoryImpl) IsExistIdempotencyKey(ctx context.Context, idempotencyKey string) (bool, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return false, err
	}
	var count int64
	err = db.
		Model(&dto.PayoutRecord{}).
		Where("idempotency_key = ?", idempotencyKey).
		Limit(1).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
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
		Where("transfer_status = ?", object.PayoutRecordStatusWaitingTransfer)
	err = query.Count(&count).Error

	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, nil
	}

	err = query.Find(&payoutRecords).Error
	if err != nil {
		return nil, err
	}

	return convert.ToPayoutRecordModels(payoutRecords), nil
}

func (r *PayoutRecordRepositoryImpl) LockPayoutRecordsTransferingByIDs(ctx context.Context, ids []int) ([]*model.PayoutRecord, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	}

	var dtos []*dto.PayoutRecord
	err = db.Model(&dto.PayoutRecord{}).
		Where("id IN ?", ids).
		Where("transfer_status = ?", object.PayoutRecordStatusWaitingTransfer).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Find(&dtos).Error

	if err != nil {
		return nil, err
	}

	return convert.ToPayoutRecordModels(dtos), nil
}

func (r *PayoutRecordRepositoryImpl) BulkUpdateStatusTransferProcessings(ctx context.Context, payoutRecords []*model.PayoutRecord, chunkSize int) error {
	var updateColumns = []string{
		"transfer_status",
		"idempotency_key",
	}

	for _, c := range utils.Chunk(payoutRecords, chunkSize) {
		err := r.bulkInsertOnDuplicate(ctx, c, updateColumns)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *PayoutRecordRepositoryImpl) BulkUpdateTransferResult(ctx context.Context, payoutRecords []*model.PayoutRecord, chunkSize int) error {

	var updateColumns = []string{
		"transfer_status",
		"transfer_executed_at",
		"idempotency_key",
		"aozora_transfer_apply_no",
		"transfer_request_error",
	}

	for _, c := range utils.Chunk(payoutRecords, chunkSize) {
		err := r.bulkInsertOnDuplicate(ctx, c, updateColumns)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PayoutRecordRepositoryImpl) bulkInsertOnDuplicate(ctx context.Context, payoutRecords []*model.PayoutRecord, updateColumns []string) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	if len(payoutRecords) == 0 {
		return nil
	}

	dtoRecords := convert.ToPayoutRecordDTOs(payoutRecords)

	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(updateColumns),
	}).Create(&dtoRecords).Error
}

// BulkCreatePayoutRecords creates multiple payout records in a single operation
func (r *PayoutRecordRepositoryImpl) BulkCreatePayoutRecords(ctx context.Context, records []*model.PayoutRecord) error {
	if len(records) == 0 {
		return nil
	}

	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	recordDTOs := make([]dto.PayoutRecord, len(records))
	for i, record := range records {
		recordDTO := dto.FromPayoutRecordModel(record)
		if recordDTO == nil {
			return fmt.Errorf("failed to convert payout record at index %d", i)
		}
		recordDTOs[i] = *recordDTO
	}

	if err := db.WithContext(ctx).CreateInBatches(recordDTOs, BATCH_SIZE).Error; err != nil {
		return err
	}

	return nil
}
