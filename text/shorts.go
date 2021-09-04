package text

import (
	"fmt"
	"strings"
	"unicode/utf16"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Char(input interface{}) string {
	//Converting integer to it's corresponding ascii character
	return string(rune(utils.ToInt(input)))
}
func Code(input interface{}) int {
	//Converting the first character of a string to its corresponding ascii character to it's ascii number
	return utils.ToInt(fmt.Sprintf("%d", int([]rune(utils.ToString(input))[0])))
}

func Len(input interface{}) int {
	//Finding the length of string using utf16 formatting
	return len(utf16.Encode([]rune(utils.ToString(input))))
}

func Upper(input interface{}) string {
	//Converting the string to UPPER case
	return strings.ToUpper(utils.ToString(input))
}
func Lower(input interface{}) string {
	//Converting the string to lower CASE
	return strings.ToLower(utils.ToString(input))
}

//TODO: Make the Proper func lower other character other than the first
func Proper(input interface{}) string {
	//Converting string to camel case
	return strings.Title(utils.ToString(input))
}

//TODO: Make DBCS language count as 1 character
func Left(input, num interface{}) string {
	//if num of character is greater than or equal to the length of string the returning the whole string
	if utils.ToInt(num) >= len(utils.ToString(input)) {
		return utils.ToString(input)
	}
	return utils.ToUTF16String(utils.ToString(input)[:utils.ToInt(num)])
}

func Right(input, num interface{}) string {
	//if num of character is greater than or equal to the length of string the returning the whole string
	if utils.ToInt(num) >= len(utils.ToString(input)) {
		return utils.ToString(input)
	}
	return utils.ToUTF16String(utils.ToString(input)[Len(input)-utils.ToInt(num):])
}

func Mid(input, start, end interface{}) string {
	//if num of character is greater than or equal to the length of string the returning the whole string
	if utils.ToInt(start) <= 0 || utils.ToInt(end) >= Len(input) {
		panic(core.ErrInvalidInput)
	}
	return utils.ToUTF16String(utils.ToString(input)[utils.ToInt(start):utils.ToInt(end)])
}

func Fixed(input, decimals interface{}) string {
	//Rounding number to nearest decimals position and then returning it to string
	return utils.ToString(calc.Round(input, decimals))
}

func Dollar(input, decimals interface{}) string {
	//Rounding number to nearest decimals position and then returning it to string and then adding a $ sign at the start
	return "$" + utils.ToString(calc.Round(input, decimals))
}

func Substitute(input, old, new interface{}) string {
	return strings.ReplaceAll(utils.ToString(input), utils.ToString(old), utils.ToString(new))
}

func Search(find, within, sPos interface{}) int {
	pos := (strings.Index(utils.ToUTF16String(utils.ToString(within)[utils.ToInt(sPos):]), utils.ToUTF16String(utils.ToString(find))))
	//if pos is returned as -1 then the "find" is not "within" string.
	if pos != -1 {
		return pos + 1 + utils.ToInt(sPos)
	}
	return pos
}

func Replace(input, start, end, new interface{}) string {
	//Encoding utf-8 to a utf-16 string
	encodedS := utils.ToUTF16String(utils.ToString(input))
	if utils.ToInt(start)+utils.ToInt(end) > Len(input) {
		return strings.Replace(encodedS, encodedS[utils.ToInt(start)-1:], utils.ToString(new), 1)
	}
	return strings.Replace(encodedS, encodedS[utils.ToInt(start)-1:(utils.ToInt(end)+utils.ToInt(end))], utils.ToString(new), 1)
}

func Rept(input, mul interface{}) string {
	//repeating the string mul number of times
	return strings.Repeat(utils.ToString(input), utils.ToInt(mul))
}
func Exact(input, input2 interface{}) string {
	//Checking if both strings are exactly the same Case sensitive
	return fmt.Sprintf("%v", utils.ToString(input) == utils.ToString(input2))
}
