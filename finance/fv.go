package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func FV(Rate, Nper, Pmt interface{}, Value ...interface{}) float64 {
	var value float64
	var ty float64
	var ans float64
	rate := utils.ToFloat64(Rate)
	periods := utils.ToFloat64(Nper)
	payment := utils.ToFloat64(Pmt)
	if len(Value) > 0 {
		value = utils.ToFloat64(Value[0])
		if len(Value) > 1 {
			ty = utils.ToFloat64(Value[1])
			if ty < 0 || ty > 1 {
				panic(core.ErrInvalidInput)
			}
		}
	}
	if rate == 0 {
		ans = value + payment*periods
	} else {
		term := math.Pow(1+rate, periods)
		if ty == 1 {
			ans = value*term + payment*(1+rate)*(term-1)/rate
		} else {
			ans = value*term + payment*(term-1)/rate
		}
	}
	return -ans
}
