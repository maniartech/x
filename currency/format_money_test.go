package currency_test

import (
	"testing"

	"github.com/maniartech/go-funcs/currency"
	"github.com/stretchr/testify/assert"
)

func TestFormatMoney(t *testing.T) {

	assert.Equal(t, "$111", formatMoney(t, "111", currency.DollorSymbol))
	assert.Equal(t, "USD1", formatMoney(t, "1", "USD"))
	assert.Equal(t, "$1,113,658", formatMoney(t, "1113658", "$"))
	assert.Equal(t, "PKR1,113,658.456", formatMoney(t, "1113658.456", "PKR"))

	// Failing cases
	assert.EqualError(t, formatMoneyErr(t, "$asdf", "$"), "invalid-input")
	assert.EqualError(t, formatMoneyErr(t, "INRas.00", "INR"), "invalid-input")
	assert.EqualError(t, formatMoneyErr(t, "54.asdfb", ""), "invalid-input")
	assert.EqualError(t, formatMoneyErr(t, "PKR025.asdf", "PKR"), "invalid-input")
	assert.EqualError(t, formatMoneyErr(t, "USD135.351.1.53", "USD"), "invalid-input")

}

func formatMoney(t *testing.T, input string, symbol string) string {
	output, err := currency.FormatMoney(input, symbol)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	return output
}

func formatMoneyErr(t *testing.T, input string, symbol string) error {
	_, err := currency.FormatMoney(input, symbol)
	return err
}
