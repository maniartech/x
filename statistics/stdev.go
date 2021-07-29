package statistics

import (
	"math"

	"github.com/maniartech/x/utils"
)

func StdevP(x ...interface{}) float64 {
	xD := Average(x...)
	var sum float64
	n := utils.ForEach(func(_ int, x interface{}) {
		switch x := x.(type) {
		case string:
			sum += 0
		case bool:
			if x {
				sum += (1 - xD) * (1 - xD)
			}
		default:
			sum += (utils.ToFloat64(x) - xD) * (utils.ToFloat64(x) - xD)
		}
	}, x...)
	return math.Sqrt(sum / utils.ToFloat64(n))
}

func StdevPA(x ...interface{}) float64 {
	xD := Average(x...)
	var sum float64
	n := utils.ForEach(func(_ int, x interface{}) {
		switch x := x.(type) {
		case string:
			sum += 0
		case bool:
			if x {
				sum += (1 - xD) * (1 - xD)
			} else {
				sum += (0 - xD) * (0 - xD)
			}
		default:
			sum += (utils.ToFloat64(x) - xD) * (utils.ToFloat64(x) - xD)
		}
	}, x...)
	return math.Sqrt(sum / utils.ToFloat64(n))
}

func StdevS(x ...interface{}) float64 {
	xD := Average(x...)
	var sum float64
	var xF float64
	n := utils.ForEach(func(_ int, x interface{}) {
		switch x.(type) {
		case string:
			xF = 0
		case bool:
			xF = 0
		default:
			xF = utils.ToFloat64(x)
		}
		sum += (xF - xD) * (xF - xD)
	}, x...)
	return math.Sqrt(sum / utils.ToFloat64(n-1))
}

func StdevA(x ...interface{}) float64 {
	xD := Average(x...)
	var sum float64
	n := utils.ForEach(func(_ int, x interface{}) {
		switch x := x.(type) {
		case string:
			sum += 0
		case bool:
			if x {
				sum += (1 - xD) * (1 - xD)
			} else {
				sum += (0 - xD) * (0 - xD)
			}
		default:
			sum += (utils.ToFloat64(x) - xD) * (utils.ToFloat64(x) - xD)
		}
	}, x...)
	return math.Sqrt(sum / utils.ToFloat64(n-1))
}
