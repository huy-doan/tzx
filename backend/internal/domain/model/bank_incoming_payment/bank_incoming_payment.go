package model

import (
	"time"

	paymentProviderModel "github.com/test-tzs/nomraeite/internal/domain/model/payment_provider"
	basedatetime "github.com/test-tzs/nomraeite/internal/domain/object/basedatetime"
)

type BankIncomingPayment struct {
	ID                       int
	PaymentProviderID        int
	MessageID                *string
	EventCreatedAt           *time.Time
	TransactionDate          *time.Time
	ValueDate                *time.Time
	VaID                     *string
	VaBranchCode             *string
	VaBranchNameKana         *string
	VaAccountNumber          *string
	VaAccountNameKana        *string
	DepositAmount            int64
	RemitterNameKana         *string
	PaymentBankName          *string
	PaymentBranchName        *string
	PartnerName              *string
	Remarks                  *string
	ItemKey                  *string
	VaDepositTransactionJson VaDepositTransactionJson
	PaymentProvider          *paymentProviderModel.PaymentProvider
	basedatetime.BaseColumnTimestamp
}
