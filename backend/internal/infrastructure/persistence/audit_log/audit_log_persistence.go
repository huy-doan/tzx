package persistence

import (
	"context"
	"math"

	model "github.com/test-tzs/nomraeite/internal/domain/model/audit_log"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/audit_log"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/audit_log/dto"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"gorm.io/gorm"
)

type AuditLogRepositoryImpl struct {
	db            *gorm.DB
	filterBuilder *persistence.GormFilterBuilder
}

func NewAuditLogRepository(db *gorm.DB) repository.AuditLogRepository {
	return &AuditLogRepositoryImpl{
		db:            db,
		filterBuilder: persistence.NewGormFilterBuilder(),
	}
}

func (r *AuditLogRepositoryImpl) Create(ctx context.Context, auditLog *model.AuditLog) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}
	auditLogDto := dto.ToDTO(auditLog)
	return db.Create(auditLogDto).Error
}

func (r *AuditLogRepositoryImpl) List(ctx context.Context, filter *model.AuditLogFilter) ([]*model.AuditLog, int, int64, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, 0, 0, err
	}

	var auditLogDtos []*dto.AuditLog
	var count int64

	query := db.WithContext(ctx).Model(&dto.AuditLog{})
	query = r.filterBuilder.ApplyBaseFilter(query, &filter.BaseFilter)

	if err := query.Count(&count).Error; err != nil {
		return nil, 0, 0, err
	}
	query = r.filterBuilder.ApplyPagination(query, filter.Pagination)
	query = query.Preload("User")

	if err := query.Find(&auditLogDtos).Error; err != nil {
		return nil, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(count) / float64(filter.Pagination.PageSize)))
	auditLogs := dto.ToModelArray(auditLogDtos)

	return auditLogs, totalPages, count, nil
}
