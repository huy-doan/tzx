package inputdata

import (
	payoutModel "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	payoutRecordModel "github.com/test-tzs/nomraeite/internal/domain/model/payout_record"
)

type TransferRequestParams struct {
	AccessToken   string
	Payout        *payoutModel.Payout
	PayoutRecords []*payoutRecordModel.PayoutRecord
}
