package datetime_test

import (
	"testing"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func TestYearFrac(t *testing.T) {
	date1 := datetime.Date(2020, 1, 1)
	date2 := datetime.Date(2020, 12, 31)
	assert.Equal(t, 1.0, datetime.YearFrac(date1, date2, 0))
	assert.Equal(t, 0.9972677595628415, datetime.YearFrac(date1, date2, 1))
	assert.Equal(t, 1.0138888888888888, datetime.YearFrac(date1, date2, 2))
	assert.Equal(t, 1.0, datetime.YearFrac(date1, date2, 3))

}
