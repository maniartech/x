package finance

import (
	"math"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func DB(Cost, Salvage, Life, Period interface{}, Month ...interface{}) float64 {
	var month float64 = 12.0
	cost := utils.ToFloat64(Cost)
	salvage := utils.ToFloat64(Salvage)
	life := utils.ToFloat64(Life)
	period := utils.ToFloat64(Period)

	if len(Month) > 0 {
		month = utils.ToFloat64(utils.ToInt(Month[0]))
		if !utils.IsValidMonth(month) {
			panic(core.ErrInvalidInput)
		}
	}

	if cost < 0 || salvage < 0 || life < 0 || period < 0 {
		panic(core.ErrInvalidInput)
	}

	if period > life {
		panic(core.ErrInvalidInput)
	}

	if salvage >= cost {
		return 0
	}
	//Rounding to 3 decimals points
	var rate float64 = calc.Round(1-math.Pow(salvage/cost, 1.0/life), 3)

	//Calculating depriciation for the first month
	var inital float64 = cost * rate * month / 12

	total := inital
	var current float64 = 0
	var ceiling float64
	if period == life {
		ceiling = life - 1
	} else {
		ceiling = period
	}
	for i := 2; i <= int(ceiling); i++ {
		current = (cost - total) * rate
		total += current
	}

	if period == 1 {
		return inital
		//last Period
	} else if period == life {
		return (cost - total) * rate
	} else {
		return current
	}
}

func DDB(Cost, Salvage, Life, Period interface{}, Factor ...interface{}) float64 {
	var factor float64 = 2.0
	cost := utils.ToFloat64(Cost)
	salvage := utils.ToFloat64(Salvage)
	life := utils.ToFloat64(Life)
	period := utils.ToFloat64(Period)

	//Checking if there is value in Factor if yes then set the value in factor
	if len(Factor) > 0 {
		factor = utils.ToFloat64(Factor[0])
		if factor < 0 || factor > 2 {
			panic(core.ErrInvalidInput)
		}
	}

	//Checking for invalid Inputs
	if cost < 0 || salvage < 0 || life < 0 || period < 0 || factor <= 0 {
		panic(core.ErrInvalidInput)
	}

	//Checking if period is greater than life
	if period > life {
		panic(core.ErrInvalidInput)
	}

	//Checking if salvage is greater than equal to cost if yes then return 0
	if salvage >= cost {
		return 0
	}

	//Calculating depreciation
	total := 0.0
	current := 0.0
	for i := 1.0; i <= period; i++ {
		current = math.Min((cost-total)*(factor/life), (cost - salvage - total))
		total += current
	}
	return current
}
