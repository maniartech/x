package text

import "github.com/maniartech/x/utils"

// Concat joins all the input into one string
func Concat(v ...interface{}) string {
	var fin string
	//Concating all the values in the v...
	utils.ForEach(func(_ int, x interface{}) {
		fin += utils.ToString(x)
	}, v...)
	return fin
}
