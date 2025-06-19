package object

type ScreenCode string

const (
	ScreenCodeUserManagementScreen ScreenCode = "USER_MANAGEMENT_SCREEN"
	ScreenCodeAuditLogScreen       ScreenCode = "AUDIT_LOG_SCREEN"
	ScreenCodeDashboardScreen      ScreenCode = "DASHBOARD_SCREEN"
	ScreenCodeRoleManagementScreen ScreenCode = "ROLE_MANAGEMENT_SCREEN"
	ScreenCodeTransferScreen       ScreenCode = "TRANSFER_SCREEN"
	ScreenCodePayinScreen          ScreenCode = "PAYIN_SCREEN"
	ScreenCodePayoutScreen         ScreenCode = "PAYOUT_SCREEN"
	ScreenCodeVirtualAccountScreen ScreenCode = "VIRTUAL_ACCOUNT_SCREEN"
	ScreenCodeMerchantScreen       ScreenCode = "MERCHANT_SCREEN"
)
