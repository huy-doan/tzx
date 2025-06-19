package repository

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/bank_incoming_payment"
)

type BankIncomingPaymentRepository interface {
	InsertWebhookIncomingPayment(ctx context.Context, service *model.BankIncomingPayment) error
}
