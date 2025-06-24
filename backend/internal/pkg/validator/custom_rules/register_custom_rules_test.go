package custom_rules

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/oapi-codegen/runtime/types"
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

func TestIsCsvFile(t *testing.T) {
	v := validator.New()
	RegisterCustomValidations(v)

	// Test cases for IsCSVFile using a struct with validation tags
	type FileUpload struct {
		File types.File `validate:"is_csv"`
	}

	tests := []struct {
		filename string
		valid    bool
	}{
		{"file.csv", true},
		{"file.txt", false},
		{"file.backup.csv", true},
		{"file.csv.backup", false},
		{"test.CSV", true},
		{"", false},
		{"a.c", false},
		{"data.csv", true},
	}

	for _, test := range tests {
		var file types.File
		file.InitFromBytes([]byte("test content"), test.filename)

		upload := FileUpload{File: file}
		err := v.Struct(upload)

		if (err == nil) != test.valid {
			t.Errorf("File '%s' validation failed. Expected valid: %v, got error: %v",
				test.filename, test.valid, err)
		}
	}
}
