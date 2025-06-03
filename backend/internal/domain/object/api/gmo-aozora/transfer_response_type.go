package object

type TransferResponseType int

const (
	TransferResponseTypeSuccess TransferResponseType = iota
	TransferResponseTypeValidationError
	TransferResponseTypeAuthError
	TransferResponseTypeMaintenanceError
	TransferResponseTypeSystemError
	TransferResponseTypeDuplicateError
)

func (t TransferResponseType) ShouldStopBatch() bool {
	switch t {
	case TransferResponseTypeAuthError, TransferResponseTypeMaintenanceError, TransferResponseTypeSystemError:
		return true
	default:
		return false
	}
}

func (t TransferResponseType) String() string {
	switch t {
	case TransferResponseTypeSuccess:
		return "SUCCESS"
	case TransferResponseTypeValidationError:
		return "VALIDATION_ERROR"
	case TransferResponseTypeAuthError:
		return "AUTH_ERROR"
	case TransferResponseTypeMaintenanceError:
		return "MAINTENANCE_ERROR"
	case TransferResponseTypeSystemError:
		return "SYSTEM_ERROR"
	case TransferResponseTypeDuplicateError:
		return "DUPLICATE_ERROR"
	default:
		return "UNKNOWN"
	}
}
