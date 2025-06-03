package custom_rules

import (
	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidations(v *validator.Validate) {
	_ = v.RegisterValidation(PasswordPolicyTag, PasswordPolicy)
}
