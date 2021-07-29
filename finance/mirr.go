package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//WIP
func Mirr(FinanceRate, ReinvestRate interface{}, Value ...interface{}) float64 {
	Frate := utils.ToFloat64(FinanceRate)
	Rrate := utils.ToFloat64(ReinvestRate)
	Pve := make([]float64, 0)
	Nve := make([]float64, 0)
	len := float64(len(Value))
	//Initalizing this to check if there is atleast one positive and one negetive value in Values
	//If not then panicing a DivideBy0 Err
	var pve bool = false
	var nve bool = false
	utils.ForEach(func(_ int, x interface{}) {
		xV := utils.ToFloat64(x)
		if xV >= 0 {
			Pve = append(Pve, xV)
			pve = true
		} else {
			Nve = append(Nve, xV)
			nve = true
		}
	}, Value...)
	if !pve || !nve {
		panic(core.ErrDivideBy0)

	}
	n := -Npv(Rrate, Pve) * math.Pow(1+Rrate, len-1)
	d := Npv(Frate, Nve) * (1 + Frate)
	return math.Pow(n/d, 1.0/(len-1.0)) - 1
}
