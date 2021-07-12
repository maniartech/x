package finance_test

import (
	"testing"

	"github.com/maniartech/x/datetime"
	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestFinanceFunc(t *testing.T) {
	assert.Equal(t, 0.05354266737075841, finance.Effect(0.0525, 4))

	//PDuration
	assert.Equal(t, 3.859866162622655, finance.PDuaration(0.025, 2000, 2200))
	assert.Equal(t, 87.6054764193714, finance.PDuaration(0.025/12.0, 1000, 1200))

	//Nper
	assert.Equal(t, 59.67386567429457, finance.Nper(0.12/12.0, -100, -1000, 10000, 1))
	assert.Equal(t, 60.08212285376166, finance.Nper(0.12/12.0, -100, -1000, 10000))
	assert.Equal(t, -9.578594039813161, finance.Nper(0.12/12.0, -100, -1000))

	//PV
	assert.Equal(t, 59777.1458511878, finance.PV(0.08/12.0, 12*20, -500, 0, 0))

	//Effect
	assert.Equal(t, 0.05354266737075841, finance.Effect(0.0525, 4))

	//Disc
	assert.Equal(t, 0.000685899220489978, finance.Disc(datetime.Date(2018, 7, 1), datetime.Date(2048, 1, 1), 97.975, 100, 1))
}
