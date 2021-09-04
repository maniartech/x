package statistics

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Skew(x ...interface{}) float64 {
	//Panicing if number of items in x... is < 3
	if len(x) < 3 {
		panic(core.ErrDivideBy0)
	}
	var sum float64
	s := StdevS(x...)
	xD := Average(x...)
	n := float64(len(x))
	//Summation of all the values of x in the formula
	utils.ForEach(func(_ int, xi interface{}) {
		sum += math.Pow((utils.ToFloat64(xi)-xD)/s, 3)
	}, x...)
	return (n * sum) / ((n - 1) * (n - 2))
}

func SkewP(x ...interface{}) float64 {
	//Panicing if number of items in x... is < 3
	if len(x) < 3 {
		panic(core.ErrDivideBy0)
	}
	s := StdevP(x...)
	//Panicing if s == 0
	if s == 0 {
		panic(core.ErrDivideBy0)
	}

	var sum float64
	xD := Average(x...)
	n := float64(len(x))
	//Summation of all the values of x in the formula
	utils.ForEach(func(_ int, xi interface{}) {
		sum += math.Pow((utils.ToFloat64(xi)-xD)/s, 3)
	}, x...)

	return sum / n
}
