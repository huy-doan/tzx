package model

import (
	apiModel "github.com/test-tzs/nomraeite/internal/domain/model/api/makeshop"
	merchant "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	reconciliationModel "github.com/test-tzs/nomraeite/internal/domain/model/reconciliation"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	reconciliationObj "github.com/test-tzs/nomraeite/internal/domain/object/reconciliation"
)

// TransferRequest represents a transfer request in the domain model
type TransferRequest struct {
	TransactionID     int
	TransactionStatus int
	Amount            int64
	AccountNumber     string
	AccountName       string
	BranchName        string
	BankBranchCode    string
	BankCode          string
	Merchant          *merchant.Merchant
	Reconciliation    *reconciliationModel.PayinReconciliation
	Shop              *apiModel.Shop

	util.BaseColumnTimestamp
}

// PaginatedTransferRequest represents a paginated result of transfer requests
type PaginatedTransferRequest struct {
	Items      []*TransferRequest
	Pagination util.Pagination
}

// GetReconciliationStatus returns the reconciliation status of the transfer request
func (tr *TransferRequest) GetReconciliationStatus() string {
	return reconciliationObj.GenerateReconciliationStatus(
		reconciliationObj.ReconciliationValue(tr.Reconciliation.PayinSummaryVsBankIncoming),
		reconciliationObj.ReconciliationValue(tr.Reconciliation.PayinSummaryVsPayinDetail),
		reconciliationObj.ReconciliationValue(tr.Reconciliation.PayinDetailSumVsPayinTransactionSum),
		reconciliationObj.ReconciliationValue(tr.Reconciliation.PayinTransactionVsMakeshopOrder))
}
