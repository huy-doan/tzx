package service

import (
	"context"

	paypayModel "github.com/test-tzs/nomraeite/internal/domain/model/paypay"
	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
	"github.com/test-tzs/nomraeite/internal/domain/model/transaction/convert"
	transaction "github.com/test-tzs/nomraeite/internal/domain/repository/transaction"
	"github.com/test-tzs/nomraeite/internal/pkg/utils"
)

// CsvReaderService provides functionality to read and process CSV data
type TransferTransactionDomainService interface {
	ImportTransferTransaction(ctx context.Context, paypayPayinDetails []*paypayModel.ConvertTransactionPayinDetail) (err error)
}

type transferTransactionDomainService struct {
	transactionRepository       transaction.TransactionRepository
	transactionRecordRepository transaction.TransactionRecordRepository
}

func NewTransferTransactionDomainService(
	tr transaction.TransactionRepository,
	trr transaction.TransactionRecordRepository,
) *transferTransactionDomainService {
	return &transferTransactionDomainService{
		transactionRepository:       tr,
		transactionRecordRepository: trr,
	}
}

// CreateTransactionRecordService is a service that handles the creation of transaction records
func (ds transferTransactionDomainService) ImportTransferTransaction(ctx context.Context, convertTransactionPayinDetails []*paypayModel.ConvertTransactionPayinDetail) (err error) {

	transactionModels := []*transactionModel.Transaction{}

	for _, convertTransactionPayinDetail := range convertTransactionPayinDetails {
		transactionModel := convert.ToTransactionModel(convertTransactionPayinDetail)
		transactionModel.SetDraft()
		transactionModels = append(transactionModels, transactionModel)
	}

	// insert transaction
	transactions, err := ds.transactionRepository.BulkCreate(ctx, transactionModels)
	if err != nil {
		return err
	}

	// ShopID => paypayPayinDetail Map
	paypayPayinDetailShopIdMap := utils.ArrayToMap(convertTransactionPayinDetails, func(t *paypayModel.ConvertTransactionPayinDetail) string { return t.Shop.ShopID })

	// 1 Transactionに対して、3 TransactionRecordを作成する
	// 入金 TransactionRecord
	// 手数料 TransactionRecord
	// 振込手数料 TransactionRecord
	transactionRecordModels := []*transactionModel.TransactionRecord{}
	for _, transaction := range transactions {

		paypayPayinDetail := paypayPayinDetailShopIdMap[transaction.ShopID]

		depositTransactionRecord := transactionModel.NewDepositTransactionRecord(
			transaction.ID,
			paypayPayinDetail.Merchant.ID,
			paypayPayinDetail.ID,
			paypayPayinDetail.Amount,
		)

		feeTransactionRecord := transactionModel.NewFeeTransactionRecord(
			transaction.ID,
			paypayPayinDetail.Merchant.ID,
			paypayPayinDetail.Amount,
		)

		transferFeeTransactionRecord := transactionModel.NewTransferFreeTransactionRecord(
			transaction.ID,
			paypayPayinDetail.Merchant.ID,
			transaction.BankCode,
		)

		transactionRecordModels = append(transactionRecordModels, depositTransactionRecord)
		transactionRecordModels = append(transactionRecordModels, feeTransactionRecord)
		transactionRecordModels = append(transactionRecordModels, transferFeeTransactionRecord)
	}

	// insert transaction record
	return ds.transactionRecordRepository.BulkCreate(ctx, transactionRecordModels)

}
