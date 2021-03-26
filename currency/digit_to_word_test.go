package currency_test

import (
	"testing"

	"github.com/maniartech/go-funcs/currency"
	"github.com/stretchr/testify/assert"
)

func TestDigit2Word(t *testing.T) {

	assert.Equal(t, "zero", digit2Word(t, "0"))
	assert.Equal(t, "one", digit2Word(t, "1"))
	assert.Equal(t, "one two", digit2Word(t, "12"))
	assert.Equal(t, "one two three", digit2Word(t, "123"))
	assert.Equal(t, "one two three point one two", digit2Word(t, "123.12"))
	assert.Equal(t, "one two three point two", digit2Word(t, "123.2"))
	assert.Equal(t, "one two three point two two", digit2Word(t, "123.223"))

	// Failing cases
	digit2WordErr(t, "asdf", currency.ErrInvalidInput)
	digit2WordErr(t, "213.100", currency.ErrInvalidInput)
	digit2WordErr(t, "INRas.00", currency.ErrInvalidInput)
	digit2WordErr(t, "54.asdfb", currency.ErrInvalidInput)
	digit2WordErr(t, "PKR025.asdf", currency.ErrInvalidInput)
	digit2WordErr(t, "USD135.351.1.53", currency.ErrInvalidInput)
}

func digit2Word(t *testing.T, input string) string {
	output, err := currency.Digit2Word(input)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	return output
}

func digit2WordErr(t *testing.T, input string, err error) {
	assert.PanicsWithValue(t, err, func() { currency.Digit2Word(input) })
}
