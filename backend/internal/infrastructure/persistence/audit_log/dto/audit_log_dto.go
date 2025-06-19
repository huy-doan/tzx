package dto

import (
	"encoding/json"
	"time"

	auditLogModel "github.com/test-tzs/nomraeite/internal/domain/model/audit_log"
	auditLogObject "github.com/test-tzs/nomraeite/internal/domain/object/audit_log"
	userDto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/user/dto"
	"gorm.io/datatypes"
)

type AuditLog struct {
	ID           int            `gorm:"column:id;primaryKey"`
	UserID       *int           `gorm:"column:user_id"`
	AuditLogType int            `gorm:"column:audit_log_type"`
	Description  string         `gorm:"column:description"`
	UserAgent    *string        `gorm:"column:user_agent"`
	IPAddress    *string        `gorm:"column:ip_address"`
	Details      datatypes.JSON `gorm:"column:details"`
	CreatedAt    time.Time      `gorm:"column:created_at"`

	// Relationships
	User *userDto.User `gorm:"foreignKey:UserID;references:ID"`
}

func (al *AuditLog) TableName() string {
	return "audit_log"
}

func (al *AuditLog) ToModel() *auditLogModel.AuditLog {
	var auditLogDetails auditLogModel.AuditLogDetails
	err := json.Unmarshal(al.Details, &auditLogDetails)
	if err != nil {
		panic(err)
	}

	auditLog := &auditLogModel.AuditLog{
		ID:           al.ID,
		UserID:       al.UserID,
		AuditLogType: auditLogObject.AuditLogType(al.AuditLogType),
		Description:  al.Description,
		UserAgent:    al.UserAgent,
		IPAddress:    al.IPAddress,
		Details:      &auditLogDetails,
		CreatedAt:    al.CreatedAt,
	}

	if al.User != nil {
		auditLog.User = al.User.ToUserModel()
	}

	return auditLog
}

func ToDTO(auditLog *auditLogModel.AuditLog) *AuditLog {
	if auditLog == nil {
		return nil
	}

	details, err := json.Marshal(auditLog.Details)
	if err != nil {
		panic(err)
	}

	auditLogDto := &AuditLog{
		ID:           auditLog.ID,
		UserID:       auditLog.UserID,
		AuditLogType: int(auditLog.AuditLogType),
		Description:  auditLog.Description,
		UserAgent:    auditLog.UserAgent,
		IPAddress:    auditLog.IPAddress,
		Details:      details,
		CreatedAt:    auditLog.CreatedAt,
	}

	if auditLog.User != nil {
		auditLogDto.User = userDto.ToUserDTO(auditLog.User)
	}

	return auditLogDto
}

func ToDTOArray(auditLogs []*auditLogModel.AuditLog) []*AuditLog {
	if auditLogs == nil {
		return nil
	}

	dtos := make([]*AuditLog, len(auditLogs))
	for i, auditLog := range auditLogs {
		dtos[i] = ToDTO(auditLog)
	}
	return dtos
}

func ToModelArray(dtos []*AuditLog) []*auditLogModel.AuditLog {
	if dtos == nil {
		return nil
	}

	auditLogs := make([]*auditLogModel.AuditLog, len(dtos))
	for i, dto := range dtos {
		auditLogs[i] = dto.ToModel()
	}
	return auditLogs
}
