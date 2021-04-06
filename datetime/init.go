package datetime

import "time"

var DayZero time.Time = time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
var DaysInMonth []int = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var MonthsInYear []string = []string{
	"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}
