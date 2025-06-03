package object

type ErrorCode string

const (
	ErrorCodeBusinessHourOutside ErrorCode = "032054"     // 業務時間外
	ErrorCodeValidation          ErrorCode = "WG_ERR_110" // バリデーションエラー
	ErrorCodeAuthValidation      ErrorCode = "WG_ERR_105" // 認証バリデーションエラー
	ErrorCodeMaintenance         ErrorCode = "WG_ERR_300" // メンテナンス中
)

func (code ErrorCode) IsBusinessHourOutside() bool {
	return code == ErrorCodeBusinessHourOutside
}

func (code ErrorCode) IsValidationError() bool {
	return code == ErrorCodeValidation
}

func (code ErrorCode) IsAuthValidationError() bool {
	return code == ErrorCodeAuthValidation
}

func (code ErrorCode) IsMaintenance() bool {
	return code == ErrorCodeMaintenance
}
