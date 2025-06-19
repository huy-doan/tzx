package adapter

import (
	"context"

	model "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora"
	transferStatusModel "github.com/test-tzs/nomraeite/internal/domain/model/api/gmo-aozora/transfer_status"
)

type ApiClient interface {
	GetAuthURL(state string) (string, error)
	Connect(ctx context.Context, code string) (*model.GmoAozoraTokenClaimResult, error)
	RefreshToken(ctx context.Context, refreshToken string) (*model.GmoAozoraTokenClaimResult, error)
	RequestTransfer(ctx context.Context, header model.TransferHeaderRequest, request model.TransferParamsRequest) *model.TransferResponse
	RequestBulkTransfer(ctx context.Context, header model.BulkTransferHeaderRequest, request model.BulkTransferParamsRequest) *model.TransferResponse
	GetTransferStatus(ctx context.Context, header model.AuthHeader, request transferStatusModel.TransferStatusRequest) (*transferStatusModel.TransferStatusResponse, error)
	GetBulkTransferStatus(ctx context.Context, header model.AuthHeader, request transferStatusModel.BulkTransferStatusRequest) (*transferStatusModel.BulkTransferStatusResponse, error)
}
