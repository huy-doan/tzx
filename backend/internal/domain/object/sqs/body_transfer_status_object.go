package object

import "time"

type BodyTransferStatus struct {
	ApplyNo          string            `json:"applyNo"`
	ShopID           string            `json:"shopId"`
	PayoutID         string            `json:"payoutId"`
	ApplyDatetime    *time.Time        `json:"applyDatetime"`
	TransferAmount   int64             `json:"transferAmount"`
	AcceptDatetime   *time.Time        `json:"acceptDatetime"`
	AcceptNo         string            `json:"acceptNo"`
	TransferStatus   string            `json:"transferStatus"`
	PayPayAttributes *PayPayAttributes `json:"paypayAttributes"`
}

type PayPayAttributes struct {
	MerchantID string `json:"merchantId"`
}

func NewBodyTransferStatus(
	applyNo string,
	shopID string,
	payoutID string,
	applyDatetime *time.Time,
	transferAmount int64,
	acceptDatetime *time.Time,
	acceptNo string,
	transferStatus string,
	payPayAttributes *PayPayAttributes,
) *BodyTransferStatus {
	return &BodyTransferStatus{
		ApplyNo:          applyNo,
		ShopID:           shopID,
		PayoutID:         payoutID,
		ApplyDatetime:    applyDatetime,
		TransferAmount:   transferAmount,
		AcceptDatetime:   acceptDatetime,
		AcceptNo:         acceptNo,
		TransferStatus:   transferStatus,
		PayPayAttributes: payPayAttributes,
	}
}
