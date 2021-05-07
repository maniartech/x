package utils

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

//TODO: Test pending!

// Loop through all the values and returns the
// interface list
func ToInterfaceSlice(v ...interface{}) []interface{} {
	list := make([]interface{}, 0)

	ForEach(func(_ int, val interface{}) {
		list = append(list, val)
	}, v...)
	return list
}

// ToFloat64 converts the generic interface value to float64
func ToFloat64(v interface{}) float64 {

	switch v := v.(type) {

	case nil:
		return 0

	// Bool
	case bool:
		if v {
			return 1
		}
		return 0

	// Bytes
	case byte:
		return float64(v)

		// Ints
	case int:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)

		// Uints
	case uint:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)

		// Floats
	case float32:
		return float64(v)

	default:
		f, e := strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
		if e != nil {
			panic(e)
		}
		return f
	}
}

// ToDecimal converts the generic interface value to decimal
func ToDecimal(v interface{}) decimal.Decimal {

	switch v := v.(type) {

	case nil:
		return decimal.Zero

	// Bool
	case bool:
		if v {
			return decimal.NewFromInt(1)
		}
		return decimal.Zero

	// Bytes
	case byte:
		return decimal.NewFromInt(int64(v))

		// Ints
	case int:
		return decimal.NewFromInt(int64(v))
	case int16:
		return decimal.NewFromInt(int64(v))
	case int32:
		return decimal.NewFromInt(int64(v))
	case int64:
		return decimal.NewFromInt(int64(v))

		// Uints
	case uint:
		return decimal.NewFromInt(int64(v))
	case uint16:
		return decimal.NewFromInt(int64(v))
	case uint32:
		return decimal.NewFromInt(int64(v))
	case uint64:
		return decimal.NewFromInt(int64(v))

		// Floats
	case float32:
		return decimal.NewFromFloat(float64(v))

	default:
		rval, err := decimal.NewFromString(fmt.Sprintf("%v", v))
		if err != nil {
			panic(err)
		}
		return rval
	}
}

// ToInt converts the interface value to the generic integer
func ToInt(v interface{}) int {
	switch v := v.(type) {

	case nil:
		return 0

	// Bool
	case bool:
		if v {
			return 1
		}
		return 0

	// Bytes
	case byte:
		return int(v)

		// Ints
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)

		// Uints
	case uint:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)

		// Floats
	case float32:
		return int(v)
	case float64:
		return int(v)

	default:
		i, e := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		if e != nil {
			panic(e)
		}
		return int(i)
	}
}

// ToInt converts the interface value to the generic integer
func ToInt64(v interface{}) int64 {
	switch v := v.(type) {

	case nil:
		return 0

	// Bool
	case bool:
		if v {
			return 1
		}
		return 0

	// Bytes
	case byte:
		return int64(v)

		// Ints
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return int64(v)

		// Uints
	case uint:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)

		// Floats
	case float32:
		return int64(v)
	case float64:
		return int64(v)

	default:
		i, e := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64)
		if e != nil {
			panic(e)
		}
		return int64(i)
	}
}
