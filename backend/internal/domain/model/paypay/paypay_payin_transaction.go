package model

import (
	"time"

	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	paypayObject "github.com/test-tzs/nomraeite/internal/domain/object/paypay"
)

// PayPayPayinTransaction represents the paypay_payin_transaction table
type PaypayPayinTransaction struct {
	ID int
	util.BaseColumnTimestamp

	PayinFileID              int
	PaymentTransactionID     string
	PaymentMerchantID        string
	MerchantBusinessName     string
	ShopID                   string
	ShopName                 string
	TerminalCode             string
	PaymentTransactionStatus paypayObject.PaypayTransactionStatus
	TransactionAt            *time.Time
	TransactionAmount        int64
	ReceiptNumber            string
	PaypayPaymentMethod      string
	SSID                     string
	MerchantOrderID          string
	PaymentDetail            PaymentDetail `gorm:"type:json;serializer:json"`
	OrderTransactionAmount   int64
	// PayinFile *model.PayinFile
}

type PaginatedPaypayPayinTransactionResult struct {
	Items []*PaypayPayinTransaction
	util.Pagination
}

type OrderTransactionAmount struct {
	ID                   int
	PaymentTransactionID int
	PaymentMethod        string
	TransactionAmount    *int64

	util.BaseColumnTimestamp
}
