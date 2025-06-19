package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type NotificationDetail struct {
	PayoutID int `json:"payout_id"`
}

func (n *NotificationDetail) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, n)
}

func (n NotificationDetail) Value() (driver.Value, error) {
	return json.Marshal(n)
}
