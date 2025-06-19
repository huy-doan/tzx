package model

import (
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

type PayinReconciliation struct {
	ID int `json:"id,omitempty"`
	util.BaseColumnTimestamp

	MerchantId                          int `json:"merchant_id,omitempty"`
	PayinDetailId                       int `json:"payin_detail_id,omitempty"`
	PayinFileGroupId                    int `json:"payin_file_group_id,omitempty"`
	PaypayPayinSummaryId                int `json:"paypay_payin_summary_id,omitempty"`
	PayinSummaryVsBankIncoming          int `json:"payin_summary_vs_bank_incoming,omitempty"`
	PayinSummaryVsPayinDetail           int `json:"payin_summary_vs_payin_detail,omitempty"`
	PayinDetailSumVsPayinTransactionSum int `json:"payin_detail_sum_vs_payin_transaction_sum,omitempty"`
	PayinTransactionVsMakeshopOrder     int `json:"payin_transaction_vs_makeshop_order,omitempty"`
	TotalTransactionRecords             int `json:"total_transaction_records,omitempty"`
	MatchedTransactionRecords           int `json:"matched_transaction_records,omitempty"`
	SkippedTransactionRecords           int `json:"skipped_transaction_records,omitempty"`
}
