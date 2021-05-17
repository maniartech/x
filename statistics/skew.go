package statistics

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Skew(x ...interface{}) float64 {
	if len(x) < 3 {
		panic(core.ErrDivideBy0)
	}
	var sum float64
	s := StdevS(x...)
	xD := Average(x...)
	n := float64(len(x))
	utils.ForEach(func(_ int, xi interface{}) {
		sum += math.Pow((utils.ToFloat64(xi)-xD)/s, 3)
	}, x...)
	return (n * sum) / ((n - 1) * (n - 2))
}

func SkewP(x ...interface{}) float64 {
	if len(x) < 3 {
		panic(core.ErrDivideBy0)
	}
	s := StdevP(x...)
	if s == 0 {
		panic(core.ErrDivideBy0)
	}

	var sum float64
	xD := Average(x...)
	n := float64(len(x))

	utils.ForEach(func(_ int, xi interface{}) {
		sum += math.Pow((utils.ToFloat64(xi)-xD)/s, 3)
	}, x...)

	return sum / n
}
