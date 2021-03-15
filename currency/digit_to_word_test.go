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
	assert.EqualError(t, digit2WordErr(t, "asdf"), "invalid-input")
	assert.EqualError(t, digit2WordErr(t, "a.0"), "invalid-input")
	assert.EqualError(t, digit2WordErr(t, "1.b"), "invalid-input")
	assert.EqualError(t, digit2WordErr(t, "1.2.3"), "invalid-input")
	assert.EqualError(t, digit2WordErr(t, ".9"), "invalid-input")
	assert.EqualError(t, digit2WordErr(t, "1."), "invalid-input")
	assert.EqualError(t, digit2WordErr(t, "1.100"), "invalid-input")
}

func digit2Word(t *testing.T, input string) string {
	output, err := currency.Digit2Word(input)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	return output
}

func digit2WordErr(t *testing.T, input string) error {
	_, err := currency.Digit2Word(input)
	return err
}
