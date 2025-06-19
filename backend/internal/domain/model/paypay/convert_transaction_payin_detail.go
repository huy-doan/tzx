package model

import (
	apiModel "github.com/test-tzs/nomraeite/internal/domain/model/api/makeshop"
	merchantModel "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// 加盟店のPaypay入金明細をTransactionに変換するためのモデル
type ConvertTransactionPayinDetail struct {
	ID int
	util.BaseColumnTimestamp
	Amount            int64
	PaymentMerchantID string

	Merchant *merchantModel.Merchant
	Shop     *apiModel.Shop
}
