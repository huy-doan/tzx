package model

import (
	"strconv"

	payoutRecordModel "github.com/makeshop-jp/master-console/internal/domain/model/payout_record"
	object "github.com/makeshop-jp/master-console/internal/domain/object/api/gmo-aozora"
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
	Transfers               []TransferParams
}

func NewTransferParamsRequest(accountID string, payoutRecord *payoutRecordModel.PayoutRecord) TransferParamsRequest {
	return TransferParamsRequest{
		AccountID:               accountID,
		TransferDesignatedDate:  payoutRecord.SendingDate.Format("2006-01-02"),
		TransferDateHolidayCode: object.TransferDateHolidayCodeNextBusinessDay,
		Transfers: []TransferParams{
			{
				BeneficiaryBankCode:   payoutRecord.BankAccount.BankCode.Value(),
				BeneficiaryBranchCode: payoutRecord.BankAccount.BranchCode.Value(),
				AccountTypeCode:       payoutRecord.BankAccount.BankAccountType.ToAccountTypeCode(),
				AccountNumber:         payoutRecord.BankAccount.AccountNo.Value(),
				BeneficiaryName:       string(payoutRecord.BankAccount.AccountName),
				TransferAmount:        strconv.Itoa(int(payoutRecord.Amount)),
			},
		},
	}
}
