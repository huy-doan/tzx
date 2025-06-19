package repository

import (
	"context"
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/order"
)

type MakeshopOrderHistoryRepository interface {
	BulkInsert(ctx context.Context, history []*model.MakeshopOrderHistory) error
	DeleteByTargetMonth(ctx context.Context, targetMonth time.Time) error
}
