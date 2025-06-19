package repository

import (
	"context"

	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
)

type TransactionRecordRepository interface {
	BulkCreate(ctx context.Context, transactionRecords []*transactionModel.TransactionRecord) error
}
