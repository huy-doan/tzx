package convert

import (
	model "github.com/test-tzs/nomraeite/internal/domain/model/paypay"
	transactionModel "github.com/test-tzs/nomraeite/internal/domain/model/transaction"
)

func ToTransactionModel(
	convertTransactionPayinDetail *model.ConvertTransactionPayinDetail,
) *transactionModel.Transaction {
	result := &transactionModel.Transaction{
		ShopID:            convertTransactionPayinDetail.Merchant.ShopID,
		BankCode:          convertTransactionPayinDetail.Shop.BankCode,
		BankBranch:        convertTransactionPayinDetail.Shop.BankBranch,
		BankBranchCode:    convertTransactionPayinDetail.Shop.BankBranchCode,
		AccountNumber:     convertTransactionPayinDetail.Shop.AccountNumber,
		AccountHolder:     convertTransactionPayinDetail.Shop.AccountHolder,
		AccountHolderKana: convertTransactionPayinDetail.Shop.AccountHolderKana,
		AccountKind:       convertTransactionPayinDetail.Shop.AccountKind,
	}

	return result
}
