package object

type PayoutStatus int

const (
	PayoutStatusWaitingApproval PayoutStatus = 1 // 1. 承認待ち
	PayoutStatusApproving       PayoutStatus = 2 // 2. 承認中
	PayoutStatusWaitingTransfer PayoutStatus = 3 // 3. 振込待ち
	PayoutStatusTransferred     PayoutStatus = 4 // 4. 振込済み
	PayoutStatusDone            PayoutStatus = 5 // 5. 完了
)

func (p PayoutStatus) String() string {
	switch p {
	case PayoutStatusWaitingApproval:
		return "承認待ち"
	case PayoutStatusApproving:
		return "承認中"
	case PayoutStatusWaitingTransfer:
		return "振込待ち"
	case PayoutStatusTransferred:
		return "振込済み"
	case PayoutStatusDone:
		return "完了"
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
func (p PayoutStatus) IsTransferred() bool {
	return p == PayoutStatusTransferred
}
func (p PayoutStatus) IsDone() bool {
	return p == PayoutStatusDone
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
		return PayoutStatusTransferred, true
	case 5:
		return PayoutStatusDone, true
	default:
		return 0, false
	}
}
