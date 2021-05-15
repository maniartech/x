package calc

import (
	"github.com/maniartech/x/utils"
)

// Floor function rounds down the given number to the nearest multiple of significance.
// That is the returned number is shifted towards the zero and closer to the multiple of significance.
//
// Arguments
//
// number : Required. The number is a value that you want to round off.
//
// significance : Required. The multiple which is used to round down the given number.
//
// Remark :
// The value is always rouded down regardless sign of a given number when adjusted away from zero.
// The value will be returned in a float.
// If the given value is an exact multiple of significance then no rounding up occurs.
// If both the number and significance are negative, the value is rounded down, towards zero.
// If number is negative, and significance is positive, the value is rounded down away from zero.
// If the input number is passed as a string the string value is converted to desired format. If sucessfully converted it retuns the expected outcome or throws an error.
//
// Example
//
//    Floor(-26, -5) // Returns -25.0
//    Floor(-27.5, 4)  // Returns -28.0
//    Floor("1.5", "0.1") // Returns 1.5
//
func Floor(input, factor interface{}) float64 {
	i := utils.ToDecimal(input)
	f := utils.ToDecimal(factor)

	// math.Floor(i / f) * f
	r, _ := i.Div(f).Floor().Mul(f).Float64()

	return r
}
