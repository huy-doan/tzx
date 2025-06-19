package bankaccount

type BankCode string

const (
	BankCodeLength = 4
)

func (bc BankCode) Value() string {
	return string(bc)
}

func (bc BankCode) IsValid() bool {
	return len(bc.Value()) == BankCodeLength
}

func FromStringToBankCode(s string) BankCode {
	return BankCode(s)
}
