package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
)

type TransferHeaderRequest struct {
	AccessToken    string
	IdempotencyKey string
}

func NewTransferHeaderRequest(accessToken, idempotencyKey string) TransferHeaderRequest {
	return TransferHeaderRequest{
		AccessToken:    accessToken,
		IdempotencyKey: idempotencyKey,
	}
}

type TransferParams struct {
	ItemID                string
	EdiInfo               string
	BeneficiaryBankCode   string
	BeneficiaryBranchCode string
	AccountTypeCode       object.AccountTypeCode
	AccountNumber         string
	BeneficiaryName       string
	TransferAmount        string
}

type TransferParamsRequest struct {
	AccountID               string
	TransferDesignatedDate  string
	TransferDateHolidayCode object.TransferDateHolidayCode
	TotalCount              string
	TotalAmount             string
	Transfers               []TransferParams
}
