package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"

	model "github.com/test-tzs/nomraeite/internal/domain/model/paypay"
)

type PaypayPayinTransactionRepository interface {
	BulkInsert(ctx context.Context, payinFileID int, transactions []*model.PaypayPayinTransaction) error
	DeleteByPayinFileID(ctx context.Context, payinFileID int) error
	ListPaypayPayinTransactions(ctx context.Context, params *inputdata.PaypayPayinTransactionListInputData) (*model.PaginatedPaypayPayinTransactionResult, error)
}
