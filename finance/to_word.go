package currency

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
)

// Num2Word convert the number string into international numbering word format.
func Num2Word(input interface{}) string {
	var word, part, centV string
	var mod, powerCounter, number, cent int
	var multiplier bool = false
	power := [6]string{"", " thousand ", " million ", " billion ", " trillion ", " quadrillion "}
	//Converting the input to int and then spliting it into two
	number, cent = ConvertToInt(input)
	if calc.NumberOfDigits(cent) == 1 {
		multiplier = true
	}

	//Checking if cent is != 0 then we compute and the cent part final output
	if cent != 0 {
		if cent > 100 {
			cent = int((calc.Round(float64(cent), -(calc.NumberOfDigits(cent) - 2))) / math.Pow10(calc.NumberOfDigits(cent)-2))
		}
		if cent == 100 {
			panic(core.ErrInvalidInput)
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
