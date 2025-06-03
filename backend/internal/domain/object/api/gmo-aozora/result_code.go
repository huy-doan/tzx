package object

type ResultCode string

const (
	ResultCodeCompleted = "1" // 完了
	ResultCodePending   = "2" // 未完了
)

func (code ResultCode) IsCompleted() bool {
	return code == ResultCodeCompleted
}

func (code ResultCode) IsPending() bool {
	return code == ResultCodePending
}
