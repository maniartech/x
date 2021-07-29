package engineering

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func BitAnd(num1, num2 interface{}) int {
	x := utils.ToInt(num1)
	y := utils.ToInt(num2)

	if x < 0 || y < 0 || x >= 281474976710656 || y >= 281474976710656 {
		panic(core.ErrInvalidInput)
	}
	return (x & y)
}

func BitOr(num1, num2 interface{}) int {
	x := utils.ToInt(num1)
	y := utils.ToInt(num2)

	if x < 0 || y < 0 || x >= 281474976710656 || y >= 281474976710656 {
		panic(core.ErrInvalidInput)
	}
	return (x | y)
}

func BitXOR(num1, num2 interface{}) int {
	x := utils.ToInt(num1)
	y := utils.ToInt(num2)

	if x < 0 || y < 0 || x >= 281474976710656 || y >= 281474976710656 {
		panic(core.ErrInvalidInput)
	}
	return (x ^ y)
}

func BitLShift(num1, num2 interface{}) int {
	x := utils.ToInt(num1)
	y := utils.ToInt(num2)

	if x < 0 || y < 0 || x >= 281474976710656 || y >= 281474976710656 {
		panic(core.ErrInvalidInput)
	}
	return (x << y)
}
func BitRShift(num1, num2 interface{}) int {
	x := utils.ToInt(num1)
	y := utils.ToInt(num2)

	if x < 0 || y < 0 || x >= 281474976710656 || y >= 281474976710656 {
		panic(core.ErrInvalidInput)
	}
	return (x >> y)
}
