package finance

// toWordTens converts numbers form 0 to 99 to word form
func toWordTens(input int, multiplier bool) string {
	var output string
	if input > 0 && input < 20 {
		output = singles[input]
		if multiplier {
			output = tys[input-1]
		}
	} else if input >= 20 && input < 100 {
		if input%10 == 0 {
			output = tys[(input/10)-1]
		} else {
			output = tys[(input/10)-1] + " " + singles[input%10]
		}
	}
	return output
}

// toWordTens converts numbers form 0 to 999 to word form
func toWordHundreds(input int) string {
	var output string
	if input == 0 {
		output = "zero"
	}
	if input/100 != 0 {
		if input%100 != 0 {
			output = toWordTens(input/100, false) + " hundred " + toWordTens(input%100, false)
		} else {
			output = toWordTens(input/100, false) + " hundred"
		}
	} else {

		output = toWordTens(input%100, false)
	}
	return output
}
