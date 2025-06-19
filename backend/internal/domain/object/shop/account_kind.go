package object

type AccountKind int

const (
	OrdinaryAccount AccountKind = 1 // (普通口座)
	CurrentAccount  AccountKind = 2 // (当座口座)
)

func (a AccountKind) String() string {
	switch a {
	case OrdinaryAccount:
		return "普通口座"
	case CurrentAccount:
		return "当座口座"
	default:
		return "不明"
	}
}

func (a AccountKind) Value() int {
	switch a {
	case OrdinaryAccount:
		return 1
	case CurrentAccount:
		return 2
	default:
		return 0
	}
}
