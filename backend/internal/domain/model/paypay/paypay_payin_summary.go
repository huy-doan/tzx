package model

import (
	"time"

	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// PayPayPayinSummary represents the paypay_payin_summary table
type PaypayPayinSummary struct {
	ID int
	util.BaseColumnTimestamp

	PayinFileID       int
	CorporateName     string
	CutoffDate        *time.Time
	PaymentDate       *time.Time
	TransactionAmount int64
	RefundAmount      int64
	UsageFee          int64
	PlatformFee       int64
	InitialFee        int64
	Tax               int64
	Cashback          int64
	Adjustment        int64
	Fee               int64
	Amount            int64

	PayinReconciliation *PayinReconciliationRelation
	// PayinFile *model.PayinFile
}

type PaginatedPaypayPayinSummaryResult struct {
	Items []*PaypayPayinSummary
	util.Pagination
}
