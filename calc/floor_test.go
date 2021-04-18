package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestFloor(t *testing.T) {
	assert.Equal(t, 120.0, calc.Floor(123.456, 10))
	assert.Equal(t, 100.0, calc.Floor(123.456, 100))
	assert.Equal(t, 0.0, calc.Floor(123.456, 1000))
	assert.Equal(t, 123.2, calc.Floor(123.456, 1.1))
	assert.Equal(t, 122.4, calc.Floor(123.456, 1.2))
	assert.Equal(t, 122.2, calc.Floor(123.456, 1.3))
	assert.Equal(t, 123.456, calc.Floor(123.456, 123.456))
	assert.Equal(t, 123.4, calc.Floor(123.456, 12.34))

	assert.Equal(t, -124.3, calc.Floor(-123.456, 1.1))
	assert.Equal(t, -123.6, calc.Floor(-123.456, 1.2))
	assert.Equal(t, -123.5, calc.Floor(-123.456, 1.3))

	assert.PanicsWithValue(t, "decimal division by 0", func() {
		calc.Floor(123.456, 0)
	})
}
