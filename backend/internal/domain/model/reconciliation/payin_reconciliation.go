package model

type PayinReconciliation struct {
	ID                                  int
	MerchantID                          int
	PaymentMerchantID                   string
	PayinFileGroupID                    int
	PaypayPayinSummaryID                int
	PayInDetailID                       int
	PayinSummaryVsBankIncoming          int
	PayinDetailSumVsPayinTransactionSum int
	PayinSummaryVsPayinDetail           int
	PayinTransactionVsMakeshopOrder     int
	TotalTransactionRecords             int
	MatchedTransactionRecords           int
	SkippedTransactionRecords           int
}
