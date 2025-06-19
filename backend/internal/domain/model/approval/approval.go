package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/approval"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// ApprovalParams contains parameters for creating a new Approval
type ApprovalParams struct {
	ID int
	util.BaseColumnTimestamp
	ApprovalWorkflowID int
	ApprovalStatus     object.ApprovalStatus
}

// Approval represents the approval entity
type Approval struct {
	ID int
	util.BaseColumnTimestamp

	ApprovalWorkflowID int
	ApprovalStatus     object.ApprovalStatus
}

// NewApproval creates a new approval instance with the given parameters
func NewApproval(params ApprovalParams) *Approval {
	return &Approval{
		ID:                  params.ID,
		ApprovalWorkflowID:  params.ApprovalWorkflowID,
		ApprovalStatus:      params.ApprovalStatus,
		BaseColumnTimestamp: params.BaseColumnTimestamp,
	}
}

func (a *Approval) SetStatus(status object.ApprovalStatus) {
	a.ApprovalStatus = status
}

type ApprovalResult struct {
	PayoutID     int
	ApprovalID   int
	Status       string
	CurrentStage int
	NextStage    *int
	Error        string
}
