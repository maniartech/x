package datetime

import (
	"math"
	"time"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//YearFrac finds the fraction of year between two dates
// Bias 0 is 30/360,
// Bias 1 is actual/actual,
// Bias 2 is actual/360,
// Bias 3 is actual/365.
func YearFrac(date1 time.Time, date2 time.Time, Bias ...interface{}) float64 {
	bias := 0
	var diff float64
	if len(Bias) > 0 {
		bias = utils.ToInt(Bias[0])
	}
	if bias == 0 {
		diff = (math.Abs(float64(Days360(date1, date2)))) / 360
	} else if bias == 1 {
		if IsLeapYear(date1.Year()) && IsLeapYear(date2.Year()) {
			diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 366
		} else {
			diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 365
		}
	} else if bias == 2 {
		diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 360
	} else if bias == 3 {
		diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 365

	} else {
		panic(core.ErrInvalidInput)
	}
	return diff
}