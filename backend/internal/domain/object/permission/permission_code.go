package object

type PermissionCode string

const (
	// User Management Screen Permissions
	PermissionCodeUserManagementScreenView   PermissionCode = "USER_MANAGEMENT_SCREEN_VIEW"
	PermissionCodeUserManagementScreenAdd    PermissionCode = "USER_MANAGEMENT_SCREEN_ADD"
	PermissionCodeUserManagementScreenEdit   PermissionCode = "USER_MANAGEMENT_SCREEN_EDIT"
	PermissionCodeUserManagementScreenDelete PermissionCode = "USER_MANAGEMENT_SCREEN_DELETE"

	// Audit Log Screen Permissions
	PermissionCodeAuditLogScreenView   PermissionCode = "AUDIT_LOG_SCREEN_VIEW"
	PermissionCodeAuditLogScreenAdd    PermissionCode = "AUDIT_LOG_SCREEN_ADD"
	PermissionCodeAuditLogScreenEdit   PermissionCode = "AUDIT_LOG_SCREEN_EDIT"
	PermissionCodeAuditLogScreenDelete PermissionCode = "AUDIT_LOG_SCREEN_DELETE"

	// Dashboard Screen Permissions
	PermissionCodeDashboardScreenView   PermissionCode = "DASHBOARD_SCREEN_VIEW"
	PermissionCodeDashboardScreenAdd    PermissionCode = "DASHBOARD_SCREEN_ADD"
	PermissionCodeDashboardScreenEdit   PermissionCode = "DASHBOARD_SCREEN_EDIT"
	PermissionCodeDashboardScreenDelete PermissionCode = "DASHBOARD_SCREEN_DELETE"

	// Role Management Screen Permissions
	PermissionCodeRoleManagementScreenView   PermissionCode = "ROLE_MANAGEMENT_SCREEN_VIEW"
	PermissionCodeRoleManagementScreenAdd    PermissionCode = "ROLE_MANAGEMENT_SCREEN_ADD"
	PermissionCodeRoleManagementScreenEdit   PermissionCode = "ROLE_MANAGEMENT_SCREEN_EDIT"
	PermissionCodeRoleManagementScreenDelete PermissionCode = "ROLE_MANAGEMENT_SCREEN_DELETE"

	// Transfer Screen Permissions
	PermissionCodeTransferScreenView   PermissionCode = "TRANSFER_SCREEN_VIEW"
	PermissionCodeTransferScreenAdd    PermissionCode = "TRANSFER_SCREEN_ADD"
	PermissionCodeTransferScreenEdit   PermissionCode = "TRANSFER_SCREEN_EDIT"
	PermissionCodeTransferScreenDelete PermissionCode = "TRANSFER_SCREEN_DELETE"

	// Payin Screen Permissions
	PermissionCodePayinScreenView   PermissionCode = "PAYIN_SCREEN_VIEW"
	PermissionCodePayinScreenAdd    PermissionCode = "PAYIN_SCREEN_ADD"
	PermissionCodePayinScreenEdit   PermissionCode = "PAYIN_SCREEN_EDIT"
	PermissionCodePayinScreenDelete PermissionCode = "PAYIN_SCREEN_DELETE"

	// Payout Screen Permissions
	PermissionCodePayoutScreenView   PermissionCode = "PAYOUT_SCREEN_VIEW"
	PermissionCodePayoutScreenAdd    PermissionCode = "PAYOUT_SCREEN_ADD"
	PermissionCodePayoutScreenEdit   PermissionCode = "PAYOUT_SCREEN_EDIT"
	PermissionCodePayoutScreenDelete PermissionCode = "PAYOUT_SCREEN_DELETE"

	// Virtual Account Screen Permissions
	PermissionCodeVirtualAccountScreenView   PermissionCode = "VIRTUAL_ACCOUNT_SCREEN_VIEW"
	PermissionCodeVirtualAccountScreenAdd    PermissionCode = "VIRTUAL_ACCOUNT_SCREEN_ADD"
	PermissionCodeVirtualAccountScreenEdit   PermissionCode = "VIRTUAL_ACCOUNT_SCREEN_EDIT"
	PermissionCodeVirtualAccountScreenDelete PermissionCode = "VIRTUAL_ACCOUNT_SCREEN_DELETE"

	// Merchant Screen Permissions
	PermissionCodeMerchantScreenView   PermissionCode = "MERCHANT_SCREEN_VIEW"
	PermissionCodeMerchantScreenAdd    PermissionCode = "MERCHANT_SCREEN_ADD"
	PermissionCodeMerchantScreenEdit   PermissionCode = "MERCHANT_SCREEN_EDIT"
	PermissionCodeMerchantScreenDelete PermissionCode = "MERCHANT_SCREEN_DELETE"

	// GMO AOZORA Connect Permissions
	PermissionCodeGmoAozoraScreenView PermissionCode = "GMO_AOZORA_CONNECT"

	// Approval Screen Permissions
	PermissionCodeApprovalTransferFirstStage  PermissionCode = "APPROVAL_TRANSFER_FIRST_STAGE"
	PermissionCodeApprovalTransferSecondStage PermissionCode = "APPROVAL_TRANSFER_SECOND_STAGE"
)

func (p PermissionCode) Value() string {
	return string(p)
}
