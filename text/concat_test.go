package text_test

import (
	"math"
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/text"
	"github.com/maniartech/x/utils"
	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	assert.Equal(t, "asdf12@590.54", text.Concat("as", "df", 12, "@", 59, 0.54))
	assert.Equal(t, "testtrue", text.Concat("test", true))
	assert.Equal(t, utils.ToString(math.Pi), text.Concat(math.Pi))
	assert.Equal(t, "0truefalse", text.Concat(00, text.Concat(true, false)))
	assert.Equal(t, "21 : XXI", text.Concat(21, " : ", calc.Roman(21)))
	assert.Equal(t, "Hello : こんにちは", text.Concat("Hello : ", "こんにちは"))
	assert.Equal(t, "Happy : 😊", text.Concat("Happy : ", "😊"))
}
