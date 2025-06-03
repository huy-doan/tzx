package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
)

type TransactionRepository interface {
	GetTransactionByID(ctx context.Context, id int) (*model.Transaction, error)
	UpdateStatus(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
}
