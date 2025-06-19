package object

type AccountTypeCode string

const (
	AccountTypeCodeOrdinary AccountTypeCode = "1" // 普通預金
	AccountTypeCodeCurrent  AccountTypeCode = "2" // 当座預金
	AccountTypeCodeOther    AccountTypeCode = "9" // その他
)

func (code AccountTypeCode) IsOrdinary() bool {
	return code == AccountTypeCodeOrdinary
}

func (code AccountTypeCode) IsCurrent() bool {
	return code == AccountTypeCodeCurrent
}

func (code AccountTypeCode) IsOther() bool {
	return code == AccountTypeCodeOther
}
