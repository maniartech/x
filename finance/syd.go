package finance

import (
	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func SYD(Cost, Salvage, Life, Per interface{}) float64 {
	cost := utils.ToFloat64(Cost)
	salvage := utils.ToFloat64(Salvage)
	life := utils.ToFloat64(Life)
	per := utils.ToFloat64(Per)

	//Panicing if the life of an asset is equal to 0
	if life == 0 {
		panic(core.ErrInvalidInput)
	}
	if per < 1 || per > life {
		panic(core.ErrInvalidInput)
	}
	return calc.Divide((cost-salvage)*(life-per+1)*2, ((life * life) + life))
}
