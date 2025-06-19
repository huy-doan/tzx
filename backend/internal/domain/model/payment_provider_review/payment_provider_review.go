package model

import (
	"time"

	sqsModel "github.com/test-tzs/nomraeite/internal/domain/model/api/sqs"
	merchantModel "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	paymentProviderModel "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider"
	basedatetime "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	commonObject "github.com/test-tzs/nomraeite/internal/domain/object/common"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payment_provider_review"
	sqsObject "github.com/test-tzs/nomraeite/internal/domain/object/sqs"
)

// PaymentProviderReview represents the payment provider review entity
type PaymentProviderReview struct {
	ID                      int
	ShopID                  string
	PaymentProviderID       object.PaymentProviderID
	IsMajor                 bool
	IdDiv                   string
	ApplicationAt           *time.Time
	CompanyName             string
	EntityName              string
	BusinessName            string
	AppliedStoreCount       int
	PayoriginID             string
	MasterAgentNote         string
	FeeRate                 float64
	ReviewInfo              PaymentProviderReviewInfo
	ApplicationReviewStatus object.ApplicationReviewStatus

	basedatetime.BaseColumnTimestamp

	// Relationships
	PaymentProvider *paymentProviderModel.PaymentProvider
	Merchant        *merchantModel.Merchant
}

func (p *PaymentProviderReview) ToSQSMessage() *sqsModel.ReviewResultMessage {
	reviewResult := sqsObject.ApplicationReviewResult{
		ShopID:       p.ShopID,
		MerchantID:   p.ReviewInfo.PaymentMerchantID,
		ReviewResult: p.ApplicationReviewStatus.SQSValue(),
	}

	messageBody := sqsObject.BodyApplicationReviewResult{
		ReviewResults: reviewResult,
	}

	payMethod := sqsObject.GetAttributePaymethodFromProvider(p.PaymentProvider.Code)
	messageAttributes := sqsObject.MessageAttribute{
		Paymethod:   &payMethod,
		MessageType: sqsObject.ApplyReviewResults,
	}

	return &sqsModel.ReviewResultMessage{
		MessageBody:       messageBody,
		MessageAttributes: messageAttributes,
	}
}

func ToSQSMessageBatch(reviews []*PaymentProviderReview) ([]*sqsModel.ReviewResultMessage, error) {
	var messages []*sqsModel.ReviewResultMessage

	for _, review := range reviews {
		message := review.ToSQSMessage()
		messages = append(messages, message)
	}

	return messages, nil
}

func (model *PaymentProviderReview) GetMerchant() *merchantModel.Merchant {
	return &merchantModel.Merchant{
		IsMajor:           model.IsMajor,
		IdDiv:             model.IdDiv,
		ShopID:            model.ShopID,
		PaymentMerchantID: model.ReviewInfo.PaymentMerchantID,
		EntityName:        model.EntityName,
		BusinessName:      model.BusinessName,
		SiteURL:           model.ReviewInfo.SiteURL,
		PaymentProviderID: model.ID,
	}
}

type PaymentProviderReviewListResult struct {
	PaymentProviderReviews []*PaymentProviderReview
	ShopIDs                []string
	ShopPagination         commonObject.Pagination
}
