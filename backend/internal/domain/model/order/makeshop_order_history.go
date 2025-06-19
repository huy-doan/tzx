package model

import (
	"errors"
	"time"

	orderObject "github.com/test-tzs/nomraeite/internal/domain/object/order"
	paypayObject "github.com/test-tzs/nomraeite/internal/domain/object/paypay"
	"gorm.io/datatypes"

	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

// MakeshopOrderHistory represents the makeshop_order_history table
type MakeshopOrderHistory struct {
	ID int `gorm:"primaryKey;autoIncrement" json:"id"`
	persistence.BaseColumnTimestamp

	ShopID                   string                               `json:"shop_id"`
	SystemOrderNumber        string                               `json:"system_order_number"`
	TransactionAmount        int64                                `json:"transaction_amount"`
	PaymentMethod            string                               `json:"payment_method"`
	OrderedAt                time.Time                            `json:"ordered_at"`
	PaymentTransactionStatus paypayObject.PaypayTransactionStatus `json:"payment_transaction_status"`
	PaymentTransactionID     string                               `json:"payment_transaction_id"`
	MerchantPaymentID        string                               `json:"merchant_payment_id"`
	TargetMonth              datatypes.Date                       `json:"target_month"`
}

func NewMakeshopOrderHistory(order *orderObject.Order, targetMonth time.Time) (*MakeshopOrderHistory, error) {
	var orderedAt time.Time
	orderedAt, err := time.Parse(time.RFC3339, order.OrderedAt)
	if err != nil {
		return nil, errors.New("error parsing OrderedAt: " + err.Error())
	}
	orderStatus := paypayObject.PaypayTransactionStatusFromString(order.PaymentTransactionStatus)
	if orderStatus == nil {
		return nil, errors.New("invalid payment transaction status")
	}

	// validate orderedAt
	if orderedAt.IsZero() || orderedAt.Month() != targetMonth.Month() || orderedAt.Year() != targetMonth.Year() {
		return nil, errors.New("orderedAt does not match the target month")
	}
	// validate transaction amount
	if order.TransactionAmount <= 0 {
		return nil, errors.New("transaction amount must be greater than zero")
	}
	// validate payment method
	if order.PaymentMethod == "" {
		return nil, errors.New("payment method cannot be empty")
	}
	// validate payment transaction ID
	if order.PaypayAttributes.PaymentID == "" {
		return nil, errors.New("payment transaction ID cannot be empty")
	}
	// validate merchant payment ID
	if order.PaypayAttributes.MerchantPaymentID == "" {
		return nil, errors.New("merchant payment ID cannot be empty")
	}
	// validate system order number
	if order.SystemOrderNumber == "" {
		return nil, errors.New("system order number cannot be empty")
	}
	// validate shop ID
	if order.ShopID == "" {
		return nil, errors.New("shop ID cannot be empty")
	}

	// Create MakeshopOrderHistory instance
	history := &MakeshopOrderHistory{
		TransactionAmount:        order.TransactionAmount,
		ShopID:                   order.ShopID,
		OrderedAt:                orderedAt,
		PaymentTransactionStatus: *orderStatus,
		SystemOrderNumber:        order.SystemOrderNumber,
		PaymentMethod:            order.PaymentMethod,
		PaymentTransactionID:     order.PaypayAttributes.PaymentID,
		MerchantPaymentID:        order.PaypayAttributes.MerchantPaymentID,
		TargetMonth:              datatypes.Date(targetMonth),
	}

	return history, nil
}
