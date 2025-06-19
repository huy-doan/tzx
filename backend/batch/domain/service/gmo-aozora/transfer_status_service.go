package service

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	"github.com/test-tzs/nomraeite/internal/datastructure/outputdata"
	adapterGmoAozora "github.com/test-tzs/nomraeite/internal/domain/adapter/gmo-aozora"
	sqsAdapter "github.com/test-tzs/nomraeite/internal/domain/adapter/sqs"
	gmoAozoraModel "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	transferStatusModel "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora/transfer_status"
	sqsModel "github.com/test-tzs/nomraeite/internal/domain/model/api/sqs"
	payoutRecordModel "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	transferStatusObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora/transfer_status"
	payoutObject "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	sqsObject "github.com/test-tzs/nomraeite/internal/domain/object/sqs"
	payoutRecordRepository "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
	utils "github.com/test-tzs/nomraeite/internal/pkg/utils"
)

type TransferStatusService interface {
	GetGmoAozoraTransferStatus(ctx context.Context, params inputdata.TransferStatusRequestParams) (outputs []outputdata.GmoAozoraTransferStatusOutput, err error)
	BulkUpdateTransferStatus(ctx context.Context, payoutRecords []*payoutRecordModel.PayoutRecord) error
	SendBankTransferMessage(sqsBodies []sqsObject.BodyTransferStatus) error
}

type transferStatusServiceImpl struct {
	logger                  logger.Logger
	apiClient               adapterGmoAozora.ApiClient
	sqsClient               sqsAdapter.SQSAdapter
	prTransferredRepository payoutRecordRepository.PayoutRecordTransferredRepository
}

func NewTransferStatusService(
	logger logger.Logger,
	apiClient adapterGmoAozora.ApiClient,
	sqsClient sqsAdapter.SQSAdapter,
	prTransferredRepository payoutRecordRepository.PayoutRecordTransferredRepository,
) TransferStatusService {
	return &transferStatusServiceImpl{
		logger:                  logger,
		apiClient:               apiClient,
		sqsClient:               sqsClient,
		prTransferredRepository: prTransferredRepository,
	}
}
func (s *transferStatusServiceImpl) GetGmoAozoraTransferStatus(ctx context.Context, params inputdata.TransferStatusRequestParams) (outputs []outputdata.GmoAozoraTransferStatusOutput, err error) {
	return s.sendRequestRecursive(ctx, params, []outputdata.GmoAozoraTransferStatusOutput{})
}

func (s *transferStatusServiceImpl) sendRequestRecursive(ctx context.Context, params inputdata.TransferStatusRequestParams, items []outputdata.GmoAozoraTransferStatusOutput) (outputs []outputdata.GmoAozoraTransferStatusOutput, err error) {
	dateStr := utils.FormatDate(params.SendingDate)

	request := transferStatusModel.TransferStatusRequest{
		AccountID:           params.AccountID,
		QueryKeyClass:       transferStatusObject.QueryKeyClassBulk,
		DateFrom:            dateStr,
		DateTo:              dateStr,
		NextItemKey:         params.NextItemKey,
		RequestTransferTerm: transferStatusObject.RequestTransferTermDesignatedDate,
	}

	header := gmoAozoraModel.AuthHeader{
		AccessToken: params.AccessToken,
	}

	response, err := s.apiClient.GetTransferStatus(ctx, header, request)
	if err != nil {
		s.logger.Error("Failed to get transfer status", map[string]any{
			"error":      err.Error(),
			"date_from":  dateStr,
			"date_to":    dateStr,
			"account_id": params.AccountID,
		})
		return outputs, err
	}

	if response.IsEmpty() {
		s.logger.Info("No transfer found in response", map[string]any{
			"date_from": dateStr,
			"date_to":   dateStr,
		})
		return outputs, nil
	}

	for _, detail := range response.TransferDetails {
		if detail.IsEmptyTransferApplies() {
			s.logger.Info("No transfer applies found for detail", nil)
			continue
		}

		transferApplies := detail.GetTransferApply()
		transferApplyDetail := transferApplies.GetTransferApplyDetail()
		transferAccept := detail.GetTransferAcceptDetail()
		transferResponse := detail.GetTransferResponse()
		transferInfo := transferResponse.GetTransferInfo()
		applyNo := transferApplies.ApplyNo

		output := outputdata.GmoAozoraTransferStatusOutput{
			GmoAozoraTransferStatus: detail.TransferStatus,
			ApplyNo:                 applyNo,
			ApplyDatetime:           transferApplyDetail.ApplyDatetime,
			AcceptNo:                transferAccept.AcceptNo,
			AcceptDatetime:          utils.ToPtr(transferAccept.AcceptDatetime),
			PayoutRecordID:          utils.ToInt(*transferInfo.EdiInfo),
		}

		if !detail.IsNotCompleted() {
			output.TransferStatus = payoutObject.PayoutRecordStatusDone
		}

		outputs = append(outputs, output)
	}
	outputs = append(outputs, items...)
	if response.HasNext() {
		s.logger.Info("There are more items to fetch", map[string]any{
			"next_item_key": response.TransferQueryBulkResponses[0].NextItemKey,
		})
		nextItemKey := response.GetNextItemKey()
		if nextItemKey != "" {
			params.NextItemKey = nextItemKey
			return s.sendRequestRecursive(ctx, params, outputs)
		}
	}

	return outputs, nil
}

func (s *transferStatusServiceImpl) BulkUpdateTransferStatus(ctx context.Context, payoutRecords []*payoutRecordModel.PayoutRecord) error {
	if len(payoutRecords) == 0 {
		s.logger.Info("No payout records to update transfer status", nil)
		return nil
	}

	s.logger.Info("Updating transfer status for payout records", map[string]any{
		"count": len(payoutRecords),
	})

	return s.prTransferredRepository.BulkUpdateTransferStatusByIDs(ctx, payoutRecords)
}

func (s *transferStatusServiceImpl) SendBankTransferMessage(sqsBodies []sqsObject.BodyTransferStatus) error {
	if len(sqsBodies) == 0 {
		s.logger.Info("No SQS bodies to send", nil)
		return nil
	}

	paymethod := sqsObject.PayPay
	messageAttr := sqsObject.MessageAttribute{
		MessageType: sqsObject.BankTransferStatus,
		Paymethod:   &paymethod,
	}
	for _, messageBody := range sqsBodies {
		err := s.sendBankTransferMessage(messageBody, messageAttr)
		if err != nil {
			s.logger.Error("Failed to send bank transfer message", map[string]any{
				"error":           err.Error(),
				"payout_id":       messageBody.PayoutID,
				"shop_id":         messageBody.ShopID,
				"transfer_status": messageBody.TransferStatus,
			})
			return err
		}
	}

	s.logger.Info("Sent bank transfer message successfully", nil)
	return nil
}

func (s *transferStatusServiceImpl) sendBankTransferMessage(messageBody sqsObject.BodyTransferStatus, messageAttr sqsObject.MessageAttribute) error {
	message := &sqsModel.BankTransferStatusMessage{
		MessageAttributes: messageAttr,
		MessageBody:       messageBody,
	}

	return s.sqsClient.SendBankTransferMessage(message)
}
