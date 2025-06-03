package object

type PayinFileType int

const (
	PayinFileTypePaymentSummary     PayinFileType = 1 // 入金レポート
	PayinFileTypePaymentDetail      PayinFileType = 2 // 入金明細
	PayinFileTypePaymentTransaction PayinFileType = 3 // 入金取引明細
)
