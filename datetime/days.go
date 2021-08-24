package datetime

import (
	"time"
)

//Days finds the difference between the two dates. date1 - date2
//
//Arguments
//
//date1: First Date
//date2: Second Date
//
//Examples
//	Days (Date(2021, 12, 31), Date(2021, 1, 1)) // returns 364
// 	Days (Date(2021, 1, 1), Date(2022, 1, 1)) // returns 365
func Days(date1 time.Time, date2 time.Time) int {
	d1 := DateValue(date1)
	d2 := DateValue(date2)
	return d1 - d2
}

//Days360 finds the difference between two dates according to a 360 day year calendar.
//
//Arguments
//
//date1: First Date
//date2: Second Date
//
//Examples
//	Days (Date(2021, 12, 31), Date(2021, 1, 1)) // returns 359
// 	Days (Date(2021, 1, 1), Date(2022, 1, 1)) // returns 360
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
