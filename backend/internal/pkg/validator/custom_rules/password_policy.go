package custom_rules

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

const PasswordPolicyTag = "password_policy"
const PasswordPolicyMinLength = 12

// PasswordPolicy validates that a password meets security requirements:
// - Minimum 12 characters
// - At least 1 uppercase letter
// - At least 1 lowercase letter
// - At least 1 number
// - At least 1 special character
func PasswordPolicy(fl validator.FieldLevel) bool {
	var (
		upperPattern   = regexp.MustCompile(`[A-Z]`)
		numberPattern  = regexp.MustCompile(`[0-9]`)
		lowerPattern   = regexp.MustCompile(`[a-z]`)
		specialPattern = regexp.MustCompile(`[!-/:-@[-` + "`" + `{-~]`)
		password       = fl.Field().String()
	)

	return len(password) >= PasswordPolicyMinLength &&
		upperPattern.MatchString(password) &&
		lowerPattern.MatchString(password) &&
		numberPattern.MatchString(password) &&
		specialPattern.MatchString(password)
}
