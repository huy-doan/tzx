package model

import (
	"time"

	object "github.com/test-tzs/nomraeite/internal/domain/object/payment_provider_review"
)

type PaymentProviderReviewCSV struct {
	IsMajor                bool       `json:"is_major" csv:"大手フラグ"`
	IdDiv                  string     `json:"id_div" csv:"ID区分"`
	ShopID                 string     `json:"shop_id" csv:"契約先申請番号"`
	ApplicationAt          *time.Time `json:"application_at" csv:"申込み完了日"`
	CompanyName            string     `json:"company_name" csv:"加盟店接続企業"`
	CorporateName          string     `json:"corporate_name" csv:"法人名"`
	EntityName             string     `json:"entity_name" csv:"法人名、個人事業主名"`
	BusinessName           string     `json:"business_name" csv:"屋号名"`
	AppliedStoreCount      int        `json:"applied_store_count" csv:"申請店舗数"`
	PayoriginID            string     `json:"payorigin_id" csv:"PayOrigin18桁Id"`
	MasterAgentNote        string     `json:"master_agent_note" csv:"包括代理店_備考"`
	FeeRate                float64    `json:"fee_rate" csv:"決済手数料率"`
	SiteURL                string     `json:"site_url" csv:"オンライン決済利用サイトURL"`
	ConfirmationChangedAt  *time.Time `json:"confirmation_changed_at" csv:"営業確認ステータス変更日時"`
	SaleReviewComment      string     `json:"sale_review_comment" csv:"[審査]営業向けコメント"`
	RepresentativeResponse string     `json:"representative_response" csv:"申込担当者_回答"`
	MerchantDate           *time.Time `json:"merchant_date" csv:"審査可決日"`
	PaymentMerchantID      string     `json:"payment_merchant_id" csv:"MID"`
	MonthlyLimitAmount     int64      `json:"monthly_limit_amount" csv:"月間決済上限金額"`
	IssuanceStatus         string     `json:"issuance_status" csv:"ID発行状況"`
	ReviewStatus           string     `json:"review_status" csv:"審査ステータス"`
	RejectDate             *time.Time `json:"reject_date" csv:"審査否決日"`
}

func FromCSVModel(csvModel PaymentProviderReviewCSV, applicationReviewStatus object.ApplicationReviewStatus) *PaymentProviderReview {
	reviewInfo := PaymentProviderReviewInfo{
		SiteURL:                csvModel.SiteURL,
		ConfirmationChangedAt:  csvModel.ConfirmationChangedAt,
		SaleReviewComment:      csvModel.SaleReviewComment,
		RepresentativeResponse: csvModel.RepresentativeResponse,
		MerchantDate:           csvModel.MerchantDate,
		MonthlyLimitAmount:     csvModel.MonthlyLimitAmount,
		IssuanceStatus:         csvModel.IssuanceStatus,
		ReviewStatus:           csvModel.ReviewStatus,
		RejectDate:             csvModel.RejectDate,
		PaymentMerchantID:      csvModel.PaymentMerchantID,
	}
	result := &PaymentProviderReview{
		IsMajor:                 csvModel.IsMajor,
		IdDiv:                   csvModel.IdDiv,
		ShopID:                  csvModel.ShopID,
		ApplicationAt:           csvModel.ApplicationAt,
		CompanyName:             csvModel.CompanyName,
		EntityName:              csvModel.EntityName,
		BusinessName:            csvModel.BusinessName,
		AppliedStoreCount:       csvModel.AppliedStoreCount,
		PayoriginID:             csvModel.PayoriginID,
		MasterAgentNote:         csvModel.MasterAgentNote,
		FeeRate:                 csvModel.FeeRate,
		ReviewInfo:              reviewInfo,
		ApplicationReviewStatus: applicationReviewStatus,
	}

	return result
}
