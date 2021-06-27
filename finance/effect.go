package finance

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Effect(NominalRate, Npery interface{}) float64 {
	NR := utils.ToFloat64(NominalRate)
	Np := utils.ToFloat64(utils.ToInt(Npery))

	if NR <= 0 && Np < 1 {
		panic(core.ErrInvalidInput)
	}
	return math.Pow((1+calc.Divide(NR, Np)), Np) - 1
}
