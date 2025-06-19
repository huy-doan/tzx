package model

import (
	approvalWorkflowStageModel "github.com/test-tzs/nomraeite/internal/domain/model/approval_workflow_stage"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// ApprovalWorkflow represents an approval workflow definition
type ApprovalWorkflow struct {
	ID                     int
	ApprovalWorkflowStages []approvalWorkflowStageModel.ApprovalWorkflowStage
	Name                   string
	util.BaseColumnTimestamp
}
