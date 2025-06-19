package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"

	model "github.com/test-tzs/nomraeite/internal/domain/model/paypay"
)

// PaypayPayinSummaryRepository handles database operations for PayPay payin summary records
type PaypayPayinSummaryRepository interface {
	BulkInsert(ctx context.Context, summaries []*model.PaypayPayinSummary) error
	DeleteByPayinFileID(ctx context.Context, payinFileID int) error
	ListPaypayPayinSummary(ctx context.Context, params *inputdata.PaypayPayinSummaryInputData) (*model.PaginatedPaypayPayinSummaryResult, error)
}
