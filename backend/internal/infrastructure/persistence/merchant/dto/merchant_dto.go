package dto

import (
	merchantModel "github.com/test-tzs/nomraeite/internal/domain/model/merchant"
	persistence "github.com/test-tzs/nomraeite/internal/infrastructure/persistence/util"
)

// Merchant is the database representation of a merchant
type Merchant struct {
	ID                      int    `gorm:"column:id;primaryKey"`
	PaymentProviderID       int    `gorm:"column:payment_provider_id"`
	EntityName              string `gorm:"column:entity_name"`
	BusinessName            string `gorm:"column:business_name"`
	SiteURL                 string `gorm:"column:site_url"`
	IsMajor                 bool   `gorm:"column:is_major"`
	IdDiv                   string `gorm:"column:id_div"`
	PaymentMerchantID       string `gorm:"column:payment_merchant_id"`
	ShopID                  string `gorm:"column:shop_id"`
	PaymentProviderReviewID int    `gorm:"column:payment_provider_review_id"`

	persistence.BaseColumnTimestamp
}

// TableName specifies the table name for Merchant
func (Merchant) TableName() string {
	return "merchant"
}

func (dto *Merchant) ToModel() *merchantModel.Merchant {
	result := &merchantModel.Merchant{
		ID:                      dto.ID,
		PaymentProviderID:       dto.PaymentProviderID,
		EntityName:              dto.EntityName,
		BusinessName:            dto.BusinessName,
		SiteURL:                 dto.SiteURL,
		IsMajor:                 dto.IsMajor,
		IdDiv:                   dto.IdDiv,
		PaymentMerchantID:       dto.PaymentMerchantID,
		ShopID:                  dto.ShopID,
		PaymentProviderReviewID: dto.PaymentProviderReviewID,
	}

	result.CreatedAt = dto.CreatedAt
	result.UpdatedAt = dto.UpdatedAt

	return result
}

func ToDTO(m *merchantModel.Merchant) *Merchant {
	result := &Merchant{
		ID:                      m.ID,
		PaymentProviderID:       m.PaymentProviderID,
		EntityName:              m.EntityName,
		BusinessName:            m.BusinessName,
		SiteURL:                 m.SiteURL,
		IsMajor:                 m.IsMajor,
		IdDiv:                   m.IdDiv,
		PaymentMerchantID:       m.PaymentMerchantID,
		ShopID:                  m.ShopID,
		PaymentProviderReviewID: m.PaymentProviderReviewID,
	}

	result.CreatedAt = m.CreatedAt
	result.UpdatedAt = m.UpdatedAt

	return result
}

func ToModelArray(dtos []*Merchant) []*merchantModel.Merchant {
	if dtos == nil {
		return nil
	}

	result := make([]*merchantModel.Merchant, len(dtos))
	for i, dto := range dtos {
		result[i] = dto.ToModel()
	}

	return result
}

func ToDTOArray(models []*merchantModel.Merchant) []*Merchant {
	if models == nil {
		return nil
	}

	result := make([]*Merchant, len(models))
	for i, model := range models {
		result[i] = ToDTO(model)
	}

	return result
}

func (dto *Merchant) ToMerchantModel() *merchantModel.Merchant {
	result := &merchantModel.Merchant{
		ID:                      dto.ID,
		PaymentProviderID:       dto.PaymentProviderID,
		EntityName:              dto.EntityName,
		BusinessName:            dto.BusinessName,
		IsMajor:                 dto.IsMajor,
		IdDiv:                   dto.IdDiv,
		PaymentMerchantID:       dto.PaymentMerchantID,
		ShopID:                  dto.ShopID,
		SiteURL:                 dto.SiteURL,
		PaymentProviderReviewID: dto.PaymentProviderReviewID,
	}

	result.CreatedAt = dto.CreatedAt
	result.UpdatedAt = dto.UpdatedAt

	return result
}
