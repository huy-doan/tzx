package inputdata

import (
	approvalObject "github.com/test-tzs/nomraeite/internal/domain/object/approval"
)

type TransferApprovalInputData struct {
	Action    approvalObject.ApprovalAction
	PayoutID  int
	UserID    int
	Note      string
	UserAgent string
	IPAddress string
}
