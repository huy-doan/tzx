package object

type BodyApplicationReviewResult struct {
	ReviewResults ApplicationReviewResult `json:"reviewResults"`
}

func NewBodyApplicationReviewResult(
	reviewResults ApplicationReviewResult,
) *BodyApplicationReviewResult {
	return &BodyApplicationReviewResult{
		ReviewResults: reviewResults,
	}
}

type ApplicationReviewResult struct {
	ShopID       string                         `json:"shopId"`
	MerchantID   string                         `json:"merchantId"`
	ReviewResult PaymentApplicationReviewResult `json:"reviewResult"`
}

func NewApplicationReviewResult(
	shopID string,
	merchantID string,
	reviewResult PaymentApplicationReviewResult,
) *ApplicationReviewResult {
	return &ApplicationReviewResult{
		ShopID:       shopID,
		MerchantID:   merchantID,
		ReviewResult: reviewResult,
	}
}
