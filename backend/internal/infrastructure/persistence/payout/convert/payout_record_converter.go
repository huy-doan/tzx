package convert

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	bankAccountObject "github.com/test-tzs/nomraeite/internal/domain/object/bank_account"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

func ToPayoutRecordDTO(record *model.PayoutRecord) *dto.PayoutRecord {
	if record == nil {
		return nil
	}

	return &dto.PayoutRecord{
		ID:                    record.ID,
		ShopID:                record.ShopID,
		PayoutID:              record.PayoutID,
		TransactionID:         record.TransactionID,
		BankName:              record.BankAccount.BankName,
		BankCode:              record.BankAccount.BankCode.Value(),
		BranchName:            record.BankAccount.BranchName,
		BranchCode:            record.BankAccount.BranchCode.Value(),
		BankAccountType:       record.BankAccount.BankAccountType,
		AccountNo:             record.BankAccount.AccountNo.Value(),
		AccountName:           string(record.BankAccount.AccountName),
		Amount:                record.Amount,
		TransferStatus:        record.TransferStatus,
		SendingDate:           record.SendingDate,
		AozoraTransferApplyNo: record.AozoraTransferApplyNo,
		AozoraTransferStatus:  record.AozoraTransferStatus,
		TransferRequestedAt:   record.TransferRequestedAt,
		TransferExecutedAt:    record.TransferExecutedAt,
		TransferRequestError:  record.TransferRequestError,
		IdempotencyKey:        record.IdempotencyKey,
		BaseColumnTimestamp: persistence.BaseColumnTimestamp{
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		},
	}
}

func ToPayoutRecordModel(dto *dto.PayoutRecord) *model.PayoutRecord {
	if dto == nil {
		return nil
	}

	return &model.PayoutRecord{
		ID:            dto.ID,
		ShopID:        dto.ShopID,
		PayoutID:      dto.PayoutID,
		TransactionID: dto.TransactionID,
		BankAccount: model.BankAccount{
			BankName:        dto.BankName,
			BranchName:      dto.BranchName,
			BankCode:        bankAccountObject.BankCode(dto.BankCode),
			BranchCode:      bankAccountObject.BankBranchCode(dto.BranchCode),
			AccountNo:       bankAccountObject.AccountNumber(dto.AccountNo),
			AccountName:     bankAccountObject.AccountHolderKana(dto.AccountName),
			BankAccountType: dto.BankAccountType,
		},
		Amount:                dto.Amount,
		TransferStatus:        dto.TransferStatus,
		SendingDate:           dto.SendingDate,
		AozoraTransferApplyNo: dto.AozoraTransferApplyNo,
		AozoraTransferStatus:  dto.AozoraTransferStatus,
		TransferRequestedAt:   dto.TransferRequestedAt,
		TransferExecutedAt:    dto.TransferExecutedAt,
		TransferRequestError:  dto.TransferRequestError,
		IdempotencyKey:        dto.IdempotencyKey,
		BaseColumnTimestamp: util.BaseColumnTimestamp{
			CreatedAt: dto.CreatedAt,
			UpdatedAt: dto.UpdatedAt,
		},
	}
}

func ToPayoutRecordDTOs(records []*model.PayoutRecord) []*dto.PayoutRecord {
	if records == nil {
		return nil
	}

	result := make([]*dto.PayoutRecord, len(records))
	for i, record := range records {
		result[i] = ToPayoutRecordDTO(record)
	}
	return result
}

func ToPayoutRecordModels(dtos []*dto.PayoutRecord) []*model.PayoutRecord {
	if dtos == nil {
		return nil
	}

	result := make([]*model.PayoutRecord, len(dtos))
	for i, dto := range dtos {
		result[i] = ToPayoutRecordModel(dto)
	}
	return result
}
