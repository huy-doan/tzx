package dto

import (
	transferTransactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transfer_transaction"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	merchantDTO "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/merchant/dto"
	convert "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/convert"
	payoutDTO "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
)

// TransferTransactionDTO represents a transfer transaction in the database
type TransferTransactionDTO struct {
	TransactionID     int    `gorm:"column:id;primaryKey"`
	TransactionStatus int    `gorm:"column:transaction_status"`
	Amount            int64  `gorm:"column:amount"`
	AccountNumber     string `gorm:"column:account_number"`
	AccountName       string `gorm:"column:account_name"`
	BranchName        string `gorm:"column:bank_branch_name"`
	BankBranchCode    string `gorm:"column:bank_branch_code"`
	BankCode          string `gorm:"column:bank_code"`

	merchantDTO.Merchant `gorm:"embedded"`
	PayoutRecord         payoutDTO.PayoutRecord `gorm:"embedded"`
	Payout               payoutDTO.Payout       `gorm:"embedded"`

	util.BaseColumnTimestamp
}

// ToModel converts a TransferTransactionDTO to a domain model
func (dto *TransferTransactionDTO) ToModel() *transferTransactionModel.TransferTransaction {
	if dto == nil {
		return nil
	}

	model := &transferTransactionModel.TransferTransaction{
		TransactionID:       dto.TransactionID,
		TransactionStatus:   dto.TransactionStatus,
		Amount:              dto.Amount,
		AccountNumber:       dto.AccountNumber,
		AccountName:         dto.AccountName,
		BranchName:          dto.BranchName,
		BankBranchCode:      dto.BankBranchCode,
		BankCode:            dto.BankCode,
		BaseColumnTimestamp: dto.BaseColumnTimestamp,
	}

	model.Merchant = dto.Merchant.ToModel()
	model.PayoutRecord = convert.ToPayoutRecordModel(&dto.PayoutRecord)
	model.Payout = dto.Payout.ToModel()

	return model
}

// FromModel converts a domain model to a TransferTransactionDTO
func FromTransferTransactionModel(model *transferTransactionModel.TransferTransaction) *TransferTransactionDTO {
	if model == nil {
		return nil
	}

	dto := &TransferTransactionDTO{
		TransactionID:       model.TransactionID,
		TransactionStatus:   model.TransactionStatus,
		Amount:              model.Amount,
		AccountNumber:       model.AccountNumber,
		AccountName:         model.AccountName,
		BranchName:          model.BranchName,
		BankBranchCode:      model.BankBranchCode,
		BankCode:            model.BankCode,
		BaseColumnTimestamp: model.BaseColumnTimestamp,
	}

	return dto
}

// ToTransferTransactionModelList converts a slice of DTOs to a slice of domain models
func ToTransferTransactionModelList(dtos []*TransferTransactionDTO) []*transferTransactionModel.TransferTransaction {
	if dtos == nil {
		return nil
	}

	models := make([]*transferTransactionModel.TransferTransaction, 0, len(dtos))
	for _, dto := range dtos {
		if dto != nil {
			models = append(models, dto.ToModel())
		}
	}

	return models
}
