package finance

import (
	"fmt"
	"time"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/datetime"
	"github.com/maniartech/x/utils"
)

//WIP
func Disc(Settlement, Maturity time.Time, Pr, Redemption interface{}, Basis ...interface{}) float64 {
	pr := utils.ToFloat64(Pr)
	redemption := utils.ToFloat64(Redemption)
	B := 360
	if len(Basis) > 0 {
		basis := utils.ToInt(Basis[0])
		if basis == 0 || basis == 2 || basis == 4 {
			B = 360
		} else if basis == 1 {
			if datetime.IsLeapYear(Settlement) && datetime.IsLeapYear(Maturity) {
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
	DSM := datetime.Days(Maturity, Settlement)
	fmt.Println(calc.Divide((redemption - pr), redemption))
	fmt.Println("DSM: ", DSM)
	return calc.Divide(redemption-pr, redemption) * calc.Divide(B, DSM)
}
