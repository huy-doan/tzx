package service

import (
	"context"

	auditLogModel "github.com/test-tzs/nomraeite/internal/domain/model/audit_log"
	userModel "github.com/test-tzs/nomraeite/internal/domain/model/user"
	auditLogRepository "github.com/test-tzs/nomraeite/internal/domain/repository/audit_log"
	userRepository "github.com/test-tzs/nomraeite/internal/domain/repository/user"
)

type AuditLogService interface {
	CreateAuditLog(ctx context.Context, auditLog *auditLogModel.AuditLog) error
	GetAuditLogs(ctx context.Context, filter *auditLogModel.AuditLogFilter) ([]*auditLogModel.AuditLog, int, int64, error)
	GetUsersWithAuditLogs(ctx context.Context) ([]*userModel.User, error)
}

type auditLogServiceImpl struct {
	auditLogRepository auditLogRepository.AuditLogRepository
	userRepository     userRepository.UserRepository
}

func NewAuditLogService(auditLogRepository auditLogRepository.AuditLogRepository, userRepository userRepository.UserRepository) AuditLogService {
	return &auditLogServiceImpl{
		auditLogRepository: auditLogRepository,
		userRepository:     userRepository,
	}
}

func (s *auditLogServiceImpl) CreateAuditLog(ctx context.Context, auditLog *auditLogModel.AuditLog) error {
	return s.auditLogRepository.Create(ctx, auditLog)
}

func (s *auditLogServiceImpl) GetAuditLogs(ctx context.Context, filter *auditLogModel.AuditLogFilter) ([]*auditLogModel.AuditLog, int, int64, error) {
	auditLogs, totalPages, totalCount, err := s.auditLogRepository.List(ctx, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	return auditLogs, totalPages, totalCount, nil
}

func (s *auditLogServiceImpl) GetUsersWithAuditLogs(ctx context.Context) ([]*userModel.User, error) {
	return s.userRepository.GetUsersWithAuditLogs(ctx)
}
