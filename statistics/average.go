package statistics

import "github.com/maniartech/x/utils"

// Average returns the average of the provided numbers.
func Average(v ...interface{}) float64 {
	var sum float64
	c := utils.ForEach(func(_ int, x interface{}) {
		sum += utils.ToFloat64(x)
	}, v...)
	return sum / float64(c)
}
