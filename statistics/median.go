package statistics

import (
	"github.com/maniartech/x/utils"
)

// Average returns the average of the provided numbers.
func Median(v ...interface{}) float64 {
	c := utils.ForEach(func(_ int, x interface{}) {
	}, v...)
	var ret float64
	if c%2 == 0 {
		ret = (utils.ToFloat64(utils.ToInt(v[c/2])) + utils.ToFloat64(utils.ToInt(v[(c/2)-1]))) / 2
	} else {
		ret = utils.ToFloat64(v[c/2])
	}
	return ret
}
 