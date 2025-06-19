package object

type ApprovalResult int

const (
	ApprovalResultApproved ApprovalResult = 1 // 承認
	ApprovalResultRejected ApprovalResult = 2 // 却下
)

// IsApproved checks if the stage is approved
func (a ApprovalResult) IsApproved() bool {
	return a == ApprovalResultApproved
}

// IsRejected checks if the stage is rejected
func (a ApprovalResult) IsRejected() bool {
	return a == ApprovalResultRejected
}

// String returns the string representation of the approval result
func (a ApprovalResult) String() string {
	switch a {
	case ApprovalResultApproved:
		return "承認"
	case ApprovalResultRejected:
		return "却下"
	default:
		return "不明"
	}
}
