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
	//Formula for calculating
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
	//Retriving value of future value if it was inputed
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
	//Calculations changing depending on the periods of the payment
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
	//Retriving value of future value if it was inputed
	if len(futureValue) > 0 {
		fv = utils.ToFloat64(futureValue[0])
		if len(futureValue) > 1 {
			ty = utils.ToFloat64(utils.ToInt(futureValue[1]))
			if ty < 0 || ty > 1 {
				panic(core.ErrInvalidInput)
			}
		}
	}
	//Calculating diffrently depending if the rate == 0
	if rate == 0 {
		ans = (payment + fv) / period
	} else {
		term := math.Pow(1+rate, period)
		//Calculating diffrently depending if the ty == 1
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
	//Retriving value of future value and ty if they were inputed
	if len(futureValue) > 0 {
		fv = utils.ToFloat64(futureValue[0])
		if len(futureValue) > 1 {
			ty = utils.ToFloat64(utils.ToInt(futureValue[1]))
			if ty < 0 || ty > 1 {
				panic(core.ErrInvalidInput)
			}
		}

	}
	//Formula for PPMT
	return PMT(rate, periods, present, fv, ty) - IPMT(rate, period, periods, present, fv, ty)
}

func CumIPMT(Rate, Nper, PresentValue, Start, End, TYPE interface{}) float64 {
	rate := utils.ToFloat64(Rate)
	periods := utils.ToFloat64(Nper)
	value := utils.ToFloat64(PresentValue)
	start := utils.ToFloat64(Start)
	end := utils.ToFloat64(End)
	ty := utils.ToFloat64(TYPE)
	//Panicing if rate <= 0 or periods <= or value <= 0
	if rate <= 0 || periods <= 0 || value <= 0 {
		panic(core.ErrInvalidInput)
	}
	//Panicing if start or end are less than 0 or start is greater than end
	if start < 0 || end < 0 || start > end {
		panic(core.ErrInvalidInput)
	}
	//If ty is neither 1 nor 0 then panicing
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
