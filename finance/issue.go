package finance

import (
	"time"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/datetime"
	"github.com/maniartech/x/utils"
)

func PriceMat(Settlement, Maturity, Issue time.Time, Rate, Yeild interface{}, Basis ...interface{}) float64 {
	rate := utils.ToFloat64(Rate)
	yeild := utils.ToFloat64(Yeild)
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

	//Checking if Date of settlement is greater than or equal to Date of maturity
	//If it is then panicing
	if datetime.DateValue(Settlement) >= datetime.DateValue(Maturity) {
		panic(core.ErrInvalidInput)
	}
	if rate < 0 || yeild <= 0 {
		panic(core.ErrInvalidInput)
	}
	DSM := 0
	DIM := 0
	A := 0
	//Calculating values depending on different basis
	if basis == 0 || basis == 4 {
		A = datetime.Days360(Issue, Settlement)
		DIM = datetime.Days360(Issue, Maturity)
		DSM = datetime.Days360(Settlement, Maturity)
	} else {
		A = datetime.Days(Issue, Settlement)
		DIM = datetime.Days(Issue, Maturity)
		DSM = datetime.Days(Settlement, Maturity)
	}
	//Calculating Numerator and Denominator
	n := 100 + (calc.Divide(DIM, B) * rate * 100)
	d := 1 + (calc.Divide(DSM, B) * yeild)
	return (n / d) - (calc.Divide(A, B) * rate * 100)
}
