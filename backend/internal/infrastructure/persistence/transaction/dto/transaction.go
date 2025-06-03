package dto

import (
	model "github.com/makeshop-jp/master-console/internal/domain/model/transaction"
	util "github.com/makeshop-jp/master-console/internal/domain/object/basedatetime"
	object "github.com/makeshop-jp/master-console/internal/domain/object/transaction"
	persistence "github.com/makeshop-jp/master-console/internal/infrastructure/persistence/util"
	"gorm.io/gorm"
)

type Transaction struct {
	ID int `json:"id"`
	persistence.BaseColumnTimestamp

	ShopID            int
	TransactionStatus object.TransactionStatus
	PayoutID          int
	PayoutRecordID    int
}

func (Transaction) TableName() string {
	return "transaction"
}

func (t *Transaction) ToModel() *model.Transaction {
	return &model.Transaction{
		ID:                t.ID,
		ShopID:            t.ShopID,
		TransactionStatus: t.TransactionStatus,
		PayoutID:          t.PayoutID,
		PayoutRecordID:    t.PayoutRecordID,
		BaseColumnTimestamp: util.BaseColumnTimestamp{
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			DeletedAt: &t.DeletedAt.Time,
		},
	}
}

func FromModel(model *model.Transaction) *Transaction {
	return &Transaction{
		ID:                model.ID,
		ShopID:            model.ShopID,
		TransactionStatus: model.TransactionStatus,
		PayoutID:          model.PayoutID,
		PayoutRecordID:    model.PayoutRecordID,
		BaseColumnTimestamp: persistence.BaseColumnTimestamp{
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
			DeletedAt: gorm.DeletedAt{Time: *model.DeletedAt, Valid: !model.DeletedAt.IsZero()},
		},
	}
}
