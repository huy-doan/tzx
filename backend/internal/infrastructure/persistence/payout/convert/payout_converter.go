package convert

import (
	modelPayout "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	objectPayout "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	"github.com/test-tzs/nomraeite/internal/infrastructure/persistence/payout/dto"
	userDto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/user/dto"
	persistenceUtil "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

func ToPayoutDTO(payout *modelPayout.Payout) *dto.Payout {
	if payout == nil {
		return nil
	}

	return &dto.Payout{
		ID:                    payout.ID,
		PayoutStatus:          payout.PayoutStatus,
		TransferType:          payout.TransferType,
		Total:                 payout.Total,
		TotalCount:            payout.TotalCount,
		SendingDate:           payout.SendingDate,
		SentDate:              payout.SentDate,
		AozoraTransferApplyNo: payout.AozoraTransferApplyNo,
		ApprovalID:            payout.ApprovalID,
		UserID:                payout.UserID,
		User:                  userDto.ToUserDTO(payout.User),
		BaseColumnTimestamp: persistenceUtil.BaseColumnTimestamp{
			CreatedAt: payout.CreatedAt,
			UpdatedAt: payout.UpdatedAt,
		},
	}
}

func ToPayoutModel(dtoObj *dto.Payout) *modelPayout.Payout {
	if dtoObj == nil {
		return nil
	}

	return &modelPayout.Payout{
		ID:                    dtoObj.ID,
		PayoutStatus:          objectPayout.PayoutStatus(dtoObj.PayoutStatus),
		TransferType:          dtoObj.TransferType,
		Total:                 dtoObj.Total,
		TotalCount:            dtoObj.TotalCount,
		SendingDate:           dtoObj.SendingDate,
		SentDate:              dtoObj.SentDate,
		AozoraTransferApplyNo: dtoObj.AozoraTransferApplyNo,
		ApprovalID:            dtoObj.ApprovalID,
		UserID:                dtoObj.UserID,
		User:                  dtoObj.User.ToUserModel(),
		BaseColumnTimestamp: util.BaseColumnTimestamp{
			CreatedAt: dtoObj.CreatedAt,
			UpdatedAt: dtoObj.UpdatedAt,
		},
	}
}

func ToPayoutDTOs(payouts []*modelPayout.Payout) []*dto.Payout {
	if payouts == nil {
		return nil
	}

	result := make([]*dto.Payout, len(payouts))
	for i, payout := range payouts {
		result[i] = ToPayoutDTO(payout)
	}
	return result
}

func ToPayoutModels(dtos []*dto.Payout) []*modelPayout.Payout {
	if dtos == nil {
		return nil
	}

	result := make([]*modelPayout.Payout, len(dtos))
	for i, dto := range dtos {
		result[i] = ToPayoutModel(dto)
	}
	return result
}
