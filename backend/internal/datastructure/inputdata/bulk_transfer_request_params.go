package inputdata

import (
	payoutModel "github.com/test-tzs/nomraeite/internal/domain/model/payout"
)

type BulkTransferRequestParams struct {
	AccessToken   string
	Payout        *payoutModel.Payout
	PayoutRecords []*payoutModel.PayoutRecord
}
