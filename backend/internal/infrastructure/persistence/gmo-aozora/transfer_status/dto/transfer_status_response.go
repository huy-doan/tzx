// backend/internal/infrastructure/persistence/gmo-aozora/dto/transfer_status_response.go
package dto

import (
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora/transfer_status"
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	transferStatusObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora/transfer_status"
)

// TransferStatusResponseDTO represents DTO for transfer status response
type TransferStatusResponseDTO struct {
	AcceptanceKeyClass         transferStatusObject.QueryKeyClass `json:"acceptanceKeyClass"`
	BaseDate                   string                             `json:"baseDate"`
	BaseTime                   string                             `json:"baseTime"`
	Count                      string                             `json:"count"`
	TransferQueryBulkResponses []TransferQueryBulkResponseDTO     `json:"transferQueryBulkResponses,omitempty"`
	TransferDetails            []TransferStatusDetailDTO          `json:"transferDetails,omitempty"`
}

// TransferQueryBulkResponseDTO represents DTO for bulk query response metadata
type TransferQueryBulkResponseDTO struct {
	DateFrom                *string                                    `json:"dateFrom,omitempty"`
	DateTo                  *string                                    `json:"dateTo,omitempty"`
	RequestNextItemKey      *string                                    `json:"requestNextItemKey,omitempty"`
	RequestTransferStatuses []RequestTransferStatusItemDTO             `json:"requestTransferStatuses,omitempty"`
	RequestTransferClass    *transferStatusObject.RequestTransferClass `json:"requestTransferClass,omitempty"`
	RequestTransferTerm     *transferStatusObject.RequestTransferTerm  `json:"requestTransferTerm,omitempty"`
	HasNext                 bool                                       `json:"hasNext"`
	NextItemKey             string                                     `json:"nextItemKey,omitempty"`
}

// RequestTransferStatusItemDTO represents DTO for individual transfer status filter
type RequestTransferStatusItemDTO struct {
	RequestTransferStatus object.GmoAozoraTransferStatus `json:"requestTransferStatus"`
}

// TransferStatusDetailDTO represents DTO for detailed transfer information
type TransferStatusDetailDTO struct {
	TransferStatus     object.GmoAozoraTransferStatus `json:"transferStatus"`
	TransferStatusName string                         `json:"transferStatusName"`
	TransferTypeName   string                         `json:"transferTypeName"`
	IsFeeFreeUse       bool                           `json:"isFeeFreeUse"`
	IsFeePointUse      bool                           `json:"isFeePointUse"`
	PointName          *string                        `json:"pointName,omitempty"`
	FeeLaterPaymentFlg *bool                          `json:"feeLaterPaymentFlg,omitempty"`
	TransferDetailFee  string                         `json:"transferDetailFee"`
	TotalDebitAmount   string                         `json:"totalDebitAmount"`
	TransferApplies    []TransferApplyInfoDTO         `json:"transferApplies,omitempty"`
	TransferAccepts    []TransferAcceptInfoDTO        `json:"transferAccepts,omitempty"`
	TransferResponses  []TransferResponseInfoDTO      `json:"transferResponses,omitempty"`
}

// TransferApplyInfoDTO represents DTO for transfer application information
type TransferApplyInfoDTO struct {
	ApplyNo              string                       `json:"applyNo"`
	TransferApplyDetails []TransferApplyDetailInfoDTO `json:"transferApplyDetails,omitempty"`
}

// TransferApplyDetailInfoDTO represents DTO for detailed transfer application info
type TransferApplyDetailInfoDTO struct {
	ApplyDatetime   *time.Time                        `json:"applyDatetime,omitempty"`
	ApplyStatus     *transferStatusObject.ApplyStatus `json:"applyStatus,omitempty"`
	ApplyUser       *string                           `json:"applyUser,omitempty"`
	ApplyComment    *string                           `json:"applyComment,omitempty"`
	ApprovalUser    *string                           `json:"approvalUser,omitempty"`
	ApprovalComment *string                           `json:"approvalComment,omitempty"`
}

// TransferAcceptInfoDTO represents DTO for transfer acceptance information
type TransferAcceptInfoDTO struct {
	AcceptNo       string    `json:"acceptNo"`
	AcceptDatetime time.Time `json:"acceptDatetime"`
}

// TransferResponseInfoDTO represents DTO for transfer response information
type TransferResponseInfoDTO struct {
	AccountID              string                  `json:"accountId"`
	RemitterName           string                  `json:"remitterName"`
	TransferDesignatedDate string                  `json:"transferDesignatedDate"`
	TransferInfos          []TransferInfoDetailDTO `json:"transferInfos,omitempty"`
}

