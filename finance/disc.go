package finance

import (
	"time"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/datetime"
	"github.com/maniartech/x/utils"
)

//actual/actual bias with leap year and non leap year not working properly
func Disc(Settlement, Maturity time.Time, Pr, Redemption interface{}, Basis ...interface{}) float64 {
	pr := utils.ToFloat64(Pr)
	redemption := utils.ToFloat64(Redemption)
	B := 360
	basis := 0
	if len(Basis) > 0 {
		basis = utils.ToInt(Basis[0])
		if basis == 0 || basis == 2 || basis == 4 {
			B = 360
		} else if basis == 1 {
			if datetime.IsLeapYear(Settlement) || datetime.IsLeapYear(Maturity) {
				B = 366
			} else {
				B = 365
			}
		} else if basis == 3 {
			B = 365
		} else {
			panic(core.ErrInvalidInput)
		}
	}
	if pr <= 0 || redemption <= 0 {
		panic(core.ErrInvalidInput)
	}

	if datetime.DateValue(Settlement) >= datetime.DateValue(Maturity) {
		panic(core.ErrInvalidInput)
	}
	DSM := 0
	if basis == 0 || basis == 4 {
		DSM = (datetime.Days360(Settlement, Maturity))
	} else {
		DSM = datetime.Days(Maturity, Settlement)
	}
	return calc.Divide((redemption-pr), redemption) * calc.Divide(B, DSM)
}
