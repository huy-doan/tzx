package notification

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	model "github.com/test-tzs/nomraeite/internal/domain/model/notification"
)

// NotificationRepository defines the interface for notification data operations
type NotificationRepository interface {
	ListNotifications(ctx context.Context, params *inputdata.NotificationListInputData) (*model.PaginatedNotificationResult, error)
	Create(ctx context.Context, notification *model.Notification) error
}
