package calc

import (
	"github.com/maniartech/x/utils"
)

func calcGCD(a int, b int) int {
	if a == 0 {
		return b
	}
	return calcGCD(b%a, a)
}

// GCD function Returns the greatest common divisor of two or more non-zero integers without a remainder.
//
// Arguments
//
// number1 : number1 is required.
// number2, ... : number2 and subsequent numbers are optional.
// If the value is not an integer, it is truncated.
//
// Remark
//
// All the values are evenly divided by 1
// A Prime number can be evenly divided by itself or 1.
//
// Example
//
//    GCD(5, 1, 45) // Returns 1
//    GCD(24, 36)  // Returns 12
//
func GCD(v ...interface{}) int {
	var result int
	utils.ForEach(func(i int, x interface{}) {
		ival := utils.ToInt(x)
		if i == 0 {
			result = ival
			return
		}
		result = calcGCD(result, ival)
	}, v...)
	return result
}
