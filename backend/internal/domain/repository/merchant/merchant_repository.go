package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	model "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
)

// MerchantRepository defines the interface for merchant data operations
type MerchantRepository interface {
	ListMerchants(ctx context.Context, params *inputdata.MerchantListInputData) ([]*model.Merchant, int, int, error)
	GetByPaymentMerchantIDs(ctx context.Context, paymentMerchantIDs []string) ([]*model.Merchant, error)
	ImportMerchants(ctx context.Context, models []*model.Merchant) ([]*model.Merchant, error)
}
