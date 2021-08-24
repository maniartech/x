package datetime

import (
	"time"

	"github.com/maniartech/x/calc"
)

//EDate Adds or subtract months from the inputed date.
//
//Arguments
//
//date: Starting date in time.Time
//months: Months to be added or subtracted
//
//Examples
//
//	EDate(Date(2021, January, 1),  3) //returns the date 2021-April-1
// 	EDate(Date(2021, January, 1), -2) //returns the date 2021-November-1
func EDate(date time.Time, months int) time.Time {
	return date.AddDate(0, months, 0)
}

//EDate returns a last day of the of the month before or after the inputed date.
//
//Arguments
//
//date: Starting date in time.Time
//months: Months to be added or subtracted
//
//Examples
//	EOMonth(Date(2021, January, 1),  1) //returns the date 2021-Febuary-28
// 	EOMonth(Date(2021, January, 1), -2) //returns the date 2021-November-30
// 	EOMonth(Date(2021, January, 1), -11) //returns the date 2020-Febuary-29
func EOMonth(date time.Time, months int) time.Time {
	var DaysInMonthC []int = []int{31, 31, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	date = date.AddDate(0, months, 0)
	if date.Month() == time.February {
		if IsLeapYear(date.Year()) {
			return time.Date(date.Year(), date.Month(), 29, 0, 0, 0, 0, time.UTC)
		} else {
			return time.Date(date.Year(), date.Month(), 28, 0, 0, 0, 0, time.UTC)
		}
	} else {
		return time.Date(date.Year(), date.Month(), DaysInMonthC[(int(date.Month())%12)], 0, 0, 0, 0, time.UTC)
	}
}

//Today returns the current date on the computer
//
//Examples (today's date being 1-January-2021)
//	Today() //returns the date 1-January-2021
func Today() time.Time { return time.Now() }

//Returns the decimal number for a particular time (WIP)
func Time(hour, min, sec int) float64 {
	cT := time.Date(1900, time.January, 1, hour, min, sec, 0, time.UTC)
	eT := time.Date(1900, time.January, 1, 23, 59, 59, 0, time.UTC)
	return calc.Round(eT.Sub(cT).Hours(), 7)
}

//Day returns the day of the date inputed
//
//Arguments
//
//date: Date of which of the day you are trying to find  in time.Time
//
//Examples
//	Day(Date(2021, January, 1)) //returns 1
// 	Day(Date(2016, March, 26)) //returns  26
func Day(date time.Time) int { return date.Day() }

//Month returns the month of the date inputed
//
//Arguments
//
//date: Date of which of the month you are trying to find  in time.Time
//
//Examples
//	Month(Date(2021, January, 1)) //returns 1 (January)
// 	Month(Date(2016, March, 26)) //returns  3 (March)
func Month(date time.Time) int { return int(date.Month()) }

//Year returns the year of the date inputed
//
//Arguments
//
//date: Date of which of the year you are trying to find  in time.Time
//
//Examples
//	Year (Date(2021, January, 1)) //returns 2021
// 	Year (Date(2016, March, 26)) //returns  2016
func Year(date time.Time) int { return date.Year() }
