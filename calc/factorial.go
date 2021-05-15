package calc

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Fact function returns the factorial of a number,
// that is 1*2*3*...* number.
//
// Arguments
//
// number : Required. The number is a non-negative number for which the factorial is calculated.
// If number is not an integer, it is truncated.
//
// Example
//
//    Fact(0) // Returns 1
//    Fact(12)  // Returns 479001600
//
func Fact(input interface{}) int {
	num := utils.ToInt(input)
	if num < 0 {
		panic(core.ErrInvalidInput)
	}
	if num == 1 || num == 0 {
		return num
	}
	return num * Fact(num-1)
}

// FactDouble func calculates the double factorial of the number.
//
// Arguments
//
// number : Required. The number is a non-negative number for which the double factorial is calculated.
// If number is not an integer, it is truncated.
//
// Remark
//
// If number is even : n!!= n(n-2)(n-4)...(4)(2)
// If number is odd : n!!= n(n-2)(n-4)...(3)(1)
//
// Example
//
//    FactDouble(0) // Returns 1
//    FactDouble(12)  // Returns 46080
//
func FactDouble(input interface{}) int {
	num := utils.ToInt(input)
	if num < 0 {
		panic(core.ErrInvalidInput)
	}
	if num <= 2 {
		return num
	}
	return num * FactDouble(num-2)
}
