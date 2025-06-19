package object

// ReconciliationValue represents individual reconciliation field comparison result
type ReconciliationValue int

// ReconciliationValue constants
const (
	ReconciliationValueMatch   ReconciliationValue = 1 // Field matches
	ReconciliationValueNoMatch ReconciliationValue = 2 // Field does not match
)

// String returns the string representation of the reconciliation value
func (r ReconciliationValue) String() string {
	switch r {
	case ReconciliationValueMatch:
		return "照合一致"
	case ReconciliationValueNoMatch:
		return "照合不一致"
	default:
		return "不明"
	}
}
