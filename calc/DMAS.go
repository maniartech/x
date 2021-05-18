package calc

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Sum returns the sum of the total numbers supplied in the input argument.
// The values can be individual numbers, arrays of numbers or slices.
//
// Arguments
//
// number1 : Required. First input to sum.
// number2 : Optional. Second input to sum.
// '...' : Optional. Subsequent numbers arrys or slices.
//
// Example
//
//    calc.Sum(10, 20, 30, 50) // Returns 110.0
//    calc.Sum(12.5, 36.2, 88, 50, []int{55, 86})  // Returns 327.7
// 	  calc.Sum(13, 4, 456.2, 25.25456, 3543, []int{13, 48}, []float64{-0.1})) // Returns 4102.35456
//
//
func Sum(v ...interface{}) float64 {
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += utils.ToFloat64(x)
	}, v...)
	return Round(sum, 15-NumberOfDigits(utils.ToInt(sum)))
}

// Minus subtracts second number from the first one.
//
// Arguments
//
// number1 : Required. The number from which you want to substract.
//
// number2 : Required. The number which is substracted from the first number.
//
// Example
//
//    calc.Minus(100, 75) // Returns 25.0
//    calc.Minus(100.456, 75.123456)  // Returns 25.332544
// 	  calc.Minus(25, -75)  // Returns 100.0
//	  calc.Minus(25.125, -75.125) // Returns 100.25
//
func Minus(num1, num2 interface{}) float64 {
	ans := utils.ToFloat64(num1) - utils.ToFloat64(num2)
	return Round(ans, 15-NumberOfDigits(utils.ToInt(ans)))
}

// Product returns the product of all input numbers supplied in the argument.
// The values can be individual numbers, arrays of numbers or slices.
//
// Arguments
//
// number1 : Required. First input to Product.
// number2 : Optional. Second input to Product.
// '...' : Optional. Subsequent numbers arrys or slices.
//
// Example
//
//    calc.Product(10) // Returns 10.0
//	  calc.Product(10, 42) // Returns 420.0
//    calc.Product(20, 5, []int{2, 10}, -4)  // Returns -8000.0
//
func Product(v ...interface{}) float64 {
	var product float64 = 1.0
	utils.ForEach(func(_ int, x interface{}) {
		product = product * utils.ToFloat64(x)
	}, v...)
	return Round(product, 15-NumberOfDigits(utils.ToInt(product)))
}

// Divide function divides the first number by a second one.
//
// Arguments
//
// number1 : Required. The Divident.
// number2 : Required. The Divisor.
//
// Example
//
//    calc.Divide(100, 5)) // Returns 20.0
//    calc.Divide(100.456, 75.123456)  // Returns 1.33721217511612
//    calc.Divide(20, -5) // Returns -4.0
// 	  calc.Divide(25.125, -75.125) // Returns -0.334442595673877
//
func Divide(num1, num2 interface{}) float64 {
	if utils.ToFloat64(num2) == 0 {
		panic(core.ErrDivideBy0)
	}
	ans := utils.ToFloat64(num1) / utils.ToFloat64(num2)
	return Round(ans, 15-NumberOfDigits(utils.ToInt(ans)))
}
