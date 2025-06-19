package object

type RequestTransferClass string

const (
	RequestTransferClassAll             RequestTransferClass = "1"
	RequestTransferClassApplicationOnly RequestTransferClass = "2"
	RequestTransferClassAcceptanceOnly  RequestTransferClass = "3"
)

func (r RequestTransferClass) String() string {
	switch r {
	case RequestTransferClassAll:
		return "ALL"
	case RequestTransferClassApplicationOnly:
		return "振込申請のみ"
	case RequestTransferClassAcceptanceOnly:
		return "振込受付情報のみ"
	default:
		return "不明"
	}
}

func (r RequestTransferClass) Value() string {
	return string(r)
}
