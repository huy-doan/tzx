package object

type ReconciliationFilterStatus int

const (
	ReconciliationFilterPerfectMatch ReconciliationFilterStatus = 1 // All reconciliation fields match
	ReconciliationFilterPartialMatch ReconciliationFilterStatus = 2 // Some reconciliation fields match
	ReconciliationFilterMismatch     ReconciliationFilterStatus = 3 // No reconciliation fields match
)

// ReconciliationFilterStatus constants
const (
	ReconciliationFilterPerfectMatchText = "照合完全一致"
	ReconciliationFilterPartialMatchText = "照合一部一致"
	ReconciliationFilterMismatchText     = "照合不一致"
)

// String returns the string representation of the reconciliation filter status
func (r ReconciliationFilterStatus) String() string {
	switch r {
	case ReconciliationFilterPerfectMatch:
		return ReconciliationFilterPerfectMatchText
	case ReconciliationFilterPartialMatch:
		return ReconciliationFilterPartialMatchText
	case ReconciliationFilterMismatch:
		return ReconciliationFilterMismatchText
	default:
		return ReconciliationFilterMismatchText
	}
}

// GenerateReconciliationStatus generates the reconciliation status based on the comparison results
func GenerateReconciliationStatus(payInSummaryVsBankIncoming ReconciliationValue, payInSummaryVsPayinDetail ReconciliationValue, payinDetailSumVsPayinTransactionSum ReconciliationValue, payinTransactionVsMakeshopOrder ReconciliationValue) string {
	if payInSummaryVsBankIncoming == ReconciliationValueMatch &&
		payInSummaryVsPayinDetail == ReconciliationValueMatch &&
		payinDetailSumVsPayinTransactionSum == ReconciliationValueMatch &&
		payinTransactionVsMakeshopOrder == ReconciliationValueMatch {
		return ReconciliationFilterPerfectMatchText
	}

	if payInSummaryVsBankIncoming == ReconciliationValueMatch ||
		payInSummaryVsPayinDetail == ReconciliationValueMatch ||
		payinDetailSumVsPayinTransactionSum == ReconciliationValueMatch ||
		payinTransactionVsMakeshopOrder == ReconciliationValueMatch {
		return ReconciliationFilterPartialMatchText
	}

	return ReconciliationFilterMismatchText
}
