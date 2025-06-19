package persistence

import (
	"time"

	"gorm.io/gorm"
)

/** ----------------------------------------------------------
 * Base Column Timestamp
 * ---------------------------------------------------------- */
type BaseColumnTimestamp struct {
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt // Soft delete
}
