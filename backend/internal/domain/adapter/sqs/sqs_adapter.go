package adapter

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/sqs"
)

type SQSAdapter interface {
	SendBankTransferMessage(message *model.BankTransferStatusMessage) error
	SendApplicationReviewResultMessage(message *model.ReviewResultMessage) error
	SendApplicationReviewResultBatchMessage(message []*model.ReviewResultMessage) error
}
