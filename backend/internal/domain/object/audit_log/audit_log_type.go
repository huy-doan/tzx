package object

type AuditLogType int

const (
	// User related audit log types
	AuditLogTypeLogin          AuditLogType = 1
	AuditLogTypeLogout         AuditLogType = 2
	AuditLogTypePasswordChange AuditLogType = 3
	AuditLogTypePasswordReset  AuditLogType = 4
	AuditLogTypeUserCreate     AuditLogType = 5
	AuditLogTypeUserUpdate     AuditLogType = 6
	AuditLogTypeUserDelete     AuditLogType = 7
	AuditLogTypeRoleChange     AuditLogType = 8

	// Two-factor authentication related
	AuditLogType2FAEnable  AuditLogType = 9
	AuditLogType2FADisable AuditLogType = 10

	// Payout related audit log types
	AuditLogTypePayoutRequest  AuditLogType = 11
	AuditLogTypePayoutApproval AuditLogType = 12
	AuditLogTypePayoutReject   AuditLogType = 13
	AuditLogTypePayoutResend   AuditLogType = 14
	AuditLogTypePayoutMarkSent AuditLogType = 15

	// Payin related audit log types
	AuditLogTypeManualPayinImport AuditLogType = 16

	// Report related audit log types
	AuditLogTypePayinReportDownload AuditLogType = 17
	AuditLogTypePayinDetailDownload AuditLogType = 18

	// Merchant related audit log types
	AuditLogTypeMerchantStatusUpload AuditLogType = 19

	// API related audit log types
	AuditLogTypeExternalAPIAccess AuditLogType = 20

	// Audit log type verify 2FA
	AuditLogTypeVerify2FA AuditLogType = 21
)

// String returns the string representation of the AuditLogType
func (a AuditLogType) String() string {
	switch a {
	case AuditLogTypeLogin:
		return "ログイン"
	case AuditLogTypeLogout:
		return "ログアウト"
	case AuditLogTypeVerify2FA:
		return "２段階認証"
	case AuditLogTypePasswordChange:
		return "パスワード変更"
	case AuditLogTypePasswordReset:
		return "パスワードリセット"
	case AuditLogTypeUserCreate:
		return "ユーザー作成"
	case AuditLogTypeUserUpdate:
		return "ユーザー編集"
	case AuditLogTypeUserDelete:
		return "ユーザー削除"
	case AuditLogTypeRoleChange:
		return "ロール変更"
	case AuditLogType2FAEnable:
		return "２段階認証有効"
	case AuditLogType2FADisable:
		return "２段階認証無効"
	case AuditLogTypePayoutRequest:
		return "出金申請"
	case AuditLogTypePayoutApproval:
		return "出金承認"
	case AuditLogTypePayoutReject:
		return "出金却下"
	case AuditLogTypePayoutResend:
		return "振込再依頼"
	case AuditLogTypePayoutMarkSent:
		return "振込を送金済みとする"
	case AuditLogTypeManualPayinImport:
		return "手動入金取り込み"
	case AuditLogTypePayinReportDownload:
		return "入金レポートをダウンロード"
	case AuditLogTypePayinDetailDownload:
		return "入金明細をダウンロード"
	case AuditLogTypeMerchantStatusUpload:
		return "加盟店審査状況をアップロード"
	case AuditLogTypeExternalAPIAccess:
		return "外部APIアクセス"
	default:
		return "不明な監査ログタイプ"
	}
}
