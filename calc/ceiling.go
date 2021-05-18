package calc

import (
	"github.com/maniartech/x/utils"
)

// Ceiling function rounds up the given number to the nearest multiple of significance.
// That is the returned number is shifted away from zero and closer to the multiple of significance.
//
// Arguments
//
// number : Required. The number is a value that you want to round off.
//
// significance : Required. The multiple is used to round up the given number.
//
// Remark
//
// The value is always rounded up regardless of any sign of a given number when adjusted away from zero.
// If the given value is an exact multiple of significance then no rounding up occurs.
// If both the number and significance are negative, the value is rounded down, away from zero.
// If the number is negative, and significance is positive, the value is rounded up towards zero.
// If the input number is passed as a string the string value is converted to the desired format. If successfully converted it returns the expected outcome or throws an error.
//
// Example
//
//    Ceiling(-42.5, 16) // Returns -32.0
//    Ceiling(-32.25, -5)  // Returns -35.0
//    Ceiling(0.234, 0.01) // Returns 0.24
//    Ceiling(22.0/7.0, 0.000009)  // Returns 3.142863
//    calc.Ceiling("4.42", "0.05")  // Returns 4.45 after converting the nmbers in string format to decimal.
//
func Ceiling(val, factor interface{}) float64 {
	v := utils.ToDecimal(val)
	f := utils.ToDecimal(factor)

	x := v.Div(f).Ceil().Mul(f)
	rval, _ := x.Float64()
	return rval
}
