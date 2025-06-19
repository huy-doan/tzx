package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	basedatetime "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	merchant "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/merchant/dto"
	reconciliationDTO "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/reconciliation/dto"
)

// TransferRequestDTO represents the data transfer object for transfer requests
type TransferTransactionRequest struct {
	TransactionID     int                                   `gorm:"column:transaction_id"`
	TransactionStatus int                                   `gorm:"column:transaction_status"`
	Amount            int64                                 `gorm:"column:amount"`
	AccountNumber     string                                `gorm:"column:account_number"`
	AccountName       string                                `gorm:"column:account_holder"`
	BranchName        string                                `gorm:"column:bank_branch_name"`
	BankBranchCode    string                                `gorm:"column:bank_branch_code"`
	BankCode          string                                `gorm:"column:bank_code"`
	Merchant          merchant.Merchant                     `gorm:"embedded"`
	Reconciliation    reconciliationDTO.PayinReconciliation `gorm:"embedded"`

	basedatetime.BaseColumnTimestamp
}

// ToModel converts TransferRequestDTO to a model
func (dto *TransferTransactionRequest) ToModel() *model.TransferRequest {
	if dto == nil {
		return nil
	}

	return &model.TransferRequest{
		Merchant:          dto.Merchant.ToModel(),
		Amount:            dto.Amount,
		TransactionID:     dto.TransactionID,
		TransactionStatus: dto.TransactionStatus,
		AccountNumber:     dto.AccountNumber,
		AccountName:       dto.AccountName,
		BranchName:        dto.BranchName,
		BankBranchCode:    dto.BankBranchCode,
		BankCode:          dto.BankCode,
		Reconciliation:    dto.Reconciliation.ToModel(),
		BaseColumnTimestamp: basedatetime.BaseColumnTimestamp{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
		},
	}
}

// ToTransferRequestModelList converts a slice of TransferRequestDTO to a slice of models
func ToTransferRequestModelList(dtos []*TransferTransactionRequest) []*model.TransferRequest {
	if dtos == nil {
		return nil
	}

	models := make([]*model.TransferRequest, 0, len(dtos))
	for _, dto := range dtos {
		models = append(models, dto.ToModel())
	}

	return models
}
