package model

import (
	"time"

	payinModel "github.com/test-tzs/nomraeite/internal/domain/model/payin"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// PayPayPayinDetail represents the paypay_payin_detail table
type PaypayPayinDetail struct {
	ID int
	util.BaseColumnTimestamp

	PayinFileID          int
	PaymentMerchantID    string
	MerchantBusinessName string
	CutoffDate           *time.Time
	TransactionAmount    int64
	RefundAmount         int64
	UsageFee             int64
	PlatformFee          int64
	InitialFee           int64
	Tax                  int64
	Cashback             int64
	Adjustment           int64
	Fee                  int64
	Amount               int64

	PayinFile           *payinModel.PayinFile
	PayinReconciliation *PayinReconciliationRelation
}

type PaginatedPaypayPayinDetailResult struct {
	Items []*PaypayPayinDetail
	util.Pagination
}
