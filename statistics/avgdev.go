package statistics

import (
	"math"

	"github.com/maniartech/x/utils"
)

// AveDev returns the average of the absolute deveation fo data points from their mean.
func AveDev(v ...interface{}) float64 {
	var c, d float64

	m := Average(v...)
	// calc the sum of diferences!

	utils.ForEach(func(_ int, item interface{}) {
		d += math.Abs(m - utils.ToFloat64(item))
		c += 1
	}, v...)

	return d / c
}
