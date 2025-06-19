package user

import (
	roleModel "github.com/test-tzs/nomraeite/internal/domain/model/role"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	passwordPkg "github.com/test-tzs/nomraeite/internal/domain/object/password"
)

// User represents the user domain model
type User struct {
	ID           int
	Email        string
	PasswordHash string `json:"-"`
	FullName     string
	EnabledMFA   bool
	MFAType      int
	Role         *roleModel.Role
	RoleID       int

	util.BaseColumnTimestamp
}

// VerifyPassword checks if the provided plain password matches the hashed password
func (u *User) VerifyPassword(plainPassword string) bool {
	return passwordPkg.ComparePassword(plainPassword, u.PasswordHash)
}

// HasEnabledMFA checks if MFA is enabled for this user
func (u *User) HasEnabledMFA() bool {
	return u.EnabledMFA
}

// Change User Password
func (u *User) ChangePassword(newPassword string) error {
	hashedPassword, err := passwordPkg.HashPassword(newPassword)
	if err != nil {
		return err
	}

	u.PasswordHash = hashedPassword

	return nil
}

// HasPermission checks if the user has a specific permission
func (u *User) HasPermission(requiredPermissionID int) bool {
	if u.Role == nil || u.Role.Permissions == nil {
		return false
	}

	for _, permission := range u.Role.Permissions {
		if permission.ID == requiredPermissionID {
			return true
		}
	}

	return false
}
