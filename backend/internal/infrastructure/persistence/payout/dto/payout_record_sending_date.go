package dto

import (
	"time"
)

type PayoutRecordDate struct {
	SendingDate time.Time `gorm:"column:sending_date"`
}
