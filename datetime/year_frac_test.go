package datetime_test

import (
	"testing"

	"github.com/maniartech/x/datetime"
	"github.com/stretchr/testify/assert"
)

func TestYearFrac(t *testing.T) {
	assert.Equal(t, 0.9972677595628415, datetime.YearFrac(datetime.Date(2020, 1, 1), datetime.Date(2020, 12, 31)))
	assert.Equal(t, 0.9972677595628415, datetime.YearFrac(datetime.Date(2020, 1, 1), datetime.Date(2020, 12, 31), 0))
	assert.Equal(t, 1.0, datetime.YearFrac(datetime.Date(2020, 1, 1), datetime.Date(2020, 12, 31), 1))
}
