package calc

import (
	"math"

	"github.com/maniartech/x/utils"
)

// Trignometric Functions

func Acos(input float64) float64      { return math.Acos(input) }
func Acosh(input interface{}) float64 { return math.Acosh(utils.ToFloat64(input)) }
func Asin(input interface{}) float64  { return math.Asin(utils.ToFloat64(input)) }
func Asinh(input interface{}) float64 { return math.Asinh(utils.ToFloat64(input)) }
func Atan(input interface{}) float64  { return math.Atan(utils.ToFloat64(input)) }
func Atanh(input interface{}) float64 { return math.Atanh(utils.ToFloat64(input)) }
func Cos(input interface{}) float64   { return math.Cos(utils.ToFloat64(input)) }
func Cosh(input interface{}) float64  { return math.Cosh(utils.ToFloat64(input)) }
func Sin(input interface{}) float64   { return math.Sin(utils.ToFloat64(input)) }
func Sinh(input interface{}) float64  { return math.Sinh(utils.ToFloat64(input)) }
func Tan(input interface{}) float64   { return math.Tan(utils.ToFloat64(input)) }
func Tanh(input interface{}) float64  { return math.Tanh(utils.ToFloat64(input)) }
func CSC(input interface{}) float64   { return 1 / math.Sin(utils.ToFloat64(input)) }

func Deg2Rad(input interface{}) float64 { return utils.ToFloat64(input) * (math.Pi / 180) }
func Rad2Deg(input interface{}) float64 { return utils.ToFloat64(input) * (180 / math.Pi) }

// Math Functions
func Abs(input interface{}) float64    { return math.Abs(utils.ToFloat64(input)) }
func Even(input interface{}) int       { return int((math.Floor(utils.ToFloat64(input)/2) + 1) * 2) }
func Odd(input interface{}) int        { return int(((math.Floor(utils.ToFloat64(input)/2) - 1) * 2) + 3) }
func INT(input interface{}) int        { return int(math.Floor(utils.ToFloat64(input))) }
func Sqrt(input interface{}) float64   { return math.Sqrt(utils.ToFloat64(input)) }
func SqrtPi(input interface{}) float64 { return utils.ToFloat64(input) * math.Pi }
func Quotient(n, d interface{}) int    { return int(math.Floor(utils.ToFloat64(n) / utils.ToFloat64(d))) }
func LN(input interface{}) float64     { return math.Log(utils.ToFloat64(input)) }
func Log10(input interface{}) float64  { return math.Log10(utils.ToFloat64(input)) }
func Exp(input interface{}) float64    { return math.Exp(utils.ToFloat64(input)) }
func Trunc(input interface{}) float64  { return math.Trunc(utils.ToFloat64(input)) }
