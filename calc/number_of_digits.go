package calc

import (
	"strings"

	"github.com/maniartech/x/utils"
)

func NumberOfDigits(Number interface{}) int {
	number := utils.ToInt(Number)
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	return count
}
func NumberOfDecimalDigits(input interface{}) int {
	s := utils.ToString(utils.ToFloat64(input))
	i := strings.IndexByte(s, '.')
	if i > -1 {
		return len(s) - i - 1
	}
	return 0
}
