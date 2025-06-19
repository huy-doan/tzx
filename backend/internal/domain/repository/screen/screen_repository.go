package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/screen"
)

type ScreenRepository interface {
	List(ctx context.Context) ([]*model.Screen, error)
}
