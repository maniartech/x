package engineering

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/text"
	"github.com/maniartech/x/utils"
)

//Convertion Table Mapping
var (
	toLiter  = map[string]float64{"l": 1, "tsp": l2Tsp, "tspm": l2Tspm, "tbs": l2Tbs, "oz": l2Oz, "cup": l2Cup, "pt": l2Pt, "uk_pt": l2UkPt, "qt": l2qt, "uk_qt": l2UkQt, "gal": l2Gal, "uk_gal": l2UkGal, "ang3": l2Ang3, "barrel": l2Barrel, "bushel": l2Bushel, "ft3": l2Ft3, "in3": l2In3, "ly3": l2Ly3, "m3": l2Ly3, "mi3": l2Mi3, "yd3": l2Yd3, "GRT": l2GRT, "MTON": l2MTON}
	toGram   = map[string]float64{"g": 1, "sg": g2Sg, "lbm": g2Lbm, "u": g2U, "ozm": g2Ozm, "gr": g2Gr, "cwt": g2Cwt, "uk_cwt": g2Uk_cwt, "stone": g2Stone, "ton": g2Ton, "uk_ton": g2Uk_ton}
	toMeter  = map[string]float64{"m": 1, "mi": m2Mi, "nmi": m2Nmi, "in": m2In, "ft": m2Ft, "yd": m2Yd, "ell": m2Ell, "ly": m2Ly, "parsec": m2Parsec, "Picapt": m2Picapt, "pica": m2Pica, "survey_mi": m2Survey_mi}
	toSec    = map[string]float64{"s": 1, "min": s2Min, "hr": s2Hr, "day": s2Day, "yr": s2Yr}
	toMps    = map[string]float64{"m/s": 1, "mph": mps2Mph, "m/h": mps2Miph, "kn": mps2Kn, "admkn": mps2Admkn}
	toNewton = map[string]float64{"N": 1, "dyn": n2Dyn, "lbf": n2Lbf, "pond": n2Pond}
	toPascal = map[string]float64{"p": 1, "atm": p2Atm, "mmHg": p2MmHg, "psi": p2Psi, "Torr": p2Torr}
	toJoule  = map[string]float64{"J": 1, "e": p2E, "c": p2C, "cal": p2Cal, "eV": p2EV, "hh": p2Hh, "wh": p2Wh, "flb": p2Flb, "btu": p2Btu}
	toBit    = map[string]float64{"bit": 1, "byte": b2Byte}
)

func Convert(number, fromUnit, toUnit interface{}) float64 {
	fr := utils.ToString(fromUnit)
	to := utils.ToString(toUnit)
	weight := []string{"g", "sg", "lbm", "u", "ozm", "gr", "cwt", "uk_cwt", "stone", "ton", "uk_ton"}
	distance := []string{"m", "mi", "nmi", "in", "ft", "yd", "ell", "ly", "parsec", "Picapt", "pica", "survey_mi"}
	time := []string{"s", "min", "hr", "day", "yr"}
	speed := []string{"m/s", "mph", "m/h", "kn", "admkn"}
	force := []string{"N", "dyn", "lbf", "pond"}
	pressure := []string{"p", "atm", "mmHg", "psi", "Torr"}
	energy := []string{"J", "e", "c", "cal", "eV", "hh", "wh", "flb", "btu"}
	temprature := []string{"C", "F", "K", "Rank", "Reau"}
	volume := []string{"l", "tsp", "tspm", "tbs", "oz", "cup", "pt", "uk_pt", "qt", "uk_qt", "gal", "uk_gal", "ang3", "barrel", "bushel", "ft3", "in3", "ly3", "m3", "mi3", "yd3", "GRT", "MTON"}
	information := []string{"bit", "byte"}

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
	} else if text.Contains(temprature, fr) && text.Contains(temprature, to) {
		return ConvertTemprature(number, fr, to)
	} else if text.Contains(volume, fr) && text.Contains(volume, to) {
		return ConvertVolume(number, fr, to)
	} else if text.Contains(information, fr) && text.Contains(information, to) {
		return ConvertInformation(number, fr, to)
	} else {
		panic(core.ErrInvalidInput)
	}
}

func ConvertMass(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	return ((1.0 / toGram[to]) * (num * toGram[fr]))
}

func ConvertDistance(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	return ((1.0 / toMeter[to]) * (num * toMeter[fr]))
}

func ConvertTime(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)

	return ((1.0 / toSec[to]) * (num * toSec[fr]))
}

func ConvertSpeed(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)

	return ((1.0 / toMps[to]) * (num * toMps[fr]))
}

func ConvertForce(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	if fr == to {
		return num
	}

	return ((1.0 / toNewton[to]) * (num * toNewton[fr]))
}
func ConvertPressure(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	if fr == to {
		return num
	}

	return ((1.0 / toPascal[to]) * (num * toPascal[fr]))
}

func ConvertEnergy(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)

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

	return ((1.0 / toBit[to]) * (num * toBit[fr]))
}

func ConvertVolume(number interface{}, fr, to string) float64 {
	num := utils.ToFloat64(number)
	return ((1.0 / toLiter[to]) * (num * toLiter[fr]))
}
