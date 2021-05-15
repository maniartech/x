package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestRound(t *testing.T) {
	//assert.Equal(t, 7.655, calc.Round(7.6549, 3))
	//assert.Equal(t, 7.25, calc.Round(7.2522, 2))

	assert.Equal(t, 15.545, calc.Round(15.545, 3))
	assert.Equal(t, 15.55, calc.Round(15.545, 2))
	assert.Equal(t, 15.5, calc.Round(15.545, 1))
	assert.Equal(t, 16.0, calc.Round(15.545, 0))
	assert.Equal(t, 20.0, calc.Round(15.545, -1))
	assert.Equal(t, 0.0, calc.Round(15.545, -2))
	assert.Equal(t, 0.0, calc.Round(15.545, -3))
}
