package currency_test

import (
	"testing"

	"github.com/maniartech/go-funcs/currency"
	"github.com/stretchr/testify/assert"
)

func TestNumWord(t *testing.T) {

	assert.Equal(t, "one", currency.Num2WordInd("1"))
	assert.Equal(t, "two", currency.Num2WordInd("2"))
	assert.Equal(t, "three", currency.Num2WordInd("3"))
	assert.Equal(t, "four", currency.Num2WordInd("4"))
	assert.Equal(t, "five", currency.Num2WordInd("5"))
	assert.Equal(t, "six", currency.Num2WordInd("6"))
	assert.Equal(t, "seven", currency.Num2WordInd("7"))
	assert.Equal(t, "eight", currency.Num2WordInd("8"))
	assert.Equal(t, "nine", currency.Num2WordInd("9"))
	assert.Equal(t, "ten", currency.Num2WordInd("10"))
	assert.Equal(t, "eleven", currency.Num2WordInd("11"))
	assert.Equal(t, "twelve", currency.Num2WordInd("12"))
	assert.Equal(t, "thirteen", currency.Num2WordInd("13"))
	assert.Equal(t, "fourteen", currency.Num2WordInd("14"))
	assert.Equal(t, "fifteen", currency.Num2WordInd("15"))
	assert.Equal(t, "sixteen", currency.Num2WordInd("16"))
	assert.Equal(t, "seventeen", currency.Num2WordInd("17"))
	assert.Equal(t, "eighteen", currency.Num2WordInd("18"))
	assert.Equal(t, "nineteen", currency.Num2WordInd("19"))
	assert.Equal(t, "twenty", currency.Num2WordInd("20"))

	assert.Equal(t, "one arab ten crore eleven lac twenty six thousand five hundred and sixty nine", currency.Num2WordInd("01101126569"))
}
