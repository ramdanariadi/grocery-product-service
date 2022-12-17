package utils

import "reflect"

func IsTypeEmpty(obj any) bool {
	return reflect.ValueOf(obj).IsZero()
}
