package service

import (
	"context"

	modelPermission "github.com/test-tzs/nomraeite/internal/domain/model/permission"
	modelRole "github.com/test-tzs/nomraeite/internal/domain/model/role"
	objectPermission "github.com/test-tzs/nomraeite/internal/domain/object/permission"
	repositoryPermission "github.com/test-tzs/nomraeite/internal/domain/repository/permission"
	repositoryRole "github.com/test-tzs/nomraeite/internal/domain/repository/role"
)

type RoleService interface {
	GetRoleByID(ctx context.Context, id int) (*modelRole.Role, error)
	GetRoleByName(ctx context.Context, name string) (*modelRole.Role, error)
	CreateRole(ctx context.Context, role *modelRole.Role) error
	UpdateRole(ctx context.Context, role *modelRole.Role) error
	DeleteRole(ctx context.Context, id int) error
	ListRoles(ctx context.Context) ([]*modelRole.Role, error)

	ReplaceAllRolePermissions(ctx context.Context, roleID int, permissionIDs []int) error
	GetPermissionsByIDs(ctx context.Context, ids []int) ([]*modelPermission.Permission, error)

	// Newly added methods that were previously in PermissionService
	HasPermission(ctx context.Context, roleID int, permissions ...objectPermission.PermissionCode) (bool, error)
}

type roleServiceImpl struct {
	roleRepository       repositoryRole.RoleRepository
	permissionRepository repositoryPermission.PermissionRepository
}

func NewRoleService(
	roleRepository repositoryRole.RoleRepository,
	permissionRepository repositoryPermission.PermissionRepository,
) RoleService {
	return &roleServiceImpl{
		roleRepository:       roleRepository,
		permissionRepository: permissionRepository,
	}
}

func (s *roleServiceImpl) GetRoleByID(ctx context.Context, id int) (*modelRole.Role, error) {
	role, err := s.roleRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	err = s.aggregateRoleUserCount(ctx, []*modelRole.Role{role})
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *roleServiceImpl) GetRoleByName(ctx context.Context, name string) (*modelRole.Role, error) {
	role, err := s.roleRepository.FindByName(ctx, name)
	if err != nil {
		return nil, err
	}
	err = s.aggregateRoleUserCount(ctx, []*modelRole.Role{role})
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *roleServiceImpl) CreateRole(ctx context.Context, role *modelRole.Role) error {
	return s.roleRepository.Create(ctx, role)
}

func (s *roleServiceImpl) UpdateRole(ctx context.Context, role *modelRole.Role) error {
	return s.roleRepository.Update(ctx, role)
}

func (s *roleServiceImpl) DeleteRole(ctx context.Context, id int) error {
	return s.roleRepository.Delete(ctx, id)
}

func (s *roleServiceImpl) ListRoles(ctx context.Context) ([]*modelRole.Role, error) {
	roles, err := s.roleRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.aggregateRoleUserCount(ctx, roles); err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *roleServiceImpl) aggregateRoleUserCount(ctx context.Context, roles []*modelRole.Role) error {
	counts, err := s.roleRepository.GetRoleUserCount(ctx, roles)
	if err != nil {
		return err
	}

	countMap := make(map[int]int64)
	for _, count := range counts {
		countMap[count.RoleID] = count.UserCount
	}

	for _, role := range roles {
		if count, exists := countMap[role.ID]; exists {
			role.UserCount = count
		} else {
			role.UserCount = 0
		}
	}

	return nil
}

func (s *roleServiceImpl) ReplaceAllRolePermissions(ctx context.Context, roleID int, permissionIDs []int) error {
	role, err := s.roleRepository.FindByID(ctx, roleID)
	if err != nil {
		return err
	}

	permissions, err := s.permissionRepository.FindByIDs(ctx, permissionIDs)
	if err != nil {
		return err
	}

	role.Permissions = permissions

	return s.roleRepository.Update(ctx, role)
}

func (s *roleServiceImpl) GetPermissionsByIDs(ctx context.Context, ids []int) ([]*modelPermission.Permission, error) {
	return s.permissionRepository.FindByIDs(ctx, ids)
}

func (s *roleServiceImpl) HasPermission(ctx context.Context, roleID int, permissions ...objectPermission.PermissionCode) (bool, error) {
	role, err := s.roleRepository.FindByID(ctx, roleID)
	if err != nil {
		return false, err
	}

	if role == nil {
		return false, nil
	}

	return role.HasPermission(permissions...), nil
}
