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

//Some issue I cant figure out
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

func PPMT(Rate, Period, Periods, PresentValue interface{}, futureValue ...interface{}) float64 {
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
	return PMT(rate, periods, present, fv, ty) - IPMT(rate, period, periods, present, fv, ty)
}

func CumIPMT(Rate, Nper, PresentValue, Start, End, TYPE interface{}) float64 {
	rate := utils.ToFloat64(Rate)
	periods := utils.ToFloat64(Nper)
	value := utils.ToFloat64(PresentValue)
	start := utils.ToFloat64(Start)
	end := utils.ToFloat64(End)
	ty := utils.ToFloat64(TYPE)

	if rate <= 0 || periods <= 0 || value <= 0 {
		panic(core.ErrInvalidInput)
	}
	if start < 0 || end < 0 || start > end {
		panic(core.ErrInvalidInput)
	}

	if ty != 0 && ty != 1 {
		panic(core.ErrInvalidInput)
	}
	var payment float64 = PMT(rate, periods, value, 0, ty)
	var interest float64 = 0

	if start == 1 {
		if ty == 0 {
			interest = -value
		}
		start += 1
	}
	for i := start; i <= end; i += 1 {
		if ty == 1 {
			interest += FV(rate, i-2, payment, value, 1) - payment
		} else {
			interest += FV(rate, i-1, payment, value, 0)
		}
	}
	interest *= rate

	return interest
}
