package finance

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/utils"
)

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
