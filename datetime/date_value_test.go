package datetime_test

import (
	"testing"
	"time"

	"github.com/maniartech/x/currency"
	"github.com/stretchr/testify/assert"
)

func TestDateValue(t *testing.T) {
	var expectedAns time.Time = time.Date(2003, time.Month(2), 17, 0, 0, 0, 0, time.UTC)

	assert.Equal(t, expectedAns, dateValue(t, "17/02/2003"))
}
func dateValue(t *testing.T, input string) string {
	output, err := currency.Digit2Word(input)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	return output
}

func dateValueErr(t *testing.T, input string, err error) {
	assert.PanicsWithValue(t, err, func() { currency.Digit2Word(input) })
}
