package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestFV(t *testing.T) {
	assert.Equal(t, 2581.4033740601185, (finance.FV(0.06/12.0, 10, -200, -500, 1)))
	assert.Equal(t, 12682.503013196976, finance.FV(0.12/12.0, 12, -1000))
	assert.Equal(t, 82846.24637190053, finance.FV(0.11/12.0, 35, -2000, 0, 1))
	assert.Equal(t, 2301.4018303408993, finance.FV(0.06/12.0, 12, -100, -1000, 1))

	assert.Equal(t, 1.3309, finance.FVSchedule(1, 0.09, 0.11, 0, 1))
}
