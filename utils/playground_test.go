package utils_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

/*
 * This is a playground, many times while developing you need to test
 * some concept. Use this test to quickly run the ideas and concept
 * using VS Code's integrated test environment!
 */
func TestPlayground(t *testing.T) {
	f := 0.12
	g := 0.12
	for i := 0; i < 20; i++ {
		f = f + 0.01
		g = g - 0.01

		// r := -1

		fmt.Printf(">>> %v\t%v\n", DigitsCount(f), f)

		// big.NewFloat(f)
		// fmt.Printf(">>> %v\t%v\t%v\t%v\t%v\n", f, calc.Round(f, r), Round(f, r), calc.Round(g, r), Round(g, r))
	}
	// println(calc.Round(0.0999999999999999, 15) * 2.0)

}

// NumberToWord converts the int64 numbers into international number word form.
// It can conver the numbers upto a billion. It also considers negative
// and franctional numbers.
func NumberToWord(number int64) string {
	var word string
	if number < 0 {
		word = "negative "
		number = -number
	}
	if number >= 1000000000 {
		word = word + NumberToWord(number/1000000000) + " billion "
		number = number % 1000000000
	}
	if number >= 1000000 {
		word = word + NumberToWord(number/1000000) + " million "
		number = number % 1000000
	}
	if number >= 1000 {
		word = word + NumberToWord(number/1000) + " thousand "
		number = number % 1000
	}
	if number >= 100 {
		word = word + NumberToWord(number/100) + " hundred "
		number = number % 100
	}
	if number >= 20 {
		if number%10 == 0 {
			word = word + NumberToWord(number/10)
		} else {
			word = word + NumberToWord(number/10) + "-" + NumberToWord(number%10)
		}
		number = number % 10
	}
	if number > 0 {
		if number == 1 {
			word = word + "one"
		} else {
			word = word + NumberToWord(number)
		}
	}
	return word
}

// NumberToRupees converts the int64 numbers into indian rupees word form.
// It can conver the numbers upto a crore.
// It also considers the fractional part of the number and converts it
// into the paise part.
// For exmple:
// NumberToRupees(1) = "one rupee"
// NumberToRupees(1.5) = "one rupee fifty paise"
// NumberToRupees(2.75) = "two rupees seventy five paise"
// NumberToRupees(123.34) = "one hundred twenty three rupees and thirty four paise"
// NumberToRupees(1234567.89) = "twelve crore thirty four lacs five hundred sixty seven rupees and eighty nine paise"
func NumberToRupees(number int64) string {
	var word string
	if number < 0 {
		word = "negative "
		number = -number
	}
	if number >= 1000000000 {
		word = word + NumberToRupees(number/1000000000) + " crore "
		number = number % 1000000000
	}
	if number >= 1000000 {
		word = word + NumberToRupees(number/1000000) + " lacs "
		number = number % 1000000
	}
	if number >= 1000 {
		word = word + NumberToRupees(number/1000) + " thousand "
		number = number % 1000
	}
	if number >= 100 {
		word = word + NumberToRupees(number/100) + " hundred "
		number = number % 100
	}
	if number >= 20 {
		if number%10 == 0 {
			word = word + NumberToRupees(number/10)
		} else {
			word = word + NumberToRupees(number/10) + "-" + NumberToRupees(number%10)
		}
		number = number % 10
	}
	if number > 0 {
		if number == 1 {
			word = word + " rupee"
		} else {
			word = word + NumberToRupees(number) + " rupee"
		}
	}
	return word
}

// Round performs the excel like rounding of the number.
// When the n is negative, it rounds the number to the left.
// When the n is positive, it rounds the fractional part of the number.
// For example:
// round(0.5, -1) = 0.0
// round(0.5, 0) = 0.0
// round(0.5, 1) = 0.5

// Ref: https://stackoverflow.com/questions/147770/how-to-round-a-number-to-significant-figures-in-go
// Ref: https://en.wikipedia.org/wiki/Rounding
// Ref: https://www.cockroachlabs.com/blog/rounding-implementations-in-go/
func Round(x float64, n int) float64 {
	pow := math.Pow(10, float64(n))
	if math.Abs(x*pow) > 1e17 {
		return x
	}
	v, frac := math.Modf(x * pow)
	if x > 0.0 {
		if frac > 0.5 || (frac == 0.5 && uint64(v)%2 != 0) {
			v += 1.0
		}
	} else {
		if frac < -0.5 || (frac == -0.5 && uint64(v)%2 != 0) {
			v -= 1.0
		}
	}
	return v / pow
}

// DigitsCount returns the total number of numerical digits in the
// left of the decimal point plug the right of the decimal point.
// DigitsCount(12345) = 5
// DigitsCount(12345.67) = 7
// DigitsCount(123.456789) = 9
// DigitsCount(-0.123456789) = 9
func DigitsCount(number float64) int {
	s := strconv.FormatFloat(number, 'f', -1, 64)
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}
	for _, c := range s {
		if c == '.' {
			return len(s) - 1
		}
	}
	return len(s)
}

// FracDigits returns the total number of digits
// after the decimal point.
// For exmple:
// FracDigits(12345) = 0
// FracDigits(12345.67) = 2
// FracDigits(123.456789) = 6
func FracDigits(x float64) int {
	return DigitsCount(x-math.Floor(x)) - 1
}

// IntDigits returns the number of digits in the non-fractional
// part of the number.
func IntDigits(x float64) int {
	return int(math.Floor(math.Log10(math.Abs(x))))
}

// Round15 performs the roudning of the number to the 15 digit from
// the left.
// For example:
// Round15(123456789.123456789) = 123456789.123457
func Round15(x float64) float64 {
	return Round(x, 15-IntDigits(x))
}

// Ceil performs the excel like ceiling of the number to the nearest factor.
// For example
// Ceil(0.5) = 1
// Ceil(1.5) = 2
// Ceil(2.5) = 2
func Ceil(x float64, factor float64) float64 {
	if x < 0 {
		return math.Ceil(x/factor) * factor
	}
	return math.Ceil(x/factor) * factor
}

// Floor performs the excel like floor of the number to the nearest factor.
// For example
// Floor(0.5) = 0
// Floor(1.5) = 1
// Floor(2.5) = 2
func Floor(x float64, factor float64) float64 {
	if x < 0 {
		return math.Floor(x/factor) * factor
	}
	return math.Floor(x/factor) * factor
}
