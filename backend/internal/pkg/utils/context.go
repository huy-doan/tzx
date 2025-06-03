package utils

import (
	"github.com/labstack/echo/v4"
)

func GetContextValue(ctx echo.Context, key string) (interface{}, bool) {
	value := ctx.Get(key)
	if value == nil {
		return nil, false
	}
	return value, true
}

func GetContextStringValue(ctx echo.Context, key string) (string, bool) {
	value, ok := GetContextValue(ctx, key)
	if !ok {
		return "", false
	}
	strValue, ok := value.(string)
	return strValue, ok
}

func GetContextIntValue(ctx echo.Context, key string) (int, bool) {
	value, ok := GetContextValue(ctx, key)
	if !ok {
		return 0, false
	}
	intValue, ok := value.(int)
	return intValue, ok
}
