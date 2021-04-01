package statistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAveDev(t *testing.T) {

	d := AveDev(11, 6, 6, 12, 12, 7, 7, 9)
	assert.Equal(t, 2.25, d)

	d = AveDev([]float64{11, 6, 6, 12, 12, 7, 7, 9})
	assert.Equal(t, 2.25, d)

	d = AveDev([]float64{11, 6, 6, 12, 12}, 7, 7, 9)
	assert.Equal(t, 2.25, d)

	d = AveDev([]float64{11, 6, 6}, []float64{12, 12, 7, 7, 9})
	assert.Equal(t, 2.25, d)

}
