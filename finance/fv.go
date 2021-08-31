package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func FV(rate, nper, pmt interface{}, value ...interface{}) float64 {
	//Initialising variables
	var Value float64
	var ty float64
	var ans float64
	Rate := utils.ToFloat64(rate)
	periods := utils.ToFloat64(nper)
	payment := utils.ToFloat64(pmt)
	if len(value) > 0 {
		Value = utils.ToFloat64(value[0])
		if len(value) > 1 {
			ty = utils.ToFloat64(value[1])
			if ty < 0 || ty > 1 {
				panic(core.ErrInvalidInput)
			}
		}
	}
	//Calcuting ans differently based on rate
	if Rate == 0 {
		ans = Value + payment*periods
	} else {
		term := math.Pow(1+Rate, periods)
		if ty == 1 {
			ans = Value*term + payment*(1+Rate)*(term-1)/Rate
		} else {
			ans = Value*term + payment*(term-1)/Rate
		}
	}
	return -ans
}

//WIP
func FVSchedule(Principal interface{}, Schedule ...interface{}) float64 {
	principal := utils.ToFloat64(Principal)
	var future float64 = principal
	utils.ForEach(func(_ int, x interface{}) {
		future *= 1 + utils.ToFloat64(x)
	}, Schedule...)
	return future
}
