package object

type PayoutStatus int

const (
	PayoutStatusWaitingApproval PayoutStatus = 1 // 承認待ち
	PayoutStatusApproving       PayoutStatus = 2 // 承認中
	PayoutStatusWaitingTransfer PayoutStatus = 3 // 振込待ち
	PayoutStatusTransferDone    PayoutStatus = 4 // 振込依頼成功
)

func (p PayoutStatus) String() string {
	switch p {
	case PayoutStatusWaitingApproval:
		return "承認待ち"
	case PayoutStatusApproving:
		return "承認中"
	case PayoutStatusWaitingTransfer:
		return "振込待ち"
	case PayoutStatusTransferDone:
		return "振込依頼成功"
	default:
		return "不明"
	}
}

func (p PayoutStatus) Value() int {
	return int(p)
}

func (p PayoutStatus) IsWaitingApproval() bool {
	return p == PayoutStatusWaitingApproval
}
func (p PayoutStatus) IsApproving() bool {
	return p == PayoutStatusApproving
}
func (p PayoutStatus) IsWaitingTransfer() bool {
	return p == PayoutStatusWaitingTransfer
}
func (p PayoutStatus) IsTransferSucceeded() bool {
	return p == PayoutStatusTransferDone
}

func GetPayoutStatusFromInt(i int) (PayoutStatus, bool) {
	switch i {
	case 1:
		return PayoutStatusWaitingApproval, true
	case 2:
		return PayoutStatusApproving, true
	case 3:
		return PayoutStatusWaitingTransfer, true
	case 4:
		return PayoutStatusTransferDone, true
	default:
		return 0, false
	}
}
