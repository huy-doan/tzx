package model

// AuditLogDetails represents the structured details of an audit log entry
type AuditLogDetails struct {
	AuditLogUserID int    `json:"audit_log_user_id,omitzero"`
	TransactionID  int    `json:"transaction_id,omitzero"`
	PayoutID       int    `json:"payout_id,omitzero"`
	PayinID        int    `json:"payin_id,omitzero"`
	OldRole        string `json:"old_role,omitempty"`
	NewRole        string `json:"new_role,omitempty"`
}

// NewAuditLogDetails creates a new instance of AuditLogDetails
func NewAuditLogDetails() *AuditLogDetails {
	return &AuditLogDetails{}
}
