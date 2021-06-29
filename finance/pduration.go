package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func PDuaration(Rate, PresentValue, FutureValue interface{}) float64 {
	rate := utils.ToFloat64(Rate)
	present := utils.ToFloat64(PresentValue)
	future := utils.ToFloat64(FutureValue)

	if rate <= 0 {
		panic(core.ErrInvalidInput)
	}
	return (math.Log(future) - math.Log(present)) / math.Log(1+rate)
}