// TransferInfoDetailDTO represents DTO for individual transfer information detail
type TransferInfoDetailDTO struct {
	TransferAmount          string                          `json:"transferAmount"`
	EdiInfo                 *string                         `json:"ediInfo,omitempty"`
	BeneficiaryBankCode     string                          `json:"beneficiaryBankCode"`
	BeneficiaryBankName     string                          `json:"beneficiaryBankName"`
	BeneficiaryBranchCode   string                          `json:"beneficiaryBranchCode"`
	BeneficiaryBranchName   string                          `json:"beneficiaryBranchName"`
	AccountTypeCode         object.AccountTypeCode          `json:"accountTypeCode"`
	AccountNumber           string                          `json:"accountNumber"`
	BeneficiaryName         string                          `json:"beneficiaryName"`
	TransferDetailResponses []TransferDetailResponseInfoDTO `json:"transferDetailResponses,omitempty"`
	UnableDetailInfos       []UnableDetailInfoDTO           `json:"unableDetailInfos,omitempty"`
}

// TransferDetailResponseInfoDTO represents DTO for transfer detail response
type TransferDetailResponseInfoDTO struct {
	BeneficiaryBankNameKanji   *string `json:"beneficiaryBankNameKanji,omitempty"`
	BeneficiaryBranchNameKanji *string `json:"beneficiaryBranchNameKanji,omitempty"`
	UsedPoint                  *string `json:"usedPoint,omitempty"`
	IsFeeFreeUsed              *bool   `json:"isFeeFreeUsed,omitempty"`
	TransferFee                *string `json:"transferFee,omitempty"`
}

// UnableDetailInfoDTO represents DTO for unable detail information
type UnableDetailInfoDTO struct {
	TransferDetailStatus *transferStatusObject.TransferDetailStatus `json:"transferDetailStatus,omitempty"`
	RefundStatus         *transferStatusObject.RefundStatus         `json:"refundStatus,omitempty"`
	IsRepayment          *bool                                      `json:"isRepayment,omitempty"`
	RepaymentDate        *string                                    `json:"repaymentDate,omitempty"`
	UnableReasonDetail   *UnableReasonDetailInfoDTO                 `json:"unableReasonDetail,omitempty"`
}

// UnableReasonDetailInfoDTO represents DTO for detailed unable reason
type UnableReasonDetailInfoDTO struct {
	UnableReasonCode    string `json:"unableReasonCode"`
	UnableReasonMessage string `json:"unableReasonMessage"`
}

// ToModel converts DTO to domain model
func (dto *TransferStatusResponseDTO) ToModel() *model.TransferStatusResponse {
	return &model.TransferStatusResponse{
		AcceptanceKeyClass:         dto.AcceptanceKeyClass,
		BaseDate:                   dto.BaseDate,
		BaseTime:                   dto.BaseTime,
		Count:                      dto.Count,
		TransferQueryBulkResponses: toTransferQueryBulkResponseModels(dto.TransferQueryBulkResponses),
		TransferDetails:            toTransferStatusDetailModels(dto.TransferDetails),
	}
}

// Helper conversion functions
func toTransferQueryBulkResponseModels(dtos []TransferQueryBulkResponseDTO) []model.TransferQueryBulkResponse {
	var result []model.TransferQueryBulkResponse
	for _, dto := range dtos {
		result = append(result, model.TransferQueryBulkResponse{
			DateFrom:                dto.DateFrom,
			DateTo:                  dto.DateTo,
			RequestNextItemKey:      dto.RequestNextItemKey,
			RequestTransferStatuses: toRequestTransferStatusItemModels(dto.RequestTransferStatuses),
			RequestTransferClass:    dto.RequestTransferClass,
			RequestTransferTerm:     dto.RequestTransferTerm,
			HasNext:                 dto.HasNext,
			NextItemKey:             dto.NextItemKey,
		})
	}
	return result
}

func toRequestTransferStatusItemModels(dtos []RequestTransferStatusItemDTO) []model.RequestTransferStatusItem {
	var result []model.RequestTransferStatusItem
	for _, dto := range dtos {
		result = append(result, model.RequestTransferStatusItem{
			RequestTransferStatus: dto.RequestTransferStatus,
		})
	}
	return result
}

func toTransferStatusDetailModels(dtos []TransferStatusDetailDTO) []model.TransferStatusDetail {
	var result []model.TransferStatusDetail
	for _, dto := range dtos {
		result = append(result, model.TransferStatusDetail{
			TransferStatus:     dto.TransferStatus,
			TransferStatusName: dto.TransferStatusName,
			TransferTypeName:   dto.TransferTypeName,
			IsFeeFreeUse:       dto.IsFeeFreeUse,
			IsFeePointUse:      dto.IsFeePointUse,
			PointName:          dto.PointName,
			FeeLaterPaymentFlg: dto.FeeLaterPaymentFlg,
			TransferDetailFee:  dto.TransferDetailFee,
			TotalDebitAmount:   dto.TotalDebitAmount,
			TransferApplies:    toTransferApplyInfoModels(dto.TransferApplies),
			TransferAccepts:    toTransferAcceptInfoModels(dto.TransferAccepts),
			TransferResponses:  toTransferResponseInfoModels(dto.TransferResponses),
		})
	}
	return result
}

