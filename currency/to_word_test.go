package currency_test

import (
	"testing"

	"github.com/maniartech/go-funcs/currency"
	"github.com/stretchr/testify/assert"
)

func TestNumWord(t *testing.T) {

	assert.Equal(t, "zero", num2Word(t, "0"))
	assert.Equal(t, "one", num2Word(t, "1"))
	assert.Equal(t, "two", num2Word(t, "2"))
	assert.Equal(t, "three", num2Word(t, "3"))
	assert.Equal(t, "four", num2Word(t, "4"))
	assert.Equal(t, "five", num2Word(t, "5"))
	assert.Equal(t, "six", num2Word(t, "6"))
	assert.Equal(t, "seven", num2Word(t, "7"))
	assert.Equal(t, "eight", num2Word(t, "8"))
	assert.Equal(t, "nine", num2Word(t, "9"))
	assert.Equal(t, "ten", num2Word(t, "10"))
	assert.Equal(t, "eleven", num2Word(t, "11"))
	assert.Equal(t, "twelve", num2Word(t, "12"))
	assert.Equal(t, "thirteen", num2Word(t, "13"))
	assert.Equal(t, "fourteen", num2Word(t, "14"))
	assert.Equal(t, "fifteen", num2Word(t, "15"))
	assert.Equal(t, "sixteen", num2Word(t, "16"))
	assert.Equal(t, "seventeen", num2Word(t, "17"))
	assert.Equal(t, "eighteen", num2Word(t, "18"))
	assert.Equal(t, "nineteen", num2Word(t, "19"))
	assert.Equal(t, "twenty", num2Word(t, "20"))
	assert.Equal(t, "twenty one", num2Word(t, "21"))
	assert.Equal(t, "one hundred one", num2Word(t, "101"))
	assert.Equal(t, "one hundred ten", num2Word(t, "110"))
	assert.Equal(t, "one hundred eleven", num2Word(t, "111"))
	assert.Equal(t, "one thousand one hundred eleven", num2Word(t, "1111"))
	assert.Equal(t, "twenty one thousand one hundred eleven", num2Word(t, "21111"))
	assert.Equal(t, "twenty thousand one hundred eleven", num2Word(t, "20111"))
	assert.Equal(t, "one million one hundred eleven", num2Word(t, "1000111"))
	assert.Equal(t, "one hundred million one hundred eleven", num2Word(t, "100000111"))
	assert.Equal(t, "one billion one hundred one million one hundred twenty six thousand five hundred sixty nine", num2Word(t, "01101126569"))
	assert.Equal(t, "one hundred eleven and ten cent", num2Word(t, "111.1"))
	assert.Equal(t, "one hundred eleven and eleven cent", num2Word(t, "111.111"))
	assert.Equal(t, "one hundred eleven and fifty six cent", num2Word(t, "111.56"))
	assert.Equal(t, "one hundred eleven and fifty cent", num2Word(t, "111.50"))

	// Failing cases
	num2WordErr(t, "asdf", currency.ErrInvalidInput)
	num2WordErr(t, "123.321asdf", currency.ErrInvalidInput)
	num2WordErr(t, "123.100", currency.ErrInvalidInput)
	num2WordErr(t, "a.0", currency.ErrInvalidInput)
	num2WordErr(t, "a.0", currency.ErrInvalidInput)
	num2WordErr(t, "1.2.3", currency.ErrInvalidInput)
	num2WordErr(t, ".9", currency.ErrInvalidInput)
	num2WordErr(t, "1.", currency.ErrInvalidInput)
}

func num2Word(t *testing.T, input string) string {
	output := currency.Num2Word(input)
	return output
}
func num2WordErr(t *testing.T, input string, err error) {
	assert.PanicsWithValue(t, err, func() { currency.Num2Word(input) })
}
