package object

// ProviderOperation represents operations that payment providers can support as a domain value object
type ProviderOperation string

const (
	// OperationPayinSync represents payin data synchronization operation
	OperationPayinSync ProviderOperation = "payin_sync"
	// OperationPayinImport represents payin data import operation
	OperationPayinImport ProviderOperation = "payin_import"
	// OperationOrderHistoryImport represents order history import operation
	OperationOrderHistoryImport ProviderOperation = "order_history_import"
)

// String representation of ProviderOperation
func (o ProviderOperation) String() string {
	return string(o)
}

// Checks if the provider operation is valid
func (o ProviderOperation) IsValid() bool {
	switch o {
	case OperationPayinSync, OperationPayinImport, OperationOrderHistoryImport:
		return true
	default:
		return false
	}
}

// Creates a ProviderOperation from a string
func FromStringOperation(operation string) ProviderOperation {
	return ProviderOperation(operation)
}
