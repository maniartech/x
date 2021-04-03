package statistics

import (
	"github.com/maniartech/x/utils"
	"github.com/shopspring/decimal"
)

// Average returns the average of the provided numbers.
func Average(v ...interface{}) float64 {
	var sum decimal.Decimal
	c := utils.ForEach(func(_ int, x interface{}) {
		sum = sum.Add(utils.ToDecimal(x))
	}, v...)
	ret, _ := sum.Div(decimal.NewFromInt(int64(c))).Float64()
	return ret
}
