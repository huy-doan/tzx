package outputdata

import (
	"time"

	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	payoutObject "github.com/test-tzs/nomraeite/internal/domain/object/payout"
)

type GmoAozoraTransferStatusOutput struct {
	ApplyNo                 string
	ApplyDatetime           *time.Time
	AcceptNo                string
	AcceptDatetime          *time.Time
	GmoAozoraTransferStatus object.GmoAozoraTransferStatus
	TransferStatus          payoutObject.PayoutRecordStatus
	PayoutRecordID          int
}
