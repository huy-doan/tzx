package model

import (
	"fmt"
	"strconv"
	"time"

	paypayObject "github.com/test-tzs/nomraeite/internal/domain/object/paypay"
)

// PaypayPayinTransactionCSV is used for CSV import, all fields as *string for csvutil
// This struct is only for import mapping, not for domain logic.
type PaypayPayinTransactionCSV struct {
	PaymentTransactionID     *string        `csv:"決済番号"`
	PaymentMerchantID        *string        `csv:"加盟店ID"`
	MerchantBusinessName     *string        `csv:"屋号"`
	ShopID                   *string        `csv:"店舗ID"`
	ShopName                 *string        `csv:"店舗名"`
	TerminalCode             *string        `csv:"端末番号/PosID"`
	PaymentTransactionStatus *string        `csv:"取引ステータス"`
	TransactionAt            *string        `csv:"取引日時"`
	TransactionAmount        *string        `csv:"取引金額"`
	ReceiptNumber            *string        `csv:"レシート番号"`
	PaypayPaymentMethod      *string        `csv:"支払い方法"`
	SSID                     *string        `csv:"SSID"`
	MerchantOrderID          *string        `csv:"加盟店決済ID"`
	PaymentDetail            *PaymentDetail `csv:"支払い詳細"`
}

// validate all fields for PaypayPayinTransactionCSV exists and are not nil
func (c *PaypayPayinTransactionCSV) ValidateFields() error {
	if c.PaymentTransactionID == nil {
		return fmt.Errorf("PaymentTransactionID field not found")
	}
	if c.PaymentMerchantID == nil {
		return fmt.Errorf("PaymentMerchantID field not found")
	}
	if c.MerchantBusinessName == nil {
		return fmt.Errorf("MerchantBusinessName field not found")
	}
	if c.ShopID == nil {
		return fmt.Errorf("ShopID field not found")
	}
	if c.ShopName == nil {
		return fmt.Errorf("ShopName field not found")
	}
	if c.TerminalCode == nil {
		return fmt.Errorf("TerminalCode field not found")
	}
	if c.PaymentTransactionStatus == nil {
		return fmt.Errorf("PaymentTransactionStatus field not found")
	}
	if c.TransactionAt == nil {
		return fmt.Errorf("TransactionAt field not found")
	}
	if c.TransactionAmount == nil {
		return fmt.Errorf("TransactionAmount field not found")
	}
	if c.ReceiptNumber == nil {
		return fmt.Errorf("ReceiptNumber field not found")
	}
	if c.PaypayPaymentMethod == nil {
		return fmt.Errorf("PaypayPaymentMethod field not found")
	}
	if c.SSID == nil {
		return fmt.Errorf("SSID field not found")
	}
	if c.MerchantOrderID == nil {
		return fmt.Errorf("MerchantOrderID field not found")
	}
	if c.PaymentDetail == nil {
		return fmt.Errorf("PaymentDetail field not found")
	}
	return nil
}

// validate fields for PaypayPayinTransactionCSV
func (c *PaypayPayinTransactionCSV) ValidateRequired() error {
	if *c.PaymentTransactionStatus == "" {
		return fmt.Errorf("PaymentTransactionStatus is required")
	}

	if *c.TransactionAmount == "" {
		return fmt.Errorf("TransactionAmount is required")
	}

	if *c.MerchantOrderID == "" {
		return fmt.Errorf("MerchantOrderID is required")
	}

	return nil
}

// ToPaypayTransaction converts a CSV transaction record to a domain model
func (c *PaypayPayinTransactionCSV) ToPaypayTransaction(payinFileID int) (*PaypayPayinTransaction, error) {
	// Validate required fields
	if err := c.ValidateRequired(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	// Parse transaction amount
	var transactionAmount int64
	if *c.TransactionAmount != "" {
		amount, err := strconv.ParseInt(*c.TransactionAmount, 10, 64)
		if err != nil {
			return nil, err
		}
		transactionAmount = amount
	}

	// Parse transaction time
	var transactionAt *time.Time
	if *c.TransactionAt != "" {
		var t time.Time
		var err error

		// Try known formats
		format := "2006/01/02 15:04:05"
		t, err = time.ParseInLocation(format, *c.TransactionAt, time.Local)

		transactionAt = &t
		if err != nil {
			return nil, fmt.Errorf("failed to parse transaction time %q: %w", *c.TransactionAt, err)
		}
	}

	// Convert transaction status
	status := paypayObject.PaypayTransactionStatusFromString(*c.PaymentTransactionStatus)
	if status == nil {
		return nil, fmt.Errorf("invalid PaymentTransactionStatus: %s", *c.PaymentTransactionStatus)
	}

	// Build the transaction model
	transaction := &PaypayPayinTransaction{
		PayinFileID:              payinFileID,
		PaymentTransactionID:     *c.PaymentTransactionID,
		PaymentMerchantID:        *c.PaymentMerchantID,
		MerchantBusinessName:     *c.MerchantBusinessName,
		ShopID:                   *c.ShopID,
		ShopName:                 *c.ShopName,
		TerminalCode:             *c.TerminalCode,
		PaymentTransactionStatus: *status,
		TransactionAt:            transactionAt,
		TransactionAmount:        transactionAmount,
		ReceiptNumber:            *c.ReceiptNumber,
		PaypayPaymentMethod:      *c.PaypayPaymentMethod,
		SSID:                     *c.SSID,
		MerchantOrderID:          *c.MerchantOrderID,
		PaymentDetail:            *c.PaymentDetail,
	}

	return transaction, nil
}
