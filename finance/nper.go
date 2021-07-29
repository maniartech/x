package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Nper(Rate, Pmt, Pv interface{}, futureValue ...interface{}) float64 {
	var fv float64
	var ans float64
	var ty float64
	pv := utils.ToFloat64(Pv)
	payment := utils.ToFloat64(Pmt)
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
		ans = (-(payment + fv) / pv)
	} else {
		var num float64 = payment*(1+rate*ty) - fv*rate
		var den float64 = (pv*rate + payment*(1+rate*ty))
		ans = math.Log(num/den) / math.Log(1+rate)
	}
	return ans
}
