package dto

import (
	merchantModel "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	model "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	basedatetime "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	merchant "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/merchant/dto"
	reconciliationDTO "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/reconciliation/dto"
)

// TransactionDTO is a data transfer object representing a transaction in the database
type TransferTransactionDetail struct {
	ID                int `gorm:"column:id;primaryKey"`
	ShopID            int `gorm:"column:shop_id"`
	TransactionStatus int `gorm:"column:transaction_status"`
	basedatetime.BaseColumnTimestamp

	Merchant       merchant.Merchant                     `gorm:"embedded"`
	Reconciliation reconciliationDTO.PayinReconciliation `gorm:"embedded"`

	AccountName    string `gorm:"column:account_name"`
	AccountNumber  string `gorm:"column:account_number"`
	BankBranchCode string `gorm:"column:bank_branch_code"`
	BankCode       string `gorm:"column:bank_code"`
	BankName       string `gorm:"column:bank_name"`
	BranchName     string `gorm:"column:bank_branch_name"`
}

// ToDetailModel converts a TransactionDTO and its records to a TransactionDetail model
func (dto *TransferTransactionDetail) ToTransferDetailModel(records []TransactionRecord) *model.TransferTransactionDetail {
	if dto == nil {
		return nil
	}

	recordModels := make([]model.TransactionRecord, 0, len(records))
	for _, record := range records {
		recordModels = append(recordModels, *record.ToModel())
	}

	return &model.TransferTransactionDetail{
		ID:                 dto.ID,
		Merchant:           dto.Merchant.ToModel(),
		Reconciliation:     dto.Reconciliation.ToModel(),
		TransactionStatus:  dto.TransactionStatus,
		AccountName:        dto.AccountName,
		AccountNumber:      dto.AccountNumber,
		BankBranchCode:     dto.BankBranchCode,
		BankCode:           dto.BankCode,
		BankName:           dto.BankName,
		BranchName:         dto.BranchName,
		TransactionRecords: recordModels,
		BaseColumnTimestamp: basedatetime.BaseColumnTimestamp{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
		},
	}
}

// ToTransferTransactionDetailModelList converts a slice of TransactionDTOs to a slice of TransferTransactionDetail models
func ToTransferTransactionDetailModelList(transactionDTOs []TransferTransactionDetail, recordDTOs []TransactionRecord) []*model.TransferTransactionDetail {
	if len(transactionDTOs) == 0 {
		return nil
	}

	recordsByTransactionID := make(map[int][]TransactionRecord)
	for _, record := range recordDTOs {
		recordsByTransactionID[record.TransactionID] = append(recordsByTransactionID[record.TransactionID], record)
	}

	transactionDetails := make([]*model.TransferTransactionDetail, 0, len(transactionDTOs))
	for _, transactionDTO := range transactionDTOs {
		records := recordsByTransactionID[transactionDTO.ID]
		transactionDetails = append(transactionDetails, transactionDTO.ToTransferDetailModel(records))
	}

	return transactionDetails
}

// ConvertTransactionsToTransferDetails converts a slice of Transaction DTOs to TransferTransactionDetail models
func ConvertTransactionsToTransferDetails(transactions []Transaction) []*model.TransferTransactionDetail {
	var result []*model.TransferTransactionDetail
	for _, transaction := range transactions {
		var transactionRecords []model.TransactionRecord
		for _, recordDTO := range transaction.TransactionRecords {
			transactionRecords = append(transactionRecords, *recordDTO.ToModel())
		}
		merchant := &merchantModel.Merchant{
			ShopID: transaction.ShopID,
		}

		detail := &model.TransferTransactionDetail{
			ID:                 transaction.ID,
			TransactionStatus:  int(transaction.TransactionStatus),
			AccountName:        transaction.AccountHolder,
			AccountNumber:      transaction.AccountNumber,
			BankBranchCode:     transaction.BankBranchCode,
			BankCode:           transaction.BankCode,
			BranchName:         transaction.BankBranch,
			PayoutID:           *transaction.PayoutID,
			Merchant:           merchant,
			TransactionRecords: transactionRecords,
		}
		result = append(result, detail)
	}

	return result
}
