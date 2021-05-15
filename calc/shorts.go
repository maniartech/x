package calc

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Trignometric Functions
//
// Acos returns inverse cosine of a number.
//
// Arguments
//
// number : Required. The number is an input to Acos function, it must be from -1 to 1.
//
// Example
//
//    Acos(0.235) // Returns 1.3335777587004611
//    Acos(-0.546)  // Returns 2.1483786013238344
//
func Acos(input float64) float64 { return math.Acos(input) }

//
// Acosh returns inverse hperbolic cosine of a number.
//
// Arguments
//
// number : Required. The number is an input to hyperbolic cosine function, it must be eual to or greater than 1.
//
// Example
//
//    Acosh(561.15) // Returns 7.023134636091518
//    Acosh(1546)  // Returns 8.036573305109862
//
func Acosh(input interface{}) float64 { return math.Acosh(utils.ToFloat64(input)) }

//
// Asin returns inverse sine of a number.
//
// Arguments
//
// number : Required. The number is an input to Asin function, it must be between -1 and 1.
//
// Example
//
//    Asin(0.26) // Returns 0.2630222029084689
//    Asin(-0.546)  // Returns -0.5775822745289381
//
func Asin(input interface{}) float64 { return math.Asin(utils.ToFloat64(input)) }

//
// Asinh returns inverse hperbolic sine of a number.
//
// Arguments
//
// number : Required. The number is an input to Asinh function, it can be any real number.
//
// Example
//
//    Asinh(-10.1653) // Returns -3.014537762282636
//    Asinh(-1546)  // Returns -8.03657351430473
//

func Asinh(input interface{}) float64 { return math.Asinh(utils.ToFloat64(input)) }

//
// Atan returns inverse tagnent of a number.
//
// Arguments
//
// number : Required. The numbers is the tagnent of an angle you want.
//
// Example
//
//    Atan(0.235) // Returns 0.23081196040266755
//    Atan(-0.546)  // Returns -0.49976700744972385
//
func Atan(input interface{}) float64 { return math.Atan(utils.ToFloat64(input)) }

//
// Atanh returns inverse hyperbolic tagnent of a number.
//
// Arguments
//
// number : Required. The number is a real number between 1 and -1 for which inverse hyperbolic of tagnent is calculated.
//
// Example
//
//    Atanh(0.99) // Returns 2.6466524123622457
//    Atanh(-0.2686)  // Returns -0.275354351742258
//
func Atanh(input interface{}) float64 { return math.Atanh(utils.ToFloat64(input)) }

//
// Cos returns the cosine of the given  number.
//
// Arguments
//
// number : Required. The numbers is the angle in radian for which you want cosine.
//
// Example
//
//    Cos(0.235) // Returns 0.9725143413326429
//    Cos(-3.5)  // Returns -0.9364566872907963
//
func Cos(input interface{}) float64 { return math.Cos(utils.ToFloat64(input)) }

//
// Cosh function returns hyperbolic cosine of a number.
//
// Arguments
//
// number : Required. The number is an input to the hyperbolic cosine funtion, it can be any real number.
//
// Example
//
//    Cosh(0.99) // Returns 1.5314055816856538
//    Cosh(-0.2686)  // Returns 1.036290378879286
//
func Cosh(input interface{}) float64 { return math.Cosh(utils.ToFloat64(input)) }

//
// Sin function returns sine of a given number.
//
// Arguments
//
// number : Required. The number is an angle in radian for which you want sine.
//
// Example
//
//    Sin(0.235) // Returns 0.23284298551241678
//    Sin(-3.5)  // Returns 0.35078322768961984
//
func Sin(input interface{}) float64 { return math.Sin(utils.ToFloat64(input)) }

