package model

import (
	"fmt"
	"time"

	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
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
	Total                 int64
	TotalCount            int
	SendingDate           *time.Time
	SentDate              *time.Time
	AozoraTransferApplyNo string
	ApprovalID            *int
	UserID                int
	User                  *userModel.User

	PayoutRecordCount     int
	PayoutRecordSumAmount int64
}

type NewPayoutParams struct {
	ID int
	util.BaseColumnTimestamp
	PayoutStatus          object.PayoutStatus
	TransferType          object.PayoutTransferType
	Total                 int64
	TotalCount            int
	SendingDate           *time.Time
	SentDate              *time.Time
	AozoraTransferApplyNo string
	ApprovalID            *int
	UserID                int
	User                  *userModel.User
	PayoutRecordCount     int
	PayoutRecordSumAmount int64
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

func (p *Payout) NormalTransfer() bool {
	return p.TransferType.IsNormalTransfer()
}

func (p *Payout) BulkTransfer() bool {
	return p.TransferType.IsBulkTransfer()
}

// CreateFromTransactionsParams contains parameters for creating a payout from transactions
type CreateFromTransactionsParams struct {
	TransactionDetails []*transactionModel.TransferTransactionDetail
	UserID             int
}

func ValidateTransactionsForPayout(transactionDetails []*transactionModel.TransferTransactionDetail) error {
	for _, detail := range transactionDetails {
		if !detail.IsEligibleForPayout() {
			return fmt.Errorf("transaction %d is not eligible for payout", detail.ID)
		}
	}
	return nil
}
