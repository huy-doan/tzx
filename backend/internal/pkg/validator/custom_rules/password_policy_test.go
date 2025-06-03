package custom_rules

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestPasswordPolicy(t *testing.T) {
	validate := validator.New()
	if err := validate.RegisterValidation(PasswordPolicyTag, PasswordPolicy); err != nil {
		t.Fatalf("Failed to register validation: %v", err)
	}

	tests := []struct {
		name     string
		password string
		valid    bool
	}{
		{"ValidPassword", "StrongP@ssw0rd123", true},
		{"TooShort", "Short1!", false},
		{"NoUppercase", "nouppercase1!", false},
		{"NoLowercase", "NOLOWERCASE1!", false},
		{"NoNumber", "NoNumber!", false},
		{"NoSpecialCharacter", "NoSpecial123", false},
		{"EmptyPassword", "", false},
		{"OnlySpecialCharacters", "!@#$%^&*()", false},
		{"OnlyNumbers", "123456789012", false},
		{"OnlyUppercase", "UPPERCASEONLY", false},
		{"OnlyLowercase", "lowercaseonly", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Var(tt.password, PasswordPolicyTag)
			if (err == nil) != tt.valid {
				t.Errorf("PasswordPolicy(%q) = %v, want %v", tt.password, err == nil, tt.valid)
			}
		})
	}
}
