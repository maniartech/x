package currency

import (
	"math"
	"strconv"
)

// Num2Word convert the number string into international numbering word format.
func Num2Word(input string) (string, error) {
	var word, part, centV, centS string
	var mod1, mod2, powerCounter, number, cent int
	var err error
	var multiplier bool = false
	power := [6]string{"", " thousand ", " million ", " billion ", " trillion ", " quadrillion "}
	//Converting the input to int and then spliting it into two
	number, cent = ConvertToInt(input)

	if err == nil {
		centS = strconv.Itoa(cent)
		if len(centS) == 1 {
			multiplier = true
		}
		if cent > 100 {
			centS = centS[:2] + "." + centS[2:]
			centF, err3 := strconv.ParseFloat(centS, 64)
			if err3 == nil {
				cent = int(math.Round(centF))
			}
		}
		//Checking if paise is != 0 then we compute and the cent part final output
		if cent != 0 {
			if cent > 0 && cent < 20 {
				centV = singles[cent]
				if multiplier {
					centV = tys[cent-1]
				}
			} else if cent > 19 && cent < 100 {
				if cent%10 == 0 {
					centV = tys[(cent/10)-1]
				} else {
					centV = tys[(cent/10)-1] + " " + singles[cent%10]
				}
			} else if cent == 100 {
				number++
				cent = 0
			}
		}
	}
	//Some basic conditons
	// computing the 3 digits of the number togther
	if number == 0 {
		word = "zero"
	} else if number > 0 && number < 20 {
		word = singles[number]
	} else {
		temp := int(number)
		for temp != 0 {
			mod1 = temp % 1000
			mod2 = mod1 % 100
			if mod1 == 0 {
				part = ""
			} else {
				if mod2 == 0 {
					//Do nothing
				} else if mod2 > 0 && mod2 < 20 {
					part = singles[mod2]
				} else {
					if mod1%10 == 0 {
						part = tys[((mod1%100)/10)-1]
					} else {
						part = tys[((mod1%100)/10)-1] + " " + singles[mod1%10]
					}
				}
				if mod1/100 != 0 && mod2 != 0 {
					part = singles[mod1/100] + " hundred " + part
				} else if mod1/100 != 0 && mod2 == 0 {
					part = singles[mod1/100] + " hundred" + part
				}
			}
			if mod1 != 0 {
				word = part + power[powerCounter] + word
			} else {
				word = part + word
			}
			temp = temp / 1000
			powerCounter++
		}
	}
	if cent == 0 {
		return word, nil
	}
	return word + " and " + centV + " cent", nil
}
