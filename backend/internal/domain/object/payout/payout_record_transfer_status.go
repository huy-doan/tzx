package object

type TransferStatus int

const (
	TransferStatusDraft           TransferStatus = 0 // 草稿
	TransferStatusWaitingTransfer TransferStatus = 1 // 振込待ち
	TransferStatusTransferSuccess TransferStatus = 2 // 振込成功
	TransferStatusTransferFailed  TransferStatus = 3 // 振込失敗
	TransferStatusTransferDone    TransferStatus = 4 // 振込完了
)

func (t TransferStatus) String() string {
	switch t {
	case TransferStatusDraft:
		return "草稿"
	case TransferStatusWaitingTransfer:
		return "振込待ち"
	case TransferStatusTransferSuccess:
		return "振込成功"
	case TransferStatusTransferFailed:
		return "振込失敗"
	case TransferStatusTransferDone:
		return "振込完了"
	default:
		return "不明"
	}
}

func (t TransferStatus) IsDraft() bool {
	return t == TransferStatusDraft
}

func (t TransferStatus) IsWaitingTransfer() bool {
	return t == TransferStatusWaitingTransfer
}

func (t TransferStatus) IsTransferSuccess() bool {
	return t == TransferStatusTransferSuccess
}

func (t TransferStatus) IsTransferFailed() bool {
	return t == TransferStatusTransferFailed
}

func (t TransferStatus) IsTransferDone() bool {
	return t == TransferStatusTransferDone
}
