package calc

import (
	"github.com/maniartech/x/utils"
)

func Round(val interface{}, places interface{}) float64 {

	v := utils.ToDecimal(val)
	p := utils.ToInt(places)

	rval, _ := v.Round(int32(p)).Float64()
	return rval
}
