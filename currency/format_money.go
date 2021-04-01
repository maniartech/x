package currency

// FormatMoneySymbol formats the money with correct placement of ',' between digits according to international system
// this function also places a currency symbol before the number
func FormatMoney(input string, currencySymbol string) string {
	value := formatMoney(input, currencySymbol, 3)
	return value
}

// FormatMoneySymbolInd formats the money with correct placement of ',' between digits according to indian numbering system
// this function also places a currency symbol before the number
func FormatMoneyInd(input string, currencySymbol string) string {
	value := formatMoney(input, currencySymbol, 2)
	return value
}

// FormatMoneySymbol formats the money with correct placement of ',' between digits according to sepSpaces
// this function also places a currency symbol before the number
func formatMoney(input string, currencySymbol string, sepSpace int) string {
	var value string
	var number string
	//Spliting the input into two parts and checking for errors in input
	strArr := Split(input)
	var strArrLen int = len(strArr)
	//setting value of number to strArr[0] cause we need to slice the slice to add `,`
	number = strArr[0]
	//Checking if the lenght of the number is greater than three so if thats the case then we do the complex
	//Computation
	if len(number) > 3 {
		//Setting the value of value to the first three digits so that we can directly put a `,` when we enter the loop.
		value = number[len(number)-3:]
		for i := len(strArr[0]) - 3; i > 0; i = i - sepSpace {
			//Checking if the value of i - 3 is not less than 0 so that we don't get a out of range error
			if (i - sepSpace) <= 0 {
				value = number[:i] + "," + value
			} else {
				value = number[i-sepSpace:i] + "," + value
			}
		}
	} else {
		value = number
	}
	//If there is a `.` in the input then the lenght of the slice containing the values is 2,
	//Then we compute the 2nd value in the slice
	if strArrLen == 2 {
		value = value + "." + strArr[1]
	}
	return currencySymbol + value
}
