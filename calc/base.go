package calc

import (
	"strconv"
	"strings"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Base(number, b, minL interface{}) string {
	input := utils.ToInt(number)
	base := utils.ToInt(b)
	min_len := utils.ToInt(minL)
	if base < 2 || base > 36 || min_len < 0 || min_len > 255 {
		panic(core.ErrInvalidInput)
	}
	value := (strconv.FormatInt(int64(input), base))
	if (min_len - len(value)) <= 0 {
		return strings.ToUpper(value)
	}
	for i := min_len - (min_len - len(value)); i < min_len; i++ {
		value = "0" + value
	}
	return strings.ToUpper(value)
}
