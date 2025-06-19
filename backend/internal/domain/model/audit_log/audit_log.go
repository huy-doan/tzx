package model

import (
	"fmt"
	"time"

	userModel "github.com/test-tzs/nomraeite/internal/domain/model/user"
	objectApproval "github.com/test-tzs/nomraeite/internal/domain/object/approval"
	object "github.com/test-tzs/nomraeite/internal/domain/object/audit_log"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// Description constants for audit log events
const (
	// User-related descriptions
	DescLogin          = "ログインしました。"
	DescLogout         = "ログアウトしました。"
	DescPasswordChange = "パスワード変更しました。"
	DescPasswordReset  = "ユーザー（%d）のパスワードパスワードリセット。"
	DescUserCreate     = "ユーザー（%d）を新規登録しました。"
	DescUserUpdate     = "ユーザー（%d）の基本情報を変更しました。"
	DescUserDelete     = "ユーザー（%d）を削除しました。"
	DescRoleChange     = "ユーザー（%d）を「%s」ロールに変更しました。"

	// Two-factor authentication descriptions
	Desc2FAEnable  = "ユーザー（%d）の２段階認証を有効しました。"
	Desc2FADisable = "ユーザー（%d）の２段階認証を無効しました。"

	Desc2FAVerification = "ユーザー（#%d）の２段階認証が%sしました。"

	// Payout-related descriptions
	DescPayoutRequest  = "ユーザー（#%d）は振込依頼データ（#%d）を作成しました。"
	DescPayoutApproval = "出金承認しました。"
	DescPayoutReject   = "出金却下しました。"
	DescPayoutResend   = "振込再依頼を行いました。"
	DescPayoutMarkSent = "振込データを送金済みとしました。"

	// Transfer approval-related descriptions
	DescTransferApproval = "ユーザー#%dは振込依頼データ#%dを承認しました。"
	DescTransferReject   = "ユーザー#%dは振込依頼データ#%dを却下しました。"

	// Payin-related descriptions
	DescManualPayinImport = "手動入金取り込みを行いました。"

	// Other descriptions
	DescMerchantStatusUpload = "加盟店審査状況をアップロードしました。"
	DescExternalAPIAccess    = "振込APIを実行しました。"
)

type AuditLog struct {
	ID           int
	UserID       *int
	AuditLogType object.AuditLogType
	Description  string
	UserAgent    *string
	IPAddress    *string
	Details      *AuditLogDetails
	CreatedAt    time.Time
	// User information loaded from the User table
	User *userModel.User
}

type AuditLogUserResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type AuditLogUsersResponse struct {
	Users []*AuditLogUserResponse `json:"users"`
}

type AuditLogResponse struct {
	ID            int               `json:"id"`
	UserID        *int              `json:"user_id"`
	User          *AuditLogUserInfo `json:"user,omitempty"`
	AuditLogType  string            `json:"audit_log_type"`
	Description   string            `json:"description"`
	TransactionID int               `json:"transaction_id,omitzero"`
	PayoutID      int               `json:"payout_id,omitzero"`
	PayinID       int               `json:"payin_id,omitzero"`
	UserAgent     string            `json:"user_agent"`
	IPAddress     string            `json:"ip_address"`
	CreatedAt     string            `json:"created_at"`
}

type AuditLogUserInfo struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type AuditLogListResponse struct {
	AuditLogs  []*AuditLogResponse `json:"audit_logs"`
	TotalPages int                 `json:"total_pages"`
	Total      int64               `json:"total"`
	Page       int                 `json:"page"`
	PageSize   int                 `json:"page_size"`
}

type AuditLogGenerator struct {
	UserID            int
	AuditLogUserID    int
	AuditLogType      object.AuditLogType
	Description       string
	UserAgent         *object.UserAgent
	IPAddress         *object.IPAddress
	CreatedAt         time.Time
	OldRole           *string
	NewRole           *string
	TransactionID     int
	PayoutID          int
	PayinID           int
	IsActionSuccessed bool
	Payload           *AuditLogPayload
}

type AuditLogPayload struct {
	PayoutID int
	Action   string
	StageID  int
	// add more fields here in the future for other audit log types
}

// NewAuditLogGenerator creates a new generator with required base fields
func NewAuditLogGenerator(
	userID int,
	auditLogType object.AuditLogType,
	ipAddress *object.IPAddress,
	userAgent *object.UserAgent,
) *AuditLogGenerator {
	return &AuditLogGenerator{
		UserID:       userID,
		AuditLogType: auditLogType,
		IPAddress:    ipAddress,
		UserAgent:    userAgent,
	}
}

// Map of audit log types to their corresponding description templates
var descriptionTemplates = map[object.AuditLogType]string{
	object.AuditLogTypeLogin:                DescLogin,
	object.AuditLogTypeLogout:               DescLogout,
	object.AuditLogTypePasswordChange:       DescPasswordChange,
	object.AuditLogTypePasswordReset:        DescPasswordReset,
	object.AuditLogTypeUserCreate:           DescUserCreate,
	object.AuditLogTypeUserUpdate:           DescUserUpdate,
	object.AuditLogTypeUserDelete:           DescUserDelete,
	object.AuditLogTypeRoleChange:           DescRoleChange,
	object.AuditLogType2FAEnable:            Desc2FAEnable,
	object.AuditLogType2FADisable:           Desc2FADisable,
	object.AuditLogTypePayoutRequest:        DescPayoutRequest,
	object.AuditLogTypePayoutApproval:       DescPayoutApproval,
	object.AuditLogTypePayoutReject:         DescPayoutReject,
	object.AuditLogTypePayoutResend:         DescPayoutResend,
	object.AuditLogTypePayoutMarkSent:       DescPayoutMarkSent,
	object.AuditLogTypeManualPayinImport:    DescManualPayinImport,
	object.AuditLogTypeMerchantStatusUpload: DescMerchantStatusUpload,
	object.AuditLogTypeExternalAPIAccess:    DescExternalAPIAccess,
	object.AuditLogTypeVerify2FA:            Desc2FAVerification,
}

// getDescription returns the appropriate description based on the audit log type
func (g *AuditLogGenerator) getDescription() string {
	template, exists := descriptionTemplates[g.AuditLogType]
	if !exists {
		return ""
	}

	switch g.AuditLogType {
	case object.AuditLogTypePasswordReset,
		object.AuditLogTypeUserCreate,
		object.AuditLogTypeUserUpdate,
		object.AuditLogTypeUserDelete:
		if g.AuditLogUserID != 0 {
			return fmt.Sprintf(template, g.AuditLogUserID)
		}
	case object.AuditLogType2FAEnable,
		object.AuditLogType2FADisable:
		if g.UserID != 0 {
			return fmt.Sprintf(template, g.UserID)
		}
	case object.AuditLogTypeRoleChange:
		if g.AuditLogUserID != 0 && g.NewRole != nil {
			return fmt.Sprintf(template, g.AuditLogUserID, *g.NewRole)
		}
	case object.AuditLogTypeVerify2FA:
		if g.UserID != 0 {
			return fmt.Sprintf(template, g.UserID, func() string {
				if g.IsActionSuccessed {
					return "成功"
				}

				return "失敗"
			}())
		}

	case object.AuditLogTypePayoutRequest:
		if g.PayoutID != 0 {
			return fmt.Sprintf(template, g.UserID, g.PayoutID)
		}
	case object.AuditLogTypePayoutApproval, object.AuditLogTypePayoutReject:
		if g.Payload != nil && g.Payload.PayoutID != 0 {
			if g.Payload.Action == string(objectApproval.ApprovalActionApproved) {
				return fmt.Sprintf(template, g.UserID, g.Payload.PayoutID)
			}

			if g.Payload.Action == string(objectApproval.ApprovalActionRejected) {
				return fmt.Sprintf(template, g.UserID, g.Payload.PayoutID)
			}

			return fmt.Sprintf(template, g.UserID, g.Payload.PayoutID)
		}
	default:
		return template
	}

	return ""
}

func (g *AuditLogGenerator) Generate() *AuditLog {
	if g.Description == "" {
		defaultDesc := g.getDescription()
		g.Description = defaultDesc
	}

	details := &AuditLogDetails{}
	userId := g.UserID

	details.AuditLogUserID = g.AuditLogUserID
	details.TransactionID = g.TransactionID
	details.PayoutID = g.PayoutID
	details.PayinID = g.PayinID

	if g.NewRole != nil {
		details.NewRole = *g.NewRole
	}

	if g.OldRole != nil {
		details.OldRole = *g.OldRole
	}

	var ipAddressStr *string
	if g.IPAddress != nil {
		ipStr := g.IPAddress.String()
		ipAddressStr = &ipStr
	}

	var userAgentStr *string
	if g.UserAgent != nil {
		uaStr := g.UserAgent.String()
		userAgentStr = &uaStr
	}

	return &AuditLog{
		UserID:       &userId,
		AuditLogType: g.AuditLogType,
		Description:  g.Description,
		UserAgent:    userAgentStr,
		IPAddress:    ipAddressStr,
		Details:      details,
		CreatedAt:    g.CreatedAt,
	}
}

func NewAuditLog() *AuditLog {
	return &AuditLog{
		Details:   &AuditLogDetails{},
		CreatedAt: time.Now(),
	}
}

type NewAuditLogParams struct {
	UserID        *int                `json:"user_id"`
	User          *userModel.User     `json:"user"`
	AuditLogType  object.AuditLogType `json:"audit_log_type"`
	Description   *string             `json:"description"`
	TransactionID *int                `json:"transaction_id"`
	PayoutID      *int                `json:"payout_id"`
	PayinID       *int                `json:"payin_id"`
	UserAgent     *object.UserAgent   `json:"user_agent"`
	IPAddress     *object.IPAddress   `json:"ip_address"`
	util.BaseColumnTimestamp
}
