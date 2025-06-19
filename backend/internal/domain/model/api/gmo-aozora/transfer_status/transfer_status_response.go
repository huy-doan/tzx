package model

import (
	"time"

	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	aozoraObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora/transfer_status"
)

type TransferStatusResponse struct {
	AcceptanceKeyClass         aozoraObject.QueryKeyClass
	BaseDate                   string
	BaseTime                   string
	Count                      string
	TransferQueryBulkResponses []TransferQueryBulkResponse
	TransferDetails            []TransferStatusDetail
}

type TransferQueryBulkResponse struct {
	DateFrom                *string
	DateTo                  *string
	RequestNextItemKey      *string
	RequestTransferStatuses []RequestTransferStatusItem
	RequestTransferClass    *aozoraObject.RequestTransferClass
	RequestTransferTerm     *aozoraObject.RequestTransferTerm
	HasNext                 bool
	NextItemKey             string
}

// RequestTransferStatusItem represents individual transfer status filter
type RequestTransferStatusItem struct {
	RequestTransferStatus object.GmoAozoraTransferStatus
}

// TransfeStatusDetail represents detailed transfer information
type TransferStatusDetail struct {
	TransferStatus     object.GmoAozoraTransferStatus
	TransferStatusName string
	TransferTypeName   string
	IsFeeFreeUse       bool
	IsFeePointUse      bool
	PointName          *string
	FeeLaterPaymentFlg *bool
	TransferDetailFee  string
	TotalDebitAmount   string
	TransferApplies    []TransferApplyInfo
	TransferAccepts    []TransferAcceptInfo
	TransferResponses  []TransferResponseInfo
}

// TransferApplyInfo represents transfer application information
type TransferApplyInfo struct {
	ApplyNo              string
	TransferApplyDetails []TransferApplyDetailInfo
}

// TransferApplyDetailInfo represents detailed transfer application info
type TransferApplyDetailInfo struct {
	ApplyDatetime   *time.Time
	ApplyStatus     *aozoraObject.ApplyStatus
	ApplyUser       *string
	ApplyComment    *string
	ApprovalUser    *string
	ApprovalComment *string
}

// TransferAcceptInfo represents transfer acceptance information
type TransferAcceptInfo struct {
	AcceptNo       string
	AcceptDatetime time.Time
}

// TransferResponseInfo represents transfer response information
type TransferResponseInfo struct {
	AccountID              string
	RemitterName           string
	TransferDesignatedDate string
	TransferInfos          []TransferInfoDetail
}

// TransferInfoDetail represents individual transfer information detail
type TransferInfoDetail struct {
	TransferAmount          string
	EdiInfo                 *string
	BeneficiaryBankCode     string
	BeneficiaryBankName     string
	BeneficiaryBranchCode   string
	BeneficiaryBranchName   string
	AccountTypeCode         object.AccountTypeCode
	AccountNumber           string
	BeneficiaryName         string
	TransferDetailResponses []TransferDetailResponseInfo
	UnableDetailInfos       []UnableDetailInfo
}

// TransferDetailResponseInfo represents transfer detail response
type TransferDetailResponseInfo struct {
	BeneficiaryBankNameKanji   *string
	BeneficiaryBranchNameKanji *string
	UsedPoint                  *string
	IsFeeFreeUsed              *bool
	TransferFee                *string
}

type UnableDetailInfo struct {
	TransferDetailStatus *aozoraObject.TransferDetailStatus
	RefundStatus         *aozoraObject.RefundStatus
	IsRepayment          *bool
	RepaymentDate        *string
	UnableReasonDetail   *UnableReasonDetailInfo
}

type UnableReasonDetailInfo struct {
	UnableReasonCode    string
	UnableReasonMessage string
}

func (r *TransferStatusResponse) IsEmpty() bool {
	return r == nil || r.Count == "0" || len(r.TransferDetails) == 0
}

func (r *TransferStatusResponse) HasNext() bool {
	return r != nil && len(r.TransferQueryBulkResponses) > 0 && r.TransferQueryBulkResponses[0].HasNext
}

func (r *TransferStatusResponse) GetNextItemKey() string {
	if r == nil || len(r.TransferQueryBulkResponses) == 0 {
		return ""
	}
	return r.TransferQueryBulkResponses[0].NextItemKey
}

func (tsd *TransferStatusDetail) IsEmptyTransferApplies() bool {
	return tsd == nil || tsd.TransferApplies == nil || len(tsd.TransferApplies) == 0
}

func (tsd *TransferStatusDetail) IsNotCompleted() bool {
	notCompleted := []object.GmoAozoraTransferStatus{
		object.GmoAozoraTransferStatusApplication,
		object.GmoAozoraTransferStatusReserved,
		object.GmoAozoraTransferStatusInProcess,
		object.GmoAozoraTransferStatusRetrying,
	}
	isNotCompleted := true
	for _, status := range notCompleted {
		if tsd.TransferStatus == status {
			isNotCompleted = false
			break
		}
	}
	return isNotCompleted
}

func (tsd *TransferStatusDetail) GetTransferApply() (transferApplyInfo TransferApplyInfo) {
	if tsd == nil || len(tsd.TransferApplies) == 0 {
		return
	}
	transferApplyInfo = tsd.TransferApplies[0]
	return
}

func (tsd *TransferStatusDetail) GetTransferAcceptDetail() (transferAccept TransferAcceptInfo) {
	if tsd == nil || len(tsd.TransferAccepts) == 0 {
		return
	}
	transferAccept = tsd.TransferAccepts[0]
	return
}

func (tsd *TransferStatusDetail) GetTransferResponse() (transferResponse TransferResponseInfo) {
	if tsd == nil || len(tsd.TransferResponses) == 0 {
		return
	}
	transferResponse = tsd.TransferResponses[0]
	return
}

func (ta *TransferApplyInfo) GetTransferApplyDetail() (detail TransferApplyDetailInfo) {
	if ta == nil || len(ta.TransferApplyDetails) == 0 {
		return
	}
	detail = ta.TransferApplyDetails[0]
	return
}

func (tr *TransferResponseInfo) GetTransferInfo() (transferInfo TransferInfoDetail) {
	if tr == nil || len(tr.TransferInfos) == 0 {
		return
	}
	transferInfo = tr.TransferInfos[0]
	return
}
