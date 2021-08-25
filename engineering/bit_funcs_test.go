package engineering_test

import (
	"testing"

	"github.com/maniartech/x/engineering"
	"github.com/stretchr/testify/assert"
)

func TestBitFunc(t *testing.T) {
	assert.Equal(t, int64(12), engineering.BitAnd(60, 13))
	assert.Equal(t, int64(1), engineering.BitAnd(1, 5))
	assert.Equal(t, int64(9), engineering.BitAnd(13, 25))

	assert.Equal(t, int64(61), engineering.BitOr(60, 13))
	assert.Equal(t, int64(31), engineering.BitOr(23, 10))

	assert.Equal(t, int64(49), engineering.BitXOR(60, 13))
	assert.Equal(t, int64(06), engineering.BitXOR(5, 3))
	assert.Equal(t, int64(49), engineering.BitXOR(60, 13))

	assert.Equal(t, int64(16), engineering.BitLShift(4, 2))
	assert.Equal(t, int64(2801664), engineering.BitLShift(684, 12))

	assert.Equal(t, int64(3), engineering.BitRShift(13, 2))
	assert.Equal(t, int64(171), engineering.BitRShift(684, 2))
}
