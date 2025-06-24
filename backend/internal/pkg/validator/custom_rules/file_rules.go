package custom_rules

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/oapi-codegen/runtime/types"
)

const IsCSVFileTag = "is_csv"

func IsCSVFile(fl validator.FieldLevel) bool {
	file, ok := fl.Field().Interface().(types.File)
	if !ok {
		return false
	}

	fileName := strings.ToLower(file.Filename())

	// Check if filename is long enough to contain ".csv"
	if len(fileName) < 4 {
		return false
	}

	return fileName[len(fileName)-4:] == ".csv"
}
