package object

import (
	paymentProviderObject "github.com/test-tzs/nomraeite/internal/domain/object/payment_provider"
)

type AttributePaymethod string

const (
	PayPay AttributePaymethod = "PayPay" // PayPay決済
)

func (m AttributePaymethod) Value() string {
	return string(m)
}

func GetAttributePaymethodFromProvider(paymentProviderCode paymentProviderObject.PaymentProviderCode) AttributePaymethod {
	switch paymentProviderCode {
	case paymentProviderObject.PaymentProviderCodePayPay:
		return PayPay
	default:
		return PayPay // Default
	}
}
