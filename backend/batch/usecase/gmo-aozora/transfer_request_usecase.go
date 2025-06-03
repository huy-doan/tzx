package usecase

import (
	"context"
	"fmt"

	payoutRepo "github.com/test-tzs/nomraeite/batch/domain/repository/payout"
	transactionRepo "github.com/test-tzs/nomraeite/batch/domain/repository/transaction"
	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	adapterGmoAozora "github.com/test-tzs/nomraeite/internal/domain/adapter/gmo-aozora"
	gmoAozoraService "github.com/test-tzs/nomraeite/internal/domain/service/gmo-aozora"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
)

type TransferRequestUsecase interface {
	Execute(ctx context.Context, params inputdata.TransferRequestParams) error
}

type transferRequestUsecaseImp struct {
	logger             logger.Logger
	payoutRepo         payoutRepo.PayoutRepository
	payoutRecordRepo   payoutRepo.PayoutRecordRepository
	transactionRepo    transactionRepo.TransactionRepository
	apiClient          adapterGmoAozora.ApiClient
	transferService    gmoAozoraService.TransferService
	transactionService gmoAozoraService.TransactionService
}

type TransferResult struct {
	Success bool
}

func NewTransferRequestUsecase(
	logger logger.Logger,
	payoutRepository payoutRepo.PayoutRepository,
	payoutRecordRepository payoutRepo.PayoutRecordRepository,
	transactionRepository transactionRepo.TransactionRepository,
	client adapterGmoAozora.ApiClient,
) TransferRequestUsecase {
	transferService := gmoAozoraService.NewTransferService(
		logger,
		payoutRecordRepository,
		client,
	)
	transactionService := gmoAozoraService.NewTransactionService(
		logger,
		transactionRepository,
	)

	return &transferRequestUsecaseImp{
		logger:             logger,
		payoutRepo:         payoutRepository,
		payoutRecordRepo:   payoutRecordRepository,
		transactionRepo:    transactionRepository,
		apiClient:          client,
		transferService:    transferService,
		transactionService: transactionService,
	}
}

func (uc *transferRequestUsecaseImp) Execute(ctx context.Context, params inputdata.TransferRequestParams) error {
	for _, record := range params.PayoutRecords {
		if !record.IsBankAccountValid() {
			uc.logger.Error(fmt.Sprintf("Invalid bank account for payout record ID: %d", record.ID), nil)
			continue
		}

		tx, txErr := database.NewTx[TransferResult](ctx)
		if txErr != nil {
			uc.logger.Error("Failed to create transaction", map[string]any{
				"error": txErr.Error(),
			})
			return fmt.Errorf("failed to create transaction: %w", txErr)
		}

		result, txErr := tx.Transact(ctx, func(txCtx context.Context) (TransferResult, error) {
			payoutRecord, err := uc.payoutRecordRepo.LockPayoutRecordByID(txCtx, record.ID)
			if err != nil {
				uc.logger.Error(fmt.Sprintf("Failed to lock payout record ID: %d", record.ID), map[string]any{
					"error": err.Error(),
				})
				return TransferResult{Success: false}, err
			}

			if payoutRecord == nil {
				uc.logger.Info(fmt.Sprintf("Payout record ID: %d not found after locking", record.ID), nil)
				return TransferResult{Success: true}, nil
			}

			err = uc.transferService.RequestTransfer(txCtx, payoutRecord, params)
			if err != nil {
				uc.logger.Error(fmt.Sprintf("Failed to request transfer for payout record ID: %d", payoutRecord.ID), map[string]any{
					"error": err.Error(),
				})
				return TransferResult{Success: false}, err
			}

			err = uc.transactionService.UpdateTransaction(txCtx, payoutRecord)
			if err != nil {
				uc.logger.Error(fmt.Sprintf("Failed to update transaction for payout record ID: %d", payoutRecord.ID), map[string]any{
					"error": err.Error(),
				})
				return TransferResult{Success: false}, err
			}

			return TransferResult{Success: true}, nil
		})

		if !result.Success {
			uc.logger.Error(fmt.Sprintf("Transfer request failed for payout record ID: %d", record.ID), map[string]any{
				"error": txErr.Error(),
			})
			return txErr
		}
	}

	return nil
}
