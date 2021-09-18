package text

func Contains(a []string, x string) bool {
	//Looping through all the values and checking if x is in the list of string 'a'
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
