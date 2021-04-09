package datetime

import (
	"math"
	"time"

	"github.com/maniartech/x/calc"
)

//Finds the difference between two dates
func Days(date1 time.Time, date2 time.Time) int {
	return int(math.Abs((date1.Sub(date2).Hours() / 24)))
}

//EDate Adds or subtract months from the inputed date
func EDate(date time.Time, months int) time.Time {
	return date.AddDate(0, months, 0)
}

//EOMonth Adds or subtract months from the inputed date and then sets the day to the last day of the month
func EOMonth(date time.Time, months int) time.Time {
	var DaysInMonthC []int = []int{31, 31, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	date = date.AddDate(0, months, 0)
	if date.Month() == time.February {
		if isLeapYear(date.Year()) {
			return time.Date(date.Year(), date.Month(), 29, 0, 0, 0, 0, time.UTC)
		} else {
			return time.Date(date.Year(), date.Month(), 28, 0, 0, 0, 0, time.UTC)
		}
	} else {
		return time.Date(date.Year(), date.Month(), DaysInMonthC[(int(date.Month())%12)], 0, 0, 0, 0, time.UTC)
	}
}

//YearFrac finds the fraction of year between two dates
func YearFrac(date1 time.Time, date2 time.Time) float64 {
	diff := (math.Abs(float64(DateValue(date1) - DateValue(date2)))) / 365
	return diff
}

//Today returns todays date
func Today() time.Time { return time.Now() }

//Returns the decimal number for a particular time (WIP)
func Time(hour, min, sec int) float64 {
	cT := time.Date(1900, time.January, 1, hour, min, sec, 0, time.UTC)
	eT := time.Date(1900, time.January, 1, 23, 59, 59, 0, time.UTC)
	return calc.Round(eT.Sub(cT).Hours(), 7)
}
func Day(date time.Time) int    { return date.Day() }
func Month(date time.Time) int  { return int(date.Month()) }
func Year(date time.Time) int   { return date.Year() }
func Hour(date time.Time) int   { return date.Hour() }
func Minute(date time.Time) int { return date.Minute() }
func Second(date time.Time) int { return date.Second() }
