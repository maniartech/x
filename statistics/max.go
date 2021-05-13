package statistics

import (
	"math"

	"github.com/maniartech/x/utils"
)

//Max returns max number from the list of numbers
func Max(v ...interface{}) float64 {
	var max float64
	utils.ForEach(func(_ int, x interface{}) {
		max = math.Max(max, utils.ToFloat64(x))
	}, v...)
	return max
}

//MaxA returns Max number from the list of numbers and other values i.e. string or bool
func MaxA(v ...interface{}) float64 {
	var max float64
	utils.ForEach(func(_ int, x interface{}) {
		switch x1 := x.(type) {
		case bool:
			if x1 {
				max = math.Max(max, 1)
			} else {
				max = math.Max(max, 0)
			}
		case string:
			max = math.Max(max, 0)
		default:
			max = math.Max(max, utils.ToFloat64(x))
		}
	}, v...)
	return max
}
