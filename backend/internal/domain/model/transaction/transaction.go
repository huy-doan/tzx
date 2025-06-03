package model

import (
	util "github.com/makeshop-jp/master-console/internal/domain/object/basedatetime"
	object "github.com/makeshop-jp/master-console/internal/domain/object/transaction"
)

type Transaction struct {
	ID int `json:"id"`
	util.BaseColumnTimestamp

	ShopID            int
	TransactionStatus object.TransactionStatus
	PayoutID          int
	PayoutRecordID    int

	// Relationships - can be added later as needed
	TransactionRecords []TransactionRecord
}

func (t *Transaction) SetStatus(status object.TransactionStatus) {
	t.TransactionStatus = status
}
