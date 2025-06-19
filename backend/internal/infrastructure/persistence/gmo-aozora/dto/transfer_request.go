package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
)

type TransferHeaderRequest struct {
	AccessToken    string
	IdempotencyKey string
}

type TransferParamsRequest struct {
	AccountID               string                         `json:"accountId"`
	TransferDesignatedDate  string                         `json:"transferDesignatedDate"`
	TransferDateHolidayCode object.TransferDateHolidayCode `json:"transferDateHolidayCode"`
	TotalCount              string                         `json:"totalCount"`
	TotalAmount             string                         `json:"totalAmount"`
	Transfers               []TransferParams               `json:"transfers"`
}

type TransferParams struct {
	ItemID                string                 `json:"itemId"`
	EdiInfo               string                 `json:"ediInfo"`
	BeneficiaryBankCode   string                 `json:"beneficiaryBankCode"`
	BeneficiaryBranchCode string                 `json:"beneficiaryBranchCode"`
	AccountTypeCode       object.AccountTypeCode `json:"accountTypeCode"`
	AccountNumber         string                 `json:"accountNumber"`
	BeneficiaryName       string                 `json:"beneficiaryName"`
	TransferAmount        string                 `json:"transferAmount"`
}

func ToDTOTransferHeaderRequest(header model.TransferHeaderRequest) TransferHeaderRequest {
	return TransferHeaderRequest{
		AccessToken:    header.AccessToken,
		IdempotencyKey: header.IdempotencyKey,
	}
}

func ToDTOTransferParamsRequest(request model.TransferParamsRequest) TransferParamsRequest {
	return TransferParamsRequest{
		AccountID:               request.AccountID,
		TransferDesignatedDate:  request.TransferDesignatedDate,
		TransferDateHolidayCode: request.TransferDateHolidayCode,
		TotalCount:              request.TotalCount,
		TotalAmount:             request.TotalAmount,
		Transfers:               ToDTOTransferParams(request.Transfers),
	}
}

func ToDTOTransferParams(transfers []model.TransferParams) []TransferParams {
	var dtoTransfers []TransferParams
	for _, transfer := range transfers {
		dtoTransfers = append(dtoTransfers, TransferParams{
			ItemID:                transfer.ItemID,
			EdiInfo:               transfer.EdiInfo,
			BeneficiaryBankCode:   transfer.BeneficiaryBankCode,
			BeneficiaryBranchCode: transfer.BeneficiaryBranchCode,
			AccountTypeCode:       transfer.AccountTypeCode,
			AccountNumber:         transfer.AccountNumber,
			BeneficiaryName:       transfer.BeneficiaryName,
			TransferAmount:        transfer.TransferAmount,
		})
	}
	return dtoTransfers
}
