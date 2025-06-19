package object

type PayoutTransferType int

const (
	PayoutTransferTypeNormalTransfer PayoutTransferType = 1 // 通常振込処理
	PayoutTransferTypeBulkTransfer   PayoutTransferType = 2 // 総合振込処理
)

func (p PayoutTransferType) String() string {
	switch p {
	case PayoutTransferTypeNormalTransfer:
		return "通常振込処理"
	case PayoutTransferTypeBulkTransfer:
		return "総合振込処理"
	default:
		return "不明"
	}
}

func (p PayoutTransferType) IsNormalTransfer() bool {
	return p == PayoutTransferTypeNormalTransfer
}

func (p PayoutTransferType) IsBulkTransfer() bool {
	return p == PayoutTransferTypeBulkTransfer
}