//
// Sinh function returns hyperbolic sine of a number.
//
// Arguments
//
// number : Required. The number is an input to the hyperbolic sine funtion, it can be any real number.
//
// Example
//
//    Sinh(0.99) // Returns 1.1598288906636083
//    Sinh(-0.2686)  // Returns -0.2718414047892157
//
func Sinh(input interface{}) float64 { return math.Sinh(utils.ToFloat64(input)) }

//
// Tan function returns tagnent of a number.
//
// Arguments
//
// number : Required. The number is an input to Tan funtion, it can be any real number.
//
// Example
//
//    Tan(-36) // Returns -7.750470905699147
//    Tan(0.235)  // Returns 0.2394237037094491
//
func Tan(input interface{}) float64 { return math.Tan(utils.ToFloat64(input)) }

//
// Tanh function returns hyperbolic tagnent of a number.
//
// Arguments
//
// number : Required. The number is an input to the hyperbolic Tan funtion, it can be any real number.
//
// Example
//
//    Tahn(0.99) // Returns 0.7573623242165263
//    Tanh(-1.256)  // Returns -0.8499575891648927
//
func Tanh(input interface{}) float64 { return math.Tanh(utils.ToFloat64(input)) }

//
// CSC function returns cosecant of a number.
//
// Arguments
//
// number : Required. The number is an input to the CSC funtion, it can be any real number.
//
// Example
//
//    CSC(-36) // Returns 1.0082892940581898
//    CSC(0.235)  // Returns 4.294739640961497
//
func CSC(input interface{}) float64 { return 1 / math.Sin(utils.ToFloat64(input)) }

//
// Deg2Rad function converts degree to radian.
//
// Arguments
//
// angle : Required. The angle is a degree that you want to convert to radian.
//
// Example
//
//    Deg2Rad(10) //  Returns 0.17453292519943295
//    Deg2Rad(-10)  // Returns -0.17453292519943295
//
func Deg2Rad(input interface{}) float64 { return utils.ToFloat64(input) * (math.Pi / 180) }

//
// Rad2Deg function converts radian to degree.
//
// Arguments
//
// angle : Required. The angle is in radian that you want to convert to degree.
//
// Example
//
//    Rad2Deg(-10.15631) //  Returns -581.9136984265131
//    Rad2Dag(231.5163)  // Returns 13264.90687848462
//
func Rad2Deg(input interface{}) float64 { return utils.ToFloat64(input) * (180 / math.Pi) }

// Math Functions
//
// Abs function returns absolute value of input number, that is a value of a number without sign.
//
// Arguments
//
// number : Required. The number can be any real number for which you want absolut value.
//
// Example
//
//    Abs(-10.15631) //  Returns 10.15631
//    Abs(10.15631)  // Returns 10.15631
//
func Abs(input interface{}) float64 { return math.Abs(utils.ToFloat64(input)) }

//
// Even function returns nearest even integer rounded up to the given input number.
//
// Arguments
//
// number : Required. The number you wanted to round up.
//
// Example
//
//    Even(-11.15631) // Returns -12
//    Even(6.2)  // Returns 8
//
func Even(input interface{}) int {
	val := utils.ToFloat64(input)
	if val < 0 {
		valI := utils.ToInt(input)
		return int(Floor(valI, 2))
	}
	return int(Ceiling(val, 2))
}

// Odd function returns nearest odd integer after rounding up the input number.
//
// Arguments
//
// number : Required. The number you want to round up.
//
// Example
//
//    Odd(-122.15631) //  Returns -123
//    Odd(123.15631)  // Returns 125
//
func Odd(input interface{}) int {
	val := utils.ToInt(input)
	if val < 0 {
		return int(Floor(val, 2)) - 1
	}
	return int(Ceiling(val, 2)) + 1
	// return int(((math.Floor(Abs(utils.ToFloat64(input))/2) - 1) * 2 * Abs(input) / utils.ToFloat64(input)) + 3*Abs(input)/utils.ToFloat64(input))
}

