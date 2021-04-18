package calc

import (
	"github.com/maniartech/x/utils"
)

//
func Ceiling(val interface{}, factor interface{}) float64 {
	v := utils.ToDecimal(val)
	f := utils.ToDecimal(factor)

	x := v.Div(f).Ceil().Mul(f)
	rval, _ := x.Float64()
	return rval
}
