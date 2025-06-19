package convert

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
)

func ToApplyNoModel(dto *dto.PayoutRecordApplyNo) model.PayoutRecordApplyNo {
	return model.PayoutRecordApplyNo{
		AozoraTransferApplyNo: dto.AozoraTransferApplyNo,
	}
}

func ToApplyNoModels(applyNos []*dto.PayoutRecordApplyNo) []model.PayoutRecordApplyNo {
	models := make([]model.PayoutRecordApplyNo, len(applyNos))
	for i, applyNo := range applyNos {
		models[i] = ToApplyNoModel(applyNo)
	}
	return models
}
