package task

import (
	"context"
	"errors"
	"fmt"

	object "github.com/test-tzs/nomraeite/internal/domain/object/connected_service_token"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/connected_service_token"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
	usecase "github.com/test-tzs/nomraeite/internal/usecase/gmo-aozora"
)

type RefreshGmoAozoraTokenTask struct {
	repo    repository.ConnectedServiceTokenRepository
	usecase usecase.AuthUseCase
	logger  logger.Logger
}

func NewRefreshGmoAozoraTokenTask(
	authUseCase usecase.AuthUseCase,
	repo repository.ConnectedServiceTokenRepository,
	logger logger.Logger,
) *RefreshGmoAozoraTokenTask {
	return &RefreshGmoAozoraTokenTask{
		usecase: authUseCase,
		repo:    repo,
		logger:  logger,
	}
}

func (t *RefreshGmoAozoraTokenTask) Do(ctx context.Context) error {
	service, err := t.repo.FindByServiceName(ctx, object.ConnectedServiceNameGmoAozora)
	if err != nil {
		t.logger.Error("Failed to find Aozora service", map[string]any{
			"error": err.Error(),
		})
		return fmt.Errorf("failed to find Aozora service: %w", err)
	}

	if service == nil {
		t.logger.Error("Aozora service not found", nil)
		return errors.New("aozora service not found")
	}

	if service.RefreshToken == "" {
		t.logger.Error("Refresh token is empty", nil)
		return errors.New("refresh token is empty")
	}

	return t.usecase.RefreshToken(ctx, service)
}
