package model

import (
	userModel "github.com/test-tzs/nomraeite/internal/domain/model/user"
	object "github.com/test-tzs/nomraeite/internal/domain/object/approval_stage"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

type ApprovalStageParams struct {
	ID int
	util.BaseColumnTimestamp
	ApprovalID              int
	ApprovalWorkflowStageID int
	ApproverID              int
	Note                    string
	ApprovalResult          object.ApprovalResult
	Approver                *userModel.User
}

type ApprovalStage struct {
	ID int
	util.BaseColumnTimestamp

	ApprovalID              int
	ApprovalWorkflowStageID int
	ApproverID              int
	Note                    string
	ApprovalResult          object.ApprovalResult
	Approver                *userModel.User
}

func NewApprovalStage(params ApprovalStageParams) *ApprovalStage {
	return &ApprovalStage{
		ID:                      params.ID,
		ApprovalID:              params.ApprovalID,
		ApprovalWorkflowStageID: params.ApprovalWorkflowStageID,
		ApproverID:              params.ApproverID,
		Note:                    params.Note,
		ApprovalResult:          params.ApprovalResult,
		Approver:                params.Approver,
		BaseColumnTimestamp:     params.BaseColumnTimestamp,
	}
}

func (a *ApprovalStage) Approve() {
	a.ApprovalResult = object.ApprovalResultApproved
}

func (a *ApprovalStage) Reject() {
	a.ApprovalResult = object.ApprovalResultRejected
}
