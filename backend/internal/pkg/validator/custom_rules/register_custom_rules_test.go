package custom_rules

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestPasswordPolicyValidation(t *testing.T) {
	v := validator.New()
	RegisterCustomValidations(v)

	// Test passwordPolicyTag is registered in RegisterCustomValidations
	tests := []struct {
		password string
		valid    bool
	}{
		{"ValidPass123!", true}, // Example of a valid password
		{"short", false},        // Too short
	}

	for _, test := range tests {
		err := v.Var(test.password, PasswordPolicyTag)
		if (err == nil) != test.valid {
			t.Errorf("Password '%s' validation failed. Expected valid: %v, got error: %v", test.password, test.valid, err)
		}
	}
}
