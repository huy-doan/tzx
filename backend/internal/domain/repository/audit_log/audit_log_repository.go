package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/audit_log"
)

// AuditLogRepository defines the interface for audit log operations
type AuditLogRepository interface {
	Create(ctx context.Context, auditLog *model.AuditLog) error
	List(ctx context.Context, filter *model.AuditLogFilter) ([]*model.AuditLog, int, int64, error)
}
