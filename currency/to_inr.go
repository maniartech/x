package currency

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
)

// Num2WordInd convert the number string into indian numbering word format.
func Num2WordInd(input string) string {
	var word, part, paiseV string
	var mod1, powerCounter, number, paise int
	var multiplier bool = false
	power := [6]string{"", "thousand", "lakh", "crore", "arab", "kharab"}

	//Converting the input to int and then spliting it into two
	number, paise = ConvertToInt(input)
	if calc.NumberOfDigits(paise) == 1 {
		multiplier = true
	}
	//Checking if paise is != 0 then we compute and the paise part final output
	if paise != 0 {
		if paise > 100 {
			paise = int((calc.Round(float64(paise), -(calc.NumberOfDigits(paise) - 2))) / math.Pow10(calc.NumberOfDigits(paise)-2))
		}
		if paise == 100 {
			panic(core.ErrInvalidInput)
		}
		paiseV = ToWordTens(paise, multiplier)
	}
	//Some basic conditons
	if number == 0 {
		word = "zero"
	} else if number > 0 && number < 20 {
		word = singles[number]
	} else {
		temp := int(number)
		for temp != 0 {
			part = ""
			// computing the first 3 digits of the number togther
			if powerCounter == 0 {
				mod1 = temp % 1000
				word = ToWordHundreds(mod1)
				powerCounter++
				temp = temp / 1000
			}
			// Coumputing the remaining parts of numbers in 2
			mod1 = temp % 100
			if mod1 == 0 {
				powerCounter++
				temp = temp / 100
				continue
			}
			part = ToWordTens(mod1, false)
			word = part + " " + power[powerCounter] + " " + word
			powerCounter++
			temp = temp / 100
		}
	}
	// if paise == 0 then we dont add the paise part to the final output
	if paise == 0 {
		return word
	}
	// if paise != 0 then we add the paise part to the final output
	return word + " and " + paiseV + " paise"
}
