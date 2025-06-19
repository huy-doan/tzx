package bankaccount

type AccountKind int

const (
	OrdinaryAccount AccountKind = iota + 1 // (普通口座)
	CurrentAccount                         // (当座口座)
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
	return int(a)
}

func (a AccountKind) IsValid() bool {
	return (a == OrdinaryAccount || a == CurrentAccount)
}
