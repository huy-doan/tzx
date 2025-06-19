package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	aozoraObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora/transfer_status"
)

type BulkTransferStatusResponse struct {
	AcceptanceKeyClass         aozoraObject.QueryKeyClass
	DetailInfoNecessity        *bool
	BulktransferItemKey        *string
	BaseDate                   string
	BaseTime                   string
	Count                      string
	DetailInfoResult           *bool
	TransferQueryBulkResponses []BulkTransferQueryBulkResponse
	BulkTransferDetails        []BulkTransferStatusDetail
}

type BulkTransferQueryBulkResponse struct {
	DateFrom                *string
	DateTo                  *string
	RequestNextItemKey      *string
	RequestTransferStatuses []RequestTransferStatusItem
	RequestTransferClass    *aozoraObject.RequestTransferClass
	RequestTransferTerm     *aozoraObject.RequestTransferTerm
	HasNext                 bool
	NextItemKey             *string
}

type BulkTransferStatusDetail struct {
	TransferStatus        object.GmoAozoraTransferStatus
	TransferStatusName    string
	TransferTypeName      string
	RemitterCode          *string
	IsFeeFreeUse          bool
	IsFeePointUse         bool
	PointName             *string
	FeeLaterPaymentFlg    *bool
	TotalFee              string
	TotalDebitAmount      string
	TransferApplies       []TransferApplyInfo
	TransferAccepts       []TransferAcceptInfo
	BulktransferResponses []BulktransferResponseInfo
}

type BulktransferResponseInfo struct {
	AccountID              string
	RemitterName           string
	TransferDesignatedDate string
	TransferDataName       string
	TotalCount             string
	TotalAmount            string
	BulkTransferInfos      []BulkTransferInfoDetail
}

type BulkTransferInfoDetail struct {
	ItemID                  string
	BeneficiaryBankCode     string
	BeneficiaryBankName     *string
	BeneficiaryBranchCode   string
	BeneficiaryBranchName   *string
	ClearingHouseName       *string
	AccountTypeCode         object.AccountTypeCode
	AccountNumber           string
	BeneficiaryName         *string
	TransferAmount          string
	NewCode                 *string
	EdiInfo                 *string
	TransferDesignatedType  *string
	Identification          *string
	TransferDetailResponses []BulkTransferDetailResponseInfo
	UnableDetailInfos       []UnableDetailInfo
}

type BulkTransferDetailResponseInfo struct {
	BeneficiaryBankNameKanji   *string
	BeneficiaryBranchNameKanji *string
	UsedPoint                  *string
	IsFeeFreeUsed              *bool
	TransferFee                *string
}
