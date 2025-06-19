package object

type PayinFileStatus int

const (
	StatusPending PayinFileStatus = 0
	StatusSuccess PayinFileStatus = 1
	StatusFailed  PayinFileStatus = 2
)

// String returns the string representation of the status
func (s PayinFileStatus) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusSuccess:
		return "success"
	case StatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

// IsPending checks if the status is pending
func (s PayinFileStatus) IsPending() bool {
	return s == StatusPending
}

// IsSuccess checks if the status is success
func (s PayinFileStatus) IsSuccess() bool {
	return s == StatusSuccess
}

// IsFailed checks if the status is failed
func (s PayinFileStatus) IsFailed() bool {
	return s == StatusFailed
}
