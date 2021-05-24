package engineering_test

import (
	"testing"

	"github.com/maniartech/x/engineering"
	"github.com/stretchr/testify/assert"
)

func TestBitFunc(t *testing.T) {
	assert.Equal(t, 12, engineering.BitAnd(60, 13))
	assert.Equal(t, 1, engineering.BitAnd(1, 5))
	assert.Equal(t, 9, engineering.BitAnd(13, 25))

	assert.Equal(t, 61, engineering.BitOr(60, 13))
	assert.Equal(t, 31, engineering.BitOr(23, 10))

	assert.Equal(t, 49, engineering.BitXOR(60, 13))
	assert.Equal(t, 6, engineering.BitXOR(5, 3))

	assert.Equal(t, 49, engineering.BitXOR(60, 13))

	assert.Equal(t, 16, engineering.BitLShift(4, 2))

	assert.Equal(t, 3, engineering.BitRShift(13, 2))
}
