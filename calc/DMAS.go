package calc

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Sum returns the sum of the total numbers in the input.
//
// Arguments
//
// numbers : Required.
//
// Example
//
//    Sum(10, 20, 30, 50) // Returns 110.0
//    Sum(12.5, 36.2, 88, 50, []int{55, 86})  // Returns 327.7
//
func Sum(v ...interface{}) float64 {
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += utils.ToFloat64(x)
	}, v...)
	return Round(sum, 15-NumberOfDigits(utils.ToInt(sum)))
}

// Minus subtracts second number from the first
//
// Arguments
//
// number1 : Required. The number from which you want to substract.
//
// number2 : Required. The number which is substracted from the first number.
//
// Example
//
//    Minus(100, 75) // Returns 25.0
//    Minus(100.456, 75.123456)  // Returns 25.332544
//
func Minus(num1, num2 interface{}) float64 {
	ans := utils.ToFloat64(num1) - utils.ToFloat64(num2)
	return Round(ans, 15-NumberOfDigits(utils.ToInt(ans)))
}

// Product returns the product of all input numbers
//
// Arguments
//
// numbers : Required. All the numbers you want to multiply and get the product.
//
// Example
//
//    Product(10, 42)) // Returns 420.0
//    Product(20, 5, []int{2, 10}, -4)  // Returns -8000.0
//
func Product(v ...interface{}) float64 {
	var product float64 = 1.0
	utils.ForEach(func(_ int, x interface{}) {
		product = product * utils.ToFloat64(x)
	}, v...)
	return Round(product, 15-NumberOfDigits(utils.ToInt(product)))
}

// Divide returns the first number by a second one.
//
// Arguments
//
// number1 : Required. The Divident.
// number2 : Required. The Divisor.
//
// Example
//
//    Divide(100, 5)) // Returns 20.0
//    Divide(100.456, 75.123456)  // Returns 1.33721217511612
//
func Divide(num1, num2 interface{}) float64 {
	if utils.ToFloat64(num2) == 0 {
		panic(core.ErrDivideBy0)
	}
	ans := utils.ToFloat64(num1) / utils.ToFloat64(num2)
	return Round(ans, 15-NumberOfDigits(utils.ToInt(ans)))
}
