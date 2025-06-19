package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
)

type BulkTransferHeaderRequest struct {
	AccessToken    string
	IdempotencyKey string
}

type BulkTransferDetail struct {
	ItemID                string                 `json:"itemId"`
	EdiInfo               string                 `json:"ediInfo"`
	BeneficiaryBankCode   string                 `json:"beneficiaryBankCode"`
	BeneficiaryBranchCode string                 `json:"beneficiaryBranchCode"`
	AccountTypeCode       object.AccountTypeCode `json:"accountTypeCode"`
	AccountNumber         string                 `json:"accountNumber"`
	BeneficiaryName       string                 `json:"beneficiaryName"`
	TransferAmount        string                 `json:"transferAmount"`
}

type BulkTransferParamsRequest struct {
	AccountID               string                         `json:"accountId"`
	TransferDesignatedDate  string                         `json:"transferDesignatedDate"`
	TransferDateHolidayCode object.TransferDateHolidayCode `json:"transferDateHolidayCode"`
	TotalCount              string                         `json:"totalCount"`
	TotalAmount             string                         `json:"totalAmount"`
	BulkTransfers           []BulkTransferDetail           `json:"bulkTransfers"`
}

func ToDTOBulkTransferHeaderRequest(header model.BulkTransferHeaderRequest) BulkTransferHeaderRequest {
	return BulkTransferHeaderRequest{
		AccessToken:    header.AccessToken,
		IdempotencyKey: header.IdempotencyKey,
	}
}

func ToDTOBulkTransferParamsRequest(request model.BulkTransferParamsRequest) BulkTransferParamsRequest {
	return BulkTransferParamsRequest{
		AccountID:               request.AccountID,
		TransferDesignatedDate:  request.TransferDesignatedDate,
		TransferDateHolidayCode: request.TransferDateHolidayCode,
		TotalCount:              request.TotalCount,
		TotalAmount:             request.TotalAmount,
		BulkTransfers:           ToDTOBulkTransferDetails(request.BulkTransfers),
	}
}

func ToDTOBulkTransferDetails(details []model.BulkTransferDetail) []BulkTransferDetail {
	var dtoDetails []BulkTransferDetail
	for _, detail := range details {
		dtoDetails = append(dtoDetails, BulkTransferDetail{
			ItemID:                detail.ItemID,
			EdiInfo:               detail.EdiInfo,
			BeneficiaryBankCode:   detail.BeneficiaryBankCode,
			BeneficiaryBranchCode: detail.BeneficiaryBranchCode,
			AccountTypeCode:       detail.AccountTypeCode,
			AccountNumber:         detail.AccountNumber,
			BeneficiaryName:       detail.BeneficiaryName,
			TransferAmount:        detail.TransferAmount,
		})
	}
	return dtoDetails
}
