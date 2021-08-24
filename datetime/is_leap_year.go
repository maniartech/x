package datetime

import "time"

//IsLeapYear finds out of the year inputed is a leap year or not in boolean form.
//
//Arguments
//
//year: Year that would be checked if it is a leap year or not
//
//Examples
//	IsLeapYear (2020) // returns true
// 	IsLeapYear (2021) // returns false
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
