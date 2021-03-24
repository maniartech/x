package currency

import (

	// "fmt"

	"math"
	"strconv"
	"strings"
)

// Digit2Word convert the indiviual digits to their word form.
func Digit2Word(input string) (string, error) {
	//Variable initalization
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
	//Spliting the number based on `.` and also will check how many `.` are there in the input
	//If there is more htan one `.` then input is invalid
	stringArr := strings.Split(input, ".")
	stringArrLen := len(stringArr)
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
		point, err = strconv.Atoi(stringArr[1])
		if err != nil {
			return "", ErrInvalidInput
		}
		if len(stringArr[1]) == 1 {
			multiplier = true
		}
	} else {
		return "", ErrInvalidInput
	}
	//Converting the input from string to int
	pointS = strconv.Itoa(point)
	if point > 100 {
		pointS = pointS[:2] + "." + pointS[2:]
		pointF, err3 := strconv.ParseFloat(pointS, 64)
		if err3 == nil {
			point = int(math.Round(pointF))
		}
	}
	// If there is only one digit after the `.` then the numbers is multiplied by 10
	if point != 0 && point < 100 {
		if multiplier {
			pointV = "point " + digit[point%10]
		} else {
			pointV = "point " + digit[point/10] + " " + digit[point%10]
		}
	} else if point == 100 {
		return "", ErrInvalidInput
	}
	//If the number == 0 then return 0
	if number == 0 {
		word = "zero "
	} else {
		for number != 0 {
			word = digit[number%10] + " " + word
			number = number / 10
		}
	}
	//Checking if the value of the point != 0 if not then we also return the point part in the reutrn
	if point != 0 {
		return word + pointV, nil
	}
	return word[:len(word)-1], nil
}
