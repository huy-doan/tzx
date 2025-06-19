package model

import (
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	paymentProviderObject "github.com/test-tzs/nomraeite/internal/domain/object/payment_provider"
)

// PaymentProvider represents the payment_provider domain entity
type PaymentProvider struct {
	ID   int
	Name string
	Code paymentProviderObject.PaymentProviderCode
	util.BaseColumnTimestamp
}
