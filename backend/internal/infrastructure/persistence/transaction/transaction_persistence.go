package persistence

import (
	"context"
	"errors"
	"math"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	dashboardModel "github.com/test-tzs/nomraeite/internal/domain/model/dashboard"
	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	transferTransactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transfer_transaction"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/dashboard"
	transactionObject "github.com/test-tzs/nomraeite/internal/domain/object/transaction"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/transaction"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/transaction/dto"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
	"gorm.io/gorm"
)

// TransactionRepositoryImpl implements the TransactionRepository interface
type TransactionRepositoryImpl struct {
	db           *gorm.DB
	queryBuilder *QueryBuilder
}

// NewTransactionRepository creates a new repository implementation
func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &TransactionRepositoryImpl{
		db:           db,
		queryBuilder: NewQueryBuilder(db),
	}
}

// GetTransactionDetails retrieves transaction details by IDs
func (r *TransactionRepositoryImpl) GetTransactionDetails(ctx context.Context, transactionIDs []int) ([]*transactionModel.TransferTransactionDetail, error) {
	if len(transactionIDs) == 0 {
		return []*transactionModel.TransferTransactionDetail{}, nil
	}

	var transferTransactionDTOs []dto.TransferTransactionDetail
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	queryBuilder := NewQueryBuilder(db)
	query := queryBuilder.BuildTransactionDetailQuery(transactionIDs)
	result := query.Find(&transferTransactionDTOs)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(transferTransactionDTOs) == 0 {
		return []*transactionModel.TransferTransactionDetail{}, nil
	}

	fetchedTransactionIDs := make([]int, 0, len(transferTransactionDTOs))
	for _, dto := range transferTransactionDTOs {
		fetchedTransactionIDs = append(fetchedTransactionIDs, dto.ID)
	}

	// preload transaction records
	var recordDTOs []dto.TransactionRecord
	if err := db.WithContext(ctx).
		Where("transaction_id IN (?)", fetchedTransactionIDs).
		Find(&recordDTOs).Error; err != nil {
		return nil, err
	}

	transactionDetails := dto.ToTransferTransactionDetailModelList(transferTransactionDTOs, recordDTOs)

	return transactionDetails, nil
}

// ListTransferRequests retrieves transaction payout requests with transaction status 0 (New)
func (r *TransactionRepositoryImpl) ListTransferRequests(
	ctx context.Context,
	params *inputdata.TransferRequestListInput,
) (*transactionModel.PaginatedTransferRequest, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	queryBuilder := NewQueryBuilder(db)
	baseQuery := queryBuilder.ApplyReconciliationFilter(
		queryBuilder.BuildTransferRequestBaseQuery(),
		params.ReconciliationFilters,
	)

	countQuery := baseQuery

	var count int64
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, err
	}
	if count == 0 {
		return &transactionModel.PaginatedTransferRequest{
			Items:      []*transactionModel.TransferRequest{},
			Pagination: util.Pagination{TotalPages: 0, TotalCount: 0},
		}, nil
	}

	offset := (params.Page - 1) * params.PageSize
	query := queryBuilder.BuildTransferRequestWithDetailsQuery(
		params.PageSize,
		offset,
		params.ReconciliationFilters,
		params.SortField,
		params.SortOrder,
	)

	var results []*dto.TransferTransactionRequest
	if err := query.WithContext(ctx).Find(&results).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(params.PageSize)))

	models := dto.ToTransferRequestModelList(results)

	return &transactionModel.PaginatedTransferRequest{
		Items:      models,
		Pagination: util.Pagination{TotalPages: totalPages, TotalCount: int(count)},
	}, nil
}

// GetTransactionSummaryRecentMonth retrieves transaction summary counts for recent months
func (r *TransactionRepositoryImpl) GetTransactionSummaryRecentMonth(ctx context.Context, recentMonthCount int) ([]*dashboardModel.TransactionSummaryCount, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	if recentMonthCount <= 0 {
		return nil, errors.New("recentMonthCount must be positive")
	}

	var results []dashboardModel.TransactionSummary
	const transactionSummarySQL = `
		SELECT
			CAST(DATE_FORMAT(created_at, '%Y-%m-01') AS DATE) AS month,
			COUNT(*) AS total_count,
			COUNT(IF(transaction.transaction_status = ?, 1, NULL)) AS processing_count,
			COUNT(IF(transaction.transaction_status = ?, 1, NULL)) AS transferred_count
		FROM transaction
		WHERE created_at >= DATE_FORMAT(CURDATE() - INTERVAL ? MONTH, '%Y-%m-01')
		GROUP BY month
		ORDER BY month DESC
		LIMIT ?
	`
	if err := db.Raw(transactionSummarySQL, object.TransactionStatusProcessing,
		object.TransactionStatusTransferred, recentMonthCount-1, recentMonthCount).Scan(&results).Error; err != nil {
		return nil, err
	}

	var output []*dashboardModel.TransactionSummaryCount
	for _, summary := range results {
		output = append(output, &dashboardModel.TransactionSummaryCount{
			Month:            summary.Month,
			TotalCount:       summary.TotalCount,
			ProcessingCount:  summary.ProcessingCount,
			TransferredCount: summary.TransferredCount,
		})
	}
	return output, nil
}

