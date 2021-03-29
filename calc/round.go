package calc

import "math"

func Round(input float64, factor int) float64 {
	var value float64
	if factor == 0 {
		return math.Round(input)
	}
	value = (math.Round(input * math.Pow10(factor))) * math.Pow10(-factor)
	return value
}
