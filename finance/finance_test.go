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
	date1 := datetime.Date(2018, 7, 1)
	date2 := datetime.Date(2048, 1, 1)
	assert.Equal(t, 0.0006864406779661065, finance.Disc(date1, date2, 97.975, 100, 0))
	assert.Equal(t, 0.0006765033407572485, finance.Disc(date1, date2, 97.975, 100, 2))
	assert.Equal(t, 0.0006858992204899875, finance.Disc(date1, date2, 97.975, 100, 3))

	//SLN
	assert.Equal(t, 2250.0, finance.SLN(30000, 7500, 10))

	//SYD
	assert.Equal(t, 4090.90909090909, finance.SYD(30000, 7500, 10, 1))
	assert.Equal(t, 409.090909090909, finance.SYD(30000, 7500, 10, 10))

	//RRI
	assert.Equal(t, 0.0009933073762913303, finance.RRI(96, 10000, 11000))

	//IntRate
	date1 = datetime.Date(2008, 2, 15)
	date2 = datetime.Date(2008, 5, 15)
	assert.Equal(t, 0.05768, finance.IntRate(date1, date2, 1000000, 1014420, 0))
	assert.Equal(t, 0.058641333333333386, finance.IntRate(date1, date2, 1000000, 1014420, 1))
	assert.Equal(t, 0.05768, finance.IntRate(date1, date2, 1000000, 1014420, 2))
	assert.Equal(t, 0.05848111111111117, finance.IntRate(date1, date2, 1000000, 1014420, 3))

	//PriceDisc
	assert.Equal(t, 98.6875, finance.PriceDisc(date1, date2, 0.0525, 100, 0))
	assert.Equal(t, 98.70901639344262, finance.PriceDisc(date1, date2, 0.0525, 100, 1))
	assert.Equal(t, 98.6875, finance.PriceDisc(date1, date2, 0.0525, 100, 2))
	assert.Equal(t, 98.7054794520548, finance.PriceDisc(date1, date2, 0.0525, 100, 3))

	//PriceMat
	assert.Equal(t, 99.98449887555694, finance.PriceMat(datetime.Date(2008, 2, 15), datetime.Date(2008, 4, 13), datetime.Date(2007, 11, 11), 0.0610, 0.0610, 0))
	assert.Equal(t, 99.98468141300759, finance.PriceMat(datetime.Date(2008, 2, 15), datetime.Date(2008, 4, 13), datetime.Date(2007, 11, 11), 0.0610, 0.0610, 1))
	assert.Equal(t, 99.9841690643986, finance.PriceMat(datetime.Date(2008, 2, 15), datetime.Date(2008, 4, 13), datetime.Date(2007, 11, 11), 0.0610, 0.0610, 2))
	assert.Equal(t, 99.98459776456946, finance.PriceMat(datetime.Date(2008, 2, 15), datetime.Date(2008, 4, 13), datetime.Date(2007, 11, 11), 0.0610, 0.0610, 3))
}
