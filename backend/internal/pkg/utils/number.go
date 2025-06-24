package utils

import (
	"fmt"
	"strconv"
)

// ParseOptionalFloat parses a *string into float64.
// Returns 0 if nil or empty, and error if parsing fails.
func ParseOptionalFloat(fieldName string, s *string) (float64, error) {
	if s == nil || *s == "" {
		return 0, nil
	}
	f, err := strconv.ParseFloat(*s, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float in field %s: %w", fieldName, err)
	}
	return f, nil
}

func ParseCSVInt64Field(fieldName string, s string) (int64, error) {
	f, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid int64 in field %s: %w", fieldName, err)
	}
	return f, nil
}
