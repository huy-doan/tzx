package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
)

type PayoutRepository interface {
	List(ctx context.Context, filter *model.PayoutFilter) ([]*model.Payout, int, int64, error)
	GetByID(ctx context.Context, id int) (*model.Payout, error)

	GetFirstTransferingPayout(ctx context.Context) (*model.Payout, error)
	Update(ctx context.Context, payout *model.Payout) error
	CreatePayout(ctx context.Context, payout *model.Payout) (*model.Payout, error)
	UpdatePayout(ctx context.Context, payout *model.Payout) (*model.Payout, error)
}
