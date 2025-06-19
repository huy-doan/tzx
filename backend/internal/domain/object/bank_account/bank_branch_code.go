package bankaccount

type BankBranchCode string

const (
	bankBranchCodeLength = 3
)

func (bbc BankBranchCode) IsValid() bool {
	return len(bbc.Value()) == bankBranchCodeLength
}

func (bbc BankBranchCode) Value() string {
	return string(bbc)
}

func FromStringToBankBranchCode(s string) BankBranchCode {
	return BankBranchCode(s)
}
