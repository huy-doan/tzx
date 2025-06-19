package inputdata

import (
	payoutModel "github.com/test-tzs/nomraeite/internal/domain/model/payout"
)

type TransferRequestParams struct {
	AccessToken    string
	Payout         *payoutModel.Payout
	PayoutRecords  []*payoutModel.PayoutRecord
	IdempotencyKey string
}
