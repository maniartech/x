package currency

import (
	"math"
	"strconv"
)

func RoundOffToTens(input int) int {
	var inputS string = strconv.Itoa(input)
	inputS = inputS[:2] + "." + inputS[2:]
	inputF, err := strconv.ParseFloat(inputS, 64)
	if err != nil {
		panic(ErrInvalidInput)
	}
	return int(math.Round(inputF))
}

// ToWordTens converts numbers form 0 to 99 to word form
func ToWordTens(input int, multiplier bool) string {
	var output string
	if input > 0 && input < 20 {
		output = singles[input]
		if multiplier {
			output = tys[input-1]
		}
	} else if input >= 20 && input < 100 {
		if input%10 == 0 {
			output = tys[(input/10)-1]
		} else {
			output = tys[(input/10)-1] + " " + singles[input%10]
		}
	}
	return output
}

// ToWordTens converts numbers form 0 to 999 to word form
func ToWordHundreds(input int) string {
	var output string
	if input == 0 {
		output = "zero"
	}
	if input/100 != 0 {
		if input%100 != 0 {
			output = ToWordTens(input/100, false) + " hundred " + ToWordTens(input%100, false)
		} else {
			output = ToWordTens(input/100, false) + " hundred"
		}
	} else {

		output = ToWordTens(input%100, false)
	}
	return output
}
