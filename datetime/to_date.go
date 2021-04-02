package datetime

import (
	"strings"
	"time"

	"github.com/maniartech/x/currency"
	"github.com/maniartech/x/utils"
)

func Date(date interface{}) time.Time {
	switch date.(type) {
	case string:
		sep := ""
		valS := date.(string)
		if strings.Contains(valS, "/") {
			sep = "/"
		} else if strings.Contains(valS, "-") {
			sep = "-"
		} else {
			panic(currency.ErrInvalidInput)
		}
		inpSep := strings.Split(valS, sep)
		inpSepLen := len(inpSep)
		if inpSepLen != 2 && inpSepLen != 3 {
			panic(currency.ErrInvalidInput)
		}
		if inpSepLen == 2 {
			return time.Date(time.Now().Year(), time.Month(utils.ToInt(inpSep[1])), utils.ToInt(inpSep[0]), 0, 0, 0, 0, time.UTC)
		} else {
			if utils.ToInt(inpSep[0]) > 31 {
				return time.Date(utils.ToInt(inpSep[0]), time.Month(utils.ToInt(inpSep[1])), utils.ToInt(inpSep[2]), 0, 0, 0, 0, time.UTC)
			} else {
				return time.Date(utils.ToInt(inpSep[2]), time.Month(utils.ToInt(inpSep[1])), utils.ToInt(inpSep[0]), 0, 0, 0, 0, time.UTC)
			}
		}

	case int:
		return DayZero.AddDate(0, 0, utils.ToInt(date)-2)
	case time.Time:
		return date.(time.Time)
	default:
		panic(currency.ErrInvalidInput)
	}
}
