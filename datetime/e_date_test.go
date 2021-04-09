package datetime_test

import (
	"testing"
	"time"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func TestEDate(t *testing.T) {
	var input time.Time = time.Date(2003, time.Month(2), 17, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, time.Date(2003, time.Month(5), 17, 0, 0, 0, 0, time.UTC), eDate(t, input, 3))
	assert.Equal(t, time.Date(2003, time.Month(4), 17, 0, 0, 0, 0, time.UTC), eDate(t, input, 2))
	assert.Equal(t, time.Date(2004, time.Month(2), 17, 0, 0, 0, 0, time.UTC), eDate(t, input, 12))
	assert.Equal(t, time.Date(2005, time.Month(2), 17, 0, 0, 0, 0, time.UTC), eDate(t, input, 24))
	assert.Equal(t, time.Date(2005, time.Month(2), 17, 0, 0, 0, 0, time.UTC), eDate(t, input, 24))
}
func eDate(t *testing.T, input time.Time, month int) time.Time {
	output := datetime.EDate(input, month)
	return output
}
