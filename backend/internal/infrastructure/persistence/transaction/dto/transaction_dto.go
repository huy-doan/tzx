package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	bankaccount "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	transactionObject "github.com/test-tzs/nomraeite/internal/domain/object/transaction"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

// Transaction is a data transfer object representing a transaction in the database
type Transaction struct {
	ID                int                                 `gorm:"column:id;primaryKey"`
	ShopID            string                              `gorm:"column:shop_id"`
	TransactionStatus transactionObject.TransactionStatus `gorm:"column:transaction_status"`
	PayoutID          *int                                `gorm:"column:payout_id"`
	BankCode          string                              `gorm:"column:bank_code"`           // 銀行コード
	BankBranchCode    string                              `gorm:"column:bank_branch_code"`    // 支店コード
	BankBranch        string                              `gorm:"column:bank_branch"`         // 支店名
	AccountNumber     string                              `gorm:"column:account_number"`      // 口座番号
	AccountHolder     string                              `gorm:"column:account_holder"`      // 口座名義
	AccountHolderKana string                              `gorm:"column:account_holder_kana"` // 口座名義カナ
	AccountKind       int                                 `gorm:"column:account_kind"`        // 1:普通口座, 2:当座口座

	TransactionRecords []TransactionRecord `gorm:"foreignKey:TransactionID"`

	persistence.BaseColumnTimestamp
}

// TableName returns the table name for the TransactionDTO
func (Transaction) TableName() string {
	return "transaction"
}

// ToModel converts a TransactionDTO to a domain model
func (dto *Transaction) ToModel() *model.Transaction {
	if dto == nil {
		return nil
	}

	return &model.Transaction{
		ID:                dto.ID,
		ShopID:            dto.ShopID,
		TransactionStatus: dto.TransactionStatus,
		PayoutID:          dto.PayoutID,
		BankCode:          bankaccount.BankCode(dto.BankCode),
		BankBranchCode:    bankaccount.BankBranchCode(dto.BankBranchCode),
		BankBranch:        bankaccount.BankBranch(dto.BankBranch),
		AccountNumber:     bankaccount.AccountNumber(dto.AccountNumber),
		AccountHolder:     bankaccount.AccountHolder(dto.AccountHolder),
		AccountHolderKana: bankaccount.AccountHolderKana(dto.AccountHolderKana),
		AccountKind:       bankaccount.AccountKind(dto.AccountKind),
		BaseColumnTimestamp: util.BaseColumnTimestamp{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
		},
	}
}

func FromModel(model *model.Transaction) *Transaction {
	if model == nil {
		return nil
	}

	return &Transaction{
		ID:                model.ID,
		ShopID:            model.ShopID,
		TransactionStatus: model.TransactionStatus,
		PayoutID:          model.PayoutID,
		BankCode:          string(model.BankCode),
		BankBranchCode:    string(model.BankBranchCode),
		BankBranch:        string(model.BankBranch),
		AccountNumber:     string(model.AccountNumber),
		AccountHolder:     string(model.AccountHolder),
		AccountHolderKana: string(model.AccountHolderKana),
		AccountKind:       model.AccountKind.Value(),
		BaseColumnTimestamp: persistence.BaseColumnTimestamp{
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		},
	}
}
