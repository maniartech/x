package engineering

import (
	"math/cmplx"

	"github.com/maniartech/x/utils"
)

func Complex(realNum, imaginaryNum interface{}) complex128 {
	return complex(utils.ToFloat64(realNum), utils.ToFloat64(imaginaryNum))
}

func IMCos(x complex128) complex128   { return cmplx.Cos(x) }
func IMCosh(x complex128) complex128  { return cmplx.Cosh(x) }
func IMCot(x complex128) complex128   { return cmplx.Cot(x) }
func IMSin(x complex128) complex128   { return cmplx.Sin(x) }
func IMSinh(x complex128) complex128  { return cmplx.Sinh(x) }
func IMTan(x complex128) complex128   { return cmplx.Tan(x) }
func IMCSC(x complex128) complex128   { return 1 / cmplx.Sin(x) }
func IMSec(x complex128) complex128   { return 1 / cmplx.Cos(x) }
func IMSech(x complex128) complex128  { return 1 / cmplx.Cosh(x) }
func IM(x complex128) complex128      { return cmplx.Cosh(x) }
func IMLN(x complex128) complex128    { return cmplx.Log(x) }
func IMLog10(x complex128) complex128 { return cmplx.Log10(x) }

// func IMPower(x complex128, y complex128) complex128 { return cmplx.Pow(x, y) }
func IMDiv(x complex128, y complex128) complex128 { return x / y }
func IMSum(x complex128, y complex128) complex128 { return x + y }
func IMSub(x complex128, y complex128) complex128 { return x - y }
func IMProduct(x ...complex128) complex128 {
	var product complex128 = x[0]
	for i := 1; i < len(x); i++ {
		product *= x[i]
	}
	return product
}
