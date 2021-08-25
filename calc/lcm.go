package calc

import "github.com/maniartech/x/utils"

// LCM function returns the least common multiple of two or more integers.
// The LCM is the smallest positive integer that is a multiple of all integers in the argument.
//
// Arguments
//
// number1 :  number1 is required.
// number2, '...' : number2 and subsequent numbers are optional.
//
// Remark
//
// If the number is not an integer, it is truncated.
//
// Example
//
//    calc.LCM(15, 12) // Returns 60
//    calc.LCM(4.5, 2.9)  // Returns 4
//    calc.LCM(0, 0, 0, 0) // Returns 0
//    calc.LCM(4, 16, 8, 20) // Returns 80
//    calc.LCM(2.468, 4.135, 6.84648, 8.45, 10.11531) // Returns 120
//
func LCM(v ...interface{}) int {
	// Initialize result
	var result int
	utils.ForEach(func(i int, x interface{}) {
		ival := utils.ToInt(x)
		if ival == 0 {
			result = 0
			return
		}
		if i == 0 {
			result = ival
			return
		}
		result = (ival * result) / (calcGCD(ival, result))
	}, v...)
	return result
}
