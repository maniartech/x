package datetime

import "time"

//isLeapYear finds out of the inputed year is a leap year or not
func IsLeapYear(year interface{}) bool {
	yearVal := 0
	switch year := year.(type) {
	case int:
		yearVal = year
	case time.Time:
		yearVal = year.Year()
	}
	if yearVal%4 == 0 {
		if yearVal%100 == 0 {
			if yearVal%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}
