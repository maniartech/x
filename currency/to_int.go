package currency

import (
	"strconv"
	"strings"

	"github.com/maniartech/x/utils"
)

//ConvertToInt, converts inputs into integer type
func ConvertToInt(input string) (int, int) {
	var numberV, decimalV int
	var err error
	strArr := strings.Split(input, ".")
	strArrLen := len(strArr)
	if strArrLen != 1 && strArrLen != 2 {
		panic(ErrInvalidInput)
	}
	//Checking all the characters in the inputed string are numerical digits before coverting the number to int
	numberValidator := strings.IndexFunc(strArr[0], utils.NotDigit) == -1
	if !numberValidator {
		panic(ErrInvalidInput)
	}
	numberV, err = strconv.Atoi(strArr[0])
	if err != nil {
		panic(ErrInvalidInput)
	}
	if strArrLen == 2 {
		decimalValidator := strings.IndexFunc(strArr[1], utils.NotDigit) == -1
		if !decimalValidator {
			panic(ErrInvalidInput)
		}
		decimalV, err = strconv.Atoi(strArr[1])
		if err != nil {
			panic(ErrInvalidInput)
		}
	}
	return numberV, decimalV
}
func Split(input string) []string {
	strArr := strings.Split(input, ".")
	strArrLen := len(strArr)
	if strArrLen != 1 && strArrLen != 2 {
		panic(ErrInvalidInput)
	}
	numberValidator := strings.IndexFunc(strArr[0], utils.NotDigit) == -1
	if !numberValidator {
		panic(ErrInvalidInput)
	}

	if strArrLen == 2 {
		decimalValidator := strings.IndexFunc(strArr[1], utils.NotDigit) == -1
		if !decimalValidator {
			panic(ErrInvalidInput)
		}
	}

	return strArr
}
