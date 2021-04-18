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
