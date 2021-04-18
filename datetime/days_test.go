package datetime_test

import (
	"testing"
	"time"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func TestDays(t *testing.T) {
	var input time.Time = time.Date(2011, time.Month(12), 31, 0, 0, 0, 0, time.UTC)
	var input2 time.Time = time.Date(2011, time.Month(1), 1, 0, 0, 0, 0, time.UTC)

	var input3 time.Time = time.Date(2011, time.Month(3), 15, 0, 0, 0, 0, time.UTC)
	var input4 time.Time = time.Date(2011, time.Month(2), 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 364, days(t, input, input2))
	assert.Equal(t, 42, days(t, input3, input4))
}
func days(t *testing.T, input time.Time, input2 time.Time) int {
	output := datetime.Days(input, input2)
	return output
}
