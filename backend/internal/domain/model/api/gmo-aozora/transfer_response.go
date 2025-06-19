package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
)

type TransferResponse struct {
	Type          object.TransferResponseType
	Message       string
	SuccessResult *SuccessTransferResult
	ErrorResult   *ErrorTransferResult
}

func NewSuccessTransferResponse(result *SuccessTransferResult) *TransferResponse {
	return &TransferResponse{
		Type:          object.TransferResponseTypeSuccess,
		Message:       "",
		SuccessResult: result,
		ErrorResult:   nil,
	}
}

func NewValidationErrorTransferResponse(message string, errorResult *ErrorTransferResult) *TransferResponse {
	return &TransferResponse{
		Type:          object.TransferResponseTypeValidationError,
		Message:       message,
		SuccessResult: nil,
		ErrorResult:   errorResult,
	}
}

func NewAuthErrorTransferResponse(message string) *TransferResponse {
	return &TransferResponse{
		Type:          object.TransferResponseTypeAuthError,
		Message:       message,
		SuccessResult: nil,
		ErrorResult:   nil,
	}
}

func NewMaintenanceErrorTransferResponse(message string) *TransferResponse {
	return &TransferResponse{
		Type:          object.TransferResponseTypeMaintenanceError,
		Message:       message,
		SuccessResult: nil,
		ErrorResult:   nil,
	}
}

func NewSystemErrorTransferResponse(message string) *TransferResponse {
	return &TransferResponse{
		Type:          object.TransferResponseTypeSystemError,
		Message:       message,
		SuccessResult: nil,
		ErrorResult:   nil,
	}
}

func NewDuplicateErrorTransferResponse(message string) *TransferResponse {
	return &TransferResponse{
		Type:          object.TransferResponseTypeDuplicateError,
		Message:       message,
		SuccessResult: nil,
		ErrorResult:   nil,
	}
}

func (r *TransferResponse) IsSuccess() bool {
	return r.Type == object.TransferResponseTypeSuccess
}

func (r *TransferResponse) ShouldStopBatch() bool {
	return r.Type.ShouldStopBatch()
}

func (r *TransferResponse) GetErrorMessage() string {
	if r.ErrorResult != nil {
		if len(r.ErrorResult.TransferErrorDetails) > 0 && len(r.ErrorResult.TransferErrorDetails[0].ErrorDetails) > 0 {
			return r.ErrorResult.TransferErrorDetails[0].ErrorDetails[0].ErrorDetailsMessage
		}
		if len(r.ErrorResult.ErrorDetails) > 0 {
			errorDetailMsg := ""
			for _, errorDetail := range r.ErrorResult.ErrorDetails {
				errorDetailMsg += errorDetail.ErrorDetailsMessage
			}
			return errorDetailMsg
		}
		if r.ErrorResult.ErrorMessage != "" {
			return r.ErrorResult.ErrorMessage
		}
	}

	if r.Message != "" {
		return r.Message
	}

	return "Transfer request failed"
}
