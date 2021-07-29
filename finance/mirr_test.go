package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestMirr(t *testing.T) {
	assert.Equal(t, 0.12609413036590555, finance.Mirr(0.1, 0.12, -120000, 39000, 30000, 21000, 37000, 46000))
	assert.Equal(t, 0.13475911082831504, finance.Mirr(0.1, 0.14, -120000, 39000, 30000, 21000, 37000, 46000))
	assert.Equal(t, -0.04804465524998058, finance.Mirr(0.1, 0.12, -120000, 39000, 30000, 21000))
}
