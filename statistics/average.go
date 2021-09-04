package statistics

import (
	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/utils"
)

// Average returns the average of the provided numbers.
func Average(v ...interface{}) float64 {
	var sum float64
	//Calcuating the sum of all the values and the total number of values
	c := utils.ForEach(func(_ int, x interface{}) {
		sum += utils.ToFloat64(x)
	}, v...)
	return calc.Divide(sum, c)
}
