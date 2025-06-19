package service

import (
	"context"
	"errors"

	connectedServiceModel "github.com/test-tzs/nomraeite/internal/domain/model/connected_service_token"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/connected_service_token"
)

type ConnectedServiceTokenDomainService interface {
	RegisterService(ctx context.Context, connectedServiceToken *connectedServiceModel.ConnectedServiceToken) error
}

type connectedServiceTokenDomainService struct {
	connectedServiceTokenRepository repository.ConnectedServiceTokenRepository
}

func NewConnectedServiceTokenDomainService(
	connectedServiceTokenRepository repository.ConnectedServiceTokenRepository,
) ConnectedServiceTokenDomainService {
	return &connectedServiceTokenDomainService{
		connectedServiceTokenRepository: connectedServiceTokenRepository,
	}
}

func (s *connectedServiceTokenDomainService) RegisterService(ctx context.Context, service *connectedServiceModel.ConnectedServiceToken) error {
	existingService, err := s.connectedServiceTokenRepository.FindByServiceName(ctx, service.ServiceName)
	if err != nil {
		return err
	}

	if existingService != nil {
		return errors.New("connected service already exists")
	}

	return s.connectedServiceTokenRepository.Create(ctx, service)
}
