package service

import (
	"context"

	modelScreen "github.com/test-tzs/nomraeite/internal/domain/model/screen"
	repositoryScreen "github.com/test-tzs/nomraeite/internal/domain/repository/screen"
)

type ScreenService interface {
	ListScreens(ctx context.Context) ([]*modelScreen.Screen, error)
}

type screenServiceImpl struct {
	screenRepository repositoryScreen.ScreenRepository
}

func NewScreenService(
	screenRepository repositoryScreen.ScreenRepository,
) ScreenService {
	return &screenServiceImpl{
		screenRepository: screenRepository,
	}
}

func (s *screenServiceImpl) ListScreens(
	ctx context.Context,
) ([]*modelScreen.Screen, error) {
	return s.screenRepository.List(ctx)
}
