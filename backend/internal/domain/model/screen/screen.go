package model

import (
	permissionModel "github.com/test-tzs/nomraeite/internal/domain/model/permission"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	"github.com/test-tzs/nomraeite/internal/pkg/api/generated"
)

// Screen represents a system screen
type Screen struct {
	ID          int
	Name        string
	ScreenCode  string
	ScreenPath  string
	Permissions []*permissionModel.Permission
	util.BaseColumnTimestamp
}

type ListScreenResponse struct {
	Screens []generated.Screen `json:"screens"`
	Total   int64              `json:"total"`
}
