package finance

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//WIP
func Mirr(financeRate, reinvestRate interface{}, value ...interface{}) float64 {
	Frate := utils.ToFloat64(financeRate)
	Rrate := utils.ToFloat64(reinvestRate)
	Pve := make([]float64, 0)
	Nve := make([]float64, 0)
	len := float64(len(value))
	//Initalizing this to check if there is atleast one positive and one negetive value in values.
	//If not then panicing a DivideBy0 Err
	var pve bool = false
	var nve bool = false
	//Going through all the values in value and sorting them into +ve and -ve slices
	//If there is neither a +ve nor -ve then throwing panicing.
	utils.ForEach(func(_ int, x interface{}) {
		xV := utils.ToFloat64(x)
		if xV >= 0 {
			Pve = append(Pve, xV)
			pve = true
		} else {
			Nve = append(Nve, xV)
			nve = true
		}
	}, value...)
	//Checking if there are both -ve or +ve values then panicing
	if !pve || !nve {
		panic(core.ErrDivideBy0)
	}

	//Calculating Numerator and Denominator
	n := -Npv(Rrate, Pve) * math.Pow(1+Rrate, len-1)
	d := Npv(Frate, Nve) * (1 + Frate)
	return math.Pow(n/d, 1.0/(len-1.0)) - 1
}