func toTransferApplyInfoModels(dtos []TransferApplyInfoDTO) []model.TransferApplyInfo {
	var result []model.TransferApplyInfo
	for _, dto := range dtos {
		result = append(result, model.TransferApplyInfo{
			ApplyNo:              dto.ApplyNo,
			TransferApplyDetails: toTransferApplyDetailInfoModels(dto.TransferApplyDetails),
		})
	}
	return result
}

func toTransferApplyDetailInfoModels(dtos []TransferApplyDetailInfoDTO) []model.TransferApplyDetailInfo {
	var result []model.TransferApplyDetailInfo
	for _, dto := range dtos {
		result = append(result, model.TransferApplyDetailInfo{
			ApplyDatetime:   dto.ApplyDatetime,
			ApplyStatus:     dto.ApplyStatus,
			ApplyUser:       dto.ApplyUser,
			ApplyComment:    dto.ApplyComment,
			ApprovalUser:    dto.ApprovalUser,
			ApprovalComment: dto.ApprovalComment,
		})
	}
	return result
}

func toTransferAcceptInfoModels(dtos []TransferAcceptInfoDTO) []model.TransferAcceptInfo {
	var result []model.TransferAcceptInfo
	for _, dto := range dtos {
		result = append(result, model.TransferAcceptInfo{
			AcceptNo:       dto.AcceptNo,
			AcceptDatetime: dto.AcceptDatetime,
		})
	}
	return result
}

func toTransferResponseInfoModels(dtos []TransferResponseInfoDTO) []model.TransferResponseInfo {
	var result []model.TransferResponseInfo
	for _, dto := range dtos {
		result = append(result, model.TransferResponseInfo{
			AccountID:              dto.AccountID,
			RemitterName:           dto.RemitterName,
			TransferDesignatedDate: dto.TransferDesignatedDate,
			TransferInfos:          toTransferInfoDetailModels(dto.TransferInfos),
		})
	}
	return result
}

func toTransferInfoDetailModels(dtos []TransferInfoDetailDTO) []model.TransferInfoDetail {
	var result []model.TransferInfoDetail
	for _, dto := range dtos {
		result = append(result, model.TransferInfoDetail{
			TransferAmount:          dto.TransferAmount,
			EdiInfo:                 dto.EdiInfo,
			BeneficiaryBankCode:     dto.BeneficiaryBankCode,
			BeneficiaryBankName:     dto.BeneficiaryBankName,
			BeneficiaryBranchCode:   dto.BeneficiaryBranchCode,
			BeneficiaryBranchName:   dto.BeneficiaryBranchName,
			AccountTypeCode:         dto.AccountTypeCode,
			AccountNumber:           dto.AccountNumber,
			BeneficiaryName:         dto.BeneficiaryName,
			TransferDetailResponses: toTransferDetailResponseInfoModels(dto.TransferDetailResponses),
			UnableDetailInfos:       toUnableDetailInfoModels(dto.UnableDetailInfos),
		})
	}
	return result
}

func toTransferDetailResponseInfoModels(dtos []TransferDetailResponseInfoDTO) []model.TransferDetailResponseInfo {
	var result []model.TransferDetailResponseInfo
	for _, dto := range dtos {
		result = append(result, model.TransferDetailResponseInfo{
			BeneficiaryBankNameKanji:   dto.BeneficiaryBankNameKanji,
			BeneficiaryBranchNameKanji: dto.BeneficiaryBranchNameKanji,
			UsedPoint:                  dto.UsedPoint,
			IsFeeFreeUsed:              dto.IsFeeFreeUsed,
			TransferFee:                dto.TransferFee,
		})
	}
	return result
}

func toUnableDetailInfoModels(dtos []UnableDetailInfoDTO) []model.UnableDetailInfo {
	var result []model.UnableDetailInfo
	for _, dto := range dtos {
		unableReason := (*model.UnableReasonDetailInfo)(nil)
		if dto.UnableReasonDetail != nil {
			unableReason = &model.UnableReasonDetailInfo{
				UnableReasonCode:    dto.UnableReasonDetail.UnableReasonCode,
				UnableReasonMessage: dto.UnableReasonDetail.UnableReasonMessage,
			}
		}

		result = append(result, model.UnableDetailInfo{
			TransferDetailStatus: dto.TransferDetailStatus,
			RefundStatus:         dto.RefundStatus,
			IsRepayment:          dto.IsRepayment,
			RepaymentDate:        dto.RepaymentDate,
			UnableReasonDetail:   unableReason,
		})
	}
	return result
}
