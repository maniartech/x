package parser

import (
	"strconv"
	"strings"
)

type D map[string]interface{}

// Find recursively searches for the given key in the dictionary.
// If the key is found, the corresponding value is returned otherwise
// the function returns nil. It also returns the second value if the
// key is found the second value is true, otherwise it is false.
func (d D) Find(key string) (interface{}, bool) {
	return find(d, key)
}

func find(val interface{}, key string) (interface{}, bool) {
	keyArr := strings.Split(key, ".")
	k1 := keyArr[0]

	switch v := val.(type) {
	case D:
		if len(keyArr) == 1 {
			return v[k1], true
		}
		return find(v[k1], strings.Join(keyArr[1:], "."))
	case map[string]interface{}:
		if len(keyArr) == 1 {
			return v[k1], true
		}
		return find(v[k1], strings.Join(keyArr[1:], "."))
	case []interface{}:
		index, err := strconv.ParseInt(k1, 10, 64)
		if err != nil {
			return nil, false
		}
		val := v[int(index)]
		if len(keyArr) == 1 {
			return val, true
		}
		return find(val, strings.Join(keyArr[1:], "."))
	}
	return nil, false
}
