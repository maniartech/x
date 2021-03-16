package is

// NotDigit checks if supplied character is a numerical digit or not.
// Eeturns true if the character is a digit otherwise returns false.
func NotDigit(c rune) bool { return c < '0' || c > '9' }
