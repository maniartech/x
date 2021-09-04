package statistics

import (
	"math"

	"github.com/maniartech/x/utils"
)

//Max returns max number from the list of numbers
func Max(v ...interface{}) float64 {
	//setting the max value as the 1st item in the v...
	var max float64 = utils.ToFloat64(v[0])
	//Using foreach to loop through all the items in v...
	utils.ForEach(func(_ int, x interface{}) {
		//using math.Max to find max number from the current max and current item in v...
		max = math.Max(max, utils.ToFloat64(x))
	}, v[1:]...)
	return max
}

//MaxA returns Max number from the list of numbers and other values i.e. string or bool
func MaxA(v ...interface{}) float64 {
	//setting the max value as the 1st item in the v...
	var max float64 = utils.ToFloat64(v[0])
	utils.ForEach(func(_ int, x interface{}) {
		//Considering values of boolean and string
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
	}, v[1:]...)
	return max
}
