package persistence

import (
	"context"

	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/transaction"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/transaction/dto"

	"github.com/test-tzs/nomraeite/internal/pkg/database"
	"gorm.io/gorm"
)

type transactionRecordPersistence struct{}

func NewTransactionRecordRepository(db *gorm.DB) repository.TransactionRecordRepository {
	return &transactionRecordPersistence{}
}

func (r transactionRecordPersistence) BulkCreate(ctx context.Context, transactionRecords []*transactionModel.TransactionRecord) error {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return err
	}
	dtos := []*dto.TransactionRecord{}

	for _, transactionRecord := range transactionRecords {
		dto := dto.FromTransactionRecordModel(transactionRecord)
		dtos = append(dtos, dto)
	}

	return db.Create(&dtos).Error
}
