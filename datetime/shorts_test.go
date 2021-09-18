package datetime_test

import (
	"testing"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func TestShorts(t *testing.T) {
	assert.Equal(t, 4, datetime.Day(datetime.Date(2003, 12, 4)))
	assert.Equal(t, 31, datetime.Day(datetime.Date(1900, 15, 31)))
	assert.Equal(t, 5, datetime.Day(datetime.Date(9856, 5, 36)))
	assert.Equal(t, 12, datetime.Month(datetime.Date(2003, 12, 4)))
	assert.Equal(t, 3, datetime.Month(datetime.Date(1900, 15, 31)))
	assert.Equal(t, 6, datetime.Month(datetime.Date(9856, 5, 36)))
	assert.Equal(t, 2003, datetime.Year(datetime.Date(2003, 12, 4)))
	assert.Equal(t, 1901, datetime.Year(datetime.Date(1900, 13, 31)))
}
