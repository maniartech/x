package calc

import "math"

func Floor(input float64, factor float64) float64 {
	if input == factor {
		return input
	}
	return ((math.Floor(input / factor)) * factor)
}
