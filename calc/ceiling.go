package calc

import "math"

func Ceiling(input float64, factor float64) float64 {
	if input == factor {
		return input
	}
	return ((math.Floor(input/factor) + 1) * factor)
}
