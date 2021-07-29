package finance

import (
	"fmt"
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/text"
	"github.com/maniartech/x/utils"
)

func Dollarde(Dollar, Fraction interface{}) float64 {
	dollar := utils.ToFloat64(Dollar)
	//Truncating Fraction
	fraction := math.Floor(utils.ToFloat64(Fraction))

	//Checking if Fraction is not a -ve number
	if fraction < 0 {
		panic(core.ErrInvalidInput)
	}

	//Checking if fraction values will cause a Division by 0 Errror
	if fraction >= 0 && fraction < 1 {
		panic(core.ErrDivideBy0)
	}

	//Setting the value of numerator to the deciamal digits
	//(rounding numbers to its 2nd last decimal digit so it doesn't cause floating point errors)
	//i.e 1.16 so, numerator is 16
	var numerator float64 = calc.Round(dollar-math.Floor(dollar), 15-calc.NumberOfDigits(utils.ToInt(dollar)))
	//If there is one decimal digit then its multiplied by 10
	//i.e. 1.1 == 10 and not 1
	fmt.Println(text.Len(numerator) == 3)
	if text.Len(numerator) == 3 {
		numerator *= 100
	} else {
		numerator = numerator * math.Pow10(calc.NumberOfDecimalDigits(numerator))
	}

	val := math.Floor(dollar) + calc.Divide(numerator, fraction)
	return calc.Round(val, 15-calc.NumberOfDigits(utils.ToInt(val)))

}

func Dollarfr(Dollar, Fraction interface{}) float64 {
	dollar := utils.ToFloat64(Dollar)
	fraction := math.Floor(utils.ToFloat64(Fraction))

	//Checking if
	if dollar < 0 {
		panic(core.ErrInvalidInput)
	}
	if dollar == 0 {
		panic(core.ErrDivideBy0)
	}

	//Setting the value of numerator to the deciamal digits
	//(rounding numbers to its 2nd last decimal digit so it doesn't cause floating point errors)
	//i.e 1.16 so, numerator is 0.16
	var decimal float64 = calc.Round(dollar-math.Floor(dollar), 15-calc.NumberOfDigits(utils.ToInt(dollar)))
	decimal = decimal * fraction

	return math.Floor(dollar) + (0.01 * decimal)
}
