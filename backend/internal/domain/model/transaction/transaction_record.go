package model

import (
	"math"

	object "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	basedatetime "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	transactionRecordObject "github.com/test-tzs/nomraeite/internal/domain/object/transaction_record"
)

// TransactionRecordType constants defined from the database comments
const (
	GmoAozoraBankCode = "0310" // GMOあおぞらネット銀行
	BankTransferFee   = 150    // 振込手数料
	PaypayCommission  = 0.0345 // メイクショップの手数料
)

// TransactionRecord represents a transaction detail record entity in the system
type TransactionRecord struct {
	ID                  int `json:"id"`
	BaseColumnTimestamp basedatetime.BaseColumnTimestamp

	TransactionID         int
	MerchantID            int
	PayinDetailID         int
	TransactionRecordType transactionRecordObject.TransactionRecordType
	Title                 string
	Amount                int64

	Transaction *Transaction
}

// GetRecordTypeText returns the string representation of the transaction record type
func (tr *TransactionRecord) GetRecordTypeText() string {
	return tr.TransactionRecordType.String()
}

// NewTransactionRecord creates a new transaction record with the specified parameters
func NewTransactionRecord(
	transactionID int,
	merchantID int,
	payinDetailID int,
	amount int64,
	recordType transactionRecordObject.TransactionRecordType,
) *TransactionRecord {
	return &TransactionRecord{
		TransactionID:         transactionID,
		MerchantID:            merchantID,
		PayinDetailID:         payinDetailID,
		Title:                 generateTitle(int(recordType)),
		TransactionRecordType: recordType,
		Amount:                amount,
	}
}

// NewDepositTransactionRecord creates a new transaction record for deposits
func NewDepositTransactionRecord(transactionID int, merchantID int, payinDetailID int, amount int64) *TransactionRecord {
	return &TransactionRecord{
		TransactionID:         transactionID,
		MerchantID:            merchantID,
		PayinDetailID:         payinDetailID,
		Title:                 generateTitle(int(transactionRecordObject.TransactionRecordTypeDeposit)),
		TransactionRecordType: transactionRecordObject.TransactionRecordTypeDeposit,
		Amount:                amount,
	}
}

// NewFeeTransactionRecord creates a new transaction record for fees
func NewFeeTransactionRecord(transactionID int, merchantID int, amount int64) *TransactionRecord {
	return &TransactionRecord{
		TransactionID:         transactionID,
		MerchantID:            merchantID,
		Title:                 generateTitle(int(transactionRecordObject.TransactionRecordTypeFee)),
		TransactionRecordType: transactionRecordObject.TransactionRecordTypeFee,
		Amount:                getPaypayCommission(amount),
	}
}

// NewTransferFreeTransactionRecord creates a new transaction record for transfer fees
func NewTransferFreeTransactionRecord(transactionID int, merchantID int, bankCode object.BankCode) *TransactionRecord {
	return &TransactionRecord{
		TransactionID:         transactionID,
		MerchantID:            merchantID,
		Title:                 generateTitle(int(transactionRecordObject.TransactionRecordTypeTransferFee)),
		TransactionRecordType: transactionRecordObject.TransactionRecordTypeTransferFee,
		Amount:                getBankTransferFee(bankCode),
	}
}

// generateTitle generates a title based on transaction record type
func generateTitle(transactionRecordType int) string {
	switch transactionRecordType {
	case int(transactionRecordObject.TransactionRecordTypeDeposit):
		return "入金"
	case int(transactionRecordObject.TransactionRecordTypeFee):
		return "Paypay手数料"
	case int(transactionRecordObject.TransactionRecordTypeTransferFee):
		return "振込手数料"
	default:
		return ""
	}
}

// getPaypayCommission calculates Paypay commission
func getPaypayCommission(amount int64) int64 {
	result := float64(amount) * PaypayCommission
	return -int64(math.Floor(result))
}

// getBankTransferFee returns the bank transfer fee
func getBankTransferFee(bankCode object.BankCode) int64 {
	if !bankCode.IsValid() || bankCode == object.BankCode(GmoAozoraBankCode) {
		return 0
	}
	return -BankTransferFee // 振込手数料はマイナス値で表現
}
