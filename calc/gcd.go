package calc

import "github.com/maniartech/go-funcs/currency"

func calcGCD(a int, b int) int {
	if a == 0 {
		return b
	}
	return calcGCD(b%a, a)
}

// GCD Finds the GCD from a list of arrayed numbers
// If the numbers are not of type int then it truncates the number
// If the type is non-numerical then throws array
func GCD(arr []interface{}) int {
	//Creating a array of the type int with the size of inputed array
	arrInt := make([]int, len(arr))
	//Checking type of each value and storing them in new array
	for i, value := range arr {
		switch typedValue := value.(type) {
		//If the type of the `arr` element is int then it gets added to the `arrInt`
		case int:
			arrInt[i] = typedValue
			break
			//If the type of the `arr` element is float64 then it gets truncated
		case float64:
			arrInt[i] = INT(value.(float64))
			break
			//If the type of `arr` is neither int nor float64 then it creates a panic
		default:
			panic(currency.ErrInvalidInput)
		}
	}
	//Initalizing results with a default value
	result := arrInt[0]
	//Finding the GCD of 2 elements at a time
	for i := 1; i < len(arrInt); i++ {
		//Using the GCD value of the prev numbers we compute it with the next value
		result = calcGCD(int(arrInt[i]), result)
		//If the results == 1 then we return 1 
		if result == 1 {
			return 1
		}
	}
	return result
}
