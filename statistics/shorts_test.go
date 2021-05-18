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

	assert.Equal(t, -1.2, statistics.Kurt(1, 2, 3, 4, 5))

	assert.Equal(t, 3.3057189502100446, statistics.STEYX([]interface{}{2, 3, 9, 1, 8, 7, 5}, []interface{}{6, 5, 11, 7, 5, 4, 4}))

	assert.Equal(t, 1.0, statistics.Slope([]interface{}{1, 3, 5}, []interface{}{2, 4, 6}))
	assert.Equal(t, 0.6693548387096776, statistics.Slope([]interface{}{2, 3, 9, 1, 8}, []interface{}{6, 5, 11, 7, 5}))

	assert.Equal(t, -1.0, statistics.Intercept([]interface{}{1, 3, 5}, []interface{}{2, 4, 6}))
	assert.Equal(t, 0.04838709677419217, statistics.Intercept([]interface{}{2, 3, 9, 1, 8}, []interface{}{6, 5, 11, 7, 5}))

	assert.Equal(t, 0.8646647167633873, statistics.ExponDist(0.2, 10, true))
	assert.Equal(t, 1.353352832366127, statistics.ExponDist(0.2, 10, false))

	assert.Equal(t, 0.3595430714067974, statistics.Skew(3, 4, 5, 2, 3, 4, 5, 6, 4, 7))
	assert.Equal(t, 0.3031933393541438, statistics.SkewP(3, 4, 5, 2, 3, 4, 5, 6, 4, 7))

	assert.Equal(t, 10.607253086419755, statistics.Forecast(30, []interface{}{6, 7, 9, 15, 21}, []interface{}{20, 28, 31, 38, 40}))

	assert.Equal(t, 0.9970544855015815, statistics.Correl([]interface{}{3, 2, 4, 5, 6}, []interface{}{9, 7, 12, 15, 17}))

	assert.Equal(t, 1.329340388179137, statistics.Gamma(2.5))
	assert.Equal(t, 0.26786612886141653, statistics.Gamma(-3.75))

	assert.Equal(t, 0.205078125, statistics.BinomDist(6, 10, 0.5, false))
	assert.Equal(t, 0.828125, statistics.BinomDist(6, 10, 0.5, true))
}
