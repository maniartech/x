package currency

import (

	// "fmt"

	"math"
	"strconv"
	"strings"
)

// Num2Word convert the number string into international numbering word format.
func Num2Char(input string) (string, error) {
	var word string
	var number int
	var point int
	var err error
	var pointS string = ""
	var pointV string = ""
	var multiplier bool = false
	stringArr := strings.Split(input, ".")
	stringArrLen := len(stringArr)
	if stringArrLen == 1 {
		number, err = strconv.Atoi(input)
		if err != nil {
			return "", errInvalidInput
		}
	} else if stringArrLen == 2 {
		number, err = strconv.Atoi(stringArr[0])
		if err != nil {
			return "", errInvalidInput
		}
		point, err = strconv.Atoi(stringArr[1])
		if err != nil {
			return "", errInvalidInput
		}
		if len(stringArr[1]) == 1 {
			multiplier = true
		}
	} else {
		return "", errInvalidInput
	}
	pointS = strconv.Itoa(point)
	if point > 100 {
		pointS = pointS[:2] + "." + pointS[2:]
		pointF, err3 := strconv.ParseFloat(pointS, 64)
		if err3 == nil {
			point = int(math.Round(pointF))
		}
	}
	if point != 0 && point < 100 {
		if multiplier == true {
			pointV = singles[point-1] + " zero"
		}
		pointV = " point " + singles[point/10] + " " + singles[point%10]
		word = pointV
	} else if point == 100 {
		return "", errInvalidInput
	}
	for number != 0 {
		if number%10 == 0 {
			word = "zero " + word
		} else {
			word = singles[number%10] + " " + word
		}
		number = number / 10
	}
	if point == 0 {
		return word + pointV, nil
	}
	return word, nil
}
