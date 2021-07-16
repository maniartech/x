package finance

import (
	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//SLN is used to find the straight line depreciation for a asset for one period
func SLN(Cost, Salvage, Life interface{}) float64 {
	cost := utils.ToFloat64(Cost)
	salvage := utils.ToFloat64(Salvage)
	life := utils.ToFloat64(Life)

	//Panicing if the life of the product is 0
	if life == 0 {
		panic(core.ErrInvalidInput)
	}
	return calc.Divide(cost-salvage, life)
}
