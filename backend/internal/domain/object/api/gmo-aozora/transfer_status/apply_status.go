package object

type ApplyStatus string

const (
	ApplyStatusNotApplied       ApplyStatus = "0" // 未申請
	ApplyStatusApplying         ApplyStatus = "1" // 申請中
	ApplyStatusReturned         ApplyStatus = "2" // 差戻
	ApplyStatusWithdrawn        ApplyStatus = "3" // 取下げ
	ApplyStatusExpired          ApplyStatus = "4" // 期限切れ
	ApplyStatusApproved         ApplyStatus = "5" // 承認済
	ApplyStatusApprovalCanceled ApplyStatus = "6" // 承認取消
	ApplyStatusAutoApproved     ApplyStatus = "7" // 自動承認
)

func (a ApplyStatus) String() string {
	switch a {
	case ApplyStatusNotApplied:
		return "未申請"
	case ApplyStatusApplying:
		return "申請中"
	case ApplyStatusReturned:
		return "差戻"
	case ApplyStatusWithdrawn:
		return "取下げ"
	case ApplyStatusExpired:
		return "期限切れ"
	case ApplyStatusApproved:
		return "承認済"
	case ApplyStatusApprovalCanceled:
		return "承認取消"
	case ApplyStatusAutoApproved:
		return "自動承認"
	default:
		return "不明"
	}
}

func (a ApplyStatus) IsApplying() bool {
	return a == ApplyStatusApplying
}

func (a ApplyStatus) IsApproved() bool {
	return a == ApplyStatusApproved || a == ApplyStatusAutoApproved
}

func (a ApplyStatus) IsCanceled() bool {
	return a == ApplyStatusApprovalCanceled || a == ApplyStatusWithdrawn
}
