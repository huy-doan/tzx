package model

import (
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/permission"
)

// Permission represents a permission in the system
type Permission struct {
	ID   int
	Name string
	Code object.PermissionCode
	util.BaseColumnTimestamp
}

type NewPermissionParams struct {
	ID int
	util.BaseColumnTimestamp
	Name string
	Code object.PermissionCode
}

func NewPermission(params NewPermissionParams) *Permission {
	return &Permission{
		Name:                params.Name,
		Code:                params.Code,
		BaseColumnTimestamp: params.BaseColumnTimestamp,
	}
}

type PermissionListResponse struct {
	Permissions []*Permission `json:"permissions"`
	Total       int64         `json:"total"`
}
