package currency

import (

	// "fmt"

	"math"
	"strconv"
	"strings"
)

// Num2Word convert the number string into international numbering word format.
func Num2Word(input string) (string, error) {
	var word string
	var part string
	var centV string
	var mod1 int
	var mod2 int
	var powerCounter int
	var number int
	var cent int
	var err error
	var centS string = ""
	var multiplier bool = false
	power := [6]string{"", " thousand ", " million ", " billion ", " trillion ", " quadrillion "}
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
		cent, err = strconv.Atoi(stringArr[1])
		if err != nil {
			return "", ErrInvalidInput
		}
		if len(stringArr[1]) == 1 {
			multiplier = true
		}
	} else {
		return "", ErrInvalidInput
	}
	if err == nil {
		centS = strconv.Itoa(cent)
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
