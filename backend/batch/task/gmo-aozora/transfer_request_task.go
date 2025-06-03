package task

import (
	"context"
	"errors"
	"fmt"

	payoutRepo "github.com/makeshop-jp/master-console/batch/domain/repository/payout"
	usecase "github.com/makeshop-jp/master-console/batch/usecase/gmo-aozora"
	"github.com/makeshop-jp/master-console/internal/datastructure/inputdata"
	object "github.com/makeshop-jp/master-console/internal/domain/object/connected_service_token"
	payoutObject "github.com/makeshop-jp/master-console/internal/domain/object/payout"
	connectedServiceTokenRepo "github.com/makeshop-jp/master-console/internal/domain/repository/connected_service_token"
	"github.com/makeshop-jp/master-console/internal/pkg/logger"
)

type TransferRequestTask struct {
	transferRequestUC         usecase.TransferRequestUsecase
	connectedServiceTokenRepo connectedServiceTokenRepo.ConnectedServiceTokenRepository
	payoutRepo                payoutRepo.PayoutRepository
	payoutRecordRepo          payoutRepo.PayoutRecordRepository
	logger                    logger.Logger
}

func NewTransferRequestTask(
	transferRequestUsecase usecase.TransferRequestUsecase,
	connectedServiceTokenRepository connectedServiceTokenRepo.ConnectedServiceTokenRepository,
	payoutRepository payoutRepo.PayoutRepository,
	payoutRecordRepository payoutRepo.PayoutRecordRepository,
	logger logger.Logger,
) *TransferRequestTask {
	return &TransferRequestTask{
		transferRequestUC:         transferRequestUsecase,
		connectedServiceTokenRepo: connectedServiceTokenRepository,
		payoutRepo:                payoutRepository,
		payoutRecordRepo:          payoutRecordRepository,
		logger:                    logger,
	}
}

func (t *TransferRequestTask) Do(ctx context.Context) error {
	serviceToken, err := t.connectedServiceTokenRepo.FindByServiceName(ctx, object.ConnectedServiceNameGmoAozora)
	if err != nil {
		return err
	}
	if serviceToken == nil || !serviceToken.IsAccessTokenValid() {
		t.logger.Info("No connected service token found for GMO Aozora", nil)
		return nil
	}

	payout, err := t.payoutRepo.GetFirstTransferingPayout(ctx)
	if err != nil {
		return err
	}
	if payout == nil {
		t.logger.Info("No payout record found for transfer request", nil)
		return nil
	}
	payoutRecords, err := t.payoutRecordRepo.GetListTransferingByPayoutID(ctx, payout.ID)
	if err != nil {
		return err
	}
	if payoutRecords == nil {
		t.logger.Info("No payout records found", nil)
		return nil
	}

	switch true {
	case payout.IndividualTransfer():
		err = t.transferRequestUC.Execute(
			ctx, inputdata.TransferRequestParams{
				AccessToken:   serviceToken.AccessToken,
				Payout:        payout,
				PayoutRecords: payoutRecords,
			})
		if err != nil {
			t.logger.Error(fmt.Sprintf("Failed to execute transfer request for payout ID: %d", payout.ID), map[string]any{
				"error": err.Error(),
			})
		}
	default:
		return errors.New("unsupported payout type for transfer request")
	}

	payout.SetStatus(payoutObject.PayoutStatusTransferred)
	err = t.payoutRepo.UpdateByID(ctx, payout.ID, payout)
	if err != nil {
		t.logger.Error(fmt.Sprintf("Failed to update payout ID: %d as processed", payout.ID), map[string]any{
			"error": err.Error(),
		})
		return err
	}

	return nil
}
