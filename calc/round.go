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
// num_digits : Required. The number of decimal points to which you want to round the number argument.
//
// Remark
//
// If the value of num_digits is greater than 0 (zero), then number is rounded to the specified number of decimal places.
// If the value of num_digits is 0, the number is rounded to the nearest integer.
// If the value of num_digits is less than 0, the number is rounded to the left of the decimal point.
//
// Example
//
//    Round(7.6549, 3) // Returns 7.655
//    Round(7.2522, 2)  // Returns 7.25
//
func Round(val interface{}, places interface{}) float64 {

	v := utils.ToDecimal(val)
	p := utils.ToInt(places)

	rval, _ := v.Round(int32(p)).Float64()
	return rval
}
