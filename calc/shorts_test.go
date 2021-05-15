package calc_test

import (
	"math"
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestTrignometicFunc(t *testing.T) {
	//Acos
	assert.Equal(t, 0.0, calc.Acos(1))
	assert.Equal(t, math.Pi, calc.Acos(-1))
	assert.Equal(t, 1.3335777587004611, calc.Acos(0.235))
	assert.Equal(t, 2.1483786013238344, calc.Acos(-0.546))

	//Acosh
	assert.Equal(t, 0.0, calc.Acosh(1))
	assert.Equal(t, 1.7627471740390859, calc.Acosh(3))
	assert.Equal(t, 7.023134636091518, calc.Acosh(561.15))
	assert.Equal(t, 8.036573305109862, calc.Acosh(1546))

	//Asin
	assert.Equal(t, 0.0, calc.Asin(0))
	assert.Equal(t, -math.Pi/2, calc.Asin(-1))
	assert.Equal(t, 0.2630222029084689, calc.Asin(0.26))
	assert.Equal(t, -0.5775822745289381, calc.Asin(-0.546))

	//Asinh
	assert.Equal(t, 0.0, calc.Asinh(0))
	assert.Equal(t, 1.8618125572133835, calc.Asinh(3.14))
	assert.Equal(t, -3.014537762282636, calc.Asinh(-10.1653))
	assert.Equal(t, -8.03657351430473, calc.Asinh(-1546))

	//Atan
	assert.Equal(t, math.Pi/4, calc.Atan(1))
	assert.Equal(t, -1.56279649745501, calc.Atan(-125))
	assert.Equal(t, 0.23081196040266755, calc.Atan(0.235))
	assert.Equal(t, -0.49976700744972385, calc.Atan(-0.546))

	//Atanh
	assert.Equal(t, 0.0, calc.Atanh(0))
	assert.Equal(t, 0.7770322608588505, calc.Atanh(0.651))
	assert.Equal(t, 2.6466524123622457, calc.Atanh(0.99))
	assert.Equal(t, -0.275354351742258, calc.Atanh(-0.2686))

	//Cosh
	assert.Equal(t, 1.0, calc.Cosh(0))
	assert.Equal(t, 3.7621956910836314, calc.Cosh(2))
	assert.Equal(t, 1.5314055816856538, calc.Cosh(0.99))
	assert.Equal(t, 1.0362903788792859, calc.Cosh(-0.2686))

	//Cos
	assert.Equal(t, 1.0, calc.Cos(0))
	assert.Equal(t, 0.5403023058681398, calc.Cos(1))
	assert.Equal(t, -0.12796368962740468, calc.Cos(-36))
	assert.Equal(t, 0.9725143413326429, calc.Cos(0.235))
	assert.Equal(t, -0.9364566872907963, calc.Cos(-3.5))

	//Sin
	assert.Equal(t, 0.0, calc.Sin(0))
	assert.Equal(t, 0.9917788534431158, calc.Sin(-36))
	assert.Equal(t, 0.23284298551241678, calc.Sin(0.235))
	assert.Equal(t, 0.35078322768961984, calc.Sin(-3.5))

	//Sinh
	assert.Equal(t, 0.0, calc.Sinh(0))
	assert.Equal(t, 3.626860407847019, calc.Sinh(2))
	assert.Equal(t, 1.1598288906636083, calc.Sinh(0.99))
	assert.Equal(t, -0.2718414047892157, calc.Sinh(-0.2686))

	//Tan
	assert.Equal(t, 0.0, calc.Tan(0))
	assert.Equal(t, -7.750470905699147, calc.Tan(-36))
	assert.Equal(t, 0.2394237037094491, calc.Tan(0.235))
	assert.Equal(t, -0.3745856401585947, calc.Tan(-3.5))

	//Tanh
	assert.Equal(t, 0.0, calc.Tanh(0))
	assert.Equal(t, 0.9640275800758169, calc.Tanh(2))
	assert.Equal(t, 0.7573623242165263, calc.Tanh(0.99))
	assert.Equal(t, -0.8499575891648927, calc.Tanh(-1.256))

	//CSC
	assert.Equal(t, 1.1883951057781212, calc.CSC(1))
	assert.Equal(t, 1.0082892940581898, calc.CSC(-36))
	assert.Equal(t, 4.294739640961497, calc.CSC(0.235))
	assert.Equal(t, 2.850763437540464, calc.CSC(-3.5))
}

func TestConvFunc(t *testing.T) {
	//Deg2Rad
	assert.Equal(t, 0.17453292519943295, calc.Deg2Rad(10))
	assert.Equal(t, -0.17453292519943295, calc.Deg2Rad(-10))
	assert.Equal(t, -0.1772610493532253, calc.Deg2Rad(-10.15631))
	assert.Equal(t, 4.040721707034948, calc.Deg2Rad(231.5163))

	//Rad2Deg
	assert.Equal(t, 57.29577951308232, calc.Rad2Deg(1))
	assert.Equal(t, -286.4788975654116, calc.Rad2Deg(-5))
	assert.Equal(t, -581.9136984265131, calc.Rad2Deg(-10.15631))
	assert.Equal(t, 13264.90687848462, calc.Rad2Deg(231.5163))
}

func TestMathFunc(t *testing.T) {
	//Abs
	assert.Equal(t, 10.0, calc.Abs(10))
	assert.Equal(t, 10.0, calc.Abs(-10))
	assert.Equal(t, 10.15631, calc.Abs(-10.15631))
	assert.Equal(t, 10.15631, calc.Abs(10.15631))

	//Even
	assert.Equal(t, 12, calc.Even(11))
	assert.Equal(t, -12, calc.Even(-11))
	assert.Equal(t, -12, calc.Even(-11.15631))
	assert.Equal(t, 8, calc.Even(6.2))

	//Odd
	assert.Equal(t, 17, calc.Odd(16))
	assert.Equal(t, -123, calc.Odd(-122))
	assert.Equal(t, -123, calc.Odd(-122.15631))
	assert.Equal(t, 125, calc.Odd(123.15631))

	//Int
	assert.Equal(t, 16, calc.Int(16))
	assert.Equal(t, -16, calc.Int(-16))
	assert.Equal(t, -123, calc.Int(-122.15631))
	assert.Equal(t, 123, calc.Int(123.15631))

	//Sqrt
	assert.Equal(t, 4.0, calc.Sqrt(16))
	assert.Equal(t, 11.097581268006106, calc.Sqrt(123.15631))

	//SqrtPi
	assert.Equal(t, 7.0898154036220635, calc.SqrtPi(16))
	assert.Equal(t, 19.669950654214343, calc.SqrtPi(123.15631))

	//Quotient
	assert.Equal(t, 3, calc.Quotient(16, 5))
	assert.Equal(t, -5, calc.Quotient(-26, 5))
	assert.Equal(t, -47, calc.Quotient(-122.15631, 2.546))
	assert.Equal(t, 2, calc.Quotient(123.15631, 48))

	//LN
	assert.Equal(t, 4.605170185988092, calc.LN(100))
	assert.Equal(t, 2.772588722239781, calc.LN(16))
	assert.Equal(t, 4.805301454167321, calc.LN(122.15631))

	//Log10
	assert.Equal(t, 4.0, calc.Log10(10000))
	assert.Equal(t, 2.0863598306747484, calc.Log10(122))
	assert.Equal(t, 2.086915905426539, calc.Log10(122.15631))
	assert.Equal(t, 0.0, calc.Log10(1))

	//Exp
	assert.Equal(t, 1.0, calc.Exp(0))
	assert.Equal(t, 1.0020020013340003, calc.Exp(0.002))

	//Trunc
	assert.Equal(t, 16.0, calc.Trunc(16))
	assert.Equal(t, -16.0, calc.Trunc(-16))
	assert.Equal(t, -122.0, calc.Trunc(-122.15631))
	assert.Equal(t, 123.0, calc.Trunc(123.15631))
}
