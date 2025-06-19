package model

import (
	"errors"
)

type GMOAozoraWebhookIncoming struct {
	AccessToken      string        `json:"access_token"`
	WebhookSignature string        `json:"webhook_signature"`
	MessageID        string        `json:"messageId"`
	Timestamp        string        `json:"timestamp"`
	Account          Accounts      `json:"account"`
	VaTransaction    VATransaction `json:"vaTransaction"`
}
type GMOAozoraWebhookBodyData struct {
	MessageID     string        `json:"messageId"`
	Timestamp     string        `json:"timestamp"`
	Account       Accounts      `json:"account"`
	VaTransaction VATransaction `json:"vaTransaction"`
}

type Accounts struct {
	RaID             string `json:"raId"`
	RaBranchCode     string `json:"raBranchCode"`
	RaBranchNameKana string `json:"raBranchNameKana"`
	RaAccountNumber  string `json:"raAccountNumber"`
	RaHolderName     string `json:"raHolderName"`
	BaseDate         string `json:"baseDate"`
	BaseTime         string `json:"baseTime"`
}

type VATransaction struct {
	VaID              string `json:"vaId"`
	TransactionDate   string `json:"transactionDate"`
	ValueDate         string `json:"valueDate"`
	VaBranchCode      string `json:"vaBranchCode"`
	VaBranchNameKana  string `json:"vaBranchNameKana"`
	VaAccountNumber   string `json:"vaAccountNumber"`
	VaAccountNameKana string `json:"vaAccountNameKana"`
	DepositAmount     string `json:"depositAmount"`
	RemitterNameKana  string `json:"remitterNameKana"`
	PaymentBankName   string `json:"paymentBankName"`
	PaymentBranchName string `json:"paymentBranchName"`
	PartnerName       string `json:"partnerName"`
	Remarks           string `json:"remarks"`
	ItemKey           string `json:"itemKey"`
}

func (g *GMOAozoraWebhookIncoming) ValidateWebhookToken(accessToken string) error {

	if g.AccessToken != accessToken {
		return errors.New("webhook token invalid")
	}
	return nil
}
