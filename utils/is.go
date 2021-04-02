package utils

import "reflect"

// IsSliceOrArray returns true if the supplied variable is
// a slice or array. It returns false otherwise.
// If the firat return value is true, the second return value
// specifies whether it is slice or not.
func IsSliceOrArray(v interface{}) (bool, int) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice:
	case reflect.Array:
		return true, reflect.ValueOf(v).Len()
	}
	return false, -1
}

// IsFunc returns `true` if the supplied value is a function.
// Returns `false` othwise.
func IsFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

// NotDigit checks if supplied character is a numerical digit or not.
// Eeturns true if the character is a digit otherwise returns false.
func NotDigit(c rune) bool { return c < '0' || c > '9' }