// Int function rounds up a number to the nearest integer.
//
// Arguments
//
// number : Required. The number you want to round up to the nearest integer.
//
// Example
//
//    Int(-122.15631) //  Returns -123
//    Int(123.15631)  // Returns 123
//
func Int(input interface{}) int { return int(math.Floor(utils.ToFloat64(input))) }

// Sqrt function returns a positive square root for an input value.
//
// Arguments
//
// number : Required. The number for which you want the square root.
//
// Remark
//
// If the number is less than zero then the function returns Invalid Input Error message.
//
// Example
//
//    Sqrt(16) //  Returns 4.0
//    Sqrt(123.15631)  // Returns 11.097581268006106
//
func Sqrt(input interface{}) float64 {
	if utils.ToFloat64(input) < 0 {
		panic(core.ErrInvalidInput)
	}
	return math.Sqrt(utils.ToFloat64(input))
}

// SqrtPi function returns the value of number multiplied by Pi.
//
// Arguments
//
// number : Required. The number to which Pi is multiplied.
//
// Remark
//
// If the number is less than zero then the function returns Invalid Input Error message.
//
//Example
//
//    SqrtPi(16) //  Returns 7.0898154036220635
//    SqrtPi(123.15631)  // Returns 19.669950654214343
//
func SqrtPi(input interface{}) float64 {
	if utils.ToFloat64(input) < 0 {
		panic(core.ErrInvalidInput)
	}
	return math.Sqrt(utils.ToFloat64(input) * math.Pi)
}

// Quotient function returns only the integer portion of the result for division operation and discard the remainder.
//
// Arguments
//
// numerator : Required. The divident
//
// denominator: Required. The divisor
//
// Example
//
//    Quotient(16, 5) //  Returns 3
//    Quotient(-26, 5)  // Returns -5
//
func Quotient(n, d interface{}) int {
	if d == 0 {
		panic(core.ErrDivideBy0)
	}
	return int(utils.ToFloat64(n) / utils.ToFloat64(d))
}

// LN function returns the natural logarithm of a number that is inverse of an exponential function.
//
// Arguments
//
// number: Required. The number for which you want the natural logarithm. It can be any real number.
//
// Remark
//
// LN function is inverse of exponetial function.
//
// Example
//
//    LN(100) //  Returns 4.605170185988092
//    LN(16)  // Returns 2.772588722239781
//
func LN(input interface{}) float64 { return math.Log(utils.ToFloat64(input)) }

// Log10 function returns the base 10 logarithm of an input number.
//
// Arguments
//
// number: Required. The number for which you want the base-10 logarithm. It must be a positive real number.
//
// Example
//
//    Log10 (122) //  Returns 2.0863598306747484
//    Log10(122.15631)  // Returns 2.086915905426539
//
func Log10(input interface{}) float64 { return math.Log10(utils.ToFloat64(input)) }

// Exp function returns the result of the value rised to the power of a number.
//
// Arguments
//
// number: Required. The number is an exponent  applied to the constant base e.
//
// Remark
//
// Exp function is inverse of LN function.
//
// Example
//
//    Exp (0) //  Returns 1.0
//    Exp(0.002)  // Returns 1.0020020013340003
//
func Exp(input interface{}) float64 { return math.Exp(utils.ToFloat64(input)) }

// Trunc function returns a truncated number to the given precision.
//
// Arguments
//
// number: Required. The number you want to truncate.
//
// Remark
//
// Trunc and Int looks similar functions as they return integers.
// Trunc removes fractional part of number and Int rounds down number to closest integer according to number's fractional part.
// But Trunc and Int works differently for negative numbers.
// Trunc(-1.6) returns -1 and Int(-1.6) returns -2 as the value of -2 is more lower than -1.
//
// Example
//
//    Trunc(-16) //  Returns -16.0
//    Trunc(-122.15631)  // Returns -122.0
//
func Trunc(input interface{}) float64 { return math.Trunc(utils.ToFloat64(input)) }
