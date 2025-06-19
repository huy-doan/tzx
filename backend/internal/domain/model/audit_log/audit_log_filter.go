package model

import (
	"fmt"
	"time"

	filterModel "github.com/test-tzs/nomraeite/internal/domain/model/filter"
)

type AuditLogFilter struct {
	filterModel.BaseFilter
	UserID       *int
	AuditLogType *int
	CreatedAt    *time.Time
	Description  *string
}

func NewAuditLogFilter() *AuditLogFilter {
	filter := &AuditLogFilter{}
	filter.ValidSortFields = map[string]bool{
		"id":          true,
		"created_at":  true,
		"description": true,
		"user_id":     true,
		"ip_address":  true,
		"user_agent":  true,
	}

	filter.SetPagination(1, 10)

	return filter
}

func (f *AuditLogFilter) ApplyFilters() {
	f.AddDateFilter("created_at", filterModel.Equal, f.CreatedAt)
	if f.UserID != nil {
		f.AddCondition("user_id", filterModel.Equal, f.UserID)
	}
	if f.AuditLogType != nil {
		f.AddCondition("audit_log_type", filterModel.Equal, f.AuditLogType)
	}
	if f.Description != nil {
		f.AddCondition("description", filterModel.Like, fmt.Sprintf("%%%s%%", *f.Description))
	}
}
