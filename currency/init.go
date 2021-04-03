package currency

var singles []string = []string{
	// The element 0 contains no value because the number zero isn't
	// rendered in the output of most numbers.
	// For example: 101 is read as one hundred and one and not one zero one
	"", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten", "eleven", "twelve",
	"thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
	"eighteen", "nineteen",
}
var tys []string = []string{
	"ten", "twenty", "thirty", "fourty", "fifty", "sixty",
	"seventy", "eighty", "ninety",
}
