package statistics

import "github.com/maniartech/x/utils"

// Sum returns the sum of the total numbers
func Sum(v ...interface{}) float64 {
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += utils.ToFloat64(x)
	}, v...)
	return sum
}
