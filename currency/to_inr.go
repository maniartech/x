package currency

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// func main() {
// 	var num string = "0.01"
// 	fmt.Println("Enter a number for which you want in word form ?")
// 	// fmt.Scanf("%d", &num)
// 	fmt.Printf("The word form of %s is %s", num, numWordINT(num))
// }

// Num2WordInd convert the number string into indian numbering word format.
func Num2WordInd(input string) (string, error) {
	word := ""
	part := ""
	paiseV := ""
	mod1 := 0
	mod2 := 0
	powerCounter := 0
	var number int = 0
	var paise int = 0
	var err error
	// err3B := false
	// fmt.Println(strings.Index(input, ".") == -1)
	stringArr := strings.Split(input, ".")
	stringArrLen := len(stringArr)

	if stringArrLen == 1 {
		number, err = strconv.Atoi(input)
		if err != nil {
			return "", err
		}
	} else if stringArrLen == 2 {
		number, err = strconv.Atoi(stringArr[0])
		if err != nil {
			return "", err
		}
		paise, err = strconv.Atoi(stringArr[1])
		if err != nil {
			return "", err
		}

	} else {
		return "", errors.New("invalid-string")
	}

	singles := [10]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tees := [10]string{"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tys := [9]string{"ten", "twenty", "thirty", "fourty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	power := [6]string{"", "thousand", "lakh", "crore", "arab", "kharab"}

	paiseS := strconv.Itoa(paise)
	if paise > 100 {
		paiseS = paiseS[:2] + "." + paiseS[2:]
		paiseF, err3 := strconv.ParseFloat(paiseS, 64)
		if err3 == nil {
			paise = int(math.Round(paiseF))
		}
	}
	if paise != 0 {
		if paise > 0 && paise < 10 {
			paiseV = singles[paise]
		} else if paise > 10 && paise < 20 {
			paiseV = tees[paise-10]
		} else if paise < 100 && paise > 19 {
			mod1 = paise % 10
			mod2 = paise / 10
			if mod2 != 0 {
				paiseV = tys[mod2-1] + " " + singles[mod1]
			} else {
				paiseV = singles[mod1]
			}
		} else if paise == 100 {
			number++
			paise = 0

		}

	}
	if number == 0 {
		word = "zero"
	} else if number > 0 && number < 10 {
		word = singles[number]
	} else if number == 10 {
		word = tys[0]
	} else if number > 10 && number < 20 {
		word = tees[number-11]
	} else {
		temp := int(number)
		for temp != 0 {
			part = ""
			if powerCounter == 0 {
				mod1 = temp % 100
				mod2 = temp % 1000
				word = singles[mod2/100]
				if mod2 != 0 {
					word = word + " hundred"
				}
				if mod1 > 10 && mod1 < 20 {
					word = word + tees[mod1-10]
				} else if mod1 == 10 {
					part = tys[0]
				} else {
					mod2 = mod1 / 10
					mod1 = mod1 % 10
					if mod1 != 0 {
						word = word + " " + tys[mod2-1]
					}
					if mod2 != 0 {
						word = word + " " + singles[mod1]
					}
				}
				powerCounter++
				temp = temp / 1000
			}
			mod1 = temp % 100
			if mod1 != 0 {
				mod2 = mod1 / 10
				if mod1 > 0 && mod1 < 10 {
					part = singles[mod1]
				}
				if mod1 > 0 && mod1 < 10 {
					part = singles[mod1]
				} else if mod1 == 10 {
					part = tys[0]
				} else if mod1 > 10 && mod1 < 20 {
					part = tees[mod1-11]
				} else {
					mod1 = mod1 % 10
					if mod1 != 0 {
						part = tys[mod2-1]
					}
					if mod2 != 0 {
						part = part + " " + singles[mod1]
					}
				}
				part = part + " " + power[powerCounter]
				word = part + " " + word
			}
			powerCounter++
			temp = temp / 100
		}
	}

	if paise == 0 {
		return word, nil
	}

	return word + " and " + paiseV + " paise", nil
}

func numWordINT(input string) string {
	word := ""
	part := ""
	centV := ""
	mod1 := 0
	mod2 := 0
	powerCounter := 0
	var number int = 0
	var cent int = 0
	var err error
	var err2 error
	var centS string = ""
	var multiplier bool = false
	singles := [10]string{"", "one ", "two ", "three ", "four ", "five ", "six ", "seven ", "eight ", "nine "}
	tees := [10]string{"eleven ", "twelve ", "thirteen ", "fourteen ", "fifteen ", "sixteen ", "seventeen ", "eighteen ", "nineteen "}
	tys := [9]string{"ten ", "twenty ", "thirty ", "fourty ", "fifty ", "sixty ", "seventy ", "eighty ", "ninety "}
	power := [6]string{"", "thousand ", "million ", "billion ", "trillion ", "quadrillion "}

	if strings.Index(input, ".") != -1 {
		number, err = strconv.Atoi(strings.Split(input, ".")[0])
		cent, err2 = strconv.Atoi(strings.Split(input, ".")[1])
		if len(strings.Split(input, ".")[1]) == 1 {
			multiplier = true
		}
	} else {
		number, err = strconv.Atoi(input)
	}
	if err == nil && err2 == nil {
		centS = strconv.Itoa(cent)
		if cent > 100 {
			centS = centS[:2] + "." + centS[2:]
			centF, err3 := strconv.ParseFloat(centS, 64)
			if err3 == nil {
				cent = int(math.Round(centF))
			}
		}
		if cent != 0 {
			if cent > 0 && cent < 10 {
				fmt.Println(multiplier)
				if multiplier == true {
					centV = tys[cent-1]
				} else {
					centV = singles[cent]
				}
			} else if cent > 10 && cent < 20 {
				centV = tees[cent-10]
			} else if cent < 100 && cent > 19 {
				mod1 = cent % 10
				mod2 = cent / 10
				if mod2 != 0 {
					centV = tys[mod2-1] + singles[mod1]
				} else {
					centV = singles[mod1]
				}
			} else if cent == 100 {
				number++
				cent = 0
			}
		}
	}
	if number == 0 {
		word = "zero "
	} else if number > 0 && number < 10 {
		word = singles[number]
	} else if number == 10 {
		word = tys[0]
	} else if number > 10 && number < 20 {
		word = singles[number]
	} else {
		temp := int(number)
		for temp != 0 {
			mod1 = temp % 1000
			word = power[powerCounter] + word
			mod2 = mod1 % 100
			if mod1 == 0 {
				continue
			} else {
				if mod2 == 0 {
					continue
				} else if mod2 > 0 && mod2 < 10 {
					part = singles[mod2]
				} else if mod2 == 10 {
					part = "ten "
				} else if mod2 > 10 && mod2 < 20 {
					part = tees[mod2-10]
				} else {
					part = tys[((mod1%100)/10)-1] + singles[mod1%10]
				}
				if mod1/100 != 0 {
					part = singles[mod1/100] + "hundred and " + part
				}
			}
			word = part + word
			temp = temp / 1000
			powerCounter++
		}
	}
	if cent == 0 {
		return word
	} else {
		return word + "and " + centV + "cent(s)"
	}
	return ""
}
