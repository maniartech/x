package calc

import "github.com/maniartech/x/utils"

//LCM func finds the Least common Multiple using the calcGCD function
func LCM(v ...interface{}) int {
	// Initialize result
	var result int
	utils.ForEach(func(i int, x interface{}) {
		ival := utils.ToInt(x)
		if ival == 0 {
			result = 0
			return
		}
		if i == 0 {
			result = ival
			return
		}
		result = (ival * result) / (calcGCD(ival, result))
	}, v...)
	return result
}
