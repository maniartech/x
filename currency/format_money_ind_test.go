package currency_test

import (
	"testing"

	"github.com/maniartech/go-funcs/currency"
	"github.com/stretchr/testify/assert"
)

func TestFormatMoneyInd(t *testing.T) {

	assert.Equal(t, "₨111", formatMoneyInd(t, "111", currency.RupeeSymbol))
	assert.Equal(t, "INR1", formatMoneyInd(t, "1", "INR"))
	assert.Equal(t, "$11,13,658", formatMoneyInd(t, "1113658", "$"))
	assert.Equal(t, "Rs.11,13,658.456", formatMoneyInd(t, "1113658.456", "Rs."))

	// Failing cases
	assert.EqualError(t, formatMoneyIndErr(t, "$asdf", "$"), "invalid-input")
	assert.EqualError(t, formatMoneyIndErr(t, "INRas.00", "INR"), "invalid-input")
	assert.EqualError(t, formatMoneyIndErr(t, "54.asdfb", ""), "invalid-input")
	assert.EqualError(t, formatMoneyIndErr(t, "Rs.025.asdf", "Rs."), "invalid-input")
	assert.EqualError(t, formatMoneyIndErr(t, "INR135.351.1.53", "INR"), "invalid-input")

}

func formatMoneyInd(t *testing.T, input string, symbol string) string {
	output, err := currency.FormatMoneyInd(input, symbol)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	return output
}

func formatMoneyIndErr(t *testing.T, input string, symbol string) error {
	_, err := currency.FormatMoneyInd(input, symbol)
	return err
}
