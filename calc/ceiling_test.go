package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestCeling(t *testing.T) {

	assert.Equal(t, -32.0, calc.Ceiling(-42.5, 16))
	assert.Equal(t, -35.0, calc.Ceiling(-32.25, -5))

	assert.Equal(t, 16.0, calc.Ceiling(13, 4))
	assert.Equal(t, 2.0, calc.Ceiling(1.2, 1))

	assert.Equal(t, 4.45, calc.Ceiling("4.42", "0.05"))
	assert.Equal(t, 4.45, calc.Ceiling(4.42, 0.05))

	assert.Equal(t, 0.24, calc.Ceiling(0.234, 0.01))

	assert.Equal(t, 0.234, calc.Ceiling(0.234, 0.001))
	assert.Equal(t, 3.144, calc.Ceiling(3.1434643, 0.001))

	assert.Equal(t, 3.142863, calc.Ceiling(22.0/7.0, 0.000009))
	assert.Equal(t, 3.142858, calc.Ceiling(22.0/7.0, "0.000001"))
	assert.Equal(t, 0.23460, calc.Ceiling(0.23456, 0.00005))

	assert.Equal(t, 0.2345674444, calc.Ceiling(0.2345674444, 0.0000000001))
	assert.Equal(t, 0.234567456789, calc.Ceiling(0.234567456789, 0.000000000001))
}
