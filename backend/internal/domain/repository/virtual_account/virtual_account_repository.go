package repository

import (
	"context"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	model "github.com/test-tzs/nomraeite/internal/domain/model/virtual_account"
)

type VirtualAccountRepository interface {
	ListVirtualAccounts(ctx context.Context, params *inputdata.VirtualAccountListInputData) (*model.PaginatedVirtualAccountResult, error)
	CreateVirtualAccount(ctx context.Context, va *model.VirtualAccount) error
	UpdateVirtualAccount(ctx context.Context, va *model.VirtualAccount) error
	FindByAccountNumber(ctx context.Context, name string) (*model.VirtualAccount, error)
	FindByID(ctx context.Context, id int) (*model.VirtualAccount, error)
	FindByAccountNumberAndBranchCode(ctx context.Context, vaAccountNumber string, vaBranchCode string) (*model.VirtualAccount, error)
}
