package object

type ApprovalAction string

const (
	ApprovalActionApproved ApprovalAction = "approved"
	ApprovalActionRejected ApprovalAction = "rejected"
)

func (a ApprovalAction) IsValid() bool {
	return a == ApprovalActionApproved || a == ApprovalActionRejected
}

func (a ApprovalAction) String() string {
	return string(a)
}
