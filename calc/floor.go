package calc

import (
	"github.com/maniartech/x/utils"
)

// Floor function returns numeric value that is rounded down towords the zero for the input number,
// that is toward the zero, to the nearest multiple of significance.
//
// Arguments
//
// number : Required. The number is a value that you want to round off.
//
// significance : Required. The multiple to which you want to round.
//
// Remark :
//
// For the positively signed number, a value is rounded down and adjusted toward zero.
// For the negatively signed number, a value is rounded down and adjusted away from zero.
// No rounding occurs if the If number is an exact multiple of significance.
//
// Example
//
//    Floor(-26, -2) // Returns -26.0
//    Floor(1.5, 0.1)  // Returns 1.5
//
func Floor(input, factor interface{}) float64 {
	i := utils.ToDecimal(input)
	f := utils.ToDecimal(factor)

	// math.Floor(i / f) * f
	r, _ := i.Div(f).Floor().Mul(f).Float64()

	return r
}
