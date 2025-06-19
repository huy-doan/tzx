package model

import "time"

type TransactionSummaryCount struct {
	Month            time.Time
	TotalCount       int
	ProcessingCount  int
	TransferredCount int
}

type TransactionSummary struct {
	Month            time.Time `gorm:"column:month"`
	TotalCount       int       `gorm:"column:total_count"`
	ProcessingCount  int       `gorm:"column:processing_count"`
	TransferredCount int       `gorm:"column:transferred_count"`
}
