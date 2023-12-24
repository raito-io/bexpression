package utils

import "reflect"

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
	v := reflect.ValueOf(value)
	return v.Kind() == reflect.Ptr && v.IsNil()
}
