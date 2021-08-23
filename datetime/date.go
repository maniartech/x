package datetime

import (
	"time"

	"github.com/maniartech/x/utils"
)

//Date returns a date in go's time.Time form.
//
//Arguments
//
//year: Year of the date
//month: Month of the date
//day: Day of the date
//
//Examples
//	Date (2021, 1, 1)// returns a date of 1-1-2021 (1-Jan-2021)
// 	Date (2000, 12, 2)// returns a date of 2-12-2000 (2-Dec-2000)
func Date(year, month, day interface{}) time.Time {
	Year := utils.ToInt(year)
	Month := utils.ToInt(month)
	Day := utils.ToInt(day)
	return time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.UTC)
}
