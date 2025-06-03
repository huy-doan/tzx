package dto

import (
	"errors"
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/connected_service_token"
	modelUtil "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/connected_service_token"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
	config "github.com/test-tzs/nomraeite/internal/pkg/config"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
	"gorm.io/gorm"
)

type ConnectedServiceToken struct {
	ID int
	persistence.BaseColumnTimestamp

	ServiceName          object.ConnectedServiceName
	UserID               int
	RefreshToken         string
	AccessToken          string
	AccessTokenExpiredAt *time.Time
}

func (ConnectedServiceToken) TableName() string {
	return "connected_service_token"
}

func (dto *ConnectedServiceToken) BeforeSave(*gorm.DB) (err error) {
	cfg := config.GetConfig()
	if cfg.EncryptionKey == "" {
		return errors.New("encryption key is not configured")
	}
	refreshToken := dto.RefreshToken
	if refreshToken != "" {
		dto.RefreshToken, err = utils.EncryptAESGCMString(cfg.EncryptionKey, refreshToken)
		if err != nil {
			return err
		}
	}
	accessToken := dto.AccessToken
	if accessToken != "" {
		dto.AccessToken, err = utils.EncryptAESGCMString(cfg.EncryptionKey, accessToken)
		if err != nil {
			return err
		}
	}

	return
}

func (dto *ConnectedServiceToken) AfterFind(*gorm.DB) (err error) {
	cfg := config.GetConfig()
	if cfg.EncryptionKey == "" {
		return errors.New("encryption key is not configured")
	}

	refreshToken := dto.RefreshToken
	if refreshToken != "" {
		dto.RefreshToken, err = utils.DecryptAESGCMString(cfg.EncryptionKey, refreshToken)
		if err != nil {
			return err
		}
	}

	accessToken := dto.AccessToken
	if accessToken != "" {
		dto.AccessToken, err = utils.DecryptAESGCMString(cfg.EncryptionKey, accessToken)
		if err != nil {
			return err
		}
	}

	return
}

func (dto *ConnectedServiceToken) ToModel() *model.ConnectedServiceToken {
	return &model.ConnectedServiceToken{
		ID:                   dto.ID,
		ServiceName:          dto.ServiceName,
		UserID:               dto.UserID,
		RefreshToken:         dto.RefreshToken,
		AccessToken:          dto.AccessToken,
		AccessTokenExpiredAt: dto.AccessTokenExpiredAt,
		BaseColumnTimestamp: modelUtil.BaseColumnTimestamp{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
		},
	}
}

func FromModel(m *model.ConnectedServiceToken) *ConnectedServiceToken {
	if m == nil {
		return nil
	}

	return &ConnectedServiceToken{
		ID:                   m.ID,
		ServiceName:          m.ServiceName,
		UserID:               m.UserID,
		RefreshToken:         m.RefreshToken,
		AccessToken:          m.AccessToken,
		AccessTokenExpiredAt: m.AccessTokenExpiredAt,
		BaseColumnTimestamp: persistence.BaseColumnTimestamp{
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
		},
	}
}
