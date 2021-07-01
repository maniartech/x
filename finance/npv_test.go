package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestNpv(t *testing.T) {
	assert.Equal(t, 1188.4434123352303, finance.Npv(0.1, -10000, 3000, 4200, 6800))
	assert.Equal(t, 1922.0615549323848, finance.Npv(0.08, 8000, 9200, 10000, 12000, 14500)+(-40000))
	assert.Equal(t, -3749.465087015553, finance.Npv(0.08, 8000, 9200, 10000, 12000, 14500, -9000)+(-40000))
}
