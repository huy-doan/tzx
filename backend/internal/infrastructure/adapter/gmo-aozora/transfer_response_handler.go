package adapter

import (
	"encoding/json"
	"net/http"

	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/gmo-aozora/dto"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
	ms "github.com/test-tzs/nomraeite/internal/pkg/utils/messages"
)

type TransferResponseHandler struct {
	logger logger.Logger
}

func NewTransferResponseHandler(logger logger.Logger) *TransferResponseHandler {
	return &TransferResponseHandler{
		logger: logger,
	}
}

func (h *TransferResponseHandler) HandleResponse(statusCode int, bodyBytes []byte) *model.TransferResponse {
	if statusCode == http.StatusCreated {
		return h.handleSuccessResponse(bodyBytes)
	}
	return h.handleErrorResponse(statusCode, bodyBytes)
}

func (h *TransferResponseHandler) handleSuccessResponse(bodyBytes []byte) *model.TransferResponse {
	var result dto.SuccessTransferResponse
	err := json.Unmarshal(bodyBytes, &result)
	if err != nil {
		h.logger.Error("Failed to decode API response", map[string]any{
			"error": err.Error(),
		})
		return model.NewDuplicateErrorTransferResponse(ms.MsgTransferRequestAvoidDuplicate)
	}
	return model.NewSuccessTransferResponse(result.ToModel())
}

func (h *TransferResponseHandler) handleErrorResponse(statusCode int, bodyBytes []byte) *model.TransferResponse {
	if statusCode != http.StatusBadRequest {
		h.logger.Error("API returned unexpected status code", map[string]any{
			"status_code": statusCode,
			"body":        string(bodyBytes),
		})
		return model.NewSystemErrorTransferResponse("System error occurred")
	}

	var errResp dto.ErrorTransferResponse
	err := json.Unmarshal(bodyBytes, &errResp)
	if err != nil {
		h.logger.Error("Failed to decode error response", map[string]any{
			"status_code": statusCode,
			"error":       err.Error(),
		})
		return model.NewSystemErrorTransferResponse("Failed to decode error response")
	}

	return h.categorizeError(&errResp)
}

func (h *TransferResponseHandler) categorizeError(errResp *dto.ErrorTransferResponse) *model.TransferResponse {
	errorResult := &model.ErrorTransferResult{
		ErrorCode:            errResp.ErrorCode,
		ErrorMessage:         errResp.ErrorMessage,
		ErrorDetails:         h.convertErrorDetails(errResp.ErrorDetails),
		TransferErrorDetails: h.convertTransferErrorDetails(errResp.TransferErrorDetails),
	}

	switch errResp.ErrorCode {
	case object.ErrorCodeAuthValidation:
		h.logger.Error("Authorization error", nil)
		return model.NewAuthErrorTransferResponse("Authorization failed")

	case object.ErrorCodeMaintenance:
		h.logger.Error("API is in maintenance mode", nil)
		return model.NewMaintenanceErrorTransferResponse("API is in maintenance mode")

	case object.ErrorCodeValidation:
		h.logger.Error("Transfer request failed validation", map[string]any{
			"error_code":    errResp.ErrorCode,
			"error_message": errResp.ErrorMessage,
		})
		return model.NewValidationErrorTransferResponse(ms.MsgTransferRequestFailedValidation, errorResult)

	default:
		message := h.extractDetailedErrorMessage(errResp)
		return model.NewValidationErrorTransferResponse(message, errorResult)
	}
}

func (h *TransferResponseHandler) extractDetailedErrorMessage(errResp *dto.ErrorTransferResponse) string {
	if len(errResp.TransferErrorDetails) > 0 && len(errResp.TransferErrorDetails[0].ErrorDetails) > 0 {
		return errResp.TransferErrorDetails[0].ErrorDetails[0].ErrorDetailsMessage
	}
	if len(errResp.ErrorDetails) > 0 {
		return errResp.ErrorDetails[0].ErrorDetailsMessage
	}
	return ms.MsgTransferRequestFailed
}

func (h *TransferResponseHandler) convertErrorDetails(dtoDetails []dto.ErrorDetail) []model.ErrorDetail {
	var result []model.ErrorDetail
	for _, detail := range dtoDetails {
		result = append(result, model.ErrorDetail{
			ErrorDetailsCode:    detail.ErrorDetailsCode,
			ErrorDetailsMessage: detail.ErrorDetailsMessage,
		})
	}
	return result
}

func (h *TransferResponseHandler) convertTransferErrorDetails(dtoDetails []dto.TransferErrorDetails) []model.TransferErrorDetails {
	var result []model.TransferErrorDetails
	for _, detail := range dtoDetails {
		result = append(result, model.TransferErrorDetails{
			ItemID:       detail.ItemID,
			ErrorDetails: h.convertErrorDetails(detail.ErrorDetails),
		})
	}
	return result
}
