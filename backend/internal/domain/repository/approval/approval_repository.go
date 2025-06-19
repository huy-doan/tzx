package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/approval"
)

type ApprovalRepository interface {
	GetApprovalByID(ctx context.Context, approvalID int) (*model.ApprovalInfo, error)
	CreateApproval(ctx context.Context, approval *model.Approval) (*model.Approval, error)
	Create(ctx context.Context, approval *model.Approval) error
	Update(ctx context.Context, approval *model.Approval) error
	GetByPayoutID(ctx context.Context, payoutID int) (*model.Approval, error)
}
