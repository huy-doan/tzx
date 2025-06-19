package task

import (
	"context"
	"fmt"

	usecase "github.com/test-tzs/nomraeite/batch/usecase/gmo-aozora"
	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	object "github.com/test-tzs/nomraeite/internal/domain/object/connected_service_token"
	connectedServiceTokenRepo "github.com/test-tzs/nomraeite/internal/domain/repository/connected_service_token"
	prTransferredRepository "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
	appConfig "github.com/test-tzs/nomraeite/internal/pkg/config"

	"github.com/test-tzs/nomraeite/internal/pkg/logger"
)

type TransferStatusTask struct {
	transferStatusUC          usecase.TransferStatusUsecase
	prTransferredRepository   prTransferredRepository.PayoutRecordTransferredRepository
	connectedServiceTokenRepo connectedServiceTokenRepo.ConnectedServiceTokenRepository
	logger                    logger.Logger
	config                    *appConfig.Config
}

func NewTransferStatusTask(
	transferStatusUsecase usecase.TransferStatusUsecase,
	prTransferredRepository prTransferredRepository.PayoutRecordTransferredRepository,
	connectedServiceTokenRepository connectedServiceTokenRepo.ConnectedServiceTokenRepository,
	logger logger.Logger,
	config *appConfig.Config,
) *TransferStatusTask {
	return &TransferStatusTask{
		transferStatusUC:          transferStatusUsecase,
		prTransferredRepository:   prTransferredRepository,
		connectedServiceTokenRepo: connectedServiceTokenRepository,
		logger:                    logger,
		config:                    config,
	}
}

func (t *TransferStatusTask) Do(ctx context.Context) error {
	accountID := t.config.GmoAozoraNetBankPrimaryAccountID
	t.logger.Info("Starting GMO Aozora transfer status task", nil)

	serviceToken, err := t.connectedServiceTokenRepo.FindByServiceName(ctx, object.ConnectedServiceNameGmoAozora)
	if err != nil {
		return err
	}
	if serviceToken == nil || !serviceToken.IsAccessTokenValid() {
		t.logger.Info("No connected service token found for GMO Aozora or token invalid", nil)
		return nil
	}

	sendingDates, err := t.prTransferredRepository.GetPayoutRecordTransferredDates(ctx)
	if err != nil {
		t.logger.Error("Failed to get payout record transferred dates", map[string]any{
			"error": err.Error(),
		})
		return err
	}
	if len(sendingDates) == 0 {
		t.logger.Info("No transferred payout records found", nil)
		return nil
	}

	for _, sendingDate := range sendingDates {
		t.logger.Info(fmt.Sprintf("Processing sending date: %s", sendingDate), nil)
		err = t.transferStatusUC.Execute(ctx, inputdata.TransferStatusRequestParams{
			AccessToken:         serviceToken.AccessToken,
			AccountID:           accountID,
			SendingDate:         sendingDate.SendingDate,
		})

		if err != nil {
			t.logger.Error("Failed to execute transfer status task", map[string]any{
				"error": err.Error(),
			})
			return err
		}
	}

	t.logger.Info("Completed GMO Aozora transfer status task", nil)
	return nil
}
