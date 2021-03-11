package currency

import (

	// "fmt"

	"math"
	"strconv"
	"strings"
)

// Digit2Word convert the indiviual digits to their word form.
func Digit2Word(input string) (string, error) {
	var word string
	var number int
	var point int
	var err error
	var pointS string = ""
	var pointV string = ""
	var multiplier bool = false
	var digit []string = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
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
			pointV = "point " + digit[point%10]
		} else {
			pointV = "point " + digit[point/10] + " " + digit[point%10]
		}
	} else if point == 100 {
		return "", errInvalidInput
	}
	if number == 0 {
		word = "zero "
	} else {
		for number != 0 {
			word = digit[number%10] + " " + word
			number = number / 10
		}
	}
	if point != 0 {
		return word + pointV, nil
	}
	return word[:len(word)-1], nil
}
