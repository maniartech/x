package calc

import "github.com/maniartech/x/utils"

// LCM function returns the least common multiple of two or more integers.
// The LCM is the smallest positive integer that is a multiple of all integer in the argument.
//
// Arguments
//
// number1 :  number1 is required,
// number2, '...' : number2 and subsequent numbers are optional.
// If the number is not an integer, it is truncated.
//
// Example
//
//    LCM(15, 12) // Returns 60
//    LCM(4.5, 2.9)  // Returns 4
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
