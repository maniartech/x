package engineering_test

import (
	"testing"

	"github.com/maniartech/x/engineering"
	"github.com/stretchr/testify/assert"
)

func TestComplex(t *testing.T) {
	assert.Equal(t, (1 + 2i), engineering.Complex(1, 2))
	assert.Equal(t, (1.1 + 2i), engineering.Complex(1.1, 2))
	assert.Equal(t, (1 + 2.134i), engineering.Complex(1, 2.134))
	assert.Equal(t, (1.123 + 2.8468i), engineering.Complex(1.123, 2.8468))
	assert.Equal(t, (-1 - 2i), engineering.Complex(-1, -2))

	assert.Equal(t, (-27.034945603074224 + 3.851153334811777i), engineering.IMCosh((4 + 3i)))
	assert.Equal(t, (-0.0754898329158637 + 0.0648774713706355i), engineering.IMCSC(4+3i))
}
