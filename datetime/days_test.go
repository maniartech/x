package datetime_test

import (
	"testing"
	"time"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func TestDays(t *testing.T) {

	date1 := datetime.Date(2017, 2, 6)
	date2 := datetime.Date(2018, 3, 5)
	var input time.Time = time.Date(2011, time.Month(12), 31, 0, 0, 0, 0, time.UTC)
	var input2 time.Time = time.Date(2011, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	var input3 time.Time = time.Date(2011, time.Month(3), 15, 0, 0, 0, 0, time.UTC)
	var input4 time.Time = time.Date(2011, time.Month(2), 1, 0, 0, 0, 0, time.UTC)
	assert.Equal(t, 364, days(t, input, input2))
	assert.Equal(t, 42, days(t, input3, input4))

	//Days360
	assert.Equal(t, 389, datetime.Days360(date1, date2))
	assert.Equal(t, 390, datetime.Days360(datetime.Date(2017, 2, 6), datetime.Date(2018, 3, 6)))
	assert.Equal(t, 391, datetime.Days360(datetime.Date(2017, 2, 6), datetime.Date(2018, 3, 7)))

	date1 = datetime.Date(2020, 1, 1)
	date2 = datetime.Date(2020, 12, 30)
	assert.Equal(t, 359, datetime.Days360(date1, date2))
	assert.Equal(t, -180, datetime.Days360(datetime.Date(2021, 1, 1), datetime.Date(2020, 6, 31)))

	date1 = datetime.Date(2020, 1, 1)
	date2 = datetime.Date(2020, 12, 31)
	assert.Equal(t, 360, datetime.Days360(date1, date2))
	assert.Equal(t, -360, datetime.Days360(date2, date1))
}

func days(t *testing.T, input time.Time, input2 time.Time) int {
	output := datetime.Days(input, input2)
	return output
}
