package information

import (
	"fmt"
	"reflect"

	"github.com/maniartech/x/utils"
)

func IsLogical(input interface{}) bool {
	//Returns true if the input is Logical Boolean Value
	switch input.(type) {
	case bool:
		return true
	}
	return false
}

func IsNonText(input interface{}) bool {
	//Returns true to all non text data types
	switch input.(type) {
	case string:
	case rune:
		return false
	}
	return true
}

func IsText(input interface{}) bool {
	switch input.(type) {
	//Returns true to all strings and runes
	case string:
	case rune:
		return true
	}
	return false
}

func IsNumber(input interface{}) bool {
	switch input.(type) {
	//Returns false to all non numeric data types
	case bool:
	case string:
	case rune:
		return false
	default:
		return true
	}
	return false
}

func N(input interface{}) float64 {
	//Returns Numeric values of data types
	switch input := input.(type) {
	case bool:
		if input {
			return 1
		} else {
			return 0
		}
	case float64:
		return input
	case string:
		return -1
	default:
		return utils.ToFloat64(input)
	}
}

func IsEven(input interface{}) bool { return utils.ToInt(input)%2 == 0 }
func IsOdd(input interface{}) bool  { return utils.ToInt(input)%2 != 0 }

func Type(input interface{}) string { return fmt.Sprintf("%v", reflect.TypeOf(input)) }
