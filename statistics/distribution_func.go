package statistics

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func BinomDist(num, trials, probablity interface{}, cumulative bool) float64 {
	n := utils.ToInt(trials)
	x := utils.ToInt(num)
	p := utils.ToFloat64(probablity)

	if n < 0 || x > n {
		panic(core.ErrInvalidInput)
	}
	if p < 0 || p > 1 {
		panic(core.ErrInvalidInput)
	}

	var sum float64

	if cumulative {
		for y := 0; y <= x; y++ {
			sum += utils.ToFloat64(calc.Combin(n, y)) * math.Pow(p, utils.ToFloat64(y)) * math.Pow(1-p, utils.ToFloat64(n-y))
		}
	} else {
		sum = utils.ToFloat64(calc.Combin(n, x)) * math.Pow(p, utils.ToFloat64(x)) * math.Pow(1-p, utils.ToFloat64(n-x))
	}
	return sum
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

func GammaDist(X, Alpha, Beta interface{}, Cumulative bool) float64 {
	a := utils.ToFloat64(Alpha)
	b := utils.ToFloat64(Beta)
	x := utils.ToFloat64(X)
	if x < 0 || a <= 0 || b <= 0 {
		panic(core.ErrInvalidInput)
	}

	if a == 1 {
		return ExponDist(x, (1 / b), false)
	}
	if Cumulative {
		var sum float64
		for X := 1; X <= utils.ToInt(x); X++ {
			sum += calc.Divide(1, X)
		}
		return sum
	}
	return calc.Divide(1, math.Pow(b, a)*Gamma(a)) * math.Pow(x, a-1) * math.Pow(math.E, -calc.Divide(x, b))
}

func NormDist(X, mean, StDev interface{}, Cumulative bool) float64 {
	x := utils.ToFloat64(X)
	m := utils.ToFloat64(mean)
	d := utils.ToFloat64(StDev)

	if d <= 0 {
		panic(core.ErrInvalidInput)
	}
	if mean == 0 && StDev == 1 && Cumulative {
		NormSDist(x, true)
	}
	if !Cumulative {
		return calc.Divide(1, math.Sqrt(2*math.Pi)*d) * math.Pow(math.E, calc.Divide((x-m)*(x-m), 2*d*d))
	}
	return 0.0
}
func NormSDist(X interface{}, Cumulative bool) float64 {
	z := utils.ToFloat64(X)
	if Cumulative {
		return 0.0
	}
	return calc.Divide(1, math.Sqrt(2*math.Pi)) * math.Pow(math.E, -(calc.Divide(z*z, 2)))
}
