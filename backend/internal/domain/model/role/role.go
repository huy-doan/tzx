package model

import (
	permission "github.com/test-tzs/nomraeite/internal/domain/model/permission"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	permissionObject "github.com/test-tzs/nomraeite/internal/domain/object/permission"
	"github.com/test-tzs/nomraeite/internal/pkg/api/generated"
)

// Role represents a user role in the system
type Role struct {
	ID   int
	Name string
	util.BaseColumnTimestamp

	UserCount int64

	// Relationships
	Permissions []*permission.Permission
}

type CountResult struct {
	RoleID    int
	UserCount int64
}

// HasPermission checks if the role has the specified permission
func (r *Role) HasPermission(permissions ...permissionObject.PermissionCode) bool {
	if len(r.Permissions) == 0 {
		return false
	}

	// Convert permissions to map for faster lookup
	requiredPerms := make(map[permissionObject.PermissionCode]bool)
	for _, p := range permissions {
		requiredPerms[p] = true
	}

	// Check if the role has any of the required permissions
	for _, perm := range r.Permissions {
		if requiredPerms[perm.Code] {
			return true
		}
	}

	return false
}

func (r *Role) HasNoUser() bool {
	return r.UserCount == 0
}

func (r *Role) IsDeleteable() bool {
	// A role is deleteable if it has no users associated with it
	return r.HasNoUser()
}

type NewRoleParams struct {
	ID          int
	Name        string
	Permissions []*permission.Permission
	util.BaseColumnTimestamp
}

func NewRole(params NewRoleParams) *Role {
	return &Role{
		ID:                  params.ID,
		Name:                params.Name,
		Permissions:         params.Permissions,
		BaseColumnTimestamp: params.BaseColumnTimestamp,
	}
}

type RolePermissionUpdateItem struct {
	ID            int   `json:"id" binding:"required"`
	PermissionIDs []int `json:"permission_ids" binding:"required"`
}

type BatchUpdateRolePermissionsRequest []RolePermissionUpdateItem

// RoleListResponse represents the API response for a list of roles
type RoleListResponse struct {
	Roles []generated.Role `json:"roles"`
}

// BatchUpdateRolePermissionsResponse represents the API response for batch updating role permissions
type BatchUpdateRolePermissionsResponse struct {
	UpdatedRoles []int `json:"updated_roles"`
	TotalUpdated int   `json:"total_updated"`
}
