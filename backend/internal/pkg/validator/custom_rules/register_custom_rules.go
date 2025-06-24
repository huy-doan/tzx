package custom_rules

import (
	"github.com/go-playground/validator/v10"
	"github.com/test-tzs/nomraeite/internal/pkg/logger"
)

func RegisterCustomValidations(v *validator.Validate) {
	logger := logger.GetLogger()
	err := v.RegisterValidation(PasswordPolicyTag, PasswordPolicy)
	if err != nil {
		logger.Error("Failed to register custom validation for PasswordPolicy", map[string]any{
			"error": err.Error(),
		})
	}

	// File rules
	err = v.RegisterValidation(IsCSVFileTag, IsCSVFile)
	if err != nil {
		logger.Error("Failed to register custom validation for IsCSVFile", map[string]any{
			"error": err.Error(),
		})
	}
}
