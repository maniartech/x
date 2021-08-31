package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Nper(rate, pmt, pv interface{}, futureValue ...interface{}) float64 {
	var fv float64
	var ans float64
	var ty float64
	Pv := utils.ToFloat64(pv)
	payment := utils.ToFloat64(pmt)
	Rate := utils.ToFloat64(rate)
	//Getting the futureValue if it was inputted
	if len(futureValue) > 0 {
		fv = utils.ToFloat64(futureValue[0])
		if len(futureValue) > 1 {
			ty = utils.ToFloat64(utils.ToInt(futureValue[1]))
			if ty < 0 || ty > 1 {
				panic(core.ErrInvalidInput)
			}
		}
	}
	//Calculating ans depending differently if the rate == 0
	if Rate == 0 {
		ans = (-(payment + fv) / Pv)
	} else {
		//Calculating numerator and denominator
		var num float64 = payment*(1+Rate*ty) - fv*Rate
		var den float64 = (Pv*Rate + payment*(1+Rate*ty))
		ans = math.Log(num/den) / math.Log(1+Rate)
	}
	return ans
}
