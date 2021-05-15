package calc

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Fact function returns the factorial of a number, that is 1*2*3*...*number.
//
// Arguments
//
// number : Required. The number is a non-negative value for which the factorial is calculated.
//
// Remark
// If the number is negative, the function will return Invalid Input Error message.
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

// FactDouble func calculates the double factorial of the number. It is represented using double exclamation mark (!!).
//
// Arguments
//
// number : Required. The number is a non-negative number for which the double factorial is calculated.
//
// Remark
// If number is not an integer, it is truncated.
// If number is even : n!!= n(n-2)(n-4)...(4)(2)
// If number is odd : n!!= n(n-2)(n-4)...(3)(1)
// The double factorial for both zero and -1 are defined as 1.
// For numbers less than -1, a double factorial is not defined and FactDouble function will throw an error.
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
