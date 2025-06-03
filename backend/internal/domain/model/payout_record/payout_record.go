package model

import (
	"time"

	bankAccountObject "github.com/makeshop-jp/master-console/internal/domain/object/bank_account"
	timeUtil "github.com/makeshop-jp/master-console/internal/domain/object/basedatetime"
	object "github.com/makeshop-jp/master-console/internal/domain/object/payout"
	utils "github.com/makeshop-jp/master-console/internal/pkg/utils"
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

// PayoutRecord represents a record of a payout transaction
type PayoutRecord struct {
	ID int
	timeUtil.BaseColumnTimestamp

	ShopID                int
	PayoutID              int
	TransactionID         int
	BankAccount           BankAccount
	Amount                float64
	TransferStatus        object.TransferStatus
	SendingDate           *time.Time
	AozoraTransferApplyNo string
	TransferRequestedAt   *time.Time
	TransferExecutedAt    *time.Time
	TransferRequestError  string
	IdempotencyKey        string
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

func (p *PayoutRecord) GenerateIdempotencyKey() {
	if p.IdempotencyKey == "" {
		p.IdempotencyKey = utils.GenerateIdempotencyKey()
	}
}

func (p *PayoutRecord) ClearIdempotencyKey() {
	p.IdempotencyKey = ""
}

func (p *PayoutRecord) SetTransferStatus(status object.TransferStatus) {
	p.TransferStatus = status
}

func (p *PayoutRecord) SetAozoraTransferApplyNo(applyNo string) {
	p.AozoraTransferApplyNo = applyNo
}

func (p *PayoutRecord) SetTransferExecutedAt() {
	p.TransferExecutedAt = utils.ToPtr(time.Now())
}

func (p *PayoutRecord) SetPayoutRecordTransferFailed(errorMessage string) {
	p.TransferRequestError = errorMessage
	p.SetTransferStatus(object.TransferStatusTransferFailed)
	p.SetTransferExecutedAt()
	p.ClearIdempotencyKey()
}

func (p *PayoutRecord) SetPayoutRecordTransferSuccess(applyNo string) {
	p.SetTransferStatus(object.TransferStatusTransferSuccess)
	p.SetTransferExecutedAt()
	p.SetAozoraTransferApplyNo(applyNo)
}

func (p *PayoutRecord) IsTransferStatusInProgress() bool {
	return p.TransferStatus == object.TransferStatusWaitingTransfer
}
