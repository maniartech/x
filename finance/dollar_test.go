package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestDollarFunc(t *testing.T) {
	assert.Equal(t, 1.125, finance.Dollarde(1.02, 16))
	assert.Equal(t, 1.125, finance.Dollarde(1.04, 32))
	assert.Equal(t, 1.3125, finance.Dollarde(1.1, 32))

	assert.Equal(t, 1.02, finance.Dollarfr(1.125, 16))
	assert.Equal(t, 1.04, finance.Dollarfr(1.125, 32))
	assert.Equal(t, 1.1, finance.Dollarfr(1.3125, 32))
}
