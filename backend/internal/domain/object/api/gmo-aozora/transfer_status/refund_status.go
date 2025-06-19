package object

type RefundStatus string

const (
	RefundStatusProcessing RefundStatus = "1" // 組戻手続中
	RefundStatusCompleted  RefundStatus = "2" // 組戻済
	RefundStatusFailed     RefundStatus = "3" // 組戻不成立
)
