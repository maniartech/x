package engineering

import (
	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/text"
	"github.com/maniartech/x/utils"
)

func Convert(number, fromUnit, toUnit interface{}) float64 {
	fr := utils.ToString(fromUnit)
	to := utils.ToString(toUnit)
	weight := []string{"g", "sg", "lbm", "u", "ozm", "grain", "cwt", "uk_cwt", "stone", "ton", "uk_ton"}
	distance := []string{"m", "mi", "nmi", "in", "ft", "yd", "ang", "ell", "ly", "parsec", "Picapt", "pica", "survey_mi"}
	time := []string{"s", "min", "hr", "day", "yr"}
	speed := []string{"m/s", "mph", "m/h", "kn", "admkn"}
	force := []string{"N", "dyn", "lbf", "pond"}
	pressure := []string{"p", "atm", "mmHg", "psi", "Torr"}
	if text.Contains(weight, fr) {
		return ConvertMass(number, fr, to)
	} else if text.Contains(distance, fr) {
		return ConvertDistance(number, fr, to)
	} else if text.Contains(time, fr) {
		ConvertTime(number, fr, to)
	} else if text.Contains(speed, fr) {
		ConvertSpeed(number, fr, to)
	} else if text.Contains(force, fr) {
		return ConvertForce(number, fr, to)
	} else if text.Contains(pressure, fr) {
		return ConvertPressure(number, fr, to)
	}

	return 0.0
}

func ConvertMass(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	var gr float64
	var ans float64
	if fr == to {
		return num
	}
	switch fr {
	case "g":
		gr = num
	case "sg":
		gr = num * 14593.90294
	case "lbm":
		gr = num * 453.59237
	case "u":
		gr = num * 1.660540199e-24
	case "ozm":
		gr = num * 28.34952313
	case "grain":
		gr = num * 0.06479891
	case "cwt":
		gr = num * 45359.237
	case "uk_cwt":
		gr = num * 50802.34544
	case "stone":
		gr = num * 6350.29318
	case "ton":
		gr = num * 907184.74
	case "uk_ton":
		gr = num * 1016046.909
	}
	switch to {
	case "g":
		ans = gr
	case "sg":
		ans = gr * 0.00006852176586
	case "lbm":
		ans = gr * 0.002204622622
	case "u":
		ans = gr * 6.02214e+23
	case "ozm":
		ans = gr * 0.03527396195
	case "grain":
		ans = gr * 15.43235835
	case "cwt":
		ans = gr * 0.00002204622622
	case "uk_cwt":
		ans = gr * 0.00001968413055
	case "stone":
		ans = gr * 0.0001574730444
	case "ton":
		ans = gr * 0.000001102311311
	case "uk_ton":
		ans = gr * 0.0000009842065276
	default:
		panic(core.ErrInvalidInput)
	}
	return calc.Round(ans, 7)
}

func ConvertDistance(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	var m float64
	var ans float64
	if fr == to {
		return num
	}
	switch fr {
	case "m":
		m = num
	case "mi":
		m = num * 1609.344
	case "nmi":
		m = num * 1852
	case "in":
		m = num * 0.0254
	case "ft":
		m = num * 0.3048
	case "yd":
		m = num * 0.9144
	case "ell":
		m = num * 1.143
	case "ly":
		m = num * 9.46073e+15
	case "parsec":
		m = num * 3.08568e+16
	case "Picapt":
		m = num * 0.0003527777778
	case "pica":
		m = num * 0.004233333333
	case "survey_mi":
		m = num * 1609.347219
	}
	switch to {
	case "m":
		ans = m
	case "mi":
		ans = m * 0.0006213711922
	case "nmi":
		ans = m * 0.0005399568035
	case "in":
		ans = m * 39.37007874
	case "ft":
		ans = m * 3.280839895
	case "yd":
		ans = m * 1.093613298
	case "ell":
		ans = m * 0.8748906387
	case "ly":
		ans = m * 0
	case "parsec":
		ans = m * 0
	case "Picapt":
		ans = m * 2834.645669
	case "pica":
		ans = m * 236.2204724
	case "survey_mi":
		ans = m * 0.0006213699495
	default:
		panic(core.ErrInvalidInput)
	}
	return ans
}

func ConvertTime(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	var sec float64
	var ans float64
	if fr == to {
		return num
	}
	switch fr {
	case "s":
		sec = num
	case "min":
		sec = num * 60
	case "hr":
		sec = num * 3600
	case "day":
		sec = num * 86400
	case "yr":
		sec = num * 31557600
	}
	switch to {
	case "s":
		ans = sec
	case "min":
		ans = sec * 0.01666666667
	case "hr":
		ans = sec * 0.0002777777778
	case "day":
		ans = sec * 0.00001157407407
	case "yr":
		ans = sec * 0.00000003168808781
	}

	return ans
}

func ConvertSpeed(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	var mps float64
	var ans float64
	if fr == to {
		return num
	}
	switch fr {
	case "m/s":
		mps = num
	case "mph":
		mps = num * 0.44704
	case "m/h":
		mps = num * 0.0002777777778
	case "kn":
		mps = num * 0.5144444444
	case "admkn":
		mps = num * 0.5147733333
	}
	switch to {
	case "m/s":
		ans = mps
	case "mph":
		ans = mps * 2.236936292
	case "m/h":
		ans = mps * 3600
	case "kn":
		ans = mps * 1.943844492
	case "admkn":
		ans = mps * 1.942602569
	}

	return ans
}

func ConvertForce(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	var N float64
	var ans float64
	if fr == to {
		return num
	}
	switch fr {
	case "N":
		N = num
	case "dyn":
		N = num * 100000
	case "lbf":
		N = num * 0.2248089431
	case "pond":
		N = num * 101.9716213
	}
	switch to {
	case "N":
		ans = N
	case "dyn":
		ans = N * 0.00001
	case "lbf":
		ans = N * 4.448221615
	case "pond":
		ans = N * 0.00980665
	}
	return ans
}

func ConvertPressure(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	var pascal float64
	var ans float64
	if fr == to {
		return num
	}
	switch fr {
	case "p":
		pascal = num
	case "atm":
		pascal = num * 101325
	case "mmHg":
		pascal = num * 133.322
	case "psi":
		pascal = num * 6894.757293
	case "Torr":
		pascal = num * 133.3223684
	}
	switch fr {
	case "p":
		ans = pascal
	case "atm":
		ans = pascal * 0.000009869232667
	case "mmHg":
		ans = pascal * 0.007500637554
	case "psi":
		ans = pascal * 0.0001450377377
	case "Torr":
		ans = pascal * 0.007500616827
	}
	return ans
}
