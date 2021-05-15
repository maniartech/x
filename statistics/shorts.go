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

func ExponDist(x, lamda interface{}, cumulative bool) float64 {
	xF := utils.ToFloat64(x)
	lamdaF := utils.ToFloat64(lamda)

	if lamdaF <= 0 || xF < 0 {
		panic(core.ErrInvalidInput)
	}

	var val float64

	if cumulative {
		val = 1 - math.Pow(math.E, -lamdaF*xF)
	} else {
		val = lamdaF * math.Pow(math.E, -lamdaF*xF)
	}
	return val
}

//WIP
func Kurt(x ...interface{}) float64 {
	n := len(x)
	if n <= 3 {
		panic(core.ErrDivideBy0)
	}
	s := Vara(x...)
	xD := Average(x...)
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += math.Pow((utils.ToFloat64(x)-xD)/s, 4)
	}, x...)
	//*utils.ToFloat64((3*(n-1)*(n-1))/((n-2)*(n-3)))
	return calc.Round(utils.ToFloat64((n*(n+1))/(n-1)*(n-2)*(n-3))*sum, 15)

}

func Intercept(x, y []interface{}) float64 {
	if len(x) != len(y) {
		panic(core.ErrInvalidInput)
	}

	var xD float64 = Average(x)
	var yD float64 = Average(y)
	b := Slope(x, y)
	return yD - (b * xD)
}

func Slope(x, y []interface{}) float64 {

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

func Standardize(x, mean, sDev interface{}) float64 {
	return (utils.ToFloat64(x) - utils.ToFloat64(mean)) / utils.ToFloat64(sDev)
}

func StdevP(x ...interface{}) float64 {
	xD := Average(x...)
	var sum float64
	n := utils.ForEach(func(_ int, x interface{}) {
		sum += (utils.ToFloat64(x) - xD) * (utils.ToFloat64(x) - xD)
	}, x...)
	return math.Sqrt(sum / utils.ToFloat64(n))
}

func Vara(v ...interface{}) float64 {
	vA := Average(v...)
	var sum float64
	c := utils.ForEach(func(_ int, x interface{}) {
		sum += (utils.ToFloat64(x) - vA) * (utils.ToFloat64(x) - vA)
	}, v...)
	return calc.Divide(sum, (c - 1))
}

func DevSQ(v ...interface{}) float64 {
	vA := Average(v...)
	var sum float64
	utils.ForEach(func(_ int, x interface{}) {
		sum += (utils.ToFloat64(x) - vA) * (utils.ToFloat64(x) - vA)
	}, v...)
	return sum
}
