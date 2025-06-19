package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora/transfer_status"
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	transferStatusObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora/transfer_status"
)

type BulkTransferStatusResponseDTO struct {
	AcceptanceKeyClass         transferStatusObject.QueryKeyClass `json:"acceptanceKeyClass"`
	DetailInfoNecessity        *bool                              `json:"detailInfoNecessity,omitempty"`
	BulktransferItemKey        *string                            `json:"bulktransferItemKey,omitempty"`
	BaseDate                   string                             `json:"baseDate"`
	BaseTime                   string                             `json:"baseTime"`
	Count                      string                             `json:"count"`
	DetailInfoResult           *bool                              `json:"detailInfoResult,omitempty"`
	TransferQueryBulkResponses []BulkTransferQueryBulkResponseDTO `json:"transferQueryBulkResponses,omitempty"`
	BulkTransferDetails        []BulkTransferStatusDetailDTO      `json:"bulkTransferDetails,omitempty"`
}

type BulkTransferQueryBulkResponseDTO struct {
	DateFrom                *string                                    `json:"dateFrom,omitempty"`
	DateTo                  *string                                    `json:"dateTo,omitempty"`
	RequestNextItemKey      *string                                    `json:"requestNextItemKey,omitempty"`
	RequestTransferStatuses []RequestTransferStatusItemDTO             `json:"requestTransferStatuses,omitempty"`
	RequestTransferClass    *transferStatusObject.RequestTransferClass `json:"requestTransferClass,omitempty"`
	RequestTransferTerm     *transferStatusObject.RequestTransferTerm  `json:"requestTransferTerm,omitempty"`
	HasNext                 bool                                       `json:"hasNext"`
	NextItemKey             *string                                    `json:"nextItemKey,omitempty"`
}

type BulkTransferStatusDetailDTO struct {
	TransferStatus        object.GmoAozoraTransferStatus `json:"transferStatus"`
	TransferStatusName    string                         `json:"transferStatusName"`
	TransferTypeName      string                         `json:"transferTypeName"`
	RemitterCode          *string                        `json:"remitterCode,omitempty"`
	IsFeeFreeUse          bool                           `json:"isFeeFreeUse"`
	IsFeePointUse         bool                           `json:"isFeePointUse"`
	PointName             *string                        `json:"pointName,omitempty"`
	FeeLaterPaymentFlg    *bool                          `json:"feeLaterPaymentFlg,omitempty"`
	TotalFee              string                         `json:"totalFee"`
	TotalDebitAmount      string                         `json:"totalDebitAmount"`
	TransferApplies       []TransferApplyInfoDTO         `json:"transferApplies,omitempty"`
	TransferAccepts       []TransferAcceptInfoDTO        `json:"transferAccepts,omitempty"`
	BulktransferResponses []BulktransferResponseInfoDTO  `json:"bulktransferResponses,omitempty"`
}

type BulktransferResponseInfoDTO struct {
	AccountID              string                      `json:"accountId"`
	RemitterName           string                      `json:"remitterName"`
	TransferDesignatedDate string                      `json:"transferDesignatedDate"`
	TransferDataName       string                      `json:"transferDataName"`
	TotalCount             string                      `json:"totalCount"`
	TotalAmount            string                      `json:"totalAmount"`
	BulkTransferInfos      []BulkTransferInfoDetailDTO `json:"bulkTransferInfos,omitempty"`
}

type BulkTransferInfoDetailDTO struct {
	ItemID                  string                              `json:"itemId"`
	BeneficiaryBankCode     string                              `json:"beneficiaryBankCode"`
	BeneficiaryBankName     *string                             `json:"beneficiaryBankName,omitempty"`
	BeneficiaryBranchCode   string                              `json:"beneficiaryBranchCode"`
	BeneficiaryBranchName   *string                             `json:"beneficiaryBranchName,omitempty"`
	ClearingHouseName       *string                             `json:"clearingHouseName,omitempty"`
	AccountTypeCode         object.AccountTypeCode              `json:"accountTypeCode"`
	AccountNumber           string                              `json:"accountNumber"`
	BeneficiaryName         *string                             `json:"beneficiaryName,omitempty"`
	TransferAmount          string                              `json:"transferAmount"`
	NewCode                 *string                             `json:"newCode,omitempty"`
	EdiInfo                 *string                             `json:"ediInfo,omitempty"`
	TransferDesignatedType  *string                             `json:"transferDesignatedType,omitempty"`
	Identification          *string                             `json:"identification,omitempty"`
	TransferDetailResponses []BulkTransferDetailResponseInfoDTO `json:"transferDetailResponses,omitempty"`
	UnableDetailInfos       []UnableDetailInfoDTO               `json:"unableDetailInfos,omitempty"`
}

type BulkTransferDetailResponseInfoDTO struct {
	BeneficiaryBankNameKanji   *string `json:"beneficiaryBankNameKanji,omitempty"`
	BeneficiaryBranchNameKanji *string `json:"beneficiaryBranchNameKanji,omitempty"`
	UsedPoint                  *string `json:"usedPoint,omitempty"`
	IsFeeFreeUsed              *bool   `json:"isFeeFreeUsed,omitempty"`
	TransferFee                *string `json:"transferFee,omitempty"`
}

