package currency_test

import (
	"testing"

	"github.com/maniartech/go-funcs/currency"
	"github.com/stretchr/testify/assert"
)

func TestNumWordIND(t *testing.T) {

	assert.Equal(t, "zero", num2WordInd(t, "0"))
	assert.Equal(t, "one", num2WordInd(t, "1"))
	assert.Equal(t, "two", num2WordInd(t, "2"))
	assert.Equal(t, "three", num2WordInd(t, "3"))
	assert.Equal(t, "four", num2WordInd(t, "4"))
	assert.Equal(t, "five", num2WordInd(t, "5"))
	assert.Equal(t, "six", num2WordInd(t, "6"))
	assert.Equal(t, "seven", num2WordInd(t, "7"))
	assert.Equal(t, "eight", num2WordInd(t, "8"))
	assert.Equal(t, "nine", num2WordInd(t, "9"))
	assert.Equal(t, "ten", num2WordInd(t, "10"))
	assert.Equal(t, "eleven", num2WordInd(t, "11"))
	assert.Equal(t, "twelve", num2WordInd(t, "12"))
	assert.Equal(t, "thirteen", num2WordInd(t, "13"))
	assert.Equal(t, "fourteen", num2WordInd(t, "14"))
	assert.Equal(t, "fifteen", num2WordInd(t, "15"))
	assert.Equal(t, "sixteen", num2WordInd(t, "16"))
	assert.Equal(t, "seventeen", num2WordInd(t, "17"))
	assert.Equal(t, "eighteen", num2WordInd(t, "18"))
	assert.Equal(t, "nineteen", num2WordInd(t, "19"))
	assert.Equal(t, "twenty", num2WordInd(t, "20"))
	assert.Equal(t, "twenty one", num2WordInd(t, "21"))
	assert.Equal(t, "one hundred and one", num2WordInd(t, "101"))
	assert.Equal(t, "one hundred and ten", num2WordInd(t, "110"))
	assert.Equal(t, "one hundred and eleven", num2WordInd(t, "111"))
	assert.Equal(t, "one thousand one hundred and eleven", num2WordInd(t, "1111"))
	assert.Equal(t, "twenty one thousand one hundred and eleven", num2WordInd(t, "21111"))
	assert.Equal(t, "twenty thousand one hundred and eleven", num2WordInd(t, "20111"))
	assert.Equal(t, "one arab ten crore eleven lakh twenty six thousand five hundred and sixty nine", num2WordInd(t, "01101126569"))
	assert.Equal(t, "one hundred and eleven and ten paise", num2WordInd(t, "111.1"))
	assert.Equal(t, "one hundred and eleven and eleven paise", num2WordInd(t, "111.111"))
	assert.Equal(t, "one hundred and twelve", num2WordInd(t, "111.100"))
	assert.Equal(t, "one hundred and eleven and fifty six paise", num2WordInd(t, "111.56"))
	assert.Equal(t, "one hundred and eleven and fifty paise", num2WordInd(t, "111.50"))

	// Failing cases
	assert.EqualError(t, num2WordIndErr(t, "asdf"), "invalid-input")
	assert.EqualError(t, num2WordIndErr(t, "a.0"), "invalid-input")
	assert.EqualError(t, num2WordIndErr(t, "1.b"), "invalid-input")
	assert.EqualError(t, num2WordIndErr(t, "1.2.3"), "invalid-input")
	assert.EqualError(t, num2WordIndErr(t, ".9"), "invalid-input")
	assert.EqualError(t, num2WordIndErr(t, "1."), "invalid-input")
}

func num2WordInd(t *testing.T, input string) string {
	output, err := currency.Num2WordInd(input)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	return output
}

func num2WordIndErr(t *testing.T, input string) error {
	_, err := currency.Num2WordInd(input)
	return err
}
