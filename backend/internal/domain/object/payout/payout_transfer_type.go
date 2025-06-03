package object

type PayoutTransferType int

const (
	PayoutTransferTypeIndividualTransfer PayoutTransferType = 1 // 個別振込処理
	PayoutTransferTypeBulkTransfer       PayoutTransferType = 2 // 総合振込処理
)

func (p PayoutTransferType) String() string {
	switch p {
	case PayoutTransferTypeIndividualTransfer:
		return "個別振込処理"
	case PayoutTransferTypeBulkTransfer:
		return "総合振込処理"
	default:
		return "不明"
	}
}

func (p PayoutTransferType) IsIndividualTransfer() bool {
	return p == PayoutTransferTypeIndividualTransfer
}

func (p PayoutTransferType) IsBulkTransfer() bool {
	return p == PayoutTransferTypeBulkTransfer
}
