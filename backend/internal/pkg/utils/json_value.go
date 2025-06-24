package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONValue[T any] struct {
	Data T
}

func (jv JSONValue[T]) Value() (driver.Value, error) {
	return json.Marshal(jv.Data)
}

func (jv *JSONValue[T]) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &jv.Data)
}
