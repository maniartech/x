package text

import "github.com/maniartech/x/utils"

// Concat joins all the input into one string
func Concat(v ...interface{}) string {
	var fin string
	utils.ForEach(func(_ int, x interface{}) {
		fin += utils.ToString(x)
	}, v...)
	return fin
}
