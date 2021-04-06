package currency_test

import (
	"testing"

<<<<<<< HEAD
	"github.com/maniartech/x/currency"
=======
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/currency"

>>>>>>> 8151667aa13b9e11970a2a53282d2337ac5b29de
	// "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestFormatMoney(t *testing.T) {

	assert.Equal(t, "$111", formatMoney(t, "111", currency.DollorSymbol))
	assert.Equal(t, "USD1", formatMoney(t, "1", "USD"))
	assert.Equal(t, "$1,113,658", formatMoney(t, "1113658", "$"))
	assert.Equal(t, "PKR1,113,658.456", formatMoney(t, "1113658.456", "PKR"))

	// Failing cases
	formatMoneyErr(t, "asdf", "", core.ErrInvalidInput)
	formatMoneyErr(t, "INRas.00", "INR", core.ErrInvalidInput)
	formatMoneyErr(t, "54.asdfb", "", core.ErrInvalidInput)
	formatMoneyErr(t, "PKR025.asdf", "PKR", core.ErrInvalidInput)
	formatMoneyErr(t, "USD135.351.1.53", "USD", core.ErrInvalidInput)
}

func formatMoney(t *testing.T, input string, symbol string) string {
	output := currency.FormatMoney(input, symbol)

	return output
}

func formatMoneyErr(t *testing.T, input string, symbol string, err error) {
	assert.PanicsWithValue(t, err, func() { currency.FormatMoney(input, symbol) })
}
