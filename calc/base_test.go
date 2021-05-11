package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/core"
	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {

	assert.Equal(t, "1010", calc.Base(10, 2))
	assert.Equal(t, "0000001010", calc.Base(10, 2, 10))
	assert.Equal(t, "000000000F", (calc.Base(15, 16, 10)))

	assert.Equal(t, "4H", (calc.Base(101, 21, 0)))
	assert.Equal(t, "94", calc.Base(202, 22, 0))
	assert.Equal(t, "D4", calc.Base(303, 23, 0))
	assert.Equal(t, "GK", calc.Base(404, 24, 0))
	assert.Equal(t, "K5", calc.Base(505, 25, 0))
	assert.Equal(t, "N8", calc.Base(606, 26, 0))
	assert.Equal(t, "Q5", calc.Base(707, 27, 0))
	assert.Equal(t, "10O", calc.Base(808, 28, 0))
	assert.Equal(t, "12A", calc.Base(909, 29, 0))
	assert.Equal(t, "B3K", calc.Base(10010, 30, 0))
	assert.Equal(t, "00B3K", calc.Base(10010, 30, 5))
	assert.Equal(t, "4KB50SRDLF", calc.Base(123123412341234, 31, 0))

	assert.PanicsWithValue(t, core.ErrInvalidInput, func() { calc.Base(-23, 1, 0) })
	assert.PanicsWithValue(t, core.ErrInvalidInput, func() { calc.Base(-23, 37, 0) })
	assert.PanicsWithValue(t, core.ErrInvalidInput, func() { calc.Base(-23, 2, -34) })
	assert.PanicsWithValue(t, core.ErrInvalidInput, func() { calc.Base(-23, 2, 256) })
}
