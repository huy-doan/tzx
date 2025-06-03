package persistence

import (
	"context"

	repository "github.com/test-tzs/nomraeite/batch/domain/repository/payout"
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/convert"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"gorm.io/gorm"
)

type PayoutRepositoryImpl struct {
	db *gorm.DB
}

func NewPayoutPersistence(db *gorm.DB) repository.PayoutRepository {
	return &PayoutRepositoryImpl{
		db: db,
	}
}

func (r *PayoutRepositoryImpl) GetFirstTransferingPayout(ctx context.Context) (*model.Payout, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var payoutDTO dto.Payout

	err = db.Model(&dto.Payout{}).
		Where(dto.Payout{
			PayoutStatus: object.PayoutStatusWaitingTransfer.Value(),
		}).
		Order("id ASC").
		First(&payoutDTO).
		Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return convert.ToPayoutModel(&payoutDTO), nil
}

func (r *PayoutRepositoryImpl) UpdateByID(ctx context.Context, id int, payout *model.Payout) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	payoutDTO := convert.ToPayoutDTO(payout)
	err = db.Model(&dto.Payout{}).
		Where("id = ?", id).
		Updates(&payoutDTO).Error

	if err != nil {
		return err
	}

	return nil
}
