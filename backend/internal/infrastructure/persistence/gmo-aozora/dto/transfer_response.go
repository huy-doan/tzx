package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
)

type SuccessTransferResponse struct {
	AccountID        string            `json:"accountId"`
	ResultCode       object.ResultCode `json:"resultCode"`
	ApplyNo          string            `json:"applyNo"`
	ApplyEndDatetime string            `json:"applyEndDatetime,omitempty"`
}

type ErrorDetail struct {
	ErrorDetailsCode    string `json:"errorDetailsCode"`
	ErrorDetailsMessage string `json:"errorDetailsMessage"`
}

type TransferErrorDetails struct {
	ItemID       string        `json:"itemId"`
	ErrorDetails []ErrorDetail `json:"errorDetails"`
}

type ErrorTransferResponse struct {
	ErrorCode            object.ErrorCode       `json:"errorCode"`
	ErrorMessage         string                 `json:"errorMessage"`
	ErrorDetails         []ErrorDetail          `json:"errorDetails"`
	TransferErrorDetails []TransferErrorDetails `json:"transferErrorDetails"`
}

func (dto *SuccessTransferResponse) ToModel() *model.SuccessTransferResult {
	return &model.SuccessTransferResult{
		AccountID:        dto.AccountID,
		ResultCode:       dto.ResultCode,
		ApplyNo:          dto.ApplyNo,
		ApplyEndDatetime: dto.ApplyEndDatetime,
	}
}
