package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	transactionObject "github.com/test-tzs/nomraeite/internal/domain/object/transaction"
)

// Transaction represents a transaction entity in the system
type Transaction struct {
	ID int `json:"id"`
	util.BaseColumnTimestamp

	ShopID            string
	TransactionStatus transactionObject.TransactionStatus
	PayoutID          *int
	BankCode          object.BankCode          // 銀行コード
	BankBranch        object.BankBranch        // 支店名
	BankBranchCode    object.BankBranchCode    // 支店コード
	AccountNumber     object.AccountNumber     // 口座番号
	AccountHolder     object.AccountHolder     // 口座名義
	AccountHolderKana object.AccountHolderKana // 口座名義カナ
	AccountKind       object.AccountKind       // 1:普通口座, 2:当座口座

	// Relationships - can be added later as needed
	TransactionRecords []TransactionRecord
}
type NewTransactionStatusParams struct {
	TransactionStatus transactionObject.TransactionStatus
}

func NewTransactionStatus(params NewTransactionStatusParams) *Transaction {
	return &Transaction{
		TransactionStatus: params.TransactionStatus,
	}
}

func (t *Transaction) SetStatus(status transactionObject.TransactionStatus) {
	t.TransactionStatus = status
}

func (t *Transaction) SetDraft() {
	t.TransactionStatus = transactionObject.TransactionStatusDraft
}

// GetFieldsToUpdate returns the list of fields that have values set for updating
func (t *Transaction) GetFieldsToUpdate() []string {
	fieldsToUpdate := []string{}

	if t.TransactionStatus != 0 {
		fieldsToUpdate = append(fieldsToUpdate, "transaction_status")
	}

	if t.PayoutID != nil {
		fieldsToUpdate = append(fieldsToUpdate, "payout_id")
	}

	return fieldsToUpdate
}
