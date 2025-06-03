package service

import (
	"context"
	"errors"
	"fmt"

	payoutRepo "github.com/test-tzs/nomraeite/batch/domain/repository/payout"
	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	adapterGmoAozora "github.com/test-tzs/nomraeite/internal/domain/adapter/gmo-aozora"
	gmoAozoraModel "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	payoutRecordModel "github.com/test-tzs/nomraeite/internal/domain/model/payout_record"
	"github.com/test-tzs/nomraeite/internal/pkg/config"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
)

type TransferService interface {
	RequestTransfer(ctx context.Context, payoutRecord *payoutRecordModel.PayoutRecord, params inputdata.TransferRequestParams) error
}

type transferServiceImpl struct {
	logger           logger.Logger
	payoutRecordRepo payoutRepo.PayoutRecordRepository
	apiClient        adapterGmoAozora.ApiClient
}

func NewTransferService(
	logger logger.Logger,
	payoutRecordRepo payoutRepo.PayoutRecordRepository,
	apiClient adapterGmoAozora.ApiClient,
) TransferService {
	return &transferServiceImpl{
		logger:           logger,
		payoutRecordRepo: payoutRecordRepo,
		apiClient:        apiClient,
	}
}

func (s *transferServiceImpl) RequestTransfer(ctx context.Context, payoutRecord *payoutRecordModel.PayoutRecord, params inputdata.TransferRequestParams) (err error) {
	config := config.GetConfig()
	payoutRecord.GenerateIdempotencyKey()

	header := gmoAozoraModel.NewTransferHeaderRequest(params.AccessToken, payoutRecord.IdempotencyKey)
	request := gmoAozoraModel.NewTransferParamsRequest(
		config.GmoAozoraNetBankPrimaryAccountID,
		payoutRecord,
	)
	response := s.apiClient.RequestTransfer(ctx, header, request)

	if response.ShouldStopBatch() {
		s.logger.Error(fmt.Sprintf("Critical error for payout record ID: %d", payoutRecord.ID), map[string]any{
			"response_type": response.Type.String(),
			"message":       response.GetErrorMessage(),
		})
		return errors.New(response.GetErrorMessage())
	}

	if !response.IsSuccess() {
		s.logger.Error(fmt.Sprintf("Transfer request failed for payout record ID: %d", payoutRecord.ID), map[string]any{
			"response_type": response.Type.String(),
			"message":       response.GetErrorMessage(),
		})
		payoutRecord.SetPayoutRecordTransferFailed(response.GetErrorMessage())
	} else {
		s.logger.Info(fmt.Sprintf("Transfer request successful for payout record ID: %d", payoutRecord.ID), nil)
		payoutRecord.SetPayoutRecordTransferSuccess(response.SuccessResult.ApplyNo)
	}

	_, err = s.payoutRecordRepo.UpdateByID(ctx, payoutRecord.ID, payoutRecord)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to update payout record ID: %d", payoutRecord.ID), map[string]any{
			"error": err.Error(),
		})
	}

	return err
}
