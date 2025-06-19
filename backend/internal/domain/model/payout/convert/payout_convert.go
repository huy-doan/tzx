package convert

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
)

// NewPayoutFromTransactions creates a new payout from transaction details
func NewPayoutFromTransactions(params model.CreateFromTransactionsParams) *model.Payout {
	total := calculateTotalFromTransactions(params.TransactionDetails)
	totalCount := len(params.TransactionDetails)

	return &model.Payout{
		PayoutStatus:          object.PayoutStatusWaitingApproval,
		Total:                 total,
		TotalCount:            totalCount,
		SendingDate:           nil,
		SentDate:              nil,
		UserID:                params.UserID,
		PayoutRecordCount:     totalCount,
		PayoutRecordSumAmount: total,
	}
}

// calculateTotalFromTransactions calculates the total amount from transaction details
func calculateTotalFromTransactions(transactionDetails []*transactionModel.TransferTransactionDetail) int64 {
	var total int64
	for _, detail := range transactionDetails {
		total += detail.CalculateTotalAmount()
	}

	return total
}
