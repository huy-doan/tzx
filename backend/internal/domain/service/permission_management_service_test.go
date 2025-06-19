package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	permissionModel "github.com/test-tzs/nomraeite/internal/domain/model/permission"
	object "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	permissionObject "github.com/test-tzs/nomraeite/internal/domain/object/permission"
	"github.com/test-tzs/nomraeite/internal/domain/service"
	"github.com/test-tzs/nomraeite/tests"
	"github.com/test-tzs/nomraeite/tests/mocks"
)

type PermissionManagementServiceTestSuite struct {
	tests.TestSuite
	permissionRepo *mocks.PermissionRepository
	service        service.PermissionService
}

func (s *PermissionManagementServiceTestSuite) SetupTest() {
	// Call the parent SetupSuite method
	s.SetupSuite()

	// Initialize mock repository
	s.permissionRepo = mocks.NewPermissionRepository(s.T())

	// Initialize service with mock repository
	s.service = service.NewPermissionService(s.permissionRepo)
}

func (s *PermissionManagementServiceTestSuite) TearDownTest() {
	// Clean up mock expectations
	s.permissionRepo.AssertExpectations(s.T())
}

func (s *PermissionManagementServiceTestSuite) TestGetPermissionsByIDs() { // Define test cases
	testCases := []struct {
		name          string
		ids           []int
		mockSetup     func()
		expectedCount int
		expectError   bool
	}{
		{
			name: "Success - Get multiple permissions",
			ids:  []int{1, 2, 4},
			mockSetup: func() {
				s.permissionRepo.On("FindByIDs", mock.Anything, []int{1, 2, 4}).
					Return([]*permissionModel.Permission{
						createTestPermission(1), // ユーザー管理
						createTestPermission(2), // ユーザーのロール変更
						createTestPermission(4), // システム全体のログ閲覧
					}, nil)
			},
			expectedCount: 3,
			expectError:   false,
		},
		{
			name: "Success - Empty result",
			ids:  []int{99, 100},
			mockSetup: func() {
				s.permissionRepo.On("FindByIDs", mock.Anything, []int{99, 100}).
					Return([]*permissionModel.Permission{}, nil)
			},
			expectedCount: 0,
			expectError:   false,
		},
		{
			name: "Success - Get single permission",
			ids:  []int{1},
			mockSetup: func() {
				s.permissionRepo.On("FindByIDs", mock.Anything, []int{1}).
					Return([]*permissionModel.Permission{createTestPermission(1)}, nil)
			},
			expectedCount: 1,
			expectError:   false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup mock expectations
			tc.mockSetup()

			// Call the service method
			permissions, err := s.service.GetPermissionsByIDs(context.Background(), tc.ids)

			// Assert expectations
			if tc.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(permissions, tc.expectedCount)
			}
		})
	}
}

func (s *PermissionManagementServiceTestSuite) TestListPermissions() {
	// Define test cases
	testCases := []struct {
		name          string
		mockSetup     func()
		expectedCount int
		expectError   bool
	}{{
		name: "Success - List all permissions",
		mockSetup: func() {
			s.permissionRepo.On("List", mock.Anything).
				Return([]*permissionModel.Permission{
					createTestPermission(1),  // ユーザー管理
					createTestPermission(2),  // ユーザーのロール変更
					createTestPermission(4),  // システム全体のログ閲覧
					createTestPermission(5),  // 自分の個人データ変更
					createTestPermission(6),  // 自分の行動ログ確認
					createTestPermission(7),  // 管理画面の参照権限
					createTestPermission(8),  // 振込み承認（事業）
					createTestPermission(9),  // 振込み承認（経理）
					createTestPermission(10), // 手動振込機能
				}, nil).Once()
		},
		expectedCount: 9,
		expectError:   false,
	},
		{
			name: "Success - Empty list",
			mockSetup: func() {
				s.permissionRepo.On("List", mock.Anything).
					Return([]*permissionModel.Permission{}, nil).Once()
			},
			expectedCount: 0,
			expectError:   false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Setup mock expectations
			tc.mockSetup()

			// Call the service method
			permissions, err := s.service.ListPermissions(context.Background())

			// Assert expectations
			if tc.expectError {
				s.Error(err)
			} else {
				s.NoError(err)
				s.Len(permissions, tc.expectedCount)
			}
		})
	}
}

// Helper function to create test permission objects based on seeded data
func createTestPermission(id int) *permissionModel.Permission {
	// Map of seeded permissions from database/seeds/master/05_permissions.sql
	permissions := map[int]struct {
		name     string
		code     string
		screenID int
	}{
		1:  {"ユーザー管理", "USER_MANAGE", 1},
		2:  {"ユーザーのロール変更", "USER_ROLE_CHANGE", 1},
		4:  {"システム全体のログ閲覧", "SYSTEM_LOG_VIEW", 2},
		5:  {"自分の個人データ変更", "EDIT_OWN_PROFILE", 3},
		6:  {"自分の行動ログ確認", "VIEW_OWN_LOG", 3},
		7:  {"管理画面の参照権限", "VIEW_ADMIN_PANEL", 4},
		8:  {"振込み承認（事業）", "TRANSFER_APPROVE_BUSINESS", 5},
		9:  {"振込み承認（経理）", "TRANSFER_APPROVE_ACCOUNTANT", 5},
		10: {"手動振込機能", "MANUAL_TRANSFER", 6},
	}

	// Get permission data from the map
	if p, ok := permissions[id]; ok {
		return &permissionModel.Permission{
			ID:   id,
			Name: p.name,
			Code: permissionObject.PermissionCode(p.code),
			BaseColumnTimestamp: object.BaseColumnTimestamp{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
	}

	// Fallback for IDs not in seed data (should not happen in these tests)
	return &permissionModel.Permission{
		ID:                  id,
		Name:                "Unknown Permission",
		Code:                permissionObject.PermissionCode("UNKNOWN"),
		BaseColumnTimestamp: object.BaseColumnTimestamp{},
	}
}

// TestPermissionManagementService runs the test suite
func TestPermissionManagementService(t *testing.T) {
	suite.Run(t, new(PermissionManagementServiceTestSuite))
}
