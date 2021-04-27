package string

import (
	"fmt"
	"strings"

	"github.com/maniartech/x/utils"
)

func Char(input rune) string             { return fmt.Sprintf("%x", input) }
func Code(input interface{}) string      { return fmt.Sprintf("%x", utils.ToString(input)) }
func Left(input, num interface{}) string { return utils.ToString(input)[:utils.ToInt(num)] }
func Len(input interface{}) int          { return len(utils.ToString(input)) }
func Upper(input interface{}) string     { return strings.ToUpper(utils.ToString(input)) }
func Lower(input interface{}) string     { return strings.ToLower(utils.ToString(input)) }
func Substitute(s, old, new, n interface{}) string {
	return strings.Replace(utils.ToString(s), utils.ToString(old), utils.ToString(new), utils.ToInt(n))
}
func Rept(input, mul interface{}) string {
	return strings.Repeat(utils.ToString(input), utils.ToInt(mul))
}
func Exact(input, input2 interface{}) string {
	return fmt.Sprintf("%v", utils.ToString(input) == utils.ToString(input2))
}
