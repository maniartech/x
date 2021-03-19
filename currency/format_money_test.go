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
	formatMoneyErr(t, "asdf", "")
	formatMoneyErr(t, "INRas.00", "INR")
	formatMoneyErr(t, "54.asdfb", "")
	formatMoneyErr(t, "PKR025.asdf", "PKR")
	formatMoneyErr(t, "USD135.351.1.53", "USD")
}

func formatMoney(t *testing.T, input string, symbol string) string {
	output := currency.FormatMoney(input, symbol)

	return output
}

func formatMoneyErr(t *testing.T, input string, symbol string) {
	assert.Panics(t, func() { currency.FormatMoney(input, symbol) }, "The code did not panic")
}
