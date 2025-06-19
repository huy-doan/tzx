package usecase

import (
	"context"
	"fmt"

	gmoAozoraService "github.com/test-tzs/nomraeite/batch/domain/service/gmo-aozora"
	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	"github.com/test-tzs/nomraeite/internal/datastructure/outputdata"
	adapterGmoAozora "github.com/test-tzs/nomraeite/internal/domain/adapter/gmo-aozora"
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	sqsObject "github.com/test-tzs/nomraeite/internal/domain/object/sqs"
	connectedServiceTokenRepo "github.com/test-tzs/nomraeite/internal/domain/repository/connected_service_token"
	payoutRecordRepository "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
)

type TransferStatusUsecase interface {
	Execute(ctx context.Context, params inputdata.TransferStatusRequestParams) error
}

type transferStatusUsecaseImpl struct {
	logger                    logger.Logger
	prTransferredRepository   payoutRecordRepository.PayoutRecordTransferredRepository
	connectedServiceTokenRepo connectedServiceTokenRepo.ConnectedServiceTokenRepository
	apiClient                 adapterGmoAozora.ApiClient
	transferStatusService     gmoAozoraService.TransferStatusService
}

func NewTransferStatusUsecase(
	logger logger.Logger,
	prTransferredRepository payoutRecordRepository.PayoutRecordTransferredRepository,
	connectedServiceTokenRepository connectedServiceTokenRepo.ConnectedServiceTokenRepository,
	apiClient adapterGmoAozora.ApiClient,
	transferStatusService gmoAozoraService.TransferStatusService,
) TransferStatusUsecase {
	return &transferStatusUsecaseImpl{
		logger:                    logger,
		prTransferredRepository:   prTransferredRepository,
		connectedServiceTokenRepo: connectedServiceTokenRepository,
		apiClient:                 apiClient,
		transferStatusService:     transferStatusService,
	}
}

func (u *transferStatusUsecaseImpl) Execute(ctx context.Context, params inputdata.TransferStatusRequestParams) error {
	payoutRecords, err := u.prTransferredRepository.GetPayoutRecordsWithMerchantBySendingDate(ctx, params.SendingDate)
	if err != nil {
		return err
	}
	if len(payoutRecords) == 0 {
		u.logger.Info("No transferred payout records found", nil)
		return nil
	}

	u.logger.Info(fmt.Sprintf("Found %d records to check transfer status", len(payoutRecords)), nil)
	results, err := u.transferStatusService.GetGmoAozoraTransferStatus(ctx, params)
	if err != nil {
		return err
	}
	if len(results) == 0 {
		return nil
	}

	newPayoutRecords, sqsBodies := u.prepareNewPayoutRecordsAndSQSBodies(payoutRecords, results)

	if len(newPayoutRecords) == 0 {
		u.logger.Info("No payout records found with changed transfer status", nil)
		return nil
	}

	err = u.transferStatusService.BulkUpdateTransferStatus(ctx, newPayoutRecords)
	if err != nil {
		return err
	}

	err = u.transferStatusService.SendBankTransferMessage(sqsBodies)
	if err != nil {
		return err
	}

	return nil
}

func (u *transferStatusUsecaseImpl) isAozoraTransferStatusChanged(payoutRecord *model.PayoutRecord, result outputdata.GmoAozoraTransferStatusOutput) bool {
	return payoutRecord.AozoraTransferApplyNo == result.ApplyNo && payoutRecord.AozoraTransferStatus != result.GmoAozoraTransferStatus
}

func (u *transferStatusUsecaseImpl) prepareNewPayoutRecordsAndSQSBodies(
	payoutRecords []*model.PayoutRecordWithMerchantID,
	results []outputdata.GmoAozoraTransferStatusOutput,
) (newPayoutRecords []*model.PayoutRecord, sqsBodies []sqsObject.BodyTransferStatus) {
	mapPayoutRecords := utils.ArrayToMap(payoutRecords, func(pr *model.PayoutRecordWithMerchantID) int {
		return pr.ID
	})

	for _, result := range results {
		payoutRecord, exists := mapPayoutRecords[result.PayoutRecordID]
		if !exists {
			u.logger.Info(fmt.Sprintf("PayoutRecord with ID %d not found", result.PayoutRecordID), nil)
			continue
		}
		pr := payoutRecord.GetPayoutRecord()
		if u.isAozoraTransferStatusChanged(pr, result) {
			payoutRecord.SetAozoraTransferStatus(result.GmoAozoraTransferStatus)
			if !result.TransferStatus.IsDraft() {
				payoutRecord.SetTransferStatus(result.TransferStatus)
			}
			newPayoutRecords = append(newPayoutRecords, pr)

			sqsBody := sqsObject.NewBodyTransferStatus(
				result.ApplyNo,
				payoutRecord.ShopID,
				utils.ToStringFromInt(payoutRecord.PayoutID),
				result.ApplyDatetime,
				payoutRecord.Amount,
				result.AcceptDatetime,
				result.AcceptNo,
				result.GmoAozoraTransferStatus.Value(),
				&sqsObject.PayPayAttributes{
					MerchantID: fmt.Sprint(rune(payoutRecord.MerchantID)),
				},
			)
			sqsBodies = append(sqsBodies, *sqsBody)
		}
	}

	return
}
