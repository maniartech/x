package calc

import (
	"strconv"

	"github.com/maniartech/go-funcs/currency"
)

func Base(input int64, base, min_len int) string {
	if base < 2 || base > 36 || min_len < 0 || min_len > 255 {
		panic(currency.ErrInvalidInput)
	}
	value := (strconv.FormatInt(input, base))
	if (min_len - len(value)) <= 0 {
		return value
	}
	for i := min_len - (min_len - len(value)); i < min_len; i++ {
		value = "0" + value
	}
	return value
}
