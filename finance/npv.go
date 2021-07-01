package finance

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/utils"
)

//NPV
func Npv(v ...interface{}) float64 {
	var sum float64
	var i float64 = 1
	r := utils.ToFloat64(v[0])
	utils.ForEach(func(_ int, val interface{}) {
		sum += calc.Divide(utils.ToFloat64(val), math.Pow((1+r), i))
		i += 1
	}, v[1:]...)
	return sum
}
