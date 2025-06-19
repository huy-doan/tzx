package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	dashboardModel "github.com/test-tzs/nomraeite/internal/domain/model/dashboard"
	model "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	transferTransactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transfer_transaction"
	transactionObject "github.com/test-tzs/nomraeite/internal/domain/object/transaction"
)

// TransactionRepository defines the interface for transaction data access
type TransactionRepository interface {
	GetTransactionDetails(ctx context.Context, transactionIDs []int) ([]*model.TransferTransactionDetail, error)
	ListTransferRequests(ctx context.Context, params *inputdata.TransferRequestListInput) (*model.PaginatedTransferRequest, error)
	ListTransferTransactions(ctx context.Context, params *inputdata.TransferTransactionInput) (*transferTransactionModel.PaginatedTransferTransaction, error)
	UpdateTransactionFields(ctx context.Context, transactionIDs []int, update *model.Transaction) error
	BulkCreate(ctx context.Context, transactions []*model.Transaction) ([]*model.Transaction, error)
	GetTransactionSummaryRecentMonth(ctx context.Context, recentMonthCount int) ([]*dashboardModel.TransactionSummaryCount, error)
	GetTransactionByID(ctx context.Context, id int) (*model.Transaction, error)
	UpdateStatus(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error)
	GetTransactionDetailsByPayoutID(ctx context.Context, payoutID int) ([]*model.TransferTransactionDetail, error)
	UpdateStatusByChunkIDs(ctx context.Context, ids []int, status transactionObject.TransactionStatus, chunkSize int) error
}
