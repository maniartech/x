package datetime

import (
	"math"
	"time"
)

//Days Finds the difference between two dates
func Days(date1 time.Time, date2 time.Time) int {
	return int(math.Abs((date1.Sub(date2).Hours() / 24)))
}

//Days360 Finds the difference between two dates on a 360 day year calendar
func Days360(date1, date2 time.Time) int {
	//Checking if the second date is smaller than the first one
	//So it could be swaped to reduce chances of error
	neg := DateValue(date2) < DateValue(date1)
	if neg {
		temp := date1
		date1 = date2
		date2 = temp
	}
	val := 0
	//Finding the difference in year, month and days
	diffYear := date2.Year() - date1.Year()
	diffMonth := int(date2.Month()) - int(date1.Month())
	days := (date2.Day() + 30 - date1.Day()) % 30
	val = diffYear*360 + (diffMonth)*30 + days
	if date2.Day() < date1.Day() {
		val -= 30
	}
	if (date1.Day() == 1 && date2.Day() == 31) || (date1.Day() == 31 && date2.Day() == 1) {
		val += 30
	}
	if neg {
		return -val
	}
	return val
}
