package service

import (
	"context"
	"fmt"

	transactionRepo "github.com/makeshop-jp/master-console/batch/domain/repository/transaction"
	payoutRecordModel "github.com/makeshop-jp/master-console/internal/domain/model/payout_record"
	object "github.com/makeshop-jp/master-console/internal/domain/object/transaction"
	"github.com/makeshop-jp/master-console/internal/pkg/logger"
)

type TransactionService interface {
	UpdateTransaction(ctx context.Context, payoutRecord *payoutRecordModel.PayoutRecord) error
}

type transactionServiceImpl struct {
	logger          logger.Logger
	transactionRepo transactionRepo.TransactionRepository
}

func NewTransactionService(
	logger logger.Logger,
	transactionRepo transactionRepo.TransactionRepository,
) TransactionService {
	return &transactionServiceImpl{
		logger:          logger,
		transactionRepo: transactionRepo,
	}
}

func (s *transactionServiceImpl) UpdateTransaction(ctx context.Context, payoutRecord *payoutRecordModel.PayoutRecord) error {
	transaction, err := s.transactionRepo.GetTransactionByID(ctx, payoutRecord.TransactionID)
	if err != nil {
		return err
	}
	if transaction == nil {
		s.logger.Error(fmt.Sprintf("Transaction not found for payout record ID: %d", payoutRecord.ID), nil)
		return nil
	}

	transaction.SetStatus(object.TransactionStatusTransfered)
	_, err = s.transactionRepo.UpdateStatus(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}
