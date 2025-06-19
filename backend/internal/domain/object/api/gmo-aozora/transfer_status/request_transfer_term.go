package object

type RequestTransferTerm string

const (
	RequestTransferTermApplicationDate RequestTransferTerm = "1" // 振込申請受付日 (ngày tiếp nhận đơn chuyển khoản)
	RequestTransferTermDesignatedDate  RequestTransferTerm = "2" // 振込指定日 (ngày chỉ định chuyển khoản)
)

func (r RequestTransferTerm) String() string {
	switch r {
	case RequestTransferTermApplicationDate:
		return "振込申請受付日"
	case RequestTransferTermDesignatedDate:
		return "振込指定日"
	default:
		return "不明"
	}
}

func (r RequestTransferTerm) Value() string {
	return string(r)
}

func (r RequestTransferTerm) IsApplicationDate() bool {
	return r == RequestTransferTermApplicationDate
}

func (r RequestTransferTerm) IsDesignatedDate() bool {
	return r == RequestTransferTermDesignatedDate
}
