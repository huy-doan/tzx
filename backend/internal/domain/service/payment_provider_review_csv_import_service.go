package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/test-tzs/nomraeite/internal/datastructure/outputdata"
	merchantModel "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	pprModel "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider_review"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payment_provider_review"
	merchantRepository "github.com/test-tzs/nomraeite/internal/domain/repository/merchant"
	pprRepository "github.com/test-tzs/nomraeite/internal/domain/repository/payment_provider_review"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
	"github.com/test-tzs/nomraeite/internal/pkg/utils/messages"
)

type PaymentProviderReviewCSVImportService interface {
	ImportCSVFile(ctx context.Context, filename string, reader io.Reader, paymentProviderId int) (*outputdata.PaymentProviderReviewImportOutput, error)
}

type paymentProviderReviewCSVImportServiceImpl struct {
	pprRepo      pprRepository.PaymentProviderReviewRepository
	merchantRepo merchantRepository.MerchantRepository
}

func NewPaymentProviderReviewCSVImportService(
	pprRepo pprRepository.PaymentProviderReviewRepository,
	merchantRepo merchantRepository.MerchantRepository,
) PaymentProviderReviewCSVImportService {
	return &paymentProviderReviewCSVImportServiceImpl{
		pprRepo:      pprRepo,
		merchantRepo: merchantRepo,
	}
}

func (s *paymentProviderReviewCSVImportServiceImpl) parseFromCSV(filename string, reader io.Reader, paymentProviderId int) ([]*pprModel.PaymentProviderReview, error) {
	applicationReviewStatus, exists := object.DetectReviewStatusFromFileName(filename)
	if !exists {
		return nil, errors.New(messages.MsgInvalidFileNameForReview)
	}

	csvReader, err := utils.NewCSVReader[pprModel.PaymentProviderReviewCSV](reader, nil)

	if err != nil {
		return nil, err
	}

	requiredHeaders := applicationReviewStatus.GetRequiredHeaders()

	if !csvReader.HasHeader(requiredHeaders) {
		return nil, fmt.Errorf(
			messages.MsgCSVHeaderMismatch,
			applicationReviewStatus.String(),
			requiredHeaders,
			csvReader.Header(),
		)
	}

	reviews, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var pprModels []*pprModel.PaymentProviderReview
	for _, review := range reviews {
		reviewModel := pprModel.FromCSVModel(review, applicationReviewStatus)
		if paymentProviderId > 0 {
			reviewModel.PaymentProviderID = object.PaymentProviderID(paymentProviderId)
		}
		pprModels = append(pprModels, reviewModel)
	}

	return pprModels, nil
}

// filterNewReviews filters out reviews that have already been changed since the last import.
func (s *paymentProviderReviewCSVImportServiceImpl) filterNewReviews(
	ctx context.Context,
	pprs []*pprModel.PaymentProviderReview,
) ([]*pprModel.PaymentProviderReview, error) {
	if len(pprs) == 0 {
		return nil, nil
	}

	shopIDs := make([]string, len(pprs))
	for i, ppr := range pprs {
		shopIDs[i] = ppr.ShopID
	}

	lastReviews, err := s.pprRepo.GetLastReviewByShopIDs(ctx, shopIDs)
	if err != nil {
		return nil, errors.New(messages.MsgFailedToGetLastReviewByShopIDs)
	}

	var newReviews []*pprModel.PaymentProviderReview
	for _, ppr := range pprs {
		// Check if the last review for this shop exists and if it matches the current review status
		lastReview, exists := lastReviews[ppr.ShopID]
		if !exists || lastReview == nil || lastReview.ApplicationReviewStatus != ppr.ApplicationReviewStatus {
			newReviews = append(newReviews, ppr)
		}
	}

	return newReviews, nil
}

func (s *paymentProviderReviewCSVImportServiceImpl) ImportCSVFile(ctx context.Context, filename string, reader io.Reader, paymentProviderId int) (*outputdata.PaymentProviderReviewImportOutput, error) {
	pprs, err := s.parseFromCSV(filename, reader, paymentProviderId)
	if err != nil {
		return nil, err
	}

	newPprs, err := s.filterNewReviews(ctx, pprs)
	if err != nil {
		return nil, err
	}

	var merchants []*merchantModel.Merchant
	for _, pprModel := range newPprs {
		if pprModel.ApplicationReviewStatus.IsApproved() {
			merchant := pprModel.GetMerchant()
			merchants = append(merchants, merchant)
		}
	}
	if len(merchants) > 0 {
		_, err := s.merchantRepo.ImportMerchants(ctx, merchants)
		if err != nil {
			return nil, err
		}
	}

	importedReviews, err := s.pprRepo.Import(ctx, newPprs)
	if err != nil {
		return nil, err
	}

	importedCount := len(importedReviews)
	skippedCount := len(pprs) - importedCount

	importResult := &outputdata.PaymentProviderReviewImportOutput{
		PaymentProviderReviews: importedReviews,
		ImportedCount:          importedCount,
		SkippedCount:           skippedCount,
	}

	return importResult, nil
}
