package repository

import (
	"context"

	model "github.com/makeshop-jp/master-console/internal/domain/model/connected_service_token"
	object "github.com/makeshop-jp/master-console/internal/domain/object/connected_service_token"
)

type ConnectedServiceTokenRepository interface {
	FindByServiceName(ctx context.Context, serviceName object.ConnectedServiceName) (*model.ConnectedServiceToken, error)
	Update(ctx context.Context, service *model.ConnectedServiceToken) error
}
