package calc

//LCM func finds the Least common Multiple using the calcGCD function
func LCM(arr []int) int {
	// Initialize result
	ans := arr[0]

	// ans contains LCM of arr[0], ..arr[i]
	// after i'th iteration,
	for i := 1; i < len(arr); i++ {
		ans = ((arr[i] * ans) / (calcGCD(arr[i], ans)))
	}
	return ans
}
