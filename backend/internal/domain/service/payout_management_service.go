package service

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	repository "github.com/test-tzs/nomraeite/internal/domain/repository/payout"
)

type PayoutManagementService interface {
	ListPayouts(ctx context.Context, filter *model.PayoutFilter) ([]*model.Payout, int, int64, error)
}

type payoutManagementServiceImpl struct {
	payoutRepo       repository.PayoutRepository
	payoutRecordRepo repository.PayoutRecordRepository
}

func NewPayoutManagementService(
	payoutRepo repository.PayoutRepository,
	payoutRecordRepo repository.PayoutRecordRepository,
) PayoutManagementService {
	return &payoutManagementServiceImpl{
		payoutRepo:       payoutRepo,
		payoutRecordRepo: payoutRecordRepo,
	}
}

func (s *payoutManagementServiceImpl) ListPayouts(ctx context.Context, filter *model.PayoutFilter) ([]*model.Payout, int, int64, error) {
	payouts, totalPages, count, err := s.payoutRepo.List(ctx, filter)
	if err != nil {
		return nil, 0, 0, err
	}

	if len(payouts) > 0 {
		var payoutIDs []int
		for _, payout := range payouts {
			payoutIDs = append(payoutIDs, payout.ID)
		}

		counts, err := s.payoutRecordRepo.CountByPayoutIDs(ctx, payoutIDs)
		if err != nil {
			return nil, 0, 0, err
		}

		sums, err := s.payoutRecordRepo.SumAmountByPayoutIDs(ctx, payoutIDs)
		if err != nil {
			return nil, 0, 0, err
		}

		for _, payout := range payouts {
			payout.PayoutRecordCount = counts[payout.ID]
			payout.PayoutRecordSumAmount = sums[payout.ID]
		}
	}

	return payouts, totalPages, count, nil
}
