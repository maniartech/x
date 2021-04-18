package calc

import (
	"github.com/maniartech/x/utils"
)

func Floor(input, factor interface{}) float64 {
	i := utils.ToDecimal(input)
	f := utils.ToDecimal(factor)

	// math.Floor(i / f) * f
	r, _ := i.Div(f).Floor().Mul(f).Float64()

	return r
}
