package repository

import (
	"context"
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
)

type PayoutRecordTransferredRepository interface {
	GetPayoutRecordTransferredDates(ctx context.Context) ([]model.PayoutRecordDate, error)
	GetPayoutRecordsWithMerchantBySendingDate(ctx context.Context, sendingDate time.Time) ([]*model.PayoutRecordWithMerchantID, error)
	GetPayoutRecordApplyNos(ctx context.Context) ([]model.PayoutRecordApplyNo, error)
	GetPayoutRecordsWithMerchantTransferredByApplyNo(ctx context.Context, applyNo string) ([]*model.PayoutRecordWithMerchantID, error)
	BulkUpdateTransferStatusByIDs(ctx context.Context, payoutRecords []*model.PayoutRecord) error
}
