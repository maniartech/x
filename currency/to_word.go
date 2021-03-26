package currency

import (
	"strconv"
)

// Num2Word convert the number string into international numbering word format.
func Num2Word(input string) string {
	var word, part, centV, centS string
	var mod, powerCounter, number, cent int
	var multiplier bool = false
	power := [6]string{"", " thousand ", " million ", " billion ", " trillion ", " quadrillion "}
	//Converting the input to int and then spliting it into two
	number, cent = ConvertToInt(input)

	centS = strconv.Itoa(cent)
	if len(centS) == 1 {
		multiplier = true
	}

	//Checking if paise is != 0 then we compute and the cent part final output
	if cent != 0 {
		if cent > 100 {
			cent = RoundOffToTens(cent)
		}
		if cent == 100 {
			panic(ErrInvalidInput)
		}
		centV = ToWordTens(cent, multiplier)
	}
	//Some basic conditons
	// computing the 3 digits of the number togther
	if number == 0 {
		word = "zero"
	} else if number > 0 && number < 20 {
		word = ToWordTens(number, false)
	} else {
		temp := int(number)
		for temp != 0 {
			mod = temp % 1000
			part = ToWordHundreds(mod)
			if mod != 0 {
				word = part + power[powerCounter] + word
			} else {
				word = part + word
			}
			temp = temp / 1000
			powerCounter++
		}
	}
	if cent == 0 {
		return word
	}
	return word + " and " + centV + " cent"
}
