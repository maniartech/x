package datetime

import (
	"math"
	"time"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//YearFrac finds the fraction of year between two dates
// Bias 0 is actual/actual
// Bias 1 is actual/365
func YearFrac(date1 time.Time, date2 time.Time, Bias ...interface{}) float64 {
	bias := 0
	var diff float64
	if len(Bias) > 0 {
		bias = utils.ToInt(Bias[0])
	}
	if bias == 0 {
		if IsLeapYear(date1.Year()) && IsLeapYear(date2.Year()) {
			diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 366
		} else {
			diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 365
		}
	} else if bias == 1 {
		diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 365
	} else {
		panic(core.ErrInvalidInput)
	}
	return diff
}
