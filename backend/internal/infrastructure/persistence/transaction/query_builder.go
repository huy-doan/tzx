package persistence

import (
	"strings"

	"github.com/test-tzs/nomraeite/internal/datastructure/inputdata"
	reconciliationObject "github.com/test-tzs/nomraeite/internal/domain/object/reconciliation"
	transactionStatusObj "github.com/test-tzs/nomraeite/internal/domain/object/transaction"
	transactionRecordType "github.com/test-tzs/nomraeite/internal/domain/object/transaction_record"
	"gorm.io/gorm"
)

type QueryBuilder struct {
	db *gorm.DB
}

func NewQueryBuilder(db *gorm.DB) *QueryBuilder {
	return &QueryBuilder{db: db}
}

func (qb *QueryBuilder) BuildTransferRequestBaseQuery() *gorm.DB {
	return qb.db.Table("transaction t").
		Joins("INNER JOIN transaction_record tr ON t.id = tr.transaction_id AND tr.transaction_record_type = ?", transactionRecordType.TransactionRecordTypeDeposit).
		Joins("INNER JOIN merchant m ON m.id = tr.merchant_id").
		Joins("LEFT JOIN payin_reconciliation pr ON pr.merchant_id = m.id AND pr.payin_detail_id = tr.payin_detail_id").
		Where("t.payout_id IS NULL AND t.transaction_status = ?", transactionStatusObj.TransactionStatusDraft)
}

func (qb *QueryBuilder) BuildTransactionDetailQuery(transactionIDs []int) *gorm.DB {
	return qb.db.Table("transaction t").
		Select(`
			t.*, 
			m.*,
			pr.*,
			t.bank_code,
			t.bank_branch,
			t.bank_branch_code,
			t.account_number,
			t.account_holder as account_name,
			t.account_kind,
			t.created_at,
			t.updated_at
		`).
		Joins("INNER JOIN transaction_record tr ON t.id = tr.transaction_id AND tr.transaction_record_type = ?", transactionRecordType.TransactionRecordTypeDeposit).
		Joins("INNER JOIN merchant m ON m.id = tr.merchant_id").
		Joins("LEFT JOIN payin_reconciliation pr ON pr.merchant_id = m.id AND pr.payin_detail_id = tr.payin_detail_id").
		Where("t.id IN (?)", transactionIDs)
}

func (qb *QueryBuilder) GetReconciliationFilterCondition(filter int) (string, []any) {
	switch reconciliationObject.ReconciliationFilterStatus(filter) {
	case reconciliationObject.ReconciliationFilterPerfectMatch:
		return `(
			pr.payin_summary_vs_bank_incoming = ? AND
			pr.payin_summary_vs_payin_detail = ? AND
			pr.payin_detail_sum_vs_payin_transaction_sum = ? AND
			pr.payin_transaction_vs_makeshop_order = ?
		)`, []any{
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
			}

	case reconciliationObject.ReconciliationFilterPartialMatch:
		return `(
			(pr.payin_summary_vs_bank_incoming = ? OR
			pr.payin_summary_vs_payin_detail = ? OR
			pr.payin_detail_sum_vs_payin_transaction_sum = ? OR
			pr.payin_transaction_vs_makeshop_order = ?)
			AND NOT
			(pr.payin_summary_vs_bank_incoming = ? AND
			pr.payin_summary_vs_payin_detail = ? AND
			pr.payin_detail_sum_vs_payin_transaction_sum = ? AND
			pr.payin_transaction_vs_makeshop_order = ?)
		)`, []any{
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
				int(reconciliationObject.ReconciliationValueMatch),
			}

	case reconciliationObject.ReconciliationFilterMismatch:
		return `(
			(pr.payin_summary_vs_bank_incoming = ? AND
			pr.payin_summary_vs_payin_detail = ? AND
			pr.payin_detail_sum_vs_payin_transaction_sum = ? AND
			pr.payin_transaction_vs_makeshop_order = ?)
			OR pr.id IS NULL
		)`, []any{
				int(reconciliationObject.ReconciliationValueNoMatch),
				int(reconciliationObject.ReconciliationValueNoMatch),
				int(reconciliationObject.ReconciliationValueNoMatch),
				int(reconciliationObject.ReconciliationValueNoMatch),
			}

	default:
		return "", nil
	}
}

