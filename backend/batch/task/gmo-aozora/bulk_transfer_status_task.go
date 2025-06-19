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

type BulkTransferStatusTask struct {
	transferStatusUC          usecase.TransferStatusUsecase
	prTransferredRepository   prTransferredRepository.PayoutRecordTransferredRepository
	connectedServiceTokenRepo connectedServiceTokenRepo.ConnectedServiceTokenRepository
	logger                    logger.Logger
	config                    *appConfig.Config
}

func NewBulkTransferStatusTask(
	transferStatusUsecase usecase.TransferStatusUsecase,
	prTransferredRepository prTransferredRepository.PayoutRecordTransferredRepository,
	connectedServiceTokenRepository connectedServiceTokenRepo.ConnectedServiceTokenRepository,
	logger logger.Logger,
	config *appConfig.Config,
) *BulkTransferStatusTask {
	return &BulkTransferStatusTask{
		transferStatusUC:          transferStatusUsecase,
		prTransferredRepository:   prTransferredRepository,
		connectedServiceTokenRepo: connectedServiceTokenRepository,
		logger:                    logger,
		config:                    config,
	}
}

func (t *BulkTransferStatusTask) Do(ctx context.Context) error {
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

	applyNos, err := t.prTransferredRepository.GetPayoutRecordApplyNos(ctx)
	if err != nil {
		t.logger.Error("Failed to get payout record Apply No", map[string]any{
			"error": err.Error(),
		})
		return err
	}
	if len(applyNos) == 0 {
		t.logger.Info("No transferred payout records found", nil)
		return nil
	}

	for _, applyNo := range applyNos {
		t.logger.Info(fmt.Sprintf("Processing apply no: %s", applyNo), nil)
		// err = t.transferStatusUC.Execute(ctx, )
		input := inputdata.BulkTransferStatusRequestParams{
			AccessToken:         serviceToken.AccessToken,
			AccountID:           accountID,
			ApplyNo:         applyNo.AozoraTransferApplyNo,
		}
		
		// TODO: Implement the logic to handle the bulk transfer status request

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
