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
	//If n < 0 or x > n then panicing
	if n < 0 || x > n {
		panic(core.ErrInvalidInput)
	}
	//If p is neither 1 nor 0 then throwing error
	if p < 0 || p > 1 {
		panic(core.ErrInvalidInput)
	}

	var sum float64
	//If cumulative is true then using different formula for calculation
	if cumulative {
		//Summation of all the values from y = 0 to x
		for y := 0; y <= x; y++ {
			sum += utils.ToFloat64(calc.Combin(n, y)) * math.Pow(p, utils.ToFloat64(y)) * math.Pow(1-p, utils.ToFloat64(n-y))
		}
	} else {
		//Forumula if !cumulative
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

func GammaDist(x, Alpha, Beta interface{}, Cumulative bool) float64 {
	a := utils.ToFloat64(Alpha)
	b := utils.ToFloat64(Beta)
	X := utils.ToFloat64(x)
	if X < 0 || a <= 0 || b <= 0 {
		panic(core.ErrInvalidInput)
	}

	if a == 1 {
		return ExponDist(X, (1 / b), false)
	}
	if Cumulative {
		var sum float64
		for y := 1; y <= utils.ToInt(X); y++ {
			sum += calc.Divide(1, y)
		}
		return sum
	}
	return calc.Divide(1, math.Pow(b, a)*Gamma(a)) * math.Pow(X, a-1) * math.Pow(math.E, -calc.Divide(X, b))
}

// func NormDist(X, mean, StDev interface{}, Cumulative bool) float64 {
// 	x := utils.ToFloat64(X)
// 	m := utils.ToFloat64(mean)
// 	d := utils.ToFloat64(StDev)

// 	if d <= 0 {
// 		panic(core.ErrInvalidInput)
// 	}
// 	if mean == 0 && StDev == 1 && Cumulative {
// 		NormSDist(x, true)
// 	}
// 	if !Cumulative {
// 		return calc.Divide(1, math.Sqrt(2*math.Pi)*d) * math.Pow(math.E, calc.Divide((x-m)*(x-m), 2*d*d))
// 	}
// 	return 0.0
// }
func NormSDist(X interface{}, cumulative bool) float64 {
	z := utils.ToFloat64(X)
	//If cumulative then returning 0
	if cumulative {
		return 0.0
	}
	//Forumla for NormSDist if !cumulative
	return calc.Divide(1, math.Sqrt(2*math.Pi)) * math.Pow(math.E, -(calc.Divide(z*z, 2)))
}
