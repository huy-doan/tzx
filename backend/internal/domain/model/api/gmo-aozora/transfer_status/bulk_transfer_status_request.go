package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	aozoraObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora/transfer_status"
)

type BulkTransferStatusRequest struct {
	AccountID             string
	QueryKeyClass         aozoraObject.QueryKeyClass         // 1: 振込申請照会対象指定, 2: 振込一括照会対象指定
	DetailInfoNecessity   *bool                              // 総合振込明細情報の取得要否
	BulktransferItemKey   *string                            // 総合振込明細情報取得対象キー
	ApplyNo               *string                            // Required when QueryKeyClass = 1
	DateFrom              *string                            // YYYY-MM-DD format, available when QueryKeyClass = 2
	DateTo                *string                            // YYYY-MM-DD format, available when QueryKeyClass = 2
	NextItemKey           *string                            // Available when QueryKeyClass = 2
	RequestTransferStatus []object.GmoAozoraTransferStatus   // Transfer status array for filtering
	RequestTransferClass  *aozoraObject.RequestTransferClass // 1: ALL, 2: 振込申請のみ, 3: 振込受付情報のみ
	RequestTransferTerm   *aozoraObject.RequestTransferTerm  // 1: 振込申請受付日, 2: 振込指定日
}

func NewBulkTransferStatusRequestForApply(accessToken, accountID, applyNo string, detailInfoNecessity bool) BulkTransferStatusRequest {
	return BulkTransferStatusRequest{
		AccountID:           accountID,
		QueryKeyClass:       aozoraObject.QueryKeyClassApplication,
		ApplyNo:             &applyNo,
		DetailInfoNecessity: &detailInfoNecessity,
	}
}

func NewBulkTransferStatusRequestForBulk(accessToken, accountID string) BulkTransferStatusRequest {
	return BulkTransferStatusRequest{
		AccountID:     accountID,
		QueryKeyClass: aozoraObject.QueryKeyClassBulk,
	}
}
