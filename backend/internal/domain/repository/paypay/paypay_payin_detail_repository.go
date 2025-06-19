package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	model "github.com/test-tzs/nomraeite/internal/domain/model/paypay"
)

type PaypayPayinDetailRepository interface {
	BulkInsert(ctx context.Context, details []*model.PaypayPayinDetail) error
	DeleteByPayinFileID(ctx context.Context, payinFileID int) error
	GetNotConvertToTransactionRecord(ctx context.Context) (paypayPayinDetails []*model.ConvertTransactionPayinDetail, err error)
	// GetNotImportedPaginatedShopIDsByMonth gets paginated shop IDs for a specific month from paypay_payin_detail joined with merchant table
	GetNotImportedPaginatedShopIDsByMonth(ctx context.Context, targetMonth string, page int, pageSize int) ([]string, error)
	GetAllPaginatedShopIDsByMonth(ctx context.Context, targetMonth string, page int, pageSize int) ([]string, error)
	ListPaypayPayinDetails(ctx context.Context, params *inputdata.PaypayPayinDetailListInputData) (*model.PaginatedPaypayPayinDetailResult, error)
}
