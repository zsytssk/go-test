package utils

import "reflect"

func IsZero(v interface{}) bool {
	val := reflect.ValueOf(v)
	return val.IsZero() // Go 1.13+
}
