package model

import (
	apiModel "github.com/test-tzs/nomraeite/internal/domain/model/api/makeshop"
	approvalModel "github.com/test-tzs/nomraeite/internal/domain/model/approval"
	merchant "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	payoutModel "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

type TransferTransaction struct {
	TransactionID     int
	TransactionStatus int
	Amount            int64
	AccountNumber     string
	AccountName       string
	BranchName        string
	BankBranchCode    string
	BankCode          string

	PayoutRecord *payoutModel.PayoutRecord
	Payout       *payoutModel.Payout
	Merchant     *merchant.Merchant
	Shop         *apiModel.Shop
	Approval     *approvalModel.ApprovalInfo
	util.BaseColumnTimestamp
}

// PaginatedTransferTransaction represents a paginated result of transfer transactions
type PaginatedTransferTransaction struct {
	Items      []*TransferTransaction
	Approval   *approvalModel.ApprovalInfo
	Pagination util.Pagination
}
