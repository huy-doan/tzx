package inputdata

import (
	payoutModel "github.com/makeshop-jp/master-console/internal/domain/model/payout"
	payoutRecordModel "github.com/makeshop-jp/master-console/internal/domain/model/payout_record"
)

type TransferRequestParams struct {
	AccessToken   string
	Payout        *payoutModel.Payout
	PayoutRecords []*payoutRecordModel.PayoutRecord
}
