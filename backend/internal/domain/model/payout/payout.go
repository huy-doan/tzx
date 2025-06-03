package model

import (
	"time"

	userModel "github.com/test-tzs/nomraeite/internal/domain/model/user"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
)

// Update from existing model to include the TransferType field
type Payout struct {
	ID int
	util.BaseColumnTimestamp

	PayoutStatus          object.PayoutStatus
	TransferType          object.PayoutTransferType
	Total                 float64
	TotalCount            int
	SendingDate           time.Time
	SentDate              time.Time
	AozoraTransferApplyNo string
	ApprovalID            *int
	UserID                int
	User                  *userModel.User

	PayoutRecordCount     int
	PayoutRecordSumAmount float64
}

type NewPayoutParams struct {
	ID int
	util.BaseColumnTimestamp
	PayoutStatus          object.PayoutStatus
	TransferType          object.PayoutTransferType
	Total                 float64
	TotalCount            int
	SendingDate           time.Time
	SentDate              time.Time
	AozoraTransferApplyNo string
	ApprovalID            *int
	UserID                int
	User                  *userModel.User
	PayoutRecordCount     int
	PayoutRecordSumAmount float64
}

func NewPayout(params NewPayoutParams) *Payout {
	return &Payout{
		ID:                    params.ID,
		PayoutStatus:          params.PayoutStatus,
		TransferType:          params.TransferType,
		Total:                 params.Total,
		TotalCount:            params.TotalCount,
		SendingDate:           params.SendingDate,
		SentDate:              params.SentDate,
		AozoraTransferApplyNo: params.AozoraTransferApplyNo,
		ApprovalID:            params.ApprovalID,
		UserID:                params.UserID,
		User:                  params.User,
		PayoutRecordCount:     params.PayoutRecordCount,
		PayoutRecordSumAmount: params.PayoutRecordSumAmount,
		BaseColumnTimestamp:   params.BaseColumnTimestamp,
	}
}

func (p *Payout) SetStatus(status object.PayoutStatus) {
	p.PayoutStatus = status
}

func (p *Payout) IndividualTransfer() bool {
	return p.TransferType.IsIndividualTransfer()
}

func (p *Payout) BulkTransfer() bool {
	return p.TransferType.IsBulkTransfer()
}
