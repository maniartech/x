package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func IsPMT(Rate, Nper, Period, Value interface{}) float64 {
	rate := utils.ToFloat64(Rate)
	period := utils.ToFloat64(Period)
	per := utils.ToFloat64(Nper)
	value := utils.ToFloat64(Value)

	return value * rate * (per/period - 1)
}

func IPMT(Rate, Period, Periods, PresentValue interface{}, futureValue ...interface{}) float64 {
	var fv float64
	var ty float64
	period := utils.ToFloat64(Period)
	periods := utils.ToFloat64(Periods)
	rate := utils.ToFloat64(Rate)
	present := utils.ToFloat64(PresentValue)
	if len(futureValue) > 0 {
		fv = utils.ToFloat64(futureValue[0])
		if len(futureValue) > 1 {
			ty = utils.ToFloat64(utils.ToInt(futureValue[1]))
			if ty < 0 || ty > 1 {
				panic(core.ErrInvalidInput)
			}
		}
	}

	var payment float64 = PMT(rate, periods, present, fv, ty)
	var interest float64
	if period == 1 {
		if ty == 1 {
			interest = 0
		} else {
			interest = -present
		}
	} else {
		if ty == 1 {
			interest = FV(rate, period-2, payment, present, 1) - payment
		} else {
			interest = FV(rate, period-1, payment, present, 0)
		}
	}
	return interest * rate
}

func PMT(Rate, Nper, Pmt interface{}, futureValue ...interface{}) float64 {
	var fv float64
	var ans float64
	var ty float64
	period := utils.ToFloat64(Nper)
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
		ans = (payment + fv) / period
	} else {
		term := math.Pow(1+rate, period)
		if ty == 1 {
			ans = (fv*rate/(term-1) + payment*rate/(1-1/term)) / (1 + rate)
		} else {
			ans = fv*rate/(term-1) + payment*rate/(1-1/term)
		}
	}
	return ans
}