// BulkCreate creates multiple transactions at once
func (r *TransactionRepositoryImpl) BulkCreate(ctx context.Context, transactions []*transactionModel.Transaction) ([]*transactionModel.Transaction, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	dtos := []*dto.Transaction{}
	for _, transaction := range transactions {
		dto := dto.FromModel(transaction)
		dtos = append(dtos, dto)
	}

	if err = db.Create(&dtos).Error; err != nil {
		return nil, err
	}
	result := []*transactionModel.Transaction{}

	for _, dto := range dtos {
		result = append(result, dto.ToModel())
	}

	return result, nil
}

func (r *TransactionRepositoryImpl) GetTransactionByID(ctx context.Context, id int) (*transactionModel.Transaction, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var transactionDTO dto.Transaction
	if err := db.Where("id = ?", id).First(&transactionDTO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return transactionDTO.ToModel(), nil
}

func (r *TransactionRepositoryImpl) UpdateStatus(ctx context.Context, transaction *transactionModel.Transaction) (*transactionModel.Transaction, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	transactionDTO := dto.FromModel(transaction)
	err = db.Where(dto.Transaction{
		ID: transactionDTO.ID,
	}).
		Updates(&transactionDTO).Error

	if err != nil {
		return nil, err
	}

	return transactionDTO.ToModel(), nil
}

func (r *TransactionRepositoryImpl) UpdateStatusByChunkIDs(ctx context.Context, ids []int, status transactionObject.TransactionStatus, chunkSize int) error {
	if len(ids) == 0 {
		return nil
	}

	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	for _, c := range utils.Chunk(ids, chunkSize) {
		err := db.Model(&dto.Transaction{}).
			Where("id IN ?", c).
			Updates(map[string]any{
				"transaction_status": status,
			}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateTransactionFields updates specified fields for a list of transactions
func (r *TransactionRepositoryImpl) UpdateTransactionFields(ctx context.Context, transactionIDs []int, update *transactionModel.Transaction) error {
	if len(transactionIDs) == 0 || update == nil {
		return nil
	}

	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	updateDTO := dto.FromModel(update)
	fieldsToUpdate := update.GetFieldsToUpdate()

	if len(fieldsToUpdate) == 0 {
		return nil
	}

	if err := db.WithContext(ctx).Model(&dto.Transaction{}).
		Where("id IN ?", transactionIDs).
		Select(fieldsToUpdate).
		Updates(updateDTO).Error; err != nil {
		return err
	}

	return nil
}

// ListTransferTransactions retrieves processed transfers with pagination and filtering
func (r *TransactionRepositoryImpl) ListTransferTransactions(ctx context.Context, params *inputdata.TransferTransactionInput) (*transferTransactionModel.PaginatedTransferTransaction, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	queryBuilder := NewQueryBuilder(db)
	countQuery := queryBuilder.BuildTransferTransactionBaseQuery(params)

	var count int64
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, err
	}
	if count == 0 {
		return &transferTransactionModel.PaginatedTransferTransaction{
			Items:      []*transferTransactionModel.TransferTransaction{},
			Pagination: util.Pagination{TotalPages: 0, TotalCount: 0},
		}, nil
	}

	offset := (params.Page - 1) * params.PageSize
	query := queryBuilder.BuildTransferTransactionWithDetailsQuery(
		params.PageSize,
		offset,
		params,
	)

	var results []*dto.TransferTransactionDTO
	if err := query.WithContext(ctx).Find(&results).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(params.PageSize)))

	models := dto.ToTransferTransactionModelList(results)

	return &transferTransactionModel.PaginatedTransferTransaction{
		Items:      models,
		Pagination: util.Pagination{TotalPages: totalPages, TotalCount: int(count)},
	}, nil
}

// GetTransactionDetailsByPayoutID retrieves transaction details by payout ID
func (r *TransactionRepositoryImpl) GetTransactionDetailsByPayoutID(ctx context.Context, payoutID int) ([]*transactionModel.TransferTransactionDetail, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var transactions []dto.Transaction
	if err := db.
		Preload("TransactionRecords").
		Where("payout_id = ?", payoutID).
		Find(&transactions).Error; err != nil {
		return nil, err
	}

	result := dto.ConvertTransactionsToTransferDetails(transactions)

	return result, nil
}
