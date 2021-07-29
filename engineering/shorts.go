package engineering

import (
	"math"
	"strconv"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/utils"
)

func Bin2Dec(x interface{}) int64 {
	val, _ := strconv.ParseInt(utils.ToString(utils.ToInt(x)), 2, 64)
	return val
}

func Oct2Dec(x interface{}) int64 {
	val, _ := strconv.ParseInt(utils.ToString(utils.ToInt(x)), 8, 64)
	return val
}

func Hex2Dec(x interface{}) int64 {
	val, _ := strconv.ParseInt(utils.ToString(utils.ToInt(x)), 16, 64)
	return val
}

func Dec2Bin(x interface{}) string { return calc.Base(Bin2Dec(x), 2, 0) }
func Dec2Oct(x interface{}) string { return calc.Base(Bin2Dec(x), 8, 0) }
func Dec2Hex(x interface{}) string { return calc.Base(Bin2Dec(x), 16, 0) }

func Bin2Oct(x interface{}) string { return Dec2Oct(Bin2Dec(x)) }
func Bin2Hex(x interface{}) string { return Dec2Hex(Bin2Dec(x)) }

func Oct2Bin(x interface{}) string { return Dec2Bin(Oct2Dec(x)) }
func Oct2Hex(x interface{}) string { return Dec2Hex(Oct2Dec(x)) }

func Delta(x, y interface{}) int { return utils.ToInt(utils.ToFloat64(x) == utils.ToFloat64(y)) }
func Erf(x interface{}) float64  { return math.Erf(utils.ToFloat64(x)) }
func Erfc(x interface{}) float64 { return math.Erfc(utils.ToFloat64(x)) }
func GetStep(x interface{}, y interface{}) int {
	return utils.ToInt(math.Abs(utils.ToFloat64(x)) <= math.Abs(utils.ToFloat64(y)))
}
