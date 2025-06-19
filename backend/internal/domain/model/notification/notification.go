package model

import (
	"fmt"
	"time"

	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	notificationObj "github.com/test-tzs/nomraeite/internal/domain/object/notification"
)

type Notification struct {
	ID               int
	NotificationType int
	UserID           int
	Description      string
	Detail           *NotificationDetail

	util.BaseColumnTimestamp
}

type NotificationGenerator struct {
	NotificationType int
	UserID           int
	PayoutID         int
	StageID          int
}

// Description constants for audit log events
const (
	DescPayoutApproval = "%s、ユーザー%dにより、振込依頼データ「#%d」は承認ステージ%dで承認されました。"
	DescPayoutReject   = "%s、ユーザー%dにより、振込依頼データ「#%d」は承認ステージ%dで却下されました。"
)

// GetDescriptionByType generates description based on notification type and parameters
func (n *NotificationGenerator) GetDescriptionByType() string {
	switch n.NotificationType {
	case notificationObj.NotificationTypePayoutApproved:
		currentTime := time.Now().Format("2006年1月2日 15:04:05")
		return fmt.Sprintf(DescPayoutApproval, currentTime, n.UserID, n.PayoutID, n.StageID)
	case notificationObj.NotificationTypePayoutRejected:
		currentTime := time.Now().Format("2006年1月2日 15:04:05")
		return fmt.Sprintf(DescPayoutReject, currentTime, n.UserID, n.PayoutID, n.StageID)
	default:
		return ""
	}
}

type PaginatedNotificationResult struct {
	Items []*Notification
	util.Pagination
}
