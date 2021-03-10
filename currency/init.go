package currency

import "errors"

var errInvalidInput error = errors.New("invalid-input")

var singles []string = []string{
	"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

var tees []string = []string{
	"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
}

var tys []string = []string{
	"ten", "twenty", "thirty", "fourty", "fifty", "sixty", "seventy", "eighty", "ninety",
}
