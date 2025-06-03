package adapter

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
)

type ApiClient interface {
	RefreshToken(ctx context.Context, refreshToken string) (*model.GmoAozoraTokenClaimResult, error)
	RequestTransfer(ctx context.Context, header model.TransferHeaderRequest, request model.TransferParamsRequest) *model.TransferResponse
}
