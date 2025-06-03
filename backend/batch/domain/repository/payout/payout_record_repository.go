package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout_record"
)

type PayoutRecordRepository interface {
	GetListTransferingByPayoutID(ctx context.Context, payoutID int) ([]*model.PayoutRecord, error)
	LockPayoutRecordByID(ctx context.Context, id int) (*model.PayoutRecord, error)
	UpdateByID(ctx context.Context, id int, payoutRecord *model.PayoutRecord) (*model.PayoutRecord, error)
}
