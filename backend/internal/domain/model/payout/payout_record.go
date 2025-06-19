package model

import (
	"time"

	aozoraObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"

	bankAccountObject "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"

	merchant "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	transaction "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	utils "github.com/test-tzs/nomraeite/internal/pkg/utils"

	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
)

type BankAccount struct {
	BankName        string
	BankCode        bankAccountObject.BankCode
	BranchName      string
	BranchCode      bankAccountObject.BankBranchCode
	AccountNo       bankAccountObject.AccountNumber
	AccountName     bankAccountObject.AccountHolderKana
	BankAccountType object.BankAccountType
}

type PayoutRecord struct {
	ID int
	util.BaseColumnTimestamp

	ShopID                string
	PayoutID              int
	TransactionID         int
	BankAccount           BankAccount
	Amount                int64
	TransferStatus        object.PayoutRecordStatus
	AozoraTransferStatus  aozoraObject.GmoAozoraTransferStatus
	SendingDate           *time.Time
	AozoraTransferApplyNo string
	TransferRequestedAt   *time.Time
	TransferExecutedAt    *time.Time
	TransferRequestError  string
	IdempotencyKey        string

	Shop        *merchant.Merchant
	Payout      *Payout
	Transaction *transaction.Transaction
}

func (b BankAccount) IsValid() bool {
	return b.AccountName.IsValid() &&
		b.BankCode.IsValid() &&
		b.BranchCode.IsValid() &&
		b.AccountNo.IsValid()
}

func (p *PayoutRecord) IsBankAccountValid() bool {
	return p.BankAccount.IsValid()
}

func (p *PayoutRecord) SetIdempotencyKey(key string) {
	if p.IdempotencyKey == "" {
		p.IdempotencyKey = key
	}
}

func (p *PayoutRecord) ClearIdempotencyKey() {
	p.IdempotencyKey = ""
}

func (p *PayoutRecord) SetTransferStatus(status object.PayoutRecordStatus) {
	p.TransferStatus = status
}

func (p *PayoutRecord) SetAozoraTransferStatus(status aozoraObject.GmoAozoraTransferStatus) {
	p.AozoraTransferStatus = status
}

func (p *PayoutRecord) SetAozoraTransferApplyNo(applyNo string) {
	p.AozoraTransferApplyNo = applyNo
}

func (p *PayoutRecord) SetTransferExecutedAt() {
	p.TransferExecutedAt = utils.ToPtr(time.Now())
}

func (p *PayoutRecord) SetTransferRequestError(errorMessage string) {
	p.TransferRequestError = errorMessage
}

func (p *PayoutRecord) SetPayoutRecordTransferFailed(errorMessage string) {
	p.SetTransferRequestError(errorMessage)
	p.SetTransferStatus(object.PayoutRecordStatusTransferFailed)
	p.SetTransferExecutedAt()
	p.ClearIdempotencyKey()
}

func (p *PayoutRecord) SetPayoutRecordTransferSuccess(applyNo string) {
	p.SetTransferRequestError("")
	p.SetTransferStatus(object.PayoutRecordStatusTransferSuccess)
	p.SetTransferExecutedAt()
	p.SetAozoraTransferApplyNo(applyNo)
}

func (p *PayoutRecord) IsTransferStatusInProgress() bool {
	return p.TransferStatus == object.PayoutRecordStatusWaitingTransfer
}
