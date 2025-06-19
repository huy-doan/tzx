package object

type TransactionStatus int

const (
	TransactionStatusDraft           TransactionStatus = 0 // ドラフト
	TransactionStatusWaitingTransfer TransactionStatus = 1 // 振込待ち
	TransactionStatusTransferred     TransactionStatus = 2 // 振込成功
)

func (t TransactionStatus) IsDraft() bool {
	return t == TransactionStatusDraft
}

func (t TransactionStatus) IsWaitingTransfer() bool {
	return t == TransactionStatusWaitingTransfer
}

func (t TransactionStatus) IsTransferred() bool {
	return t == TransactionStatusTransferred
}

func (t TransactionStatus) String() string {
	switch t {
	case TransactionStatusDraft:
		return "ドラフト"
	case TransactionStatusWaitingTransfer:
		return "振込待ち"
	case TransactionStatusTransferred:
		return "振込成功"
	default:
		return "承認待ち"
	}
}
