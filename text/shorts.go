package text

import (
	"fmt"
	"strings"
	"unicode/utf16"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Char(input interface{}) string { return string(rune(utils.ToInt(input))) }
func Code(input interface{}) int {
	return utils.ToInt(fmt.Sprintf("%d", int([]rune(utils.ToString(input))[0])))
}
func Len(input interface{}) int { return len(utf16.Encode([]rune(utils.ToString(input)))) }

// func LenB(input interface{}) int     { return }
func Upper(input interface{}) string { return strings.ToUpper(utils.ToString(input)) }
func Lower(input interface{}) string { return strings.ToLower(utils.ToString(input)) }

//TODO: Make the Proper func lower other character other than the first
func Proper(input interface{}) string { return strings.Title(utils.ToString(input)) }

//TODO: Make DBCS language count as 1 character
func Left(input, num interface{}) string {
	if utils.ToInt(num) >= len(utils.ToString(input)) {
		return utils.ToString(input)
	}
	return utils.ToUTF16String(utils.ToString(input)[:utils.ToInt(num)])
}

func Right(input, num interface{}) string {
	if utils.ToInt(num) >= len(utils.ToString(input)) {
		return utils.ToString(input)
	}
	return utils.ToUTF16String(utils.ToString(input)[Len(input)-utils.ToInt(num):])
}

func Mid(input, start, end interface{}) string {
	if utils.ToInt(start) <= 0 || utils.ToInt(end) >= Len(input) {
		panic(core.ErrInvalidInput)
	}
	return utils.ToUTF16String(utils.ToString(input)[utils.ToInt(start):utils.ToInt(end)])
}

func Fixed(input, decimals interface{}) string {
	return utils.ToString(calc.Round(input, decimals))
}

func Dollar(input, decimals interface{}) string {
	return "$" + utils.ToString(calc.Round(input, decimals))
}

func Substitute(input, old, new interface{}) string {
	return strings.ReplaceAll(utils.ToString(input), utils.ToString(old), utils.ToString(new))
}

func Search(find, within, sPos interface{}) int {
	pos := (strings.Index(utils.ToUTF16String(utils.ToString(within)[utils.ToInt(sPos):]), utils.ToUTF16String(utils.ToString(find))))
	if pos != -1 {
		return pos + 1 + utils.ToInt(sPos)
	}
	return pos
}

func Replace(input, start, end, new interface{}) string {
	encodedS := utils.ToUTF16String(utils.ToString(input))
	if utils.ToInt(start)+utils.ToInt(end) > Len(input) {
		return strings.Replace(encodedS, encodedS[utils.ToInt(start)-1:], utils.ToString(new), 1)
	}
	return strings.Replace(encodedS, encodedS[utils.ToInt(start)-1:(utils.ToInt(end)+utils.ToInt(end))], utils.ToString(new), 1)
}

func Rept(input, mul interface{}) string {
	return strings.Repeat(utils.ToString(input), utils.ToInt(mul))
}
func Exact(input, input2 interface{}) string {
	return fmt.Sprintf("%v", utils.ToString(input) == utils.ToString(input2))
}
