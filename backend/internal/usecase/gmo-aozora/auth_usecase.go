package usecase

import (
	"context"
	"fmt"

	adapter "github.com/makeshop-jp/master-console/internal/domain/adapter/gmo-aozora"
	model "github.com/makeshop-jp/master-console/internal/domain/model/connected_service_token"
	repository "github.com/makeshop-jp/master-console/internal/domain/repository/connected_service_token"
	"github.com/makeshop-jp/master-console/internal/pkg/logger"
)

type AuthUseCase interface {
	RefreshToken(ctx context.Context, connectedServiceToken *model.ConnectedServiceToken) error
}

type authUseCaseImpl struct {
	repo   repository.ConnectedServiceTokenRepository
	client adapter.ApiClient
	logger logger.Logger
}

func NewAuthUseCase(
	repo repository.ConnectedServiceTokenRepository,
	client adapter.ApiClient,
	logger logger.Logger,
) AuthUseCase {
	return &authUseCaseImpl{
		repo:   repo,
		client: client,
		logger: logger,
	}
}

func (s *authUseCaseImpl) RefreshToken(ctx context.Context, connectedServiceToken *model.ConnectedServiceToken) error {
	response, err := s.client.RefreshToken(ctx, connectedServiceToken.RefreshToken)
	if err != nil {
		s.logger.Error("Failed to refresh Aozora token", map[string]any{
			"error": err.Error(),
		})
		return fmt.Errorf("failed to refresh Aozora token: %w", err)
	}

	connectedServiceToken.SetGmoAozoraToken(response)
	return s.repo.Update(ctx, connectedServiceToken)
}
