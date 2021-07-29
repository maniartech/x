package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestPMTFunc(t *testing.T) {
	//PMT
	assert.Equal(t, 1037.0320893591606, finance.PMT(0.08/12.0, 10, 10000))
	assert.Equal(t, 1030.1643271779742, finance.PMT(0.08/12.0, 10, 10000, 0, 1))
	assert.Equal(t, 129.0811608679973, finance.PMT(0.06/12.0, 18*12, 0, 50000))

	//IPMT
	assert.Equal(t, 66.66666666666667, finance.IPMT(0.1/12.0, 1, 3*12, -8000))
	assert.Equal(t, 292.45, finance.IPMT(0.1, 3, 3, -8000))

	//PPMT
	assert.Equal(t, 75.62318600836673, finance.PPMT(0.1/12.0, 1, 2*12, 0, 2000))
	// assert.Equal(t, 27598.05, finance.PPMT(8.0/100.0, 10, 10, 200000))
}
