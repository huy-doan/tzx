package persistence

import (
	"context"
	"errors"

	model "github.com/test-tzs/nomraeite/internal/domain/model/connected_service_token"
	object "github.com/test-tzs/nomraeite/internal/domain/object/connected_service_token"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/connected_service_token"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/connected_service_token/dto"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"gorm.io/gorm"
)

type ConnectedServiceTokenRepositoryImpl struct {
	db *gorm.DB
}

func NewConnectedServiceTokenPersistence(db *gorm.DB) repository.ConnectedServiceTokenRepository {
	return &ConnectedServiceTokenRepositoryImpl{db: db}
}

func (r *ConnectedServiceTokenRepositoryImpl) FindByServiceName(ctx context.Context, serviceName object.ConnectedServiceName) (*model.ConnectedServiceToken, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var serviceDTO dto.ConnectedServiceToken
	err = db.Where("service_name = ?", serviceName).
		First(&serviceDTO).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return serviceDTO.ToModel(), nil
}

func (r *ConnectedServiceTokenRepositoryImpl) Update(ctx context.Context, service *model.ConnectedServiceToken) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	serviceDTO := dto.FromModel(service)
	return db.Save(serviceDTO).Error
}
