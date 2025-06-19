package model

import (
	"fmt"
	"time"

	"github.com/test-tzs/nomraeite/internal/pkg/utils"
)

// PayPayPayinSummaryCSV is used for CSV import, all fields as *string for csvutil
// This struct is only for import mapping, not for domain logic.
type PaypayPayinSummaryCSV struct {
	CorporateName     *string `csv:"法人名"`
	CutoffDate        *string `csv:"締め日"`
	PaymentDate       *string `csv:"支払日"`
	TransactionAmount *string `csv:"取引額"`
	RefundAmount      *string `csv:"返金額"`
	UsageFee          *string `csv:"利用料"`
	PlatformFee       *string `csv:"プラットフォーム使用料"`
	InitialFee        *string `csv:"初期費用"`
	Tax               *string `csv:"税"`
	Cashback          *string `csv:"キャッシュバック"`
	Adjustment        *string `csv:"調整額"`
	Fee               *string `csv:"入金手数料"`
	Amount            *string `csv:"支払金額"`
}

// validate all fields for PaypayPayinSummaryCSV exists and are not nil
func (c *PaypayPayinSummaryCSV) ValidateFields() error {
	if c.CorporateName == nil {
		return fmt.Errorf("CorporateName field not found")
	}
	if c.CutoffDate == nil {
		return fmt.Errorf("CutoffDate field not found")
	}
	if c.PaymentDate == nil {
		return fmt.Errorf("PaymentDate field not found")
	}
	if c.TransactionAmount == nil {
		return fmt.Errorf("TransactionAmount field not found")
	}
	if c.RefundAmount == nil {
		return fmt.Errorf("RefundAmount field not found")
	}
	if c.UsageFee == nil {
		return fmt.Errorf("UsageFee field not found")
	}
	if c.PlatformFee == nil {
		return fmt.Errorf("PlatformFee field not found")
	}
	if c.InitialFee == nil {
		return fmt.Errorf("InitialFee field not found")
	}
	if c.Tax == nil {
		return fmt.Errorf("tax field not found")
	}
	if c.Cashback == nil {
		return fmt.Errorf("cashback field not found")
	}
	if c.Adjustment == nil {
		return fmt.Errorf("adjustment field not found")
	}
	if c.Fee == nil {
		return fmt.Errorf("fee field not found")
	}
	if c.Amount == nil {
		return fmt.Errorf("amount field not found")
	}
	return nil
}

// validate fields for PaypayPayinTransactionCSV
func (c *PaypayPayinSummaryCSV) Validate() error {

	if *c.TransactionAmount == "" {
		return fmt.Errorf("TransactionAmount is required")
	}

	if *c.RefundAmount == "" {
		return fmt.Errorf("RefundAmount is required")
	}

	if *c.UsageFee == "" {
		return fmt.Errorf("UsageFee is required")
	}

	if *c.PlatformFee == "" {
		return fmt.Errorf("PlatformFee is required")
	}

	if *c.Tax == "" {
		return fmt.Errorf("tax is required")
	}

	if *c.Cashback == "" {
		return fmt.Errorf("cashback is required")
	}

	if *c.Adjustment == "" {
		return fmt.Errorf("adjustment is required")
	}

	if *c.Fee == "" {
		return fmt.Errorf("fee is required")
	}

	if *c.Amount == "" {
		return fmt.Errorf("amount is required")
	}

	return nil
}

// ToPayPayPayinSummary converts the CSV record to a domain model
func (c *PaypayPayinSummaryCSV) ToPayPayPayinSummary(payinFileID int) (*PaypayPayinSummary, error) {
	// Validate required fields
	if err := c.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	var cutoffDate *time.Time
	if *c.CutoffDate != "" {
		t, err := utils.ParseFlexibleDate(*c.CutoffDate)
		if err != nil {
			return nil, err
		}
		cutoffDate = &t
	}

	var paymentDate *time.Time
	if *c.PaymentDate != "" {
		t, err := utils.ParseFlexibleDate(*c.PaymentDate)
		if err != nil {
			return nil, err
		}
		paymentDate = &t
	}

	transactionAmount, err := utils.ParseCSVInt64Field("TransactionAmount", *c.TransactionAmount)
	if err != nil {
		return nil, err
	}
	refundAmount, err := utils.ParseCSVInt64Field("RefundAmount", *c.RefundAmount)
	if err != nil {
		return nil, err
	}
	usageFee, err := utils.ParseCSVInt64Field("UsageFee", *c.UsageFee)
	if err != nil {
		return nil, err
	}
	platformFee, err := utils.ParseCSVInt64Field("PlatformFee", *c.PlatformFee)
	if err != nil {
		return nil, err
	}
	initialFee, err := utils.ParseCSVInt64Field("InitialFee", *c.InitialFee)
	if err != nil {
		return nil, err
	}
	tax, err := utils.ParseCSVInt64Field("Tax", *c.Tax)
	if err != nil {
		return nil, err
	}
	cashback, err := utils.ParseCSVInt64Field("Cashback", *c.Cashback)
	if err != nil {
		return nil, err
	}
	adjustment, err := utils.ParseCSVInt64Field("Adjustment", *c.Adjustment)
	if err != nil {
		return nil, err
	}
	fee, err := utils.ParseCSVInt64Field("Fee", *c.Fee)
	if err != nil {
		return nil, err
	}
	amount, err := utils.ParseCSVInt64Field("Amount", *c.Amount)
	if err != nil {
		return nil, err
	}

	return &PaypayPayinSummary{
		PayinFileID:       payinFileID,
		CorporateName:     *c.CorporateName,
		CutoffDate:        cutoffDate,
		PaymentDate:       paymentDate,
		TransactionAmount: transactionAmount,
		RefundAmount:      refundAmount,
		UsageFee:          usageFee,
		PlatformFee:       platformFee,
		InitialFee:        initialFee,
		Tax:               tax,
		Cashback:          cashback,
		Adjustment:        adjustment,
		Fee:               fee,
		Amount:            amount,
	}, nil
}
