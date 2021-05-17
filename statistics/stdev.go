package statistics

import (
	"math"

	"github.com/maniartech/x/core"
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

func Correl(x, y []interface{}) float64 {
	if len(x) == 0 || len(y) == 0 {
		panic(core.ErrDivideBy0)
	}
	xD := Average(x)
	yD := Average(y)

	var n float64
	var d1 float64
	var d2 float64
	var xF float64
	var yF float64
	for i := 0; i < len(x); i++ {
		xF = utils.ToFloat64(x[i])
		yF = utils.ToFloat64(y[i])

		n += (xF - xD) * (yF - yD)
		d1 += math.Pow((xF - xD), 2)
		d2 += math.Pow((yF - yD), 2)

	}
	return n / math.Sqrt(d1*d2)
}
