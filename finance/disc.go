package finance

import (
	"time"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/datetime"
	"github.com/maniartech/x/utils"
)

//Disc returns discount percentage for a security.
//
//Arguments
//
//Settlement: The settlement for the security.
//
//Maturity: The maturity date for the expiration of the security.
//
//Pr: Price of security per $100 face value.
//
//Redemption: Value of redemption of security per $100 face value.
//
//Basis: OPTIONAL, Day count basis used for calculations
//
//Examples:
//	date1 = 2018-Jul-1
//	date2 = 2048-Jan-1
//	Disc(date1, date2, 97.975, 100, 0) //Returns 0.0006864406779661065
//	Disc(date1, date2, 97.975, 100, 2) //Returns 0.0006765033407572485
func Disc(settlement, maturity time.Time, pr, redemption interface{}, basis ...interface{}) float64 {
	Pr := utils.ToFloat64(pr)
	Redemption := utils.ToFloat64(redemption)
	//Setting Default Basis
	B := 360
	Basis := 0
	//getting the value of basis if it was inputted
	if len(basis) > 0 {
		Basis = utils.ToInt(basis[0])
		if Basis == 0 || Basis == 2 || Basis == 4 {
			B = 360
		} else if Basis == 1 {
			if datetime.IsLeapYear(settlement) || datetime.IsLeapYear(maturity) {
				B = 366
			} else {
				B = 365
			}
		} else if Basis == 3 {
			B = 365
		} else {
			panic(core.ErrInvalidInput)
		}
	}
	if Pr <= 0 || Redemption <= 0 {
		panic(core.ErrInvalidInput)
	}

	if datetime.DateValue(settlement) >= datetime.DateValue(maturity) {
		panic(core.ErrInvalidInput)
	}
	DSM := 0
	if Basis == 0 || Basis == 4 {
		DSM = (datetime.Days360(settlement, maturity))
	} else {
		DSM = datetime.Days(maturity, settlement)
	}
	return calc.Divide((Redemption-Pr), Redemption) * calc.Divide(B, DSM)
}

//PriceDisc returns the price of a discounted security per $100 face value.
//
//Arguments
//
//Settlement: The settlement for the security.
//
//Maturity: The maturity date for the expiration of the security.
//
//Discount: Discount rate on the security.
//
//Redemption: Value of redemption of security per $100 face value.
//
//Basis: OPTIONAL, Day count basis used for calculations
//
//Examples:
//	date1 = 2018-Jul-1
//	date2 = 2048-Jan-1
//	Disc(date1, date2, 97.975, 100, 0) //Returns 0.0006864406779661065
//	Disc(date1, date2, 97.975, 100, 2) //Returns 0.0006765033407572485
func PriceDisc(settlement, maturity time.Time, discount, redemption interface{}, basis ...interface{}) float64 {
	Discount := utils.ToFloat64(discount)
	Redemption := utils.ToFloat64(redemption)
	//Setting Default Basis
	B := 360
	Basis := 0
	//getting the value of basis if it was inputted
	if len(basis) > 0 {
		Basis = utils.ToInt(basis[0])
		if Basis == 0 || Basis == 2 || Basis == 4 {
			B = 360
		} else if Basis == 1 {
			if datetime.IsLeapYear(settlement) || datetime.IsLeapYear(maturity) {
				B = 366
			} else {
				B = 365
			}
		} else if Basis == 3 {
			B = 365
		} else {
			panic(core.ErrInvalidInput)
		}
	}

	//Checking if Date of settlement is greater than or equal to Date of maturity
	//If it is then panicing
	if datetime.DateValue(settlement) >= datetime.DateValue(maturity) {
		panic(core.ErrInvalidInput)
	}
	if Discount <= 0 || Redemption <= 0 {
		panic(core.ErrInvalidInput)
	}
	DSM := 0
	if Basis == 0 || Basis == 4 {
		DSM = (datetime.Days360(settlement, maturity))
	} else {
		DSM = datetime.Days(maturity, settlement)
	}
	return Redemption - (Discount*Redemption)*calc.Divide(DSM, B)
}
