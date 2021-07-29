package information

import (
	"fmt"
	"reflect"

	"github.com/maniartech/x/utils"
)

func IsLogical(input interface{}) bool {
	switch input.(type) {
	case bool:
		return true
	}
	return false
}

func IsNonText(input interface{}) bool {
	switch input.(type) {
	case string:
		return false
	}
	return true
}

func IsText(input interface{}) bool {
	switch input.(type) {
	case string:
		return true
	}
	return false
}

func IsNumber(input interface{}) bool {
	switch input.(type) {
	case byte:
		return true

		// Ints
	case int:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true

		// Uints
	case uint:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true

		// Floats
	case float32:
		return true
	case float64:
		return true
	}
	return false
}

func N(input interface{}) float64 {
	switch input := input.(type) {
	case bool:
		if input {
			return 1
		} else {
			return 0
		}
	case byte:
		return utils.ToFloat64(input)

		// Ints
	case int:
		return utils.ToFloat64(input)
	case int16:
		return utils.ToFloat64(input)
	case int32:
		return utils.ToFloat64(input)
	case int64:
		return utils.ToFloat64(input)

		// Uints
	case uint:
		return utils.ToFloat64(input)
	case uint16:
		return utils.ToFloat64(input)
	case uint32:
		return utils.ToFloat64(input)
	case uint64:
		return utils.ToFloat64(input)

		// Floats
	case float32:
		return utils.ToFloat64(input)
	case float64:
		return input
	case string:
		return -1
	}
	return 0
}

func IsEven(input interface{}) bool { return utils.ToInt(input)%2 == 0 }
func IsOdd(input interface{}) bool  { return utils.ToInt(input)%2 != 0 }

func Type(input interface{}) string { return fmt.Sprintf("%v", reflect.TypeOf(input)) }
