package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"

	"github.com/test-tzs/nomraeite/internal/pkg/utils"
)

type PaymentDetailItem struct {
	PaymentMethod string  `json:"paymentMethod"`
	Amount        float64 `json:"amount"`
}

// PaymentDetail represents PayPay payment details stored as JSON in the database
type PaymentDetail []PaymentDetailItem
type PayinReconciliationRelation = utils.JSONValue[PayinReconciliation]

func (pd *PaymentDetail) UnmarshalText(text []byte) error {
	// Handle empty or whitespace-only input
	textStr := strings.TrimSpace(string(text))
	if textStr == "" || textStr == "null" {
		*pd = PaymentDetail{}
		return nil
	}

	// Try to unmarshal as JSON into a temporary variable
	var temp []PaymentDetailItem
	if err := json.Unmarshal([]byte(textStr), &temp); err != nil {
		// If JSON parsing fails, treat as empty and return the error for debugging
		*pd = PaymentDetail{}
		return err
	}

	// Assign the temporary variable to the actual slice
	*pd = PaymentDetail(temp)
	return nil
}

// Value implements driver.Valuer interface for database storage
func (pd PaymentDetail) Value() (driver.Value, error) {
	if len(pd) == 0 {
		return "[]", nil
	}
	return json.Marshal(pd)
}

// Scan implements sql.Scanner interface for database retrieval
func (pd *PaymentDetail) Scan(value any) error {
	if value == nil {
		*pd = PaymentDetail{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("cannot scan non-string value into PaymentDetail")
	}

	if len(bytes) == 0 {
		*pd = PaymentDetail{}
		return nil
	}

	var temp []PaymentDetailItem
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return err
	}

	*pd = PaymentDetail(temp)
	return nil
}
