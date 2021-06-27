package finance

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func PV(Rate, Nper, Pmt interface{}, futureValue ...interface{}) float64 {
	var fv float64
	var ans float64
	var ty float64
	np := utils.ToFloat64(Nper)
	pmt := utils.ToFloat64(Pmt)
	rate := utils.ToFloat64(Rate)
	if len(futureValue) > 0 {
		fv = utils.ToFloat64(futureValue[0])
		if len(futureValue) > 1 {
			ty = utils.ToFloat64(utils.ToInt(futureValue[1]))
			if ty < 0 || ty > 1 {
				panic(core.ErrInvalidInput)
			}
		}
	}
	if rate == 0 {
		ans = -pmt*np - fv

	} else {
		ans = calc.Divide((((1-math.Pow(1+rate, np))/rate)*pmt*(1+rate*ty) - fv), math.Pow(1+rate, np))
	}
	return ans
}
