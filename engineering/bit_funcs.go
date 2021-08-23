package engineering

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

const limit int64 = 281474976710656 // 2 ^ 48

func BitAnd(num1, num2 interface{}) int64 {
	x := utils.ToInt64(num1)
	y := utils.ToInt64(num2)

	if x < 0 || y < 0 || x >= limit || y >= limit {
		panic(core.ErrInvalidInput)
	}
	return (x & y)
}

func BitOr(num1, num2 interface{}) int64 {
	x := utils.ToInt64(num1)
	y := utils.ToInt64(num2)

	if x < 0 || y < 0 || x >= limit || y >= limit {
		panic(core.ErrInvalidInput)
	}
	return (x | y)
}

func BitXOR(num1, num2 interface{}) int64 {
	x := utils.ToInt64(num1)
	y := utils.ToInt64(num2)

	if x < 0 || y < 0 || x >= limit || y >= limit {
		panic(core.ErrInvalidInput)
	}
	return (x ^ y)
}

func BitLShift(num1, num2 interface{}) int64 {
	x := utils.ToInt64(num1)
	y := utils.ToInt64(num2)

	if x < 0 || y < 0 || x >= limit || y >= limit {
		panic(core.ErrInvalidInput)
	}
	return (x << y)
}
func BitRShift(num1, num2 interface{}) int64 {
	x := utils.ToInt64(num1)
	y := utils.ToInt64(num2)

	if x < 0 || y < 0 || x >= limit || y >= limit {
		panic(core.ErrInvalidInput)
	}
	return (x >> y)
}
