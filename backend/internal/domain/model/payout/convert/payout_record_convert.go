package convert

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	transaction "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	bankAccountObject "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
)

// NewPayoutRecordsFromTransactions creates payout records from transaction details
func NewPayoutRecordsFromTransactions(transactionDetails []*transaction.TransferTransactionDetail, payoutID int) []*model.PayoutRecord {
	payoutRecords := make([]*model.PayoutRecord, 0, len(transactionDetails))

	for _, detail := range transactionDetails {
		payoutRecord := NewPayoutRecordFromTransaction(detail, payoutID)
		payoutRecords = append(payoutRecords, payoutRecord)
	}

	return payoutRecords
}

func NewPayoutRecordFromTransaction(detail *transaction.TransferTransactionDetail, payoutID int) *model.PayoutRecord {
	return &model.PayoutRecord{
		ShopID:         detail.Merchant.ShopID,
		PayoutID:       payoutID,
		TransactionID:  detail.ID,
		Amount:         int64(detail.CalculateTotalAmount()),
		TransferStatus: object.PayoutRecordStatusDraft,
		BankAccount: model.BankAccount{
			BankAccountType: object.BankAccountTypeOrdinary,
			BankName:        detail.BankName,
			BankCode:        bankAccountObject.FromStringToBankCode(detail.BankCode),
			BranchName:      detail.BranchName,
			BranchCode:      bankAccountObject.FromStringToBankBranchCode(detail.BankBranchCode),
			AccountNo:       bankAccountObject.FromStringToAccountNumber(detail.AccountNumber),
			AccountName:     bankAccountObject.FromStringToAccountHolderKana(detail.AccountName),
		},
		IdempotencyKey: "",
	}
}
