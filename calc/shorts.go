package calc

import (
	"math"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Trignometric Functions
//
// Acos calculates the inverse cosine of a number and returns an angle in the radians between 0 and π.
// Geometrically it is the ratio of a triangle's adjacent side over its hypotenuse.
//
// Arguments
//
// number : Required. The inverse cosine of an angle you want, it must be between -1 to 1.
//
// Remark
//
// If the input value of an argument is outside the range of -1 to 1 then it will throw an error.
//
// Example
//
//    calc.Acos(0.235) // Returns 1.3335777587004611
//    calc.Acos(-0.546)  // Returns 2.1483786013238344
//    calc.Acos(-1) // Returns 3.141593
//    calc.Acos(0) // Returns 1.570796
//    calc.Acos(1) // Returns 0
//    calc.Acos(-1.5689) // Returns Error Message
//    calc.Acos(1.5698)  // Returns Error Message
//
func Acos(input float64) float64 { return math.Acos(input) }

//
// Acosh calculates the inverse hyperbolic cosine of a number and returns the angle in the radian.
//
// Arguments
//
// number : Required. The hyperbolic cosine of an angle you want, it must be greater than or equal to 1.
//
// Remark
//
// If the input value of an argument is less than 1 then it will throw an error.
//
// Example
//
//    calc.Acosh(561.15) // Returns 7.023134636091518
//    calc.Acosh(1546)  // Returns 8.036573305109862
//    calc.Acosh(1) // Returns 0.0
//    calc.Acosh(0) // Returns Invalid Input Error Message
//    calc.Acosh(-1) // Returns Invalid Input Error Message
//
func Acosh(input interface{}) float64 { return math.Acosh(utils.ToFloat64(input)) }

//
// Asin calculates the inverse sine of a number and returns an angle in the radians between -π/2 and π/2.
// Geometrically it is the ratio of a triangle's opposite side over its hypotenuse.
//
// Arguments
//
// number : Required. The number is an input to Asin function, it must be between -1 and 1.
//
// Remark
//
// If the input value of an argument is outside the range of -1 to 1 then it will throw an error.
//
// Example
//
//    calc.Asin(0.26) // Returns 0.2630222029084689
//    calc.Asin(-0.546)  // Returns -0.5775822745289381
//    calc.Asin(0) // Returns 0.0
//    calc.Asin(-1) // Returns -1.5708
//    calc.Asin(1) // Returns 1.570796
//
func Asin(input interface{}) float64 { return math.Asin(utils.ToFloat64(input)) }

//
// Asinh function returns the inverse hyperbolic sine of a number and returns the angle in the radian.
//
// Arguments
//
// number : Required. The number is an input to Asinh function, it can be any real number.
//
// Example
//
//    calc.Asinh(-10.1653) // Returns -3.014537762282636
//    calc.Asinh(-1546)  // Returns -8.03657351430473
//    calc.Asinh(0) // Returns 0.0
//    calc.Asinh(3.14) // Returns(1.8618125572133835)
//
func Asinh(input interface{}) float64 { return math.Asinh(utils.ToFloat64(input)) }

//
// Atan calculates the inverse tangent of a number and returns an angle in the radians between -π/2 and π/2.
// Geometrically it is the ratio of the right triangle's opposite side over its adjacent side.
//
// Arguments
//
// number : Required. The number is a tangent of an angle you want. It can be any real number.
//
// Example
//
//    calc.Atan(0.235) // Returns 0.23081196040266755
//    calc.Atan(-0.546)  // Returns -0.49976700744972385
//    calc.Atan(1) // Returns 0.785398
//    calc.Atan(-1) // Returns 0.7854
//    calc.Atan(0) // Returns 0.0
//
func Atan(input interface{}) float64 { return math.Atan(utils.ToFloat64(input)) }

//
// Atanh returns the inverse hyperbolic tangent of a number and returns the angle in radian.
//
// Arguments
//
// number : Required.  The inverse hyperbolic tangent of the angle you want must be between -1 to 1.
//
// Remark
//
// If the input value of an argument is outside the range of -1 to 1 then it will throw an error.
// The inverse hyperbolic tangent of 1 and -1 are Infinity and -Infinity.
//
// Example
//
//    calc.Atanh(0.99) // Returns 2.6466524123622457
//    calc.Atanh(-0.2686)  // Returns -0.275354351742258
//    calc.Atanh(0.651) // Returns 0.7770322608588505
//    calc.Atanh(-23) // Returns an error
//    calc.Atanh(23) // Returns an error
//
func Atanh(input interface{}) float64 { return math.Atanh(utils.ToFloat64(input)) }

//
// Cos function calculates the cosine of the given number and returns the angle in radian.
// Geometrically it is the ratio of a right triangle's adjacent side over its hypotenuse.
//
// Arguments
//
// number : Required. An angle in the radian for which you want cosine.
//
// Example
//
//    calc.Cos(0.235) // Returns 0.9725143413326429
//    calc.Cos(-3.5)  // Returns -0.9364566872907963
//    calc.Cos(0) // Returns 1.0
//    calc.Cos(-36) // Returns -0.12796368962740468
//    calc.Cos(1) // Returns 0.5403023058681398
//
func Cos(input interface{}) float64 { return math.Cos(utils.ToFloat64(input)) }

//
// Cosh function calculates the hyperbolic cosine of a number and returns an angle in the radian.
//
// Arguments
//
// number : Required. The number is an input to the hyperbolic cosine function, it can be any real number.
//
// Example
//
//    calc.Cosh(0.99) // Returns 1.5314055816856538
//    calc.Cosh(-0.2686)  // Returns 1.036290378879286
//    calc.Cosh(0) // Returns 1
//    calc.Cosh(0.99) // Returns 1.5314055816856538
//    calc.Cosh(1)  // Returns 1.543081
//
func Cosh(input interface{}) float64 { return math.Cosh(utils.ToFloat64(input)) }

//
// Sin function calculates the sine of a given number and returns the angle in radian.
// Geometrically it is the ratio of the right triangle's opposite side over its hypotenuse.
//
// Arguments
//
// number : Required. The number is an angle in radian for which you want sine.
//
// Example
//
//    calc.Sin(0.235) // Returns 0.23284298551241678
//    calc.Sin(-3.5)  // Returns 0.35078322768961984
//    calc.Sin(0) // Returns 0.0
//    calc.Sin(1) // Returns 0.841471
//    calc.Sin(23) // Returns 0.84622
//
func Sin(input interface{}) float64 { return math.Sin(utils.ToFloat64(input)) }

//
// Sinh function calculates the hyperbolic sine of a number and returns the angle in radian.
//
// Arguments
//
// number : Required. The number is an input to the hyperbolic sine function, it can be any real number.
//
// Example
//
//    calc.Sinh(0.99) // Returns 1.1598288906636083
//    calc.Sinh(-0.2686)  // Returns -0.2718414047892157
//    calc.Sinh(0) // Returns 0
//    calc.Sinh(1)  //Returns 1.175201
//    calc.Sinh(10) // Returns 11013.23287
//
func Sinh(input interface{}) float64 { return math.Sinh(utils.ToFloat64(input)) }

//
// Tan function calculates the tangent of a number and returns the angle in radian.
// Geometrically it is the ratio of the right triangle's opposite side over its adjacent side.
//
// Arguments
//
// number : Required. The number is an input to Tan function, it can be any real number.
//
// Example
//
//    calc.Tan(-36) // Returns -7.750470905699147
//    calc.Tan(0.235)  // Returns 0.2394237037094491
//    calc.Tan(0) // Returns 0.0
//    calc.Tan(1) // Returns 1.557407725
//    calc.Tan(-3.5) // Returns -0.3745856401585947
//
func Tan(input interface{}) float64 { return math.Tan(utils.ToFloat64(input)) }

//
// Tanh function calculates the hyperbolic tangent of a number and returns the angle in radian.
//
// Arguments
//
// number : Required. The number is an input to the hyperbolic Tan function, it can be any real number.
//
// Example
//
//    calc.Tanh(0.99) // Returns 0.7573623242165263
//    calc.Tanh(-1.256)  // Returns -0.8499575891648927
//    calc.Tanh(0) // Returns 0.0
//    calc.Tanh(2) // Returns 0.9640275800758169
//
func Tanh(input interface{}) float64 { return math.Tanh(utils.ToFloat64(input)) }

//
// CSC function calculates the cosecant of a number and returns the angle in the radian.
// Geometrically it is the ratio of the right triangle's hypotenuse divided by its opposite side.
//
// Arguments
//
// number : Required. The number is an input to the CSC function, it can be any real number.
//
// Example
//
//    calc.CSC(-36) // Returns 1.0082892940581898
//    calc.CSC(0.235) // Returns 4.294739640961497
//    calc.CSC(1) // Returns 1.1883951057781212
//    calc.CSC(15) // Returns 1.538
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
//    calc.Deg2Rad(10) //  Returns 0.17453292519943295
//    calc.Deg2Rad(-10)  // Returns -0.17453292519943295
//    calc.Deg2Rad(-10.15631) // Returns -0.1772610493532253
//    calc.Deg2Rad(180) // Returns 3.14159
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
//    calc.Rad2Deg(-10.15631) //  Returns -581.9136984265131
//    calc.Rad2Dag(231.5163)  // Returns 13264.90687848462
//    calc.Rad2Deg(1)  // Returns 57.29577951308232
//    calc.Rad2Deg(-5)  // Returns -286.4788975654116
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
//    calc.Abs(-10.15631) //  Returns 10.15631
//    calc.Abs(10.15631)  // Returns 10.15631
//    calc.Abs(-10) // Returns 10.0
//
func Abs(input interface{}) float64 { return math.Abs(utils.ToFloat64(input)) }

//
// Even function returns nearest even integer rounded up to the given input number.
//
// Arguments
//
// number : Required. The number you wanted to round up.
//
// Remark
//
// Even function rounds positive and negative number up (away from zero)
// so the positive number becomes larger and negative number becomes smaller.
//
// Example
//
//    calc.Even(-11.15631) // Returns -12
//    calc.Even(6.2)  // Returns 8
//    calc.Even(11) // Returns 12
//
func Even(input interface{}) int {
	val := utils.ToFloat64(input)
	if val < 0 {
		valI := utils.ToInt(input)
		return int(Floor(valI, 2))
	}
	return int(Ceiling(val, 2))
}

// Odd function returns the nearest odd integer after rounding up the input number.
//
// Arguments
//
// number : Required. The number you want to round up.
//
// Returns
//
// Even function rounds a positive and negative number up (away from zero)
// so the positive number becomes larger and the negative number becomes smaller.
//
// Example
//
//    calc.Odd(-122.15631) //  Returns -123
//    calc.Odd(123.15631)  // Returns 125
//    calc.Odd(16) // Returns 17
//
func Odd(input interface{}) int {
	val := utils.ToInt(input)
	if val < 0 {
		return int(Floor(val, 2)) - 1
	}
	return int(Ceiling(val, 2)) + 1
	// return int(((math.Floor(Abs(utils.ToFloat64(input))/2) - 1) * 2 * Abs(input) / utils.ToFloat64(input)) + 3*Abs(input)/utils.ToFloat64(input))
}

// Int function rounds down a number to the nearest integer.
//
// Arguments
//
// number : Required. The number you want to round down to the nearest integer.
//
// Example
//
//    calc.Int(-122.15631) //  Returns -123
//    calc.Int(123.15631)  // Returns 123
//    calc.Int(16) // Returns 16
//    calc.Int(-16) // Returns -16
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
// If the number is less than zero then the function returns an Invalid Input Error message.
//
// Example
//
//    calc.Sqrt(16) //  Returns 4.0
//    calc.Sqrt(123.15631)  // Returns 11.097581268006106
//    calc.sqrt(0)  // Returns 0
//    calc.Sqrt(-16) // Returns Invalid Input Error
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
// If the number is less than zero then the function returns an Invalid Input Error message.
//
//Example
//
//    calc.SqrtPi(16) //  Returns 7.0898154036220635
//    calc.SqrtPi(123.15631)  // Returns 19.669950654214343
//    calc.SqrtPi(-16) // Returns Invalid Input Error
//
func SqrtPi(input interface{}) float64 {
	if utils.ToFloat64(input) < 0 {
		panic(core.ErrInvalidInput)
	}
	return math.Sqrt(utils.ToFloat64(input) * math.Pi)
}

// Quotient function returns only the integer portion of the result for division operation and discards the remainder.
//
// Arguments
//
// numerator : Required. The dividend
//
// denominator: Required. The divisor
//
// Example
//
//    calc.Quotient(16, 5) //  Returns 3
//    calc.Quotient(-26, 5)  // Returns -5
//    calc.Quotient(-122.15631, 2.546) // Returns -47
//    calc.Quotient(123.15631, 48)  // Returns 2
//
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
// LN function is inverse of an exponential function.
//
// Example
//
//    calc.LN(100) //  Returns 4.605170185988092
//    calc.LN(16)  // Returns 2.772588722239781
//    calc.LN(122.15631) // 4.805301454167321
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
//    calc.Log10 (122) //  Returns 2.0863598306747484
//    calc.Log10(122.15631)  // Returns 2.086915905426539
//    calc.Log(10000) // Returns 4.0
//    calc.Log(1) // Returns 0.0
//
func Log10(input interface{}) float64 { return math.Log10(utils.ToFloat64(input)) }

// Exp function returns the result of a constant e raised to the power of a number.
// The value of constant e is approximately equal to 2.71828182845904, the base of the natural logarithm.
//
// Arguments
//
// number: Required. The number is an exponent applied to the constant base e.
//
// Remark
//
// Exp function is inverse of LN function.
//
// Example
//
//    calc.Exp (0) //  Returns 1.0
//    calc.Exp(0.002)  // Returns 1.0020020013340003
//    calc.Exp(2) // Returns 7.38905609893
//    calc.Exp(1) // Returns 2.71828182846
//
func Exp(input interface{}) float64 { return math.Exp(utils.ToFloat64(input)) }

//Trunc function returns a truncated number to the given precision.
//
// Arguments
//
// number: Required. The number you want to truncate.
//
// Remark
//
// Trunc and Int look similar functions as they return integers.
// Trunc removes fractional part of a number and Int rounds down number to closest integer according to number's fractional part.
// But Trunc and Int work differently for negative numbers.
// Trunc(-1.6) returns -1 and Int(-1.6) returns -2 as the value of -2 is more lower than -1.
//
// Example
//
//    calc.Trunc(-16) //  Returns -16.0
//    calc.Trunc(-122.15631)  // Returns -122.0
//    calc.Trunc(16) //  Returns 16.0
//    calc.Trunc(123.15631) //  Returns 123.0
//
func Trunc(input interface{}) float64 { return math.Trunc(utils.ToFloat64(input)) }
