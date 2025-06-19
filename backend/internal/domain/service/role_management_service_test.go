package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	modelPermission "github.com/test-tzs/nomraeite/internal/domain/model/permission"
	modelRole "github.com/test-tzs/nomraeite/internal/domain/model/role"
	objectPermission "github.com/test-tzs/nomraeite/internal/domain/object/permission"
	"github.com/test-tzs/nomraeite/internal/domain/service"
	"github.com/test-tzs/nomraeite/tests"
	"github.com/test-tzs/nomraeite/tests/mocks"
)

type RoleServiceTestSuite struct {
	tests.TestSuite
	roleRepo       *mocks.RoleRepository
	permissionRepo *mocks.PermissionRepository
	roleService    service.RoleService
	ctx            context.Context
}

func (s *RoleServiceTestSuite) SetupTest() {
	s.roleRepo = mocks.NewRoleRepository(s.T())
	s.permissionRepo = mocks.NewPermissionRepository(s.T())
	s.roleService = service.NewRoleService(s.roleRepo, s.permissionRepo)
	s.ctx = context.Background()
}

func (s *RoleServiceTestSuite) TearDownTest() {
	// Clean up any resources if needed
}

func (s *RoleServiceTestSuite) TestGetRoleByID() {
	// Test cases
	testCases := []struct {
		name          string
		roleID        int
		mockRole      *modelRole.Role
		mockError     error
		expectedRole  *modelRole.Role
		expectedError error
	}{
		{
			name:          "Success",
			roleID:        1,
			mockRole:      &modelRole.Role{ID: 1, Name: "Admin"},
			mockError:     nil,
			expectedRole:  &modelRole.Role{ID: 1, Name: "Admin"},
			expectedError: nil,
		},
		{
			name:          "Not Found",
			roleID:        99,
			mockRole:      nil,
			mockError:     errors.New("role not found"),
			expectedRole:  nil,
			expectedError: errors.New("role not found"),
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.roleRepo.On("FindByID", s.ctx, tc.roleID).Return(tc.mockRole, tc.mockError).Once()

			// Mock GetRoleUserCount if role is found successfully
			if tc.mockError == nil && tc.mockRole != nil {
				s.roleRepo.On("GetRoleUserCount", s.ctx, []*modelRole.Role{tc.mockRole}).Return([]*modelRole.CountResult{}, nil).Once()
			}

			// Execute
			result, err := s.roleService.GetRoleByID(s.ctx, tc.roleID)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}
			s.Equal(tc.expectedRole, result)

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestGetRoleByName() {
	// Test cases
	testCases := []struct {
		name          string
		roleName      string
		mockRole      *modelRole.Role
		mockError     error
		expectedRole  *modelRole.Role
		expectedError error
	}{
		{
			name:          "Success",
			roleName:      "Admin",
			mockRole:      &modelRole.Role{ID: 1, Name: "Admin"},
			mockError:     nil,
			expectedRole:  &modelRole.Role{ID: 1, Name: "Admin"},
			expectedError: nil,
		},
		{
			name:          "Not Found",
			roleName:      "Unknown Role",
			mockRole:      nil,
			mockError:     errors.New("role not found"),
			expectedRole:  nil,
			expectedError: errors.New("role not found"),
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.roleRepo.On("FindByName", s.ctx, tc.roleName).Return(tc.mockRole, tc.mockError).Once()

			// Mock GetRoleUserCount if role is found successfully
			if tc.mockError == nil && tc.mockRole != nil {
				s.roleRepo.On("GetRoleUserCount", s.ctx, []*modelRole.Role{tc.mockRole}).Return([]*modelRole.CountResult{}, nil).Once()
			}

			// Execute
			result, err := s.roleService.GetRoleByName(s.ctx, tc.roleName)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}
			s.Equal(tc.expectedRole, result)

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestCreateRole() {
	// Test cases
	testCases := []struct {
		name          string
		role          *modelRole.Role
		mockError     error
		expectedError error
	}{
		{
			name:          "Success",
			role:          &modelRole.Role{Name: "New Role"},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Failure",
			role:          &modelRole.Role{Name: "Duplicate Role"},
			mockError:     errors.New("duplicate entry"),
			expectedError: errors.New("duplicate entry"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.roleRepo.On("Create", s.ctx, tc.role).Return(tc.mockError).Once()

			// Execute
			err := s.roleService.CreateRole(s.ctx, tc.role)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestUpdateRole() {
	// Test cases
	testCases := []struct {
		name          string
		role          *modelRole.Role
		mockError     error
		expectedError error
	}{
		{
			name:          "Success",
			role:          &modelRole.Role{ID: 1, Name: "Updated Role"},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Failure",
			role:          &modelRole.Role{ID: 99, Name: "Unknown Role"},
			mockError:     errors.New("role not found"),
			expectedError: errors.New("role not found"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.roleRepo.On("Update", s.ctx, tc.role).Return(tc.mockError).Once()

			// Execute
			err := s.roleService.UpdateRole(s.ctx, tc.role)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestDeleteRole() {
	// Test cases
	testCases := []struct {
		name          string
		roleID        int
		mockError     error
		expectedError error
	}{
		{
			name:          "Success",
			roleID:        1,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Failure",
			roleID:        99,
			mockError:     errors.New("role not found"),
			expectedError: errors.New("role not found"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.roleRepo.On("Delete", s.ctx, tc.roleID).Return(tc.mockError).Once()

			// Execute
			err := s.roleService.DeleteRole(s.ctx, tc.roleID)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestListRoles() {
	// Test cases
	testCases := []struct {
		name          string
		mockRoles     []*modelRole.Role
		mockError     error
		expectedRoles []*modelRole.Role
		expectedError error
	}{
		{
			name: "Success",
			mockRoles: []*modelRole.Role{
				{ID: 1, Name: "Admin"},
				{ID: 2, Name: "User"},
			},
			mockError: nil,
			expectedRoles: []*modelRole.Role{
				{ID: 1, Name: "Admin"},
				{ID: 2, Name: "User"},
			},
			expectedError: nil,
		},
		{
			name:          "Failure",
			mockRoles:     nil,
			mockError:     errors.New("database error"),
			expectedRoles: nil,
			expectedError: errors.New("database error"),
		},
	}
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.roleRepo.On("List", s.ctx).Return(tc.mockRoles, tc.mockError).Once()

			// Mock GetRoleUserCount if roles are found successfully
			if tc.mockError == nil && tc.mockRoles != nil {
				s.roleRepo.On("GetRoleUserCount", s.ctx, tc.mockRoles).Return([]*modelRole.CountResult{}, nil).Once()
			}

			// Execute
			result, err := s.roleService.ListRoles(s.ctx)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}
			s.Equal(tc.expectedRoles, result)

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestReplaceAllRolePermissions() {
	// Test cases
	testCases := []struct {
		name              string
		roleID            int
		permissionIDs     []int
		mockRoleFindError error
		mockRole          *modelRole.Role
		mockPermissions   []*modelPermission.Permission
		mockPermFindError error
		mockUpdateError   error
		expectedError     error
	}{
		{
			name:              "Success",
			roleID:            1,
			permissionIDs:     []int{1, 2, 3},
			mockRoleFindError: nil,
			mockRole:          &modelRole.Role{ID: 1, Name: "Admin"},
			mockPermissions: []*modelPermission.Permission{
				{ID: 1, Name: "Create User"},
				{ID: 2, Name: "Update User"},
				{ID: 3, Name: "Delete User"},
			},
			mockPermFindError: nil,
			mockUpdateError:   nil,
			expectedError:     nil,
		},
		{
			name:              "Role Not Found",
			roleID:            99,
			permissionIDs:     []int{1, 2, 3},
			mockRoleFindError: errors.New("role not found"),
			mockRole:          nil,
			mockPermissions:   nil,
			mockPermFindError: nil,
			mockUpdateError:   nil,
			expectedError:     errors.New("role not found"),
		},
		{
			name:              "Permissions Not Found",
			roleID:            1,
			permissionIDs:     []int{99, 100},
			mockRoleFindError: nil,
			mockRole:          &modelRole.Role{ID: 1, Name: "Admin"},
			mockPermissions:   nil,
			mockPermFindError: errors.New("permissions not found"),
			mockUpdateError:   nil,
			expectedError:     errors.New("permissions not found"),
		},
		{
			name:              "Update Error",
			roleID:            1,
			permissionIDs:     []int{1, 2, 3},
			mockRoleFindError: nil,
			mockRole:          &modelRole.Role{ID: 1, Name: "Admin"},
			mockPermissions: []*modelPermission.Permission{
				{ID: 1, Name: "Create User"},
				{ID: 2, Name: "Update User"},
				{ID: 3, Name: "Delete User"},
			},
			mockPermFindError: nil,
			mockUpdateError:   errors.New("update error"),
			expectedError:     errors.New("update error"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations for role
			s.roleRepo.On("FindByID", s.ctx, tc.roleID).Return(tc.mockRole, tc.mockRoleFindError).Once()

			if tc.mockRoleFindError == nil {
				// Setup expectations for permissions
				s.permissionRepo.On("FindByIDs", s.ctx, tc.permissionIDs).Return(tc.mockPermissions, tc.mockPermFindError).Once()

				if tc.mockPermFindError == nil {
					// Expected role update with new permissions
					expectedUpdatedRole := tc.mockRole
					expectedUpdatedRole.Permissions = tc.mockPermissions

					// Setup expectations for role update
					s.roleRepo.On("Update", s.ctx, expectedUpdatedRole).Return(tc.mockUpdateError).Once()
				}
			}

			// Execute
			err := s.roleService.ReplaceAllRolePermissions(s.ctx, tc.roleID, tc.permissionIDs)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
			s.permissionRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestGetPermissionsByIDs() {
	// Test cases
	testCases := []struct {
		name              string
		permissionIDs     []int
		mockPermissions   []*modelPermission.Permission
		mockError         error
		expectedPermCount int
		expectedError     error
	}{
		{
			name:          "Success",
			permissionIDs: []int{1, 2, 3},
			mockPermissions: []*modelPermission.Permission{
				{ID: 1, Name: "Create User"},
				{ID: 2, Name: "Update User"},
				{ID: 3, Name: "Delete User"},
			},
			mockError:         nil,
			expectedPermCount: 3,
			expectedError:     nil,
		},
		{
			name:              "Empty ID List",
			permissionIDs:     []int{},
			mockPermissions:   []*modelPermission.Permission{},
			mockError:         nil,
			expectedPermCount: 0,
			expectedError:     nil,
		},
		{
			name:              "Error",
			permissionIDs:     []int{99, 100},
			mockPermissions:   nil,
			mockError:         errors.New("permissions not found"),
			expectedPermCount: 0,
			expectedError:     errors.New("permissions not found"),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.permissionRepo.On("FindByIDs", s.ctx, tc.permissionIDs).Return(tc.mockPermissions, tc.mockError).Once()

			// Execute
			result, err := s.roleService.GetPermissionsByIDs(s.ctx, tc.permissionIDs)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
				s.Nil(result)
			} else {
				s.NoError(err)
				s.Len(result, tc.expectedPermCount)
				s.Equal(tc.mockPermissions, result)
			}

			// Verify mocks
			s.permissionRepo.AssertExpectations(s.T())
		})
	}
}

func (s *RoleServiceTestSuite) TestHasPermission() {
	permUser := objectPermission.PermissionCode("USER_MANAGE")
	permEdit := objectPermission.PermissionCode("EDIT_OWN_PROFILE")

	// Create a test role implementation
	roleWithUserPermission := &modelRole.Role{
		ID:   1,
		Name: "Admin",
		// Initialize permissions in role directly or set a flag
		// This depends on how Role.HasPermission is implemented
	}

	// Test cases
	testCases := []struct {
		name                        string
		roleID                      int
		mockRole                    *modelRole.Role
		mockRoleErr                 error
		permissions                 []objectPermission.PermissionCode
		expectedResult              bool
		expectedError               error
		mockRoleHasPermissionResult bool // For each test case, define what HasPermission should return
	}{
		{
			name:                        "Has Permission",
			roleID:                      1,
			mockRole:                    roleWithUserPermission,
			mockRoleErr:                 nil,
			permissions:                 []objectPermission.PermissionCode{permUser},
			expectedResult:              true,
			expectedError:               nil,
			mockRoleHasPermissionResult: true,
		},
		{
			name:                        "Does Not Have Permission",
			roleID:                      1,
			mockRole:                    roleWithUserPermission,
			mockRoleErr:                 nil,
			permissions:                 []objectPermission.PermissionCode{permEdit},
			expectedResult:              false,
			expectedError:               nil,
			mockRoleHasPermissionResult: false,
		},
		{
			name:                        "Role Not Found",
			roleID:                      99,
			mockRole:                    nil,
			mockRoleErr:                 errors.New("role not found"),
			permissions:                 []objectPermission.PermissionCode{permUser},
			expectedResult:              false,
			expectedError:               errors.New("role not found"),
			mockRoleHasPermissionResult: false,
		},
		{
			name:                        "Nil Role",
			roleID:                      99,
			mockRole:                    nil,
			mockRoleErr:                 nil,
			permissions:                 []objectPermission.PermissionCode{permUser},
			expectedResult:              false,
			expectedError:               nil,
			mockRoleHasPermissionResult: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup expectations
			s.roleRepo.On("FindByID", s.ctx, tc.roleID).Return(tc.mockRole, tc.mockRoleErr).Once()

			// For valid role test cases, prepare the role to return the expected result
			if tc.mockRole != nil {
				// Get the real Role type to see if it has Permissions field we can set
				// Or create a custom Role implementation that returns a fixed result
				// Depending on the implementation, you may need to set fields or other preparation

				// Mock the HasPermission method indirectly by setting up the role with
				// appropriate permission data according to test case
				// This is done in different ways depending on the implementation:

				// For example, if Role.HasPermission checks a Permissions slice:
				if tc.mockRoleHasPermissionResult {
					// Setup permissions to make HasPermission return true
					perm := &modelPermission.Permission{
						ID:   1,
						Code: objectPermission.PermissionCode(permUser),
					}
					tc.mockRole.Permissions = []*modelPermission.Permission{perm}
				} else {
					// Setup permissions to make HasPermission return false
					tc.mockRole.Permissions = []*modelPermission.Permission{}
				}
			}

			// Execute
			result, err := s.roleService.HasPermission(s.ctx, tc.roleID, tc.permissions...)

			// Assert
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			} else {
				s.NoError(err)
			}
			s.Equal(tc.expectedResult, result)

			// Verify mocks
			s.roleRepo.AssertExpectations(s.T())
		})
	}
}

// TestRoleService runs the test suite
func TestRoleService(t *testing.T) {
	suite.Run(t, new(RoleServiceTestSuite))
}
