package calc

import (
	"strconv"
	"strings"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

// Base function returns the string representation of the provided number
// converted into a different base.
//
// Arguments
//
// number: The number that needs to be converted into another representation.
//
// radix: The base of the representation that the number should be converted to.
//
// minL: Optional, if provided, guarantees the minimum length of the resulting number
// by providing padding with leading zeros.
//
// Remark
//
// If the value of radix is outside the range of 2 to 36, it returns an Invalid Input Error.
// If the value of min_len is outside the range of 0 to 255 it returns an Invalid Input Error.
// Any non-integer value entered as input is truncated to an integer.
//
// Example
//
//    calc.Base(15, 16) // Returns F
//	  calc.Base(10, 2)  // Returns 1010
//    calc.Base(-23, 1, 0)  // Returns Error as radix is outside the range, that is less than 2.
//    calc.Base(-23, 37, 0) // Returns Error as radix is outside the range, that is greater than 37.
//    calc.Base(-23, 2, -34) // Returns Error as minL is outside the range, that is less than 0.
//    calc.Base(-23, 2, 256) // Returns Error as minL is outside the range, that is greater than 255.
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
