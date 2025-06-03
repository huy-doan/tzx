package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
)

type SuccessTransferResult struct {
	AccountID        string
	ResultCode       object.ResultCode
	ApplyNo          string
	ApplyEndDatetime string
}

type ErrorDetail struct {
	ErrorDetailsCode    string
	ErrorDetailsMessage string
}

type TransferErrorDetails struct {
	ItemID       string
	ErrorDetails []ErrorDetail
}

type ErrorTransferResult struct {
	ErrorCode            object.ErrorCode
	ErrorMessage         string
	ErrorDetails         []ErrorDetail
	TransferErrorDetails []TransferErrorDetails
}
