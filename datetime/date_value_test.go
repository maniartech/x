package datetime_test

import (
	"testing"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func TestDateValue(t *testing.T) {
	// var expectedAns time.Time = time.Date(2003, time.Month(2), 17, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, 37669, dateValue(t, "2003/02/17"))
	assert.Equal(t, 37669, dateValue(t, "2003-02-17"))
	assert.Equal(t, 37669, dateValue(t, "2003-Febuary-17"))
	assert.Equal(t, 37669, dateValue(t, "2003-Feb-17"))
	assert.Equal(t, 37669, dateValue(t, "2003/Febuary/17"))
	assert.Equal(t, 37669, dateValue(t, "2003/Feb/17"))
	assert.Equal(t, 37669, dateValue(t, "2003/Feb/17"))
	assert.Equal(t, 37669, datetime.DateValue(datetime.Date(2003, 2, 17)))

}
func dateValue(t *testing.T, input string) int {
	output := datetime.DateValue(input)
	return output
}
