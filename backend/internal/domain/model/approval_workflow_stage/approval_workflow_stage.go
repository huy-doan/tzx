package model

import (
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// ApprovalWorkflowStage represents a stage in an approval workflow
type ApprovalWorkflowStage struct {
	ID int
	util.BaseColumnTimestamp

	ApprovalWorkflowID int
	StageName          string
	Level              int
	PermissionID       int
	Note               string
}
