package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payin"
)

type PayinFileGroupRepository interface {
	Create(ctx context.Context, group *model.PayinFileGroup) (*model.PayinFileGroup, error)
	ListPayinFileGroups(ctx context.Context, params *inputdata.PayinFileGroupListInputData) (*model.PaginatedPayinFileGroupResult, error)
}
