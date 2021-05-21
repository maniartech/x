package statistics

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

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

func Kurt(x ...interface{}) float64 {
	n := utils.ToFloat64(len(x))
	if n <= 3 {
		panic(core.ErrDivideBy0)
	}
	s := StdevA(x...)
	xD := Average(x...)
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += math.Pow((utils.ToFloat64(x)-xD)/s, 4)
	}, x...)

	return calc.Round(((n*(n+1))/((n-1)*(n-2)*(n-3)))*sum-3*(n-1)*(n-1)/((n-2)*(n-3)), 14)
}

func Intercept(y, x []interface{}) float64 {
	if len(x) != len(y) {
		panic(core.ErrInvalidInput)
	}

	var xD float64 = Average(x)
	var yD float64 = Average(y)
	b := Slope(y, x)
	return yD - (b * xD)
}

func Slope(y, x []interface{}) float64 {

	if len(x) != len(y) {
		panic(core.ErrInvalidInput)
	}

	var n float64
	var d float64
	var xD float64 = Average(x)
	var yD float64 = Average(y)

	for i := 0; i < len(x); i++ {
		n += (utils.ToFloat64(x[i]) - xD) * (utils.ToFloat64(y[i]) - yD)
		d += (utils.ToFloat64(x[i]) - xD) * (utils.ToFloat64(x[i]) - xD)
	}
	return n / d
}

func STEYX(y, x []interface{}) float64 {
	if len(x) != len(y) {
		panic(core.ErrInvalidInput)
	}
	l := len(x)
	if l < 3 {
		panic(core.ErrDivideBy0)
	}
	xD := Average(x)
	yD := Average(y)

	var sumYS float64
	var n float64
	var d float64
	var xi float64
	var yi float64

	for i := 0; i < l; i++ {
		xi = utils.ToFloat64(x[i])
		yi = utils.ToFloat64(y[i])

		n += (xi - xD) * (yi - yD)
		d += math.Pow((xi - xD), 2)

		sumYS += math.Pow((yi - yD), 2)
	}

	return math.Sqrt(calc.Divide(sumYS-calc.Divide(n*n, d), l-2))
}

func Standardize(x, mean, sDev interface{}) float64 {
	return (utils.ToFloat64(x) - utils.ToFloat64(mean)) / utils.ToFloat64(sDev)
}

func Vara(v ...interface{}) float64 {
	vA := Average(v...)
	var sum float64
	c := utils.ForEach(func(_ int, x interface{}) {
		sum += (utils.ToFloat64(x) - vA) * (utils.ToFloat64(x) - vA)
	}, v...)
	return calc.Divide(sum, (c - 1))
}

func DevSQ(x ...interface{}) float64 {
	xD := Average(x...)
	var sum float64
	utils.ForEach(func(_ int, xi interface{}) {
		sum += (utils.ToFloat64(xi) - xD) * (utils.ToFloat64(xi) - xD)
	}, x...)
	return sum
}

func Forecast(b interface{}, y, x []interface{}) float64 {
	return Intercept(y, x) + Slope(y, x)*utils.ToFloat64(b)
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

func Gamma(x interface{}) float64 {
	if utils.ToFloat64(x) == 0 || utils.ToFloat64(x) < 0 && utils.ToFloat64(x)/utils.ToFloat64(utils.ToInt(x)) == 1 {
		panic(core.ErrInvalidInput)
	}
	return math.Gamma(utils.ToFloat64(x))
}
