package currency

import (

	// "fmt"

	"math"
	"strconv"
	"strings"
)

// Num2WordInd convert the number string into indian numbering word format.
func Num2WordInd(input string) (string, error) {
	word := ""
	part := ""
	paiseV := ""
	mod1 := 0
	mod2 := 0
	powerCounter := 0
	var number int = 0
	var paise int = 0
	var err error
	var multiplier bool = false
	power := [6]string{"", "thousand", "lakh", "crore", "arab", "kharab"}
	//Spliting the number and checking for `.`
	stringArr := strings.Split(input, ".")
	stringArrLen := len(stringArr)
	//Checking if there is a `.` in the number if there is more than one the input is wrong
	if stringArrLen == 1 {
		number, err = strconv.Atoi(input)
		if err != nil {
			return "", ErrInvalidInput
		}
	} else if stringArrLen == 2 {
		number, err = strconv.Atoi(stringArr[0])
		if err != nil {
			return "", ErrInvalidInput
		}
		paise, err = strconv.Atoi(stringArr[1])
		if err != nil {
			return "", ErrInvalidInput
		}
		if len(stringArr[1]) == 1 {
			multiplier = true
		}
	} else {
		return "", ErrInvalidInput
	}
	paiseS := strconv.Itoa(paise)
	//Checking if paise is != 0 then we compute and the paise part final output
	if paise != 0 {
		if paise > 100 {
			paiseS = paiseS[:2] + "." + paiseS[2:]
			paiseF, err3 := strconv.ParseFloat(paiseS, 64)
			if err3 == nil {
				paise = int(math.Round(paiseF))
			}
		}
		if paise > 0 && paise < 20 {
			paiseV = singles[paise]
			if multiplier {
				paiseV = tys[paise-1]
			}
		} else if paise < 100 {
			if (paise % 10) == 0 {
				paiseV = tys[(paise/10)-1]
			} else {
				paiseV = tys[(paise/10)-1] + " " + singles[paise%10]
			}
		} else if paise == 100 {
			number++
			paise = 0
		}
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
				mod1 = temp % 100
				mod2 = temp % 1000
				word = singles[mod2/100]
				if mod2/100 != 0 {
					word = word + " hundred and "
				}
				if mod1 > 0 && mod1 < 20 {
					word = word + singles[mod1]
				} else {
					if (mod1 % 10) == 0 {
						word = word + tys[(mod1/10)%10-1]
					} else {
						word = word + tys[(mod1/10)%10-1] + " " + singles[mod1%10]
					}
				}
				powerCounter++
				temp = temp / 1000
			}
			// Coumputing the remaining parts of numbers in 2
			mod1 = temp % 100
			if mod1 != 0 {
				if mod1 > 0 && mod1 < 20 {
					part = singles[mod1]
				} else {
					if mod1%10 == 0 {
						part = tys[(mod1/10)%10-1]
					} else {

						part = tys[(mod1/10)%10-1] + " " + singles[mod1%10]
					}
				}
				part = part + " " + power[powerCounter]
				word = part + " " + word
			}
			powerCounter++
			temp = temp / 100
		}
	}
	// if paise == 0 then we dont add the paise part to the final output
	if paise == 0 {
		return word, nil
	}
	// if paise != 0 then we add the paise part to the final output
	return word + " and " + paiseV + " paise", nil
}
