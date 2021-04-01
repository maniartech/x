package statistics

import (
	"fmt"
	"strconv"
)

func AverageA(v ...interface{}) float64 {

	var sum float64

	for _, n := range v {
		switch tval := n.(type) {
		case string:
			continue
		case bool:
			if tval {
				sum += 1
			}
		default:
			if fval, err := strconv.ParseFloat(fmt.Sprintf("%v", tval), 63); err == nil {
				sum += fval
			}
		}
	}
	return sum / float64(len(v))
}
