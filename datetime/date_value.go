package datetime

import (
	"strings"
	"time"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

//DateValue converts a date and then converts itto a serial number.
//
//Arguments
//
//date in either string form or a time.Time form
//
//Examples
//	DateValue("1/1/2021") // returns a value of 44197
// 	DateValue("1/1/2021") // returns a value of 44197
func DateValue(date interface{}) int {
	var val time.Time = time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	switch date := date.(type) {
	case string:
		valS := date
		sep := ""
		//Finding out what seperator is used in the string and using that to seperate the the string
		//For example 01-01-2001, so the sepereator is "-"
		if strings.Contains(valS, "/") {
			sep = "/"
		} else if strings.Contains(valS, "-") {
			sep = "-"
		} else {
			panic(core.ErrInvalidInput)
		}

		//Storing the sepertated values and the number of seperated values
		inpSep := strings.Split(valS, sep)
		inpSepLen := len(inpSep)

		//If the value is !2 or !3 then panicing
		if inpSepLen != 2 && inpSepLen != 3 {
			panic(core.ErrInvalidInput)
		}
		//Checking if all the characters are numerical digits or not
		isDigit := strings.IndexFunc(inpSep[1], utils.NotDigit) == -1
		var month time.Month
		if isDigit {
			month = getMonth(utils.ToInt(inpSep[1]))
		} else {
			monthS := inpSep[1]
			month = getMonth(monthS[:3])
		}
		//Checking the length and then creating a Date accordingly
		if inpSepLen == 2 {
			val = time.Date(time.Now().Year(), month, utils.ToInt(inpSep[0]), 0, 0, 0, 0, time.UTC)
		} else {
			val = time.Date(utils.ToInt(inpSep[0]), month, utils.ToInt(inpSep[2]), 0, 0, 0, 0, time.UTC)
		}
	case time.Time:
		val = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
	}
	return int(val.Sub(DayZero).Hours()/24) + 2
}

//getMonth inputs a month and then convert it into time.Month form.
//
//Arguments
//
//month in either string or a int form
//
//Examples
//	getMonth("January") // returns a value of 1
//	getMonth("Feb") // returns a value of 2
// 	getMonth(12) // returns a value of 12
func getMonth(month interface{}) time.Month {
	switch v := month.(type) {
	case string:
		for i := 0; i < 12; i++ {
			if MonthsInYear[i] == v {
				return time.Month(i + 1)
			}
		}
		panic(core.ErrInvalidInput)
	case int:
		return time.Month(month.(int))
	default:
		panic(core.ErrInvalidInput)
	}
}