func (qb *QueryBuilder) BuildTransactionSumSubquery() *gorm.DB {
	return qb.db.Table("transaction_record").
		Select("transaction_id, SUM(amount) AS total_amount").
		Group("transaction_id")
}

func (qb *QueryBuilder) BuildTransferRequestWithDetailsQuery(limit, offset int, reconciliationFilters []int, sortField, sortOrder string) *gorm.DB {
	transactionSumQuery := qb.BuildTransactionSumSubquery()

	query := qb.BuildTransferRequestBaseQuery().
		Select(`
			m.*,
			pr.*,
			t.id AS transaction_id,
			t.transaction_status,
			COALESCE(tr_sum.total_amount, 0) AS amount,
			t.created_at,
			t.updated_at,
			t.bank_code,
			t.bank_branch,
			t.bank_branch_code,
			t.account_number,
			t.account_holder AS account_name,
			t.account_kind
		`).
		Joins("LEFT JOIN (?) AS tr_sum ON t.id = tr_sum.transaction_id", transactionSumQuery).
		Limit(limit).
		Offset(offset)

	query = qb.ApplyReconciliationFilter(query, reconciliationFilters)
	query = qb.ApplySort(query, sortField, sortOrder)

	return query
}

func (qb *QueryBuilder) ApplyReconciliationFilter(query *gorm.DB, filters []int) *gorm.DB {
	if len(filters) == 0 {
		return query
	}

	var orConditions []string
	var args []any

	for _, filter := range filters {
		condition, filterArgs := qb.GetReconciliationFilterCondition(filter)
		if condition != "" {
			orConditions = append(orConditions, condition)
			args = append(args, filterArgs...)
		}
	}

	if len(orConditions) > 0 {
		whereClause := "(" + strings.Join(orConditions, " OR ") + ")"
		return query.Where(whereClause, args...)
	}

	return query
}

func (qb *QueryBuilder) ApplySort(query *gorm.DB, field, order string) *gorm.DB {
	columnMap := map[string]string{
		"merchant_name": "m.entity_name",
		"amount":        "tr_sum.total_amount",
		"provider_name": "pp.name",
	}

	if col, ok := columnMap[field]; ok {
		if order == "" {
			order = "asc"
		}
		return query.Order(col + " " + order)
	}

	return query.Order("t.id DESC")
}

// BuildTransferTransactionBaseQuery builds the base query for transfer transactions
func (qb *QueryBuilder) BuildTransferTransactionBaseQuery(params *inputdata.TransferTransactionInput) *gorm.DB {
	query := qb.db.Table("payout_record pr").
		Joins("JOIN payout p ON pr.payout_id = p.id").
		Joins("JOIN transaction t ON pr.transaction_id = t.id").
		Joins("INNER JOIN transaction_record tr ON t.id = tr.transaction_id AND tr.transaction_record_type = ?",
			transactionRecordType.TransactionRecordTypeDeposit).
		Joins("LEFT JOIN merchant m ON m.id = tr.merchant_id").
		Where("pr.payout_id IS NOT NULL").
		Where("p.id = ?", params.PayoutID)

	return query
}

// BuildTransferTransactionWithDetailsQuery creates a query for transfer transactions with details, pagination and filtering
func (qb *QueryBuilder) BuildTransferTransactionWithDetailsQuery(limit, offset int, params *inputdata.TransferTransactionInput) *gorm.DB {
	transactionSumQuery := qb.BuildTransactionSumSubquery()

	query := qb.BuildTransferTransactionBaseQuery(params).
		Select(`DISTINCT
			t.*,
			m.*,
			p.*,
			pr.*,
			COALESCE(tr_sum.total_amount, 0) AS amount
		`).
		Joins("LEFT JOIN (?) AS tr_sum ON t.id = tr_sum.transaction_id", transactionSumQuery).
		Limit(limit).
		Offset(offset)

	query = qb.ApplySort(query, params.SortField, params.SortOrder)

	return query
}
