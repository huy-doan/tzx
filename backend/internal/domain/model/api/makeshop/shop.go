package model

import (
	paymentProviderReviewModel "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider_review"
	object "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	commonObject "github.com/test-tzs/nomraeite/internal/domain/object/common"
)

const (
	MaximumRequestedShopIDs = 100
)

// Shop represents the shop entity
type Shop struct {
	ShopID            string
	ShopName          string
	ShopURL           string
	BankName          string
	BankCode          object.BankCode
	BankBranch        object.BankBranch
	BankBranchCode    object.BankBranchCode
	AccountKind       object.AccountKind
	AccountNumber     object.AccountNumber
	AccountHolder     object.AccountHolder
	AccountHolderKana object.AccountHolderKana

	// Relations
	PaymentProviderReviews []*paymentProviderReviewModel.PaymentProviderReview
}

func (s *Shop) IsEmptyPaymentProviderReviews() bool {
	return len(s.PaymentProviderReviews) == 0
}

type ShopListResult struct {
	Shops []*Shop
	commonObject.Pagination
}
