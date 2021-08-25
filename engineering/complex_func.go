package engineering

import (
	"fmt"
	"math/cmplx"
	"strings"

	"github.com/maniartech/x/utils"
)

func Complex(realNum, imaginaryNum interface{}) complex128 {
	return complex(utils.ToFloat64(realNum), utils.ToFloat64(imaginaryNum))
}

//IMCos returns the cosine of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the cosine of.
//
//Examples
//	IMCos(1+1i) //returns 0.8337300251311491-0.9888977057628651i
//	IMCos(12-3i) //returns  8.49563643031773-5.375320381963727i
func IMCos(x complex128) complex128 { return cmplx.Cos(x) }

//IMCosh returns the hyperbolic cosine of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the hyperbolic cosine of.
//
//Examples
//	IMCosh(1+1i) //returns 0.8337300251311491+0.9888977057628651i
//	IMCosh(12-3i) //returns  -80563.0111483336-11483.978737982383i
func IMCosh(x complex128) complex128 { return cmplx.Cosh(x) }

//IMCot returns the cotangent of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the cotangent of.
//
//Examples
//	IMCot(1+1i) //returns 0.21762156185440268-0.868014142895925i
//	IMCot(12-3i) //returns -0.00449884150567166+1.0020949734340645i
func IMCot(x complex128) complex128 { return cmplx.Cot(x) }

//IMSin returns the sine of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the sine of.
//
//Examples
//	IMSin(1+1i) //returns 1.2984575814159773+0.6349639147847361i
//	IMSin(12-3i) //returns -5.402034774516558-8.453623415581824i
func IMSin(x complex128) complex128 { return cmplx.Sin(x) }

//IMSinh returns the hyperbolic sine of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the hyperbolic sine of.
//
//Examples
//	IMSinh(1+1i) //returns 0.6349639147847361+1.2984575814159773i
//	IMSinh(12-3i) //returns -80563.01114225085-11483.978738849457i
func IMSinh(x complex128) complex128 { return cmplx.Sinh(x) }

//IMTan returns the tangent of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the tangent of.
//
//Examples
//	IMTan(1+1i) //returns 0.2717525853195117+1.0839233273386948i
//	IMTan(12-3i) //returns -0.00447996037507582-0.9978892938076563i
func IMTan(x complex128) complex128 { return cmplx.Tan(x) }

//IMCSC returns the cosecant of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the cosecant of.
//
//Examples
//	IMCSC(1+1i) //returns 0.6215180171704284-0.30393100162842646i
//	IMCSC(12-3i) //returns -0.05367376094284146+0.08399386180355566i
func IMCSC(x complex128) complex128 { return 1 / cmplx.Sin(x) }

//IMLN returns the Natural Logarithm of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the Natural Logarithm of.
//
//Examples
//	IMLN(1+1i) //returns 0.3465735902799727+0.7853981633974483i
//	IMLN(12-3i) //returns 2.5152189606962176-0.24497866312686414i
func IMLN(x complex128) complex128 { return cmplx.Log(x) }

//IMLog10 returns the Logarithm 10 of the complex number.
//
//Arguments
//
//x: Complex Number of which you want the Logarithm 10 of.
//
//Examples
//	IMLog10(1+1i) //returns 0.3465735902799727+0.7853981633974483i
//	IMLog10(12-3i) //returns 2.5152189606962176-0.24497866312686414i
func IMLog10(x complex128) complex128 { return cmplx.Log10(x) }

//IMSum returns summation of complex numbers.
//
//Arguments
//
//x: List of Complex numbers.
//
//Examples
//	IMSum((1+1i), (12-3i), (-3.45-5.3i)) //returns 10-7.3i
func IMSum(x ...complex128) complex128 {
	var sum complex128 = x[0]
	for i := 1; i < len(x); i++ {
		sum *= x[i]
	}
	return sum
}

//IMSub returns summation of two complex numbers (x + y).
//
//Arguments
//
//x: First Number to be subtracted.
//
//y: Second Number to be subtracting.
//
//Examples
//	IMSub((12-3i), (1+1i) ) //returns 11-2i
func IMSub(x complex128, y complex128) complex128 { return x - y }

//IMSum returns summation of complex numbers
//
//Arguments
//
//x: List of Complex numbers.
//
//Examples
//	IMSum((1+1i), (12-3i), (-3.45-5.3i)) //returns 10-7.3i
func IMProduct(x ...complex128) complex128 {
	var product complex128 = x[0]
	for i := 1; i < len(x); i++ {
		product *= x[i]
	}
	return product
}

//IMDiv returns divison of two complex numbers (x / y).
//
//Arguments
//
//x: Numerator
//
//y: Denominator
//
//Examples
//	IMDiv((1+1i), (12-3i)) //returns 0.058823529411764705+0.09803921568627451i
func IMDiv(x complex128, y complex128) complex128 { return x / y }

func complexSplit(input complex128) (float64, float64) {
	s := fmt.Sprintf("%v", input)
	ss := make([]string, 2)
	var x float64
	var y float64
	if strings.Index(s, "+") != -1 {
		ss = strings.Split(s, "+")
		y = utils.ToFloat64(ss[1][:len(ss[1])-2])
	} else {
		ss = strings.Split(s, "-")
		y = utils.ToFloat64(ss[1][:len(ss[1])-2]) * -1
	}
	x = utils.ToFloat64(ss[0][1:])
	return x, y
}

//Imaginary returns the imaginary coefficent of the complex number.
//
//Arguments
//
//input: The complex number you want to find the imaginary number.
//
//Examples
//	Imaginary((1+1i)) //returns 1i
//	Imaginary((135+5.153i)) //returns 5.153i
func Imaginary(input complex128) float64 {
	_, y := complexSplit(input)
	return y
}

//Real returns the Real Coefficient of the complex number.
//
//Arguments
//
//input: The complex number you want to find the Real Coefficient.
//
//Examples
//	Real((1+1i)) //returns 1
//	Real((135+5.153i)) //returns 135
func IMReal(input complex128) float64 {
	x, _ := complexSplit(input)
	return x
}