func (dto *BulkTransferStatusResponseDTO) ToModel() *model.BulkTransferStatusResponse {
	return &model.BulkTransferStatusResponse{
		AcceptanceKeyClass:         dto.AcceptanceKeyClass,
		DetailInfoNecessity:        dto.DetailInfoNecessity,
		BulktransferItemKey:        dto.BulktransferItemKey,
		BaseDate:                   dto.BaseDate,
		BaseTime:                   dto.BaseTime,
		Count:                      dto.Count,
		DetailInfoResult:           dto.DetailInfoResult,
		TransferQueryBulkResponses: toBulkTransferQueryBulkResponseModels(dto.TransferQueryBulkResponses),
		BulkTransferDetails:        toBulkTransferStatusDetailModels(dto.BulkTransferDetails),
	}
}

func toBulkTransferQueryBulkResponseModels(dtos []BulkTransferQueryBulkResponseDTO) []model.BulkTransferQueryBulkResponse {
	var result []model.BulkTransferQueryBulkResponse
	for _, dto := range dtos {
		result = append(result, model.BulkTransferQueryBulkResponse{
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

func toBulkTransferStatusDetailModels(dtos []BulkTransferStatusDetailDTO) []model.BulkTransferStatusDetail {
	var result []model.BulkTransferStatusDetail
	for _, dto := range dtos {
		result = append(result, model.BulkTransferStatusDetail{
			TransferStatus:        dto.TransferStatus,
			TransferStatusName:    dto.TransferStatusName,
			TransferTypeName:      dto.TransferTypeName,
			RemitterCode:          dto.RemitterCode,
			IsFeeFreeUse:          dto.IsFeeFreeUse,
			IsFeePointUse:         dto.IsFeePointUse,
			PointName:             dto.PointName,
			FeeLaterPaymentFlg:    dto.FeeLaterPaymentFlg,
			TotalFee:              dto.TotalFee,
			TotalDebitAmount:      dto.TotalDebitAmount,
			TransferApplies:       toTransferApplyInfoModels(dto.TransferApplies),
			TransferAccepts:       toTransferAcceptInfoModels(dto.TransferAccepts),
			BulktransferResponses: toBulktransferResponseInfoModels(dto.BulktransferResponses),
		})
	}
	return result
}

func toBulktransferResponseInfoModels(dtos []BulktransferResponseInfoDTO) []model.BulktransferResponseInfo {
	var result []model.BulktransferResponseInfo
	for _, dto := range dtos {
		result = append(result, model.BulktransferResponseInfo{
			AccountID:              dto.AccountID,
			RemitterName:           dto.RemitterName,
			TransferDesignatedDate: dto.TransferDesignatedDate,
			TransferDataName:       dto.TransferDataName,
			TotalCount:             dto.TotalCount,
			TotalAmount:            dto.TotalAmount,
			BulkTransferInfos:      toBulkTransferInfoDetailModels(dto.BulkTransferInfos),
		})
	}
	return result
}

func toBulkTransferInfoDetailModels(dtos []BulkTransferInfoDetailDTO) []model.BulkTransferInfoDetail {
	var result []model.BulkTransferInfoDetail
	for _, dto := range dtos {
		result = append(result, model.BulkTransferInfoDetail{
			ItemID:                  dto.ItemID,
			BeneficiaryBankCode:     dto.BeneficiaryBankCode,
			BeneficiaryBankName:     dto.BeneficiaryBankName,
			BeneficiaryBranchCode:   dto.BeneficiaryBranchCode,
			BeneficiaryBranchName:   dto.BeneficiaryBranchName,
			ClearingHouseName:       dto.ClearingHouseName,
			AccountTypeCode:         dto.AccountTypeCode,
			AccountNumber:           dto.AccountNumber,
			BeneficiaryName:         dto.BeneficiaryName,
			TransferAmount:          dto.TransferAmount,
			NewCode:                 dto.NewCode,
			EdiInfo:                 dto.EdiInfo,
			TransferDesignatedType:  dto.TransferDesignatedType,
			Identification:          dto.Identification,
			TransferDetailResponses: toBulkTransferDetailResponseInfoModels(dto.TransferDetailResponses),
			UnableDetailInfos:       toUnableDetailInfoModels(dto.UnableDetailInfos),
		})
	}
	return result
}

func toBulkTransferDetailResponseInfoModels(dtos []BulkTransferDetailResponseInfoDTO) []model.BulkTransferDetailResponseInfo {
	var result []model.BulkTransferDetailResponseInfo
	for _, dto := range dtos {
		result = append(result, model.BulkTransferDetailResponseInfo{
			BeneficiaryBankNameKanji:   dto.BeneficiaryBankNameKanji,
			BeneficiaryBranchNameKanji: dto.BeneficiaryBranchNameKanji,
			UsedPoint:                  dto.UsedPoint,
			IsFeeFreeUsed:              dto.IsFeeFreeUsed,
			TransferFee:                dto.TransferFee,
		})
	}
	return result
}
