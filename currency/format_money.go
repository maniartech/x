package currency

import (
	"strings"

	"github.com/maniartech/go-funcs/is"
)

// FormatMoneySymbol formats the money with correct placement of ',' between digits according to international system
// this function also places a currency symbol before the number
func FormatMoney(input string, currencySymbol string) string {
	var value string
	var number string
	var strArr []string = strings.Split(input, ".")
	strArrLen := len(strArr)
	//checking if there more than one `.` in the inputed value if thats the case then the input is wrong.
	if strArrLen != 1 && strArrLen != 2 {
		panic(errInvalidInput)
	}
	//Checking all the characters in the inputed string are numerical digits
	numberValidator := strings.IndexFunc(strArr[0], is.NotDigit) == -1
	if !numberValidator {
		panic(errInvalidInput)
	}
	number = strArr[0]
	//Checking if the lenght of the number is greater than three so if thats the case then we do the complex
	//Computation
	if len(number) > 3 {
		//Setting the value of value to the first three digits so that we can directly put a `,` when we enter the loop.
		value = number[len(number)-3:]
		for i := len(strArr[0]) - 3; i > 0; i = i - 3 {
			//Checking if the value of i - 3 is not less than 0 so that we don't get a out of range error
			if (i - 3) <= 0 {
				value = number[:i] + "," + value
			} else {
				value = number[i-3:i] + "," + value
			}
		}
	} else {
		value = number
	}
	//If there is a `.` in the input then the lenght of the slice containing the values is 2,
	//Then we compute the 2nd value in the slice
	if strArrLen == 2 {
		pointValidator := strings.IndexFunc(strArr[1], is.NotDigit) == -1
		if !pointValidator {
			panic(errInvalidInput)
		}
		value = value + "." + strArr[1]
	}
	return currencySymbol + value
}
