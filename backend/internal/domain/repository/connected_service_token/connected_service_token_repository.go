package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/connected_service_token"
	object "github.com/test-tzs/nomraeite/internal/domain/object/connected_service_token"
)

type ConnectedServiceTokenRepository interface {
	FindByServiceName(ctx context.Context, serviceName object.ConnectedServiceName) (*model.ConnectedServiceToken, error)
	Update(ctx context.Context, service *model.ConnectedServiceToken) error
	Create(ctx context.Context, service *model.ConnectedServiceToken) error
	RevokeByServiceName(ctx context.Context, serviceName object.ConnectedServiceName) error
}
