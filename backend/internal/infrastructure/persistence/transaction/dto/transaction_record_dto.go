package dto

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	object "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	transactionRecordObject "github.com/test-tzs/nomraeite/internal/domain/object/transaction_record"
	util "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

// TransactionRecord represents a transaction record in the database
type TransactionRecord struct {
	ID                    int    `gorm:"column:id;primaryKey"`
	TransactionID         int    `gorm:"column:transaction_id"`
	MerchantID            int    `gorm:"column:merchant_id"`
	PayinDetailID         int    `gorm:"column:payin_detail_id"`
	TransactionRecordType int    `gorm:"column:transaction_record_type"`
	Title                 string `gorm:"column:title"`
	Amount                int64  `gorm:"column:amount"`
	util.BaseColumnTimestamp
}

// TableName returns the table name for the TransactionRecordDTO
func (TransactionRecord) TableName() string {
	return "transaction_record"
}

// ToModel converts a TransactionRecordDTO to a domain model
func (dto *TransactionRecord) ToModel() *model.TransactionRecord {
	if dto == nil {
		return nil
	}

	return &model.TransactionRecord{
		ID:                    dto.ID,
		TransactionID:         dto.TransactionID,
		MerchantID:            dto.MerchantID,
		PayinDetailID:         dto.PayinDetailID,
		TransactionRecordType: transactionRecordObject.TransactionRecordType(dto.TransactionRecordType),
		Title:                 dto.Title,
		Amount:                dto.Amount,
		BaseColumnTimestamp: object.BaseColumnTimestamp{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
		},
	}
}

// FromModel converts a domain model to a TransactionRecordDTO
func FromTransactionRecordModel(in *model.TransactionRecord) *TransactionRecord {
	if in == nil {
		return nil
	}

	return &TransactionRecord{
		ID:                    in.ID,
		TransactionID:         in.TransactionID,
		MerchantID:            in.MerchantID,
		PayinDetailID:         in.PayinDetailID,
		TransactionRecordType: int(in.TransactionRecordType),
		Title:                 in.Title,
		Amount:                in.Amount,
		BaseColumnTimestamp: util.BaseColumnTimestamp{
			CreatedAt: in.BaseColumnTimestamp.CreatedAt,
			UpdatedAt: in.BaseColumnTimestamp.UpdatedAt,
		},
	}
}
