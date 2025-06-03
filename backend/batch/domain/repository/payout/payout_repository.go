package repository

import (
	"context"

	model "github.com/makeshop-jp/master-console/internal/domain/model/payout"
)

type PayoutRepository interface {
	GetFirstTransferingPayout(ctx context.Context) (*model.Payout, error)
	UpdateByID(ctx context.Context, id int, payout *model.Payout) error
}
