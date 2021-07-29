package utils

import "reflect"

// IsSliceOrArray returns true if the supplied variable is
// a slice or array. It returns `false` otherwise.
// If the value is a slice or an Array, it returns the
// length of the slice/array as second return value.
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
// Returns true if the character is a digit otherwise returns false.
func NotDigit(c rune) bool { return c < '0' || c > '9' }

// IsValidMonth accepts a "Month" and checks if that is a valid month
// Returns true if the value of the "Month" is betwern 1 and 12 otherwise returns false
func IsValidMonth(input interface{}) bool {
	month := ToInt(input)
	if month < 1 || month > 12 {
		return false
	}
	return true
}
