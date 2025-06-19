package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
)

type PayoutRecordRepository interface {
	CountByPayoutID(ctx context.Context, payoutID int) (int, error)
	CountByPayoutIDs(ctx context.Context, payoutIDs []int) (map[int]int, error)
	SumAmountByPayoutID(ctx context.Context, payoutID int) (int64, error)
	SumAmountByPayoutIDs(ctx context.Context, payoutIDs []int) (map[int]int64, error)

	IsExistIdempotencyKey(ctx context.Context, idempotencyKey string) (bool, error)
	GetListTransferingByPayoutID(ctx context.Context, payoutID int) ([]*model.PayoutRecord, error)
	LockPayoutRecordsTransferingByIDs(ctx context.Context, ids []int) ([]*model.PayoutRecord, error)
	BulkUpdateStatusTransferProcessings(ctx context.Context, payoutRecords []*model.PayoutRecord, chunkSize int) error
	BulkUpdateTransferResult(ctx context.Context, payoutRecords []*model.PayoutRecord, chunkSize int) error
	BulkCreatePayoutRecords(ctx context.Context, records []*model.PayoutRecord) error
	UpdateTransferStatusByPayoutID(ctx context.Context, payoutID int, transferStatus object.PayoutRecordStatus) error
}
