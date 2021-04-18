package datetime_test

import (
	"testing"
	"time"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var input time.Time = time.Date(2003, time.Month(2), 17, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, time.Date(2003, time.Month(5), 30, 0, 0, 0, 0, time.UTC), eoMonth(t, input, 3))
	assert.Equal(t, time.Date(2003, time.Month(4), 31, 0, 0, 0, 0, time.UTC), eoMonth(t, input, 2))
	assert.Equal(t, time.Date(2004, time.Month(2), 29, 0, 0, 0, 0, time.UTC), eoMonth(t, input, 12))
	assert.Equal(t, time.Date(2005, time.Month(2), 28, 0, 0, 0, 0, time.UTC), eoMonth(t, input, 24))
	assert.Equal(t, time.Date(2005, time.Month(2), 28, 0, 0, 0, 0, time.UTC), eoMonth(t, input, 24))
}
func eoMonth(t *testing.T, input time.Time, month int) time.Time {
	output := datetime.EOMonth(input, month)
	return output
}
