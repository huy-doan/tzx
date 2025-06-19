package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	model "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider_review"
)

type PaymentProviderReviewRepository interface {
	GetShopReviewList(ctx context.Context, input *inputdata.PaymentProviderReviewListInput) (*model.PaymentProviderReviewListResult, error)
	GetLastReviewByShopIDs(ctx context.Context, shopIDs []string) (map[string]*model.PaymentProviderReview, error)
	Import(ctx context.Context, models []*model.PaymentProviderReview) ([]*model.PaymentProviderReview, error)
}
