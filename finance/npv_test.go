package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestNpv(t *testing.T) {
	assert.Equal(t, 1188.44, finance.Npv(0.1, -10000, 3000, 4200, 6800))
}
