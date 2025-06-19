package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/permission"
)

// PermissionRepository defines the interface for permission data access
type PermissionRepository interface {
	// FindByIDs finds multiple permissions by their IDs
	FindByIDs(ctx context.Context, ids []int) ([]*model.Permission, error)

	// List retrieves all permissions
	List(ctx context.Context) ([]*model.Permission, error)
}
