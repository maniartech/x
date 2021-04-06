package currency

import (
	"math"
	"strconv"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
)

// Digit2Word convert the indiviual digits to their word form.
func Digit2Word(input string) (string, error) {
	//Variable initalization
	var word, pointV, pointS string
	var digit []string = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	var number, point int
	var single bool = false

	//Converting the input to int and then spliting it into two
	number, point = ConvertToInt(input)
	pointS = strconv.Itoa(point)
	if len(pointS) < 2 {
		single = true
	}
	//Converting the input from string to int
	if point > 100 {
		point = int((calc.Round(float64(point), -(calc.NumberOfDigits(point) - 2))) / math.Pow10(calc.NumberOfDigits(point)-2))
	}
	if point != 0 && point < 100 {
		if single {
			pointV = "point " + digit[point%10]
		} else {
			pointV = "point " + digit[point/10] + " " + digit[point%10]
		}
	} else if point == 100 {
		panic(core.ErrInvalidInput)
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
