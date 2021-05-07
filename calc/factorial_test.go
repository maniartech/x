package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	assert.Equal(t, 0, calc.Fact(0))
	assert.Equal(t, 1, calc.Fact(1))
	assert.Equal(t, 2, calc.Fact(2))
	assert.Equal(t, 6, calc.Fact(3))
	assert.Equal(t, 24, calc.Fact(4))
	assert.Equal(t, 120, calc.Fact(5))
	assert.Equal(t, 1, calc.Fact(1.4315))
	assert.Equal(t, 2, calc.Fact(2.1853))
	assert.Equal(t, 6, calc.Fact(3.153))
	assert.Equal(t, 24, calc.Fact(4.731))
	assert.Equal(t, 120, calc.Fact(5.13551354))
	assert.Equal(t, 479001600, calc.Fact(12))
	assert.Equal(t, 1307674368000, calc.Fact(15))
	assert.Equal(t, 355687428096000, calc.Fact(17))
	//assert.Equal(t, 243290200817664000000, calc.Fact(20))
	assert.PanicsWithValue(t, core.ErrInvalidInput, func() { calc.Fact(-23) })
}
func TestFactDouble(t *testing.T) {
	assert.Equal(t, 0, calc.FactDouble(0))
	assert.Equal(t, 1, calc.FactDouble(1))
	assert.Equal(t, 2, calc.FactDouble(2))
	assert.Equal(t, 3, calc.FactDouble(3))
	assert.Equal(t, 8, calc.FactDouble(4))
	assert.Equal(t, 15, calc.FactDouble(5))
	assert.Equal(t, 1, calc.FactDouble(1.4315))
	assert.Equal(t, 2, calc.FactDouble(2.1853))
	assert.Equal(t, 3, calc.FactDouble(3.153))
	assert.Equal(t, 8, calc.FactDouble(4.731))
	assert.Equal(t, 15, calc.FactDouble(5.13551354))
	assert.Equal(t, 46080, calc.FactDouble(12))
	assert.Equal(t, 2027025, calc.FactDouble(15))
	assert.Equal(t, 34459425, calc.FactDouble(17))
	assert.Equal(t, 3715891200, calc.FactDouble(20))
	assert.PanicsWithValue(t, core.ErrInvalidInput, func() { calc.FactDouble(-23) })
}
