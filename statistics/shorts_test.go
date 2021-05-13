package statistics_test

import (
	"testing"

	"github.com/maniartech/x/statistics"
	"github.com/stretchr/testify/assert"
)

func TestShorts(t *testing.T) {
	assert.Equal(t, 0.9729550745276566, statistics.Fisher(0.75))

	assert.Equal(t, 0.75, statistics.FisherInv(statistics.Fisher(0.75)))
}
