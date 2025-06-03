package dto

import (
	"time"

	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	payoutDto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

type PayoutRecord struct {
	ID int `json:"id"`
	persistence.BaseColumnTimestamp

	ShopID                int                    `json:"shop_id"`
	PayoutID              int                    `json:"payout_id"`
	TransactionID         int                    `json:"transaction_id"`
	BankName              string                 `json:"bank_name"`
	BankCode              string                 `json:"bank_code"`
	BranchName            string                 `json:"branch_name"`
	BranchCode            string                 `json:"branch_code"`
	BankAccountType       object.BankAccountType `json:"bank_account_type"`
	AccountNo             string                 `json:"account_no"`
	AccountName           string                 `json:"account_name"`
	Amount                float64                `json:"amount"`
	TransferStatus        object.TransferStatus  `json:"transfer_status"`
	SendingDate           *time.Time             `json:"sending_date"`
	AozoraTransferApplyNo string                 `json:"aozora_transfer_apply_no"`
	TransferRequestedAt   *time.Time             `json:"transfer_requested_at"`
	TransferExecutedAt    *time.Time             `json:"transfer_executed_at"`
	TransferRequestError  string                 `json:"transfer_request_error"`
	IdempotencyKey        string                 `json:"idempotency_key"`

	Payout *payoutDto.Payout `json:"payout,omitempty" gorm:"foreignKey:PayoutID"`
}

func (PayoutRecord) TableName() string {
	return "payout_record"
}
