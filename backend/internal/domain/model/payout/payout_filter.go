package model

import (
	"time"

	filterModel "github.com/test-tzs/nomraeite/internal/domain/model/filter"
	objectPayout "github.com/test-tzs/nomraeite/internal/domain/object/payout"
)

type PayoutFilter struct {
	filterModel.BaseFilter
	CreatedAt    *time.Time
	SendingDate  *time.Time
	SentDate     *time.Time
	PayoutStatus *objectPayout.PayoutStatus
}

func NewPayoutFilter() *PayoutFilter {
	filter := &PayoutFilter{}
	filter.ValidSortFields = map[string]bool{
		"id":            true,
		"created_at":    true,
		"updated_at":    true,
		"sending_date":  true,
		"sent_date":     true,
		"payout_status": true,
	}

	filter.SetPagination(1, 10)

	return filter
}

func (f *PayoutFilter) ApplyFilters() {
	f.AddDateFilter("created_at", filterModel.Equal, f.CreatedAt)
	f.AddDateFilter("sending_date", filterModel.Equal, f.SendingDate)
	f.AddDateFilter("sent_date", filterModel.Equal, f.SentDate)

	if f.PayoutStatus != nil {
		f.AddCondition("payout_status", filterModel.Equal, int(*f.PayoutStatus))
	}
}
