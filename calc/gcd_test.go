package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestGCD(t *testing.T) {
	assert.Equal(t, 2, calc.GCD(2, 4, 6, 8, 10))
	assert.Equal(t, 4, calc.GCD(4, 16, 8, 20))
	assert.Equal(t, 3, calc.GCD(6, 15, 21, 27, 30))
	assert.Equal(t, 0, calc.GCD(0, 0, 0, 0))
	assert.Equal(t, 1, calc.GCD(1, 2, 3, 0, 0))
}
