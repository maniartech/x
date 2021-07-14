package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestMirr(t *testing.T) {
	assert.Equal(t, 0.13, finance.Mirr(0.1, 0.12, -120000, 39000, 30000, 21000, 37000, 46000))
}
