package object

type TransactionStatus int

const (
	TransactionStatusDraft           TransactionStatus = 0 // 草稿
	TransactionStatusWaitingTransfer TransactionStatus = 1 // 振込待ち
	TransactionStatusTransfered      TransactionStatus = 2 // 振込成功
)

func (t TransactionStatus) IsDraft() bool {
	return t == TransactionStatusDraft
}

func (t TransactionStatus) IsWaitingTransfer() bool {
	return t == TransactionStatusWaitingTransfer
}

func (t TransactionStatus) IsTransfered() bool {
	return t == TransactionStatusTransfered
}
