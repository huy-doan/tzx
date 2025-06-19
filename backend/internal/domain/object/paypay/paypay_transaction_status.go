package object

import (
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
)

type PaypayTransactionStatus int

const (
	TransactionComplete        PaypayTransactionStatus = 1 // 取引完了
	TransactionAccepted        PaypayTransactionStatus = 2 // 取引受付完了
	RefundComplete             PaypayTransactionStatus = 3 // 返金完了
	TransactionCancelled       PaypayTransactionStatus = 4 // 取引取消
	TransactionAcceptCancelled PaypayTransactionStatus = 5 // 取引受付取消
	Adjustment                 PaypayTransactionStatus = 6 // 調整
	RemittanceComplete         PaypayTransactionStatus = 7 // 送金完了
)

// PaypayTransactionStatusFromString maps Japanese status string to enum value
func PaypayTransactionStatusFromString(s string) *PaypayTransactionStatus {
	switch s {
	case "取引完了":
		return utils.ToPtr(TransactionComplete)
	case "取引受付完了":
		return utils.ToPtr(TransactionAccepted)
	case "返金完了":
		return utils.ToPtr(RefundComplete)
	case "取引取消":
		return utils.ToPtr(TransactionCancelled)
	case "取引受付取消":
		return utils.ToPtr(TransactionAcceptCancelled)
	case "調整":
		return utils.ToPtr(Adjustment)
	case "送金完了":
		return utils.ToPtr(RemittanceComplete)
	default:
		return nil
	}
}
