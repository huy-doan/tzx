package model

import (
	"time"

	merchant "github.com/makeshop-jp/master-console/internal/domain/model/merchant"
	transaction "github.com/makeshop-jp/master-console/internal/domain/model/transaction"

	util "github.com/makeshop-jp/master-console/internal/domain/object/basedatetime"
	object "github.com/makeshop-jp/master-console/internal/domain/object/payout"
)

type PayoutRecord struct {
	ID int
	util.BaseColumnTimestamp

	ShopID                int
	PayoutID              int
	TransactionID         int
	BankName              string
	BankCode              string
	BranchName            string
	BranchCode            string
	BankAccountType       object.BankAccountType
	AccountNo             string
	AccountName           string
	Amount                float64
	TransferStatus        object.TransferStatus
	SendingDate           *time.Time
	AozoraTransferApplyNo string
	TransferRequestedAt   *time.Time
	TransferExecutedAt    *time.Time
	TransferRequestError  string
	IdempotencyKey        string

	Shop        *merchant.Merchant
	Payout      *Payout
	Transaction *transaction.Transaction
}
