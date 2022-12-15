package utils

import "reflect"

func IsStructEmpty(obj any) bool {
	return reflect.ValueOf(obj).IsZero()
}
