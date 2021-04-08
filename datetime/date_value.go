package datetime

import (
	"strings"
	"time"

	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func DateValue(date interface{}) int {
	var val time.Time = time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	sep := ""
	valS := date.(string)
	if strings.Contains(valS, "/") {
		sep = "/"
	} else if strings.Contains(valS, "-") {
		sep = "-"
	} else {
		panic(core.ErrInvalidInput)
	}
	inpSep := strings.Split(valS, sep)
	inpSepLen := len(inpSep)

	if inpSepLen != 2 && inpSepLen != 3 {
		panic(core.ErrInvalidInput)
	}
	isDigit := strings.IndexFunc(inpSep[1], utils.NotDigit) == -1
	var month time.Month
	if isDigit {
		month = getMonth(utils.ToInt(inpSep[1]))
	} else {
		monthS := inpSep[1]
		month = getMonth(monthS[:3])
	}
	if inpSepLen == 2 {
		val = time.Date(time.Now().Year(), month, utils.ToInt(inpSep[0]), 0, 0, 0, 0, time.UTC)
	} else {
		val = time.Date(utils.ToInt(inpSep[0]), month, utils.ToInt(inpSep[2]), 0, 0, 0, 0, time.UTC)
	}
	return int(val.Sub(DayZero).Hours()/24) + 2
}
func getMonth(input interface{}) time.Month {
	switch v := input.(type) {
	case string:
		for i := 0; i < 12; i++ {
			if MonthsInYear[i] == v {
				return time.Month(i + 1)
			}
		}
		panic(core.ErrInvalidInput)
	case int:
		return time.Month(input.(int))
	default:
		panic(core.ErrInvalidInput)
	}
}
