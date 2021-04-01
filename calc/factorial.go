package calc

//FACT func calculates the factorial of the number
func FACT(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * FACT(num-1)
}
//FACTDOUBLE func calculates the double factorial of the number
func FACTDOUBLE(num int) int {
	if num <= 2 {
		return num
	}
	return num * FACT(num-2)
}
