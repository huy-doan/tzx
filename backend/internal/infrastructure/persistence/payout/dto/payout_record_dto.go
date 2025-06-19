package dto

import (
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	aozoraObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	bankAccountObject "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

// PayoutRecord represents the database model for payout records
type PayoutRecord struct {
	ID                    int                       `gorm:"primaryKey;column:id"`
	ShopID                string                    `gorm:"column:shop_id"`
	PayoutID              int                       `gorm:"column:payout_id"`
	TransactionID         int                       `gorm:"column:transaction_id"`
	BankName              string                    `gorm:"column:bank_name"`
	BankCode              string                    `gorm:"column:bank_code"`
	BranchName            string                    `gorm:"column:branch_name"`
	BranchCode            string                    `gorm:"column:branch_code"`
	BankAccountType       object.BankAccountType    `gorm:"column:bank_account_type"`
	AccountNo             string                    `gorm:"column:account_no"`
	AccountName           string                    `gorm:"column:account_name"`
	Amount                int64                     `gorm:"column:amount"`
	TransferStatus        object.PayoutRecordStatus `gorm:"column:transfer_status"`
	SendingDate           *time.Time                `gorm:"column:sending_date"`
	AozoraTransferApplyNo string                    `gorm:"column:aozora_transfer_apply_no"`
	AozoraTransferStatus  aozoraObject.GmoAozoraTransferStatus
	TransferRequestedAt   *time.Time `gorm:"column:transfer_requested_at"`
	TransferExecutedAt    *time.Time `gorm:"column:transfer_executed_at"`
	TransferRequestError  string     `gorm:"column:transfer_request_error"`
	IdempotencyKey        string     `gorm:"column:idempotency_key"`

	persistence.BaseColumnTimestamp
}

// TableName returns the database table name for PayoutRecord
func (PayoutRecord) TableName() string {
	return "payout_record"
}

// ToModel converts a PayoutRecord to a domain model
func (p *PayoutRecord) ToModel() *model.PayoutRecord {
	if p == nil {
		return nil
	}

	return &model.PayoutRecord{
		ID:            p.ID,
		ShopID:        p.ShopID,
		PayoutID:      p.PayoutID,
		TransactionID: p.TransactionID,
		BankAccount: model.BankAccount{
			BankName:        p.BankName,
			BranchName:      p.BranchName,
			BankCode:        bankAccountObject.BankCode(p.BankCode),
			BranchCode:      bankAccountObject.BankBranchCode(p.BranchCode),
			AccountNo:       bankAccountObject.AccountNumber(p.AccountNo),
			AccountName:     bankAccountObject.AccountHolderKana(p.AccountName),
			BankAccountType: p.BankAccountType,
		},
		Amount:                p.Amount,
		TransferStatus:        p.TransferStatus,
		SendingDate:           p.SendingDate,
		AozoraTransferApplyNo: p.AozoraTransferApplyNo,
		TransferRequestedAt:   p.TransferRequestedAt,
		TransferExecutedAt:    p.TransferExecutedAt,
		TransferRequestError:  p.TransferRequestError,
		IdempotencyKey:        p.IdempotencyKey,
	}
}

// FromModel converts a domain model to a PayoutRecord
func FromPayoutRecordModel(p *model.PayoutRecord) *PayoutRecord {
	if p == nil {
		return nil
	}

	return &PayoutRecord{
		ID:                    p.ID,
		ShopID:                p.ShopID,
		PayoutID:              p.PayoutID,
		TransactionID:         p.TransactionID,
		BankName:              p.BankAccount.BankName,
		BankCode:              p.BankAccount.BankCode.Value(),
		BranchName:            p.BankAccount.BranchName,
		BranchCode:            p.BankAccount.BranchCode.Value(),
		BankAccountType:       p.BankAccount.BankAccountType,
		AccountNo:             p.BankAccount.AccountNo.Value(),
		AccountName:           string(p.BankAccount.AccountName),
		Amount:                p.Amount,
		TransferStatus:        p.TransferStatus,
		SendingDate:           p.SendingDate,
		AozoraTransferApplyNo: p.AozoraTransferApplyNo,
		TransferRequestedAt:   p.TransferRequestedAt,
		TransferExecutedAt:    p.TransferExecutedAt,
		TransferRequestError:  p.TransferRequestError,
		IdempotencyKey:        p.IdempotencyKey,
	}
}

func FromPayoutRecordModels(records []*model.PayoutRecord) []*PayoutRecord {
	if records == nil {
		return nil
	}

	result := make([]*PayoutRecord, len(records))
	for i, record := range records {
		result[i] = FromPayoutRecordModel(record)
	}
	return result
}

func ToPayoutRecordModels(dtos []*PayoutRecord) []*model.PayoutRecord {
	if dtos == nil {
		return nil
	}

	result := make([]*model.PayoutRecord, len(dtos))
	for i, dto := range dtos {
		result[i] = dto.ToModel()
	}
	return result
}
