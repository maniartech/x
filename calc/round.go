package calc

import (
	"github.com/maniartech/x/utils"
)

// Round function rounds a number to a specified number of digits.
//
// Arguments
//
// number : Required. The number that you want to round up.
//
// num_digits : Required. The number of decimal points to which number should be rounded.
//
// Remark
//
// If the value of num_digits is greater than 0 (zero), then the number is rounded to the specified number of decimal places.
// If the value of num_digits is 0, the number is rounded to the nearest integer.
// If the value of num_digits is less than 0, the number is rounded to the left of the decimal point.
//
// Example
//
//    calc.Round(7.6549, 3) // Returns 7.655
//    calc.Round(7.2522, 2)  // Returns 7.25
//    calc.Round(15.545, 0)  // Returns 16.0
//    calc.Round(15.545, 1)  // Returns 15.5
//    calc.Round(15.545, -1)  // Returns 20.0
//
func Round(val interface{}, places interface{}) float64 {

	v := utils.ToDecimal(val)
	p := utils.ToInt(places)

	rval, _ := v.Round(int32(p)).Float64()
	return rval
}
