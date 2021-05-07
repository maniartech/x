package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestRoman(t *testing.T) {
	assert.Equal(t, "I", calc.Roman(1))
	assert.Equal(t, "II", calc.Roman(2))
	assert.Equal(t, "III", calc.Roman(3))
	assert.Equal(t, "IV", calc.Roman(4))
	assert.Equal(t, "V", calc.Roman(5))
	assert.Equal(t, "VI", calc.Roman(6))
	assert.Equal(t, "VII", calc.Roman(7))
	assert.Equal(t, "VIII", calc.Roman(8))
	assert.Equal(t, "IX", calc.Roman(9))
	assert.Equal(t, "X", calc.Roman(10))
	assert.Equal(t, "CDXCIX", calc.Roman(499))
	assert.Equal(t, "D", calc.Roman(500))
	assert.Equal(t, "MD", calc.Roman(1500))
	assert.Equal(t, "MDI", calc.Roman(1501))
	assert.Equal(t, "MDXXI", calc.Roman(1521))
	assert.Equal(t, "MDXLI", calc.Roman(1541))
	assert.Equal(t, "MDLI", calc.Roman(1551))
	assert.Equal(t, "MMMCMXCIX", calc.Roman(3999))
}
