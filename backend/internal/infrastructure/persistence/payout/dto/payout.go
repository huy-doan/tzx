package dto

import (
	"time"

	model "github.com/test-tzs/nomraeite/internal/domain/model/payout"
	util "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
	object "github.com/test-tzs/nomraeite/internal/domain/object/payout"
	approvalDto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/approval/dto"
	userDto "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/user/dto"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

type Payout struct {
	ID int
	persistence.BaseColumnTimestamp

	PayoutStatus          object.PayoutStatus
	TransferType          object.PayoutTransferType
	Total                 int64
	TotalCount            int
	SendingDate           *time.Time
	SentDate              *time.Time
	AozoraTransferApplyNo string
	ApprovalID            *int
	UserID                int

	User     *userDto.User         `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Approval *approvalDto.Approval `gorm:"foreignKey:ApprovalID"`
}

func (Payout) TableName() string {
	return "payout"
}

// ToModel converts a Payout to a domain model
func (p *Payout) ToModel() *model.Payout {
	if p == nil {
		return nil
	}

	return &model.Payout{
		ID:                    p.ID,
		PayoutStatus:          p.PayoutStatus,
		Total:                 p.Total,
		TotalCount:            p.TotalCount,
		SendingDate:           p.SendingDate,
		SentDate:              p.SentDate,
		AozoraTransferApplyNo: p.AozoraTransferApplyNo,
		ApprovalID:            p.ApprovalID,
		UserID:                p.UserID,
		BaseColumnTimestamp: util.BaseColumnTimestamp{
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
	}
}

// FromModel converts a domain model to a Payout
func FromModel(p *model.Payout) *Payout {
	if p == nil {
		return nil
	}

	return &Payout{
		ID:                    p.ID,
		PayoutStatus:          p.PayoutStatus,
		Total:                 p.Total,
		TotalCount:            p.TotalCount,
		SendingDate:           p.SendingDate,
		SentDate:              p.SentDate,
		AozoraTransferApplyNo: p.AozoraTransferApplyNo,
		ApprovalID:            p.ApprovalID,
		UserID:                p.UserID,
		BaseColumnTimestamp: persistence.BaseColumnTimestamp{
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		},
	}
}
