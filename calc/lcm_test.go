package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestLCM(t *testing.T) {
	//assert.Equal(t, 60, calc.LCM(15, 12))
	//assert.Equal(t, 4, calc.LCM(4.5, 2.9))

	assert.Equal(t, 120, calc.LCM(2, 4, 6, 8, 10))
	assert.Equal(t, 120, calc.LCM(2.468, 4.135, 6.84648, 8.45, 10.11531))
	assert.Equal(t, 80, calc.LCM(4, 16, 8, 20))
	assert.Equal(t, 1890, calc.LCM(6, 15, 21, 27, 30))
	assert.Equal(t, 0, calc.LCM(0, 0, 0, 0))
	assert.Equal(t, 0, calc.LCM(1, 2, 3, 0, 0))
}
