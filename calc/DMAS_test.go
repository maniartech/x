package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/stretchr/testify/assert"
)

func TestDMAS(t *testing.T) {
	assert.Equal(t, 4102.0, calc.Sum(13, 4, 456, 25, 3543, []int{13, 48}))
	assert.Equal(t, 4102.55, calc.Sum(13, 4, 456.2, 25.25, 3543, []int{13, 48}, []float64{0.1}))
	assert.Equal(t, 4102.35456, calc.Sum(13, 4, 456.2, 25.25456, 3543, []int{13, 48}, []float64{-0.1}))
	assert.Equal(t, 4102.35456, calc.Sum(13, 4, 456.2, 25.25456, 3543, []int{13, 48}, []float64{-0.1}))

	assert.Equal(t, 25.0, calc.Minus(100, 75))
	assert.Equal(t, 25.332544, calc.Minus(100.456, 75.123456))
	assert.Equal(t, 100.0, calc.Minus(25, -75))
	assert.Equal(t, 100.25, calc.Minus(25.125, -75.125))

	assert.Equal(t, 10.0, calc.Product(10))
	assert.Equal(t, -8000.0, calc.Product(20, 5, []int{2, 10}, -4))
	assert.Equal(t, 4.06, calc.Product(2, 20.3, 0.1))

	assert.Equal(t, 20.06, calc.Divide(100, 5))
	assert.Equal(t, 1.33721217511612, calc.Divide(100.456, 75.123456))
	assert.Equal(t, -4.0, calc.Divide(20, -5))
	assert.Equal(t, -0.334442595673877, calc.Divide(25.125, -75.125))

	assert.PanicsWithValue(t, core.ErrDivideBy0, func() { calc.Divide(2, 0) })
}
