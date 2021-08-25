package engineering

import (
	"math"

	"github.com/maniartech/x/utils"
)

//BesselJ returns the Bessel function.
//
//Arguments
//
//x: The number required for the computation of this function
//
//n: The order of the Bassel function. n will be truncated to an Integer if it is not an Integer.
//
//Examples
//	BesselJ (1.9, 2) //returns 0.3299257276923874
// BesselJ (-25, 6) //returns  -0.15870034085651266
func BesselJ(x, n interface{}) float64 { return math.Jn(utils.ToInt(n), utils.ToFloat64(x)) }

//BesselY returns the Bessel function. Also known as Weber function or Neumann function.
//
//Arguments
//
//x: The number required for the computation of this function
//
//n: The order of the Bassel function. n will be truncated to an Integer if it is not an Integer.
//
//Remark
//
//The order of the Bessel function 'n' should always be greater than > 0 or else it will return NaN
//
//Examples
//	BesselY (2.5, 1) //returns 0.1459181379667858
//	BesselY(67, 12) //returns  -0.09511730990903257
//	BesselY(-25, 12) //returns NaN
func BesselY(x, n interface{}) float64 { return math.Yn(utils.ToInt(n), utils.ToFloat64(x)) }
