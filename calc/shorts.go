package calc

import "math"

// Trignometric Functions

func Acos(input float64) float64  { return math.Acos(input) }
func Acosh(input float64) float64 { return math.Acosh(input) }
func Asin(input float64) float64  { return math.Asin(input) }
func Asinh(input float64) float64 { return math.Asinh(input) }
func Atan(input float64) float64  { return math.Atan(input) }
func Atanh(input float64) float64 { return math.Atanh(input) }
func Cos(input float64) float64   { return math.Cos(input) }
func Cosh(input float64) float64  { return math.Cosh(input) }
func Sin(input float64) float64   { return math.Sin(input) }
func Sinh(input float64) float64  { return math.Sinh(input) }
func Tan(input float64) float64   { return math.Tan(input) }
func Tanh(input float64) float64  { return math.Tanh(input) }
func CSC(input float64) float64   { return 1 / math.Sin(input) }

func Deg2Rad(input float64) float64 { return input * (math.Pi / 180) }
func Rad2Deg(input float64) float64 { return input * (180 / math.Pi) }

// Math Functions
func Abs(input float64) float64    { return math.Abs(input) }
func Even(input float64) int       { return int((math.Floor(input/2) + 1) * 2) }
func Odd(input float64) int        { return int(((math.Floor(input/2) - 1) * 2) + 3) }
func INT(input float64) int        { return int(math.Floor(input)) }
func Sqrt(input float64) float64   { return math.Sqrt(input) }
func SqrtPi(input float64) float64 { return input * math.Pi }
func Quotient(n, d float64) int    { return int(math.Floor(n / d)) }
func LN(input float64) float64     { return math.Log(input) }
func Log10(input float64) float64  { return math.Log10(input) }
func Exp(input float64) float64    { return math.Exp(input) }
func Trunc(input float64) float64  { return math.Trunc(input) }
