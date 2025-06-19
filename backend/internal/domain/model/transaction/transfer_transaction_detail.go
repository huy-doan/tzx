package model

import (
	merchant "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	reconciliationModel "github.com/test-tzs/nomraeite/internal/domain/model/reconciliation"
	basedatetime "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	transactionStatusObj "github.com/test-tzs/nomraeite/internal/domain/object/transaction"
)

type TransferTransactionDetail struct {
	ID                int
	TransactionStatus int
	AccountName       string
	AccountNumber     string
	BankBranchCode    string
	BankCode          string
	BankName          string
	BranchName        string
	PayoutID          int
	Merchant          *merchant.Merchant
	Reconciliation    *reconciliationModel.PayinReconciliation
	basedatetime.BaseColumnTimestamp

	TransactionRecords []TransactionRecord
}

func (td *TransferTransactionDetail) GetTransactionStatusText() string {
	return transactionStatusObj.TransactionStatus(td.TransactionStatus).String()
}

func (td *TransferTransactionDetail) GetStatus() transactionStatusObj.TransactionStatus {
	return transactionStatusObj.TransactionStatus(td.TransactionStatus)
}

func (td *TransferTransactionDetail) CalculateTotalAmount() int64 {
	var total int64
	for _, record := range td.TransactionRecords {
		total += record.Amount
	}
	return total
}

func (td *TransferTransactionDetail) IsEligibleForPayout() bool {
	return td.GetStatus() == transactionStatusObj.TransactionStatusDraft && td.PayoutID == 0
}
