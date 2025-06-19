package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/approval_workflow_stage"
)

type ApprovalWorkflowStageRepository interface {
	GetByID(ctx context.Context, id int) (*model.ApprovalWorkflowStage, error)
	GetApprovalWorkflowStage(ctx context.Context, workflowID int, level int) (*model.ApprovalWorkflowStage, error)
	GetNextStage(ctx context.Context, stage *model.ApprovalWorkflowStage) (*model.ApprovalWorkflowStage, error)
}
