package convert

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	dto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
)

func ToSendingDateModel(dto *dto.PayoutRecordDate) model.PayoutRecordDate {
	return model.PayoutRecordDate{
		SendingDate: dto.SendingDate,
	}
}

func ToSendingDateModels(sendingDates []*dto.PayoutRecordDate) []model.PayoutRecordDate {
	models := make([]model.PayoutRecordDate, len(sendingDates))
	for i, sendingDate := range sendingDates {
		models[i] = ToSendingDateModel(sendingDate)
	}
	return models
}
