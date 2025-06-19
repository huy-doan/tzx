package model

type PayoutRecordWithMerchantID struct {
	PayoutRecord
	MerchantID int
}

func (p *PayoutRecordWithMerchantID) GetPayoutRecord() *PayoutRecord {
	return &p.PayoutRecord
}
