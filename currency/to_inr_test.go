package currency_test

import (
	"testing"

	"github.com/maniartech/go-funcs/currency"
	"github.com/stretchr/testify/assert"
)

func TestNumWord(t *testing.T) {

	assert.Equal(t, "one", num2WordInd(t, "1"))
	assert.Equal(t, "two", num2WordInd(t, "2"))
	assert.Equal(t, "three", num2WordInd(t, "3"))
	assert.Equal(t, "four", num2WordInd(t, "4"))
	assert.Equal(t, "five", num2WordInd(t, "5"))
	assert.Equal(t, "six", num2WordInd(t, "6"))
	assert.Equal(t, "seven", num2WordInd(t, "7"))
	assert.Equal(t, "eight", num2WordInd(t, "8"))
	assert.Equal(t, "nine", num2WordInd(t, "9"))
	assert.Equal(t, "ten", num2WordInd(t, "10"))
	assert.Equal(t, "eleven", num2WordInd(t, "11"))
	assert.Equal(t, "twelve", num2WordInd(t, "12"))
	assert.Equal(t, "thirteen", num2WordInd(t, "13"))
	assert.Equal(t, "fourteen", num2WordInd(t, "14"))
	assert.Equal(t, "fifteen", num2WordInd(t, "15"))
	assert.Equal(t, "sixteen", num2WordInd(t, "16"))
	assert.Equal(t, "seventeen", num2WordInd(t, "17"))
	assert.Equal(t, "eighteen", num2WordInd(t, "18"))
	assert.Equal(t, "nineteen", num2WordInd(t, "19"))
	assert.Equal(t, "twenty", num2WordInd(t, "20"))

	assert.Equal(t, "one arab ten crore eleven lac twenty six thousand five hundred and sixty nine", num2WordInd(t, "01101126569"))
}

func num2WordInd(t *testing.T, input string) string {
	output, err := currency.Num2WordInd(input)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	return output
}
