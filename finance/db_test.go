package finance_test

import (
	"testing"

	"github.com/maniartech/x/finance"
	"github.com/stretchr/testify/assert"
)

func TestDBFunction(t *testing.T) {
	cost := 1000000
	salvageValue := 100000
	life := 6
	assert.Equal(t, 186083.33333333334, finance.DB(cost, salvageValue, life, 1, 7))
	assert.Equal(t, 259639.41666666666, finance.DB(cost, salvageValue, life, 2, 7))
	assert.Equal(t, 176814.44275000002, finance.DB(cost, salvageValue, life, 3, 7))
	assert.Equal(t, 120410.63551274998, finance.DB(cost, salvageValue, life, 4, 7))
	assert.Equal(t, 81999.64278418274, finance.DB(cost, salvageValue, life, 5, 7))
	assert.Equal(t, 55841.75673602846, finance.DB(cost, salvageValue, life, 6, 7))
	//assert.Equal(t, 15845.098473848, finance.DB(cost, salvageValue, life, 7, 7))

	cost = 2400
	salvageValue = 300
	life = 10

	assert.Equal(t, 1.3150684931506849, finance.DDB(cost, salvageValue, life*365, 1))
	assert.Equal(t, 40.00, finance.DDB(cost, salvageValue, life*12, 1))
	assert.Equal(t, 480.00, finance.DDB(cost, salvageValue, life, 1))
	assert.Equal(t, 306.00, finance.DDB(cost, salvageValue, life, 2, 1.5))
	assert.Equal(t, 22.1225472000001, finance.DDB(cost, salvageValue, life, 10))
}
