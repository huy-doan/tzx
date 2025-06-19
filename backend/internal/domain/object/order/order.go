package object

type PaypayAttributes struct {
	PaymentID         string `json:"paymentId"`
	MerchantPaymentID string `json:"merchantPaymentId"`
}

type Order struct {
	ShopID                   string           `json:"shopId"`
	SystemOrderNumber        string           `json:"systemOrderNumber"`
	TransactionAmount        int64            `json:"amount"`
	PaymentMethod            string           `json:"payMethod"`
	PaymentTransactionStatus string           `json:"status"`
	OrderedAt                string           `json:"orderedAt"`
	PaypayAttributes         PaypayAttributes `json:"paypayAttributes"`
}
