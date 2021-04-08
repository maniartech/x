package datetime

import "time"

var DayZero time.Time = Date(1900, 1, 1)
var DaysInMonth []int = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var MonthsInYear []string = []string{
	"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
}
