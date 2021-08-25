package engineering

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

const limit int64 = 281474976710656 // 2 ^ 48
const shift_limit int = 53

//BitAnd returns bitwise AND ( & ) operation between two operands.
//
//Arguments
//
//num1: A decimal number greater than 0.
//
//num2: A decimal number greater than 0.
//
//Remark
//
//If either of the input arguments are greater than or equal to 2 ^ 48 or 281474976710656 the fucntion will panic ErrInvalidInput
//
//If either of the input argument are less than 0 then the fucntion will panic ErrInvalidInput
//
//Examples
//	BitAnd(60, 13) //returns 12
//	BitAnd(1, 5) //returns  1
func BitAnd(num1, num2 interface{}) int64 {
	Num1 := utils.ToInt64(num1)
	Num2 := utils.ToInt64(num2)

	if Num1 < 0 || Num2 < 0 || Num1 >= limit || Num2 >= limit {
		panic(core.ErrInvalidInput)
	}
	return (Num1 & Num2)
}

//BitOr returns bitwise OR ( | ) operation between two operands.
//
//Arguments
//
//num1: A decimal number greater than 0.
//
//num2: A decimal number greater than 0.
//
//Remark
//
//If either of the input arguments are greater than or equal to 2 ^ 48 or 281474976710656 the fucntion will panic ErrInvalidInput
//
//If either of the input argument are less than 0 then the fucntion will panic ErrInvalidInput
//
//Examples
//	BitOr(60, 13) //returns 61
//	BitOr(23, 10) //returns 31
func BitOr(num1, num2 interface{}) int64 {
	Num1 := utils.ToInt64(num1)
	Num2 := utils.ToInt64(num2)

	if Num1 < 0 || Num2 < 0 || Num1 >= limit || Num2 >= limit {
		panic(core.ErrInvalidInput)
	}
	return (Num1 | Num2)
}

//BitXOR returns bitwise Exclusive OR ( ^ ) operation between two operands.
//
//Arguments
//
//num1: A decimal number greater than 0.
//
//num2: A decimal number greater than 0.
//
//Remark
//
//If either of the input arguments are greater than or equal to 2 ^ 48 or 281474976710656 the fucntion will panic ErrInvalidInput
//
//If either of the input argument are less than 0 then the fucntion will panic ErrInvalidInput
//
//Examples
//	BitXOR(60, 13) //returns 49
//	BitXOR(5, 3) //returns 6
func BitXOR(num1, num2 interface{}) int64 {
	Num1 := utils.ToInt64(num1)
	Num2 := utils.ToInt64(num2)

	if Num1 < 0 || Num2 < 0 || Num1 >= limit || Num2 >= limit {
		panic(core.ErrInvalidInput)
	}
	return (Num1 ^ Num2)
}

//BitLShift returns number with its bits shifted to left by a specified amount.
//
//Arguments
//
//num1: Number of which the bits will be shifted, must be an Integer and greater than or equal to 0.
//
//shift: The number of bits need to be shifted, must be an integer and less than 53.
//
//Remark
//
//If either of the input arguments are greater than or equal to 2 ^ 48 or 281474976710656 the fucntion will panic ErrInvalidInput.
//
//If either of the input argument are less than 0 then the fucntion will panic ErrInvalidInput.
//
//If the shift amount is a negetive number it will shift the bit to the opposite direction.
//
//Results from a negative left shift will be same as a positive right shift.
//
//Examples
//	BitLShift(4, 2) //returns 16
//	BitLShift(684, 12) //returns 2801664
func BitLShift(num1, shift interface{}) int64 {
	Num1 := utils.ToInt64(num1)
	Shift := utils.ToInt(shift)

	if Num1 < 0 || Num1 >= limit || Shift >= shift_limit {
		panic(core.ErrInvalidInput)
	}
	return (Num1 << Shift)
}

//BitXOR returns bitwise Exclusive OR ( ^ ) operation between two operands.
//
//Arguments
//
//num1: Number of which the bits will be shifted, must be an Integer and greater than or equal to 0.
//
//shift: The number of bits need to be shifted, must be an integer and less than 53.
//
//Remark
//
//If either of the input arguments are greater than or equal to 2 ^ 48 or 281474976710656 the fucntion will panic ErrInvalidInput.
//
//If either of the input argument are less than 0 then the fucntion will panic ErrInvalidInput.
//
//If the shift amount is a negetive number it will shift the bit to the opposite direction.
//
//Results from a negative right shift will be same as a positive left shift.
//
//Examples
//	BitRShift(13, 2) //returns 3
//	BitRShift(684, 2) //returns 171
func BitRShift(num1, shift interface{}) int64 {
	Num1 := utils.ToInt64(num1)
	Shift := utils.ToInt(shift)

	if Num1 < 0 || Shift < 0 || Num1 >= limit || Shift >= shift_limit {
		panic(core.ErrInvalidInput)
	}
	return (Num1 >> Shift)
}
