package statistics

import (
	"math"

	"github.com/maniartech/x/utils"
)

//Max returns max number from the list of numbers
func Min(v ...interface{}) float64 {
	//Initializing the minimum value as the first item in the v...
	var min float64 = utils.ToFloat64(v[0])
	utils.ForEach(func(_ int, x interface{}) {
		//Finding the minimum value in v... by comparing the current min value and the items in v...
		min = math.Min(min, utils.ToFloat64(x))
	}, v...)
	return min
}

//MinA returns min number from the list of numbers and other values i.e. string or bool
func MinA(v ...interface{}) float64 {
	//Initializing the minimum value as the first item in the v...
	var min float64 = utils.ToFloat64(v[0])
	//Finding the minimum value in v... by comparing the current min value and the items in v...
	utils.ForEach(func(_ int, x interface{}) {
		switch x1 := x.(type) {
		//Also considering the value of bool and strings
		case bool:
			if x1 {
				min = math.Min(min, 1)
			} else {
				min = math.Min(min, 0)
			}
		case string:
			min = math.Min(min, 0)
		default:
			min = math.Min(min, utils.ToFloat64(x))
		}
	}, v...)
	return min
}
