package persistence

import (
	"context"
	"math"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/convert"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"gorm.io/gorm"
)

type PayoutRepositoryImpl struct {
	db            *gorm.DB
	filterBuilder *persistence.GormFilterBuilder
}

func NewPayoutRepository(db *gorm.DB) repository.PayoutRepository {
	return &PayoutRepositoryImpl{
		db:            db,
		filterBuilder: persistence.NewGormFilterBuilder(),
	}
}

func (r *PayoutRepositoryImpl) List(ctx context.Context, filter *model.PayoutFilter) ([]*model.Payout, int, int64, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	var payoutDTOs []*dto.Payout
	var count int64

	if filter != nil {
		filter.ApplyFilters()
	} else {
		filter = model.NewPayoutFilter()
	}

	query := db.WithContext(ctx).Model(&dto.Payout{})

	query = r.filterBuilder.ApplyBaseFilter(query, &filter.BaseFilter)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, 0, err
	}

	query = r.filterBuilder.ApplyPagination(query, filter.Pagination)

	query = query.Preload("User").Unscoped()

	if err := query.Find(&payoutDTOs).Error; err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(filter.Pagination.PageSize)))

	payouts := convert.ToPayoutModels(payoutDTOs)

	return payouts, totalPages, int64(count), nil
}

func (r *PayoutRepositoryImpl) GetFirstTransferingPayout(ctx context.Context) (*model.Payout, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var payoutDTO dto.Payout

	err = db.Model(&dto.Payout{}).
		Where("payout_status = ?", object.PayoutStatusWaitingTransfer).
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

func (r *PayoutRepositoryImpl) Update(ctx context.Context, payout *model.Payout) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}

	payoutDTO := convert.ToPayoutDTO(payout)

	if err := db.
		Model(&dto.Payout{}).
		Where("id = ?", payout.ID).
		Updates(payoutDTO).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *PayoutRepositoryImpl) GetByID(ctx context.Context, id int) (*model.Payout, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var payoutDTO dto.Payout

	if err := db.
		Model(&dto.Payout{}).
		Where("id = ?", id).
		First(&payoutDTO).
		Error; err != nil {
		return nil, err
	}

	return convert.ToPayoutModel(&payoutDTO), nil
}

// CreatePayout creates a new payout record
func (p *PayoutRepositoryImpl) CreatePayout(ctx context.Context, payout *model.Payout) (*model.Payout, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	payoutDTO := dto.FromModel(payout)
	if err := db.WithContext(ctx).Create(payoutDTO).Error; err != nil {
		return nil, err
	}

	return payoutDTO.ToModel(), nil
}

// UpdatePayout updates an existing payout record
func (p *PayoutRepositoryImpl) UpdatePayout(ctx context.Context, payout *model.Payout) (*model.Payout, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	payoutDTO := dto.FromModel(payout)
	if err := db.WithContext(ctx).Save(payoutDTO).Error; err != nil {
		return nil, err
	}

	return payoutDTO.ToModel(), nil
}
