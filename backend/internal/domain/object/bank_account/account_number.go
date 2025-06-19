package bankaccount

type AccountNumber string

const (
	AccountNumberLength = 7
)

func (an AccountNumber) Value() string {
	return string(an)
}

func (an AccountNumber) IsValid() bool {
	return len(an.Value()) == AccountNumberLength
}

func FromStringToAccountNumber(s string) AccountNumber {
	return AccountNumber(s)
}
