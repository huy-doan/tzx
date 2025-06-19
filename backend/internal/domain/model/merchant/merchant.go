package model

import (
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

// Merchant represents the merchant entity
type Merchant struct {
	ID                      int
	PaymentProviderID       int
	EntityName              string
	BusinessName            string
	SiteURL                 string
	IsMajor                 bool
	IdDiv                   string
	PaymentMerchantID       string
	ShopID                  string
	PaymentProviderReviewID int

	util.BaseColumnTimestamp
}
