package utils

import (
	"reflect"
)

// ForEach loops through variadic arguments. If an argument
// is found to be slice or array, it further loops through it.
// For example, if the arguments like {1, 2, 3}, 4, 5, {6, 7}
// is provided, where
// v1 = {1, 2, 3}
// v2 = 4
// v3 = 5
// v4 = {6, 7}
// It loops through individual numbers and calls the provided function
// for each of 1, 2, 3, 4, 5, 6, 7 values.
func ForEach(cb func(int, interface{}), v ...interface{}) int {
	return ForEachControlled(func(i int, x interface{}) bool {
		cb(i, x)
		return true
	}, v...)
}

func ForEachControlled(cb func(int, interface{}) bool, v ...interface{}) int {
	c := 0

	for i := 0; i < len(v); i++ {
		item := v[i]

		// If the item is nil
		rt := reflect.TypeOf(item)
		if item == nil {
			if !cb(c, nil) {
				break
			}
			c += 1
			continue
		}

		kind := rt.Kind()
		if kind == reflect.Array || kind == reflect.Slice {
			val := reflect.ValueOf(item)
			for j := 0; j < val.Len(); j++ {
				// fmt.Printf(">>> %v, %v\n", c, val.Index(j).Interface())
				if !cb(c, val.Index(j).Interface()) {
					break
				}
				c += 1
			}
		} else {
			if !cb(c, item) {
				break
			}
			c += 1
		}
	}
	return c
}
