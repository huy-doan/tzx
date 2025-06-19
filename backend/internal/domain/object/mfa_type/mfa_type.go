package object

type MFAType int

const (
	MFA_TYPE_EMAIL MFAType = 1
)

// String returns the string representation of the MFAType
func (m MFAType) String() string {
	switch m {
	case MFA_TYPE_EMAIL:
		return "Email"
	default:
		return "Unknown"
	}
}

// IsValidMFAType checks if the given MFA type is valid
func IsValidMFAType(mfaType int) bool {
	switch MFAType(mfaType) {
	case MFA_TYPE_EMAIL:
		return true
	default:
		return false
	}
}

// GetDefaultMFAType returns the default MFA type
func GetDefaultMFAType() int {
	return int(MFA_TYPE_EMAIL)
}
