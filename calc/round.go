package calc

import (
	"github.com/maniartech/x/utils"
)

func Round(val interface{}, places interface{}) float64 {

	v := utils.ToDecimal(val)
	p := utils.ToInt(places)

	rval, _ := v.Round(int32(p)).Float64()
	return rval

	// var value float64
	// if factor == 0 {
	// 	return math.Round(input)
	// }
	// value = (math.Round(input * math.Pow10(factor))) * math.Pow10(-factor)
	// return value
}
