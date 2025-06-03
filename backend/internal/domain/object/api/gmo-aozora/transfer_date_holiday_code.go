package object

type TransferDateHolidayCode string

const (
	TransferDateHolidayCodeNextBusinessDay     TransferDateHolidayCode = "1" // 翌営業日
	TransferDateHolidayCodePreviousBusinessDay TransferDateHolidayCode = "2" // 前営業日
	TransferDateHolidayCodeErrorReturn         TransferDateHolidayCode = "3" // エラー返却
)

func (code TransferDateHolidayCode) Value() string {
	return string(code)
}
