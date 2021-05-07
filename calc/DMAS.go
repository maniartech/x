package calc

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Sum returns the sum of the total numbers
func Sum(v ...interface{}) float64 {
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += utils.ToFloat64(x)
	}, v...)
	return Round(sum, 15-NumberOfDigits(utils.ToInt(sum)))
}

// Minus subtracts first number from the second
func Minus(num1, num2 interface{}) float64 {
	ans := utils.ToFloat64(num1) - utils.ToFloat64(num2)
	return Round(ans, 15-NumberOfDigits(utils.ToInt(ans)))
}

// Product returns the product of the total numbers
func Product(v ...interface{}) float64 {
	var product float64 = 1.0
	utils.ForEach(func(_ int, x interface{}) {
		product = product * utils.ToFloat64(x)
	}, v...)
	return Round(product, 15-NumberOfDigits(utils.ToInt(product)))
}

//Divide divides the first number from the second
func Divide(num1, num2 interface{}) float64 {
	if utils.ToFloat64(num2) == 0 {
		panic(core.ErrDivideBy0)
	}
	ans := utils.ToFloat64(num1) / utils.ToFloat64(num2)
	return Round(ans, 15-NumberOfDigits(utils.ToInt(ans)))
}
