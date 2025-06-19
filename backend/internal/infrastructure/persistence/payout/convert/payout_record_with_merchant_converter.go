package convert

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
)

func ToPayoutRecordWithMerchantIDModel(dto *dto.PayoutRecordWithMerchantID) *model.PayoutRecordWithMerchantID {
	if dto == nil {
		return nil
	}

	return &model.PayoutRecordWithMerchantID{
		PayoutRecord: *ToPayoutRecordModel(&dto.PayoutRecord),
		MerchantID:   dto.MerchantID,
	}
}

func ToPayoutRecordWithMerchantIDModels(dtos []*dto.PayoutRecordWithMerchantID) []*model.PayoutRecordWithMerchantID {
	if dtos == nil {
		return nil
	}

	result := make([]*model.PayoutRecordWithMerchantID, len(dtos))
	for i, dto := range dtos {
		result[i] = ToPayoutRecordWithMerchantIDModel(dto)
	}
	return result
}
