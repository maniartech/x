package calc

import (
	"strconv"
	"strings"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Base function returns the string representation of the provided
// number converted into different base.
//
// Syntax
//
//    Base(number, redix, minL interface{}) string
//
// Arguments
//
// number : The number that needs to be converted into another representation
//
// redix  : The base of the reprentation that the number should be converted to
//
// minL   : Optional, if provided, guearantees the minimum length of the resulting number
// by providing padding with leading zeros.
//
// Example
//
//    Base(15, 16) // Returns F
//    Base(10, 2)  // Returns 1010
//
func Base(number, redix interface{}, minL ...interface{}) string {
	input := utils.ToInt(number)
	base := utils.ToInt(redix)
	min_len := 0
	if len(minL) > 0 {
		min_len = utils.ToInt(minL[0])
	}

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
