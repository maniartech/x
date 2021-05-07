package calc

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//FACT func calculates the factorial of the number
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

//FACTDOUBLE func calculates the double factorial of the number
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
