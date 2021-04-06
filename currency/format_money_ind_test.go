package currency_test

import (
	"testing"

	"github.com/maniartech/x/currency"
	"github.com/stretchr/testify/assert"
)

func TestFormatMoneyInd(t *testing.T) {

	assert.Equal(t, "₨111", formatMoneyInd(t, "111", currency.RupeeSymbol))
	assert.Equal(t, "INR1", formatMoneyInd(t, "1", "INR"))
	assert.Equal(t, "$11,13,658", formatMoneyInd(t, "1113658", "$"))
	assert.Equal(t, "Rs.11,13,658.456", formatMoneyInd(t, "1113658.456", "Rs."))

	// Failing cases
	formatMoneyIndErr(t, "asdf", "", currency.ErrInvalidInput)
	formatMoneyIndErr(t, "INRas.00", "INR", currency.ErrInvalidInput)
	formatMoneyIndErr(t, "54.asdfb", "", currency.ErrInvalidInput)
	formatMoneyIndErr(t, "PKR025.asdf", "PKR", currency.ErrInvalidInput)
	formatMoneyIndErr(t, "USD135.351.1.53", "USD", currency.ErrInvalidInput)
}

func formatMoneyInd(t *testing.T, input string, symbol string) string {
	output := currency.FormatMoneyInd(input, symbol)
	return output
}

func formatMoneyIndErr(t *testing.T, input string, symbol string, err error) {
	assert.PanicsWithValue(t, err, func() { currency.FormatMoneyInd(input, symbol) })
}
