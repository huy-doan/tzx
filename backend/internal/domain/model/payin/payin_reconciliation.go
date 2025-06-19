package model

import (
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

type PayinReconciliation struct {
	ID int
	util.BaseColumnTimestamp

	PaymentMerchantID                   string
	PaymentFileGroupID                  int
	PaypayPayinSummaryID                int
	PayinSummaryVsBankIncoming          int
	PayinSummaryVsPayinDetail           int
	PayinDetailSumVsPayinTransactionSum int
	PayinTransactionVsMakeShopOrder     int
	TotalTransactionRecords             int
	MatchedTransactionRecords           int
	SkippedTransactionRecords           int
}
