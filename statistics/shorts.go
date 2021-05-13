package statistics

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Count, counts the number of items passes in the func.
func Count(v ...interface{}) int {
	c := utils.ForEach(func(_ int, x interface{}) {
	}, v...)
	return c
}

// CountA, counts the number of items passes in the func.
func CountA(v ...interface{}) int {
	c := utils.ForEach(func(_ int, x interface{}) {
	}, v...)
	return c
}

func Fisher(x interface{}) float64 {
	xF := utils.ToFloat64(x)
	if xF <= -1 || xF >= 1 {
		panic(core.ErrInvalidInput)
	}
	return 0.5 * calc.LN((1+xF)/(1-xF))
}

func FisherInv(y interface{}) float64 {
	yF := utils.ToFloat64(y)
	return calc.Round((math.Pow(math.E, yF*2)-1)/(math.Pow(math.E, yF*2)+1), 15)
}

func GeoMean(v ...interface{}) float64 {
	var vN float64 = 1
	c := utils.ForEach(func(_ int, x interface{}) {
		if utils.ToFloat64(x) <= 0 {
			panic(core.ErrInvalidInput)
		}
		vN *= utils.ToFloat64(x)
	}, v...)
	val := math.Pow(vN, calc.Divide(1, c))
	return val
}

func HarMean(v ...interface{}) float64 {
	var sum float64 = 0
	c := utils.ForEach(func(_ int, x interface{}) {
		if utils.ToFloat64(x) <= 0 {
			panic(core.ErrInvalidInput)
		}
		sum += calc.Divide(1, x)
	}, v...)
	return calc.Divide(1, calc.Divide(1, c)*sum)
}

func Vara(v ...interface{}) float64 {
	vA := Average(v)
	var sum float64
	c := utils.ForEach(func(_ int, x interface{}) {
		sum += (utils.ToFloat64(x) - vA) * (utils.ToFloat64(x) - vA)
	}, v...)
	return calc.Divide(sum, (c - 1))
}

func DevSQ(v ...interface{}) float64 {
	vA := Average(v)
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += (utils.ToFloat64(x) - vA) * (utils.ToFloat64(x) - vA)
	}, v...)
	return sum
}
