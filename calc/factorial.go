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
// If the number is negative, the function will return an Invalid Input Error message.
// If the number is not an integer, it is truncated.
//
// Example
//
//    calc.Fact(0) // Returns 1
//    calc.Fact(12)  // Returns 479001600
//    calc.Fact(1.4315) // Returns 1
//    calc.Fact(4.731) // Returns 24
//	  calc.Fact(-23)  // Returns Invalid Input Error Message
//
func Fact(input interface{}) int {
	num := utils.ToInt(input)
	if num < 0 {
		panic(core.ErrInvalidInput)
	}
	if num == 0 {
		return 1
	}
	if num-1 <= 1 {
		return num
	}
	return num * Fact(num-1)
}

// FactDouble function calculates the double factorial of the number. It is represented using a double exclamation mark (!!).
//
// Arguments
//
// number : Required. A number is a non-negative number for which the double factorial is calculated.
//
// Remark
//
// If number is not an integer, it is truncated.
// If number is even : n!!= n(n-2)(n-4)...(4)(2)
// If number is odd : n!!= n(n-2)(n-4)...(3)(1)
// The double factorial for both zero and -1 are defined as 1.
// For numbers less than -1, a double factorial is not defined and FactDouble function will throw an error.
//
// Example
//
//    calc.FactDouble(0) // Returns 1
//    calc.FactDouble(12)  // Returns 46080
//    calc.FactDouble(3.153) // Returns 6
//    calc.FactDouble(5.13551354) // Returns 120
//    calc.FactDouble(-23) // Returns Invalid Input Error Message
//
func FactDouble(input interface{}) int {
	num := utils.ToInt(input)
	if num < 0 {
		panic(core.ErrInvalidInput)
	}
	if num == 0 {
		return 1
	}
	if num-2 <= 2 {
		return num
	}
	return num * FactDouble(num-2)
}
