package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestFinanceFunc(t *testing.T) {
	assert.Equal(t, 0.05354266737075841, finance.Effect(0.0525, 4))

	//PDuration
	assert.Equal(t, 3.859866162622655, finance.PDuaration(0.025, 2000, 2200))
	assert.Equal(t, 87.6054764193714, finance.PDuaration(0.025/12.0, 1000, 1200))
}
