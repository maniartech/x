package engineering

import (
	"github.com/maniartech/x/text"
	"github.com/maniartech/x/utils"
)

const (
	l2Tsp    float64 = 0.004928921594
	l2Tspm   float64 = 0.005
	l2Tbs    float64 = 0.01478676478
	l2Oz     float64 = 0.02957352956
	l2Cup    float64 = 0.2365882365
	l2pt     float64 = 0.473176473
	l2UkPt   float64 = 0.56826125
	l2qt     float64 = 0.946352946
	l2UkQt   float64 = 1.1365225
	l2Gal    float64 = 3.785411784
	l2UkGal  float64 = 4.54609
	l2Ang3   float64 = 1e-27
	l2Barrel float64 = 158.9872949
	l2Bushel float64 = 35.23907017
	l2Ft3    float64 = 28.31684659
	l2In3    float64 = 0.016387064
	l2Ly3    float64 = 8.46787e+50
	l2M3     float64 = 1000
	l2Mi3    float64 = 4168181825441
	l2Yd3    float64 = 764.554858
	l2GRT    float64 = 2831.684659
	l2MTON   float64 = 1132.673864
)

var toLiter = map[string]float64{"l": 1, "tsp": l2Tsp, "tspm": l2Tspm, "tbs": l2Tbs, "oz": l2Oz, "cup": l2Cup, "pt": l2pt, "uk_pt": l2UkPt, "qt": l2qt, "uk_qt": l2UkQt, "gal": l2Gal, "uk_gal": l2UkGal, "ang3": l2Ang3, "barrel": l2Barrel, "bushel": l2Bushel, "ft3": l2Ft3, "in3": l2In3, "ly3": l2Ly3, "m3": l2Ly3, "mi3": l2Mi3, "yd3": l2Yd3, "GRT": l2GRT, "MTON": l2MTON}

func Convert(number, fromUnit, toUnit interface{}) float64 {
	fr := utils.ToString(fromUnit)
	to := utils.ToString(toUnit)
	weight := []string{"g", "sg", "lbm", "u", "ozm", "gr", "cwt", "uk_cwt", "stone", "ton", "uk_ton"}
	distance := []string{"m", "mi", "nmi", "in", "ft", "yd", "ell", "ly", "parsec", "Picapt", "pica", "survey_mi"}
	time := []string{"s", "min", "hr", "day", "yr"}
	speed := []string{"m/s", "mph", "m/h", "kn", "admkn"}
	force := []string{"N", "dyn", "lbf", "pond"}
	pressure := []string{"p", "atm", "mmHg", "psi", "Torr"}
	energy := []string{"e", "c", "cal", "eV", "hh", "wh", "flb", "btu"}
	temprature := []string{"C", "F", "K", "Rank", "Reau"}
	information := []string{"bit", "byte"}
	volume := []string{"l", "tsp", "tspm", "tbs", "oz", "cup", "pt", "uk_pt", "qt", "uk_qt", "gal", "uk_gal", "ang3", "barrel", "bushel", "ft3", "in3", "ly3", "m3", "mi3", "yd3", "GRT", "MTON"}

	if text.Contains(weight, fr) {
		return ConvertMass(number, fr, to)
	} else if text.Contains(distance, fr) {
		return ConvertDistance(number, fr, to)
	} else if text.Contains(time, fr) {
		return ConvertTime(number, fr, to)
	} else if text.Contains(speed, fr) {
		return ConvertSpeed(number, fr, to)
	} else if text.Contains(force, fr) {
		return ConvertForce(number, fr, to)
	} else if text.Contains(pressure, fr) {
		return ConvertPressure(number, fr, to)
	} else if text.Contains(energy, fr) {
		return ConvertEnergy(number, fr, to)
	}

	return 0.0
}

func ConvertMass(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	toGram := map[string]float64{"g": 1, "sg": 14593.90294, "lbm": 453.59237, "u": 1.660540199e-24, "ozm": 28.34952313, "gr": 0.06479891, "cwt": 45359.237, "uk_cwt": 50802.34544, "stone": 6350.29318, "ton": 907184.74, "uk_ton": 1016046.909}

	return ((1.0 / toGram[to]) * (num * toGram[fr]))
}

func ConvertDistance(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	toMeter := map[string]float64{"m": 1, "mi": 1609.344, "nmi": 1852, "in": 0.0254, "ft": 0.3048, "yd": 0.9144, "ell": 1.143, "ly": 9460730472580800, "parsec": 3.08567758e+16, "Picapt": 0.0003527777778, "pica": 0.004233333333, "survey_mi": 1609.347219}

	return ((1.0 / toMeter[to]) * (num * toMeter[fr]))
}

func ConvertTime(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	toSec := map[string]float64{"s": 1, "min": 60, "hr": 3600, "day": 86400, "yr": 31557600}

	return ((1.0 / toSec[to]) * (num * toSec[fr]))
}

func ConvertSpeed(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	toMps := map[string]float64{"m/s": 1, "mph": 0.44704, "m/h": 0.0002777777778, "kn": 0.5144444444, "admkn": 0.5147733333}

	return ((1.0 / toMps[to]) * (num * toMps[fr]))
}

func ConvertForce(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	if fr == to {
		return num
	}
	toNewton := map[string]float64{"N": 1, "dyn": 100000, "lbf": 0.2248089431, "pond": 101.9716213}

	return ((1.0 / toNewton[to]) * (num * toNewton[fr]))
}
func ConvertPressure(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	if fr == to {
		return num
	}
	toPascal := map[string]float64{"p": 1, "atm": 101325, "mmHg": 133.322, "psi": 6894.757293, "Torr": 133.3223684}

	return ((1.0 / toPascal[to]) * (num * toPascal[fr]))
}

//TODO Double check value of 1 eV
func ConvertEnergy(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	toJoule := map[string]float64{"e": 0.0000001, "c": 4.184, "cal": 4.1868, "eV": 1.602176634e-19, "hh": 2684519.538, "wh": 3600, "flb": 1.355817948, "btu": 1055.055853}

	return ((1.0 / toJoule[to]) * (num * toJoule[fr]))
}

func ConvertTemprature(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	if fr == to {
		return num
	}
	toC := 0.0
	ans := 0.0

	switch fr {
	case "C":
		toC = 1
	case "F":
		toC = (num - 32) * (1.0 / 1.8)
	case "K":
		toC = num - 273.15
	case "Rank":
		toC = (num - 491.67) * 9.0 / 5.0
	case "Reau":
		toC = num * 5 / 4
	}
	switch to {
	case "C":
		ans = 1
	case "F":
		ans = (toC * 1.8) + 32
	case "K":
		ans = toC + 273.15
	case "Rank":
		ans = (toC + 273.15) * (9.0 / 5.0)
	case "Reau":
		ans = toC * 4 / 5
	}
	return ans
}

func ConvertInformation(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	toBit := map[string]float64{"bit": 1, "byte": 8}

	return ((1.0 / toBit[to]) * (num * toBit[fr]))
}

func ConvertVolume(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)

	return ((1.0 / toLiter[to]) * (num * toLiter[fr]))
}
