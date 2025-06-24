package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONSliceField[T any] []T

func (jf *JSONSliceField[T]) Scan(value any) error {
	if value == nil {
		*jf = make(JSONSliceField[T], 0)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, jf)
}

func (jf JSONSliceField[T]) Value() (driver.Value, error) {
	return json.Marshal(jf)
}
