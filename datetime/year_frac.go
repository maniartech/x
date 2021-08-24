package datetime

import (
	"math"
	"time"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//YearFrac finds the fraction of year between two dates
//
//Arguments
//
//date1: First Date  in time.Time
//date1: Second Date  in time.Time
//basis: (Optional) Type of day count basis to be used in the calculation
//
//Remark
//
// Basis 0 is 30/360.
//
// Basis 1 is actual/actual.
//
// Basis 2 is actual/360.
//
// Basis 3 is actual/365.
//
//Examples
//	Day(Date(2021, January, 1)) //returns 1
// 	Day(Date(2016, March, 26)) //returns  26
func YearFrac(date1 time.Time, date2 time.Time, basis ...interface{}) float64 {
	Basis := 0
	var diff float64
	if len(basis) > 0 {
		Basis = utils.ToInt(basis[0])
	}
	//Computing Difference according to the Basis

	switch Basis {
	case 0:
		diff = (math.Abs(float64(Days360(date1, date2)))) / 360
	case 1:
		if IsLeapYear(date1.Year()) && IsLeapYear(date2.Year()) {
			diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 366
		} else {
			diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 365
		}
	case 2:
		diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 360
	case 3:
		diff = (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 365

	default:
		panic(core.ErrInvalidInput)
	}
	return diff
}
