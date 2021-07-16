package finance

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func RRI(Nper, Pv, Fv interface{}) float64 {
	nper := utils.ToFloat64(Nper)
	pv := utils.ToFloat64(Pv)
	fv := utils.ToFloat64(Fv)

	return math.Pow(fv/pv, calc.Divide(1, nper)) - 1
}

func PDuaration(Rate, PresentValue, FutureValue interface{}) float64 {
	rate := utils.ToFloat64(Rate)
	present := utils.ToFloat64(PresentValue)
	future := utils.ToFloat64(FutureValue)

	if rate <= 0 {
		panic(core.ErrInvalidInput)
	}
	return (math.Log(future) - math.Log(present)) / math.Log(1+rate)
}

//NPV
func Npv(Rate interface{}, v ...interface{}) float64 {
	var sum float64
	var i float64 = 1
	rate := utils.ToFloat64(Rate)
	utils.ForEach(func(_ int, val interface{}) {
		sum += calc.Divide(utils.ToFloat64(val), math.Pow((1+rate), i))
		i += 1
	}, v[0:]...)
	return sum
}

func Effect(NominalRate, Npery interface{}) float64 {
	NR := utils.ToFloat64(NominalRate)
	Np := utils.ToFloat64(utils.ToInt(Npery))

	if NR <= 0 && Np < 1 {
		panic(core.ErrInvalidInput)
	}
	return math.Pow((1+calc.Divide(NR, Np)), Np) - 1
}
