package object

type PaymentApplicationReviewResult int

const (
	PaymentApplicationReviewResultInProgress    PaymentApplicationReviewResult = 1
	PaymentApplicationReviewResultApproved      PaymentApplicationReviewResult = 2
	PaymentApplicationReviewResultRejected      PaymentApplicationReviewResult = 3
	PaymentApplicationReviewResultRequestChange PaymentApplicationReviewResult = 4
)
