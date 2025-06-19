package model

import (
	"encoding/json"
	"time"
)

// PaymentProviderReviewInfo represents additional information for payment provider review
type PaymentProviderReviewInfo struct {
	SiteURL                string     `json:"site_url,omitempty"`
	ConfirmationChangedAt  *time.Time `json:"confirmation_changed_at,omitempty"`
	SaleReviewComment      string     `json:"sale_review_comment,omitempty"`
	RepresentativeResponse string     `json:"representative_response,omitempty"`
	MerchantDate           *time.Time `json:"merchant_date,omitempty"`
	MonthlyLimitAmount     int64      `json:"monthly_limit_amount,omitempty"`
	IssuanceStatus         string     `json:"issuance_status,omitempty"`
	ReviewStatus           string     `json:"review_status,omitempty"`
	RejectDate             *time.Time `json:"reject_date,omitempty"`
	PaymentMerchantID      string     `json:"payment_merchant_id"`
}

func (info *PaymentProviderReviewInfo) MarshalJson() []byte {
	jsonBytes, err := json.Marshal(info)
	if err != nil {
		// Return empty JSON object instead of nil
		return []byte("{}")
	}
	return jsonBytes
}

func UnmarshalJson(data []byte) *PaymentProviderReviewInfo {
	// Handle empty data
	if len(data) == 0 {
		return &PaymentProviderReviewInfo{}
	}

	var info PaymentProviderReviewInfo
	if err := json.Unmarshal(data, &info); err != nil {
		// Return default struct instead of nil to prevent panic
		return &PaymentProviderReviewInfo{}
	}
	return &info
}
