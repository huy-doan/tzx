package model

import (
	object "github.com/test-tzs/nomraeite/internal/domain/object/api/gmo-aozora"
)

type BulkTransferHeaderRequest struct {
	AccessToken    string
	IdempotencyKey string
}

func NewBulkTransferHeaderRequest(accessToken, idempotencyKey string) BulkTransferHeaderRequest {
	return BulkTransferHeaderRequest{
		AccessToken:    accessToken,
		IdempotencyKey: idempotencyKey,
	}
}

type BulkTransferDetail struct {
	ItemID                string                 // 明細番号 (required)
	EdiInfo               string                 // EDI情報 (optional)
	BeneficiaryBankCode   string                 // 被仕向金融機関番号 (required)
	BeneficiaryBranchCode string                 // 被仕向支店番号 (required)
	AccountTypeCode       object.AccountTypeCode // 科目コード（預金種別コード） (required)
	AccountNumber         string                 // 口座番号 (required)
	BeneficiaryName       string                 // 受取人名 (required)
	TransferAmount        string                 // 振込金額 (required)
}

type BulkTransferParamsRequest struct {
	AccountID               string                         // 口座ID (required)
	TransferDesignatedDate  string                         // 振込指定日 (required)
	TransferDateHolidayCode object.TransferDateHolidayCode // 振込指定日休日コード (optional, default 1)
	TotalCount              string                         // 合計件数 (required)
	TotalAmount             string                         // 合計金額 (required)
	BulkTransfers           []BulkTransferDetail           // 総合振込明細情報 (required)
}
