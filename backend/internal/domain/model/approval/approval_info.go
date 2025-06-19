package model

import (
	approvalStageModel "github.com/test-tzs/nomraeite/internal/domain/model/approval_stage"
	approvalWorkflowModel "github.com/test-tzs/nomraeite/internal/domain/model/approval_workflow"
	object "github.com/test-tzs/nomraeite/internal/domain/object/approval"
)

type ApprovalInfo struct {
	Id               int
	ApprovalStatus   object.ApprovalStatus
	ApprovalStages   []approvalStageModel.ApprovalStage
	ApprovalWorkflow approvalWorkflowModel.ApprovalWorkflow
}
