package statistics_test

import (
	"testing"

	"github.com/maniartech/x/statistics"
	"github.com/stretchr/testify/assert"
)

func TestShorts(t *testing.T) {

	assert.Equal(t, 0.9729550745276566, statistics.Fisher(0.75))

	assert.Equal(t, 0.75, statistics.FisherInv(statistics.Fisher(0.75)))

	assert.Equal(t, 26.054558142482477, statistics.StdevP(1345, 1301, 1368, 1322, 1310, 1370, 1350, 1303, 1318, 1299))

	// assert.Equal(t, -1.2, statistics.Kurt(1, 2, 3, 4, 5))
	assert.Equal(t, 1.0, statistics.Slope([]interface{}{2, 4, 6}, []interface{}{1, 3, 5}))
	assert.Equal(t, 0.6693548387096776, statistics.Slope([]interface{}{6, 5, 11, 7, 5}, []interface{}{2, 3, 9, 1, 8}))

	assert.Equal(t, -1.0, statistics.Intercept([]interface{}{2, 4, 6}, []interface{}{1, 3, 5}))
	assert.Equal(t, 0.04838709677419217, statistics.Intercept([]interface{}{6, 5, 11, 7, 5}, []interface{}{2, 3, 9, 1, 8}))

	assert.Equal(t, 0.8646647167633873, statistics.ExponDist(0.2, 10, true))
	assert.Equal(t, 1.353352832366127, statistics.ExponDist(0.2, 10, false))
}
