package utils

import (
	"reflect"
)

func CountNonNil(values ...interface{}) int {
	r := 0

	for i := range values {
		value := values[i]
		if value != nil {
			if !IsNull(value) {
				r++
			}
		}
	}

	return r
}

func IsNull(value interface{}) bool {
	if value == nil {
		return true
	}

	// typed nil interfaces
	v := reflect.ValueOf(value)

	return (v.Kind() == reflect.Ptr || v.Kind() == reflect.Slice) && v.IsNil()
}

func Ptr[T any](value T) *T {
	return &value
}
