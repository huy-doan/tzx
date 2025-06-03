package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/makeshop-jp/master-console/internal/pkg/validator/custom_rules"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *CustomValidator {
	v := validator.New()

	custom_rules.RegisterCustomValidations(v)
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		if name == "" {
			name = fld.Name
		}

		if comma := strings.Index(name, ","); comma != -1 {
			name = name[:comma]
		}

		return name
	})

	return &CustomValidator{
		validator: v,
	}
}

func (cv *CustomValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}
