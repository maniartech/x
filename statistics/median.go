package statistics

import (
	"github.com/maniartech/x/utils"
)

//TODO: Add sorting of numbers then only the function is complete
func Median(v ...interface{}) float64 {
	var ret float64
	//Counting the number of items in v...
	c := utils.ForEach(func(_ int, x interface{}) {
	}, v...)
	//Checking if total number of items are even number or odd number
	if c%2 == 0 {
		//if the total is an even number than applying this formula
		ret = (utils.ToFloat64(utils.ToInt(v[c/2])) + utils.ToFloat64(utils.ToInt(v[(c/2)-1]))) / 2
	} else {
		ret = utils.ToFloat64(v[c/2])
	}
	return ret
}
