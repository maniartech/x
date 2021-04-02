package datetime

import (
	"math"
	"time"
)

func DateValue(input time.Time) int {
	return int(input.Sub(DayZero).Hours()/24) + 2
}

func Days(date1 time.Time, date2 time.Time) int {
	return int(math.Abs((date1.Sub(date2).Hours() / 24)))
}
func EDate(date time.Time, months int) time.Time {
	return date.AddDate(0, months, 0)
}
func EOMonth(date time.Time, months int) time.Time {
	date = date.AddDate(0, months, 0)
	if date.Month() == time.February {
		if isLeapYear(date.Year()) {
			date = date.AddDate(0, 0, 29-date.Day())
		} else {
			date = date.AddDate(0, 0, 28-date.Day())
		}
	} else {
		date = date.AddDate(0, 0, DaysInMonth[int(date.Month())%12]-date.Day()+1)
	}
	return date
}

func Today() time.Time { return time.Now() }

func Day(date time.Time) int    { return date.Day() }
func Month(date time.Time) int  { return int(date.Month()) }
func Year(date time.Time) int   { return date.Year() }
func Hour(date time.Time) int   { return date.Hour() }
func Minute(date time.Time) int { return date.Minute() }
func Second(date time.Time) int { return date.Second() }
