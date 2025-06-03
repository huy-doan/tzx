package persistence

import (
	"context"
	"errors"

	repository "github.com/makeshop-jp/master-console/batch/domain/repository/transaction"
	model "github.com/makeshop-jp/master-console/internal/domain/model/transaction"
	dto "github.com/makeshop-jp/master-console/internal/infrastructure/persistence/transaction/dto"
	"github.com/makeshop-jp/master-console/internal/pkg/database"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionPersistence(db *gorm.DB) repository.TransactionRepository {
	return &TransactionRepositoryImpl{
		db: db,
	}
}

func (r *TransactionRepositoryImpl) GetTransactionByID(ctx context.Context, id int) (*model.Transaction, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	var transactionDTO dto.Transaction
	if err := db.Where("id = ?", id).First(&transactionDTO).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return transactionDTO.ToModel(), nil
}

func (r *TransactionRepositoryImpl) UpdateStatus(ctx context.Context, transaction *model.Transaction) (*model.Transaction, error) {
	db, err := database.GetTxOrDB(ctx)
	if err != nil {
		return nil, err
	}

	transactionDTO := dto.FromModel(transaction)
	err = db.Where(dto.Transaction{
		ID: transactionDTO.ID,
	}).
		Updates(&transactionDTO).Error

	if err != nil {
		return nil, err
	}

	return transactionDTO.ToModel(), nil
}
