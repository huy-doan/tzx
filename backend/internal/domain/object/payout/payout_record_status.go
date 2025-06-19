package object

type PayoutRecordStatus int

const (
	PayoutRecordStatusDraft           PayoutRecordStatus = 0 // 下書き
	PayoutRecordStatusWaitingTransfer PayoutRecordStatus = 1 // 送金待ち
	PayoutRecordStatusInProgress      PayoutRecordStatus = 2 // 振込処理中
	PayoutRecordStatusTransferSuccess PayoutRecordStatus = 3 // 送金済み
	PayoutRecordStatusTransferFailed  PayoutRecordStatus = 4 // 送金失敗
	PayoutRecordStatusDone            PayoutRecordStatus = 5 // 完了
)

func (p PayoutRecordStatus) String() string {
	switch p {
	case PayoutRecordStatusDraft:
		return "下書き"
	case PayoutRecordStatusWaitingTransfer:
		return "送金待ち"
	case PayoutRecordStatusTransferSuccess:
		return "送金済み"
	case PayoutRecordStatusTransferFailed:
		return "送金失敗"
	case PayoutRecordStatusDone:
		return "完了"
	default:
		return "不明"
	}
}

func (t PayoutRecordStatus) IsDraft() bool {
	return t == PayoutRecordStatusDraft
}

func (t PayoutRecordStatus) IsWaitingTransfer() bool {
	return t == PayoutRecordStatusWaitingTransfer
}

func (t PayoutRecordStatus) IsInProgress() bool {
	return t == PayoutRecordStatusInProgress
}

func (t PayoutRecordStatus) IsTransferSuccess() bool {
	return t == PayoutRecordStatusTransferSuccess
}

func (t PayoutRecordStatus) IsTransferFailed() bool {
	return t == PayoutRecordStatusTransferFailed
}

func (t PayoutRecordStatus) IsTransferDone() bool {
	return t == PayoutRecordStatusDone
}
