package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
	aozoraObject "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora/transfer_status"
)

type TransferStatusRequest struct {
	AccountID             string
	QueryKeyClass         aozoraObject.QueryKeyClass        // 1: 振込申請照会対象指定, 2: 振込一括照会対象指定
	ApplyNo               string                            // Required when QueryKeyClass = 1
	DateFrom              string                            // YYYY-MM-DD format, available when QueryKeyClass = 2
	DateTo                string                            // YYYY-MM-DD format, available when QueryKeyClass = 2
	NextItemKey           string                            // Available when QueryKeyClass = 2
	RequestTransferStatus []object.GmoAozoraTransferStatus  // Transfer status array for filtering
	RequestTransferClass  aozoraObject.RequestTransferClass // 1: ALL, 2: 振込申請のみ, 3: 振込受付情報のみ
	RequestTransferTerm   aozoraObject.RequestTransferTerm  // 1: 振込申請受付日, 2: 振込指定日
}

func NewTransferStatusRequestForApply(accountID, applyNo string) TransferStatusRequest {
	return TransferStatusRequest{
		AccountID:     accountID,
		QueryKeyClass: aozoraObject.QueryKeyClassApplication,
		ApplyNo:       applyNo,
	}
}

func NewTransferStatusRequestForBulk(accessToken, accountID string) TransferStatusRequest {
	return TransferStatusRequest{
		AccountID:     accountID,
		QueryKeyClass: aozoraObject.QueryKeyClassBulk,
	}
}
