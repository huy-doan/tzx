package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/approval_stage"
)

type ApprovalStageRepository interface {
	Create(ctx context.Context, stage *model.ApprovalStage) error
	GetCurrentStageByApprovalID(ctx context.Context, approvalID int) (*model.ApprovalStage, error)
}
