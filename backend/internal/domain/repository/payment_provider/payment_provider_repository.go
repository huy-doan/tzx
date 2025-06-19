package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider"
)

type PaymentProviderRepository interface {
	// FindByCode retrieves a payment provider by its code
	FindByCode(ctx context.Context, code string) (*model.PaymentProvider, error)

	// FindByID retrieves a payment provider by its ID
	FindByID(ctx context.Context, id int) (*model.PaymentProvider, error)

	List(ctx context.Context) ([]*model.PaymentProvider, error)
}
