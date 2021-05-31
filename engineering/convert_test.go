package engineering_test

import (
	"testing"

	"github.com/maniartech/x/engineering"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {

	assert.Equal(t, 1.0, engineering.Convert(1, "g", "g"))
	assert.Equal(t, 11.2, engineering.Convert(0.8, "stone", "lbm"))
	assert.Equal(t, 0.8, engineering.Convert(11.2, "lbm", "stone"))
	// assert.Equal(t, 6.85e-05, engineering.Convert(1, "g", "sg"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "lbm"))
	// assert.Equal(t, 6.02214E+23, engineering.Convert(1, "g", "u"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "ozm"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "grain"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "cwt"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "uk_cwt"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "stone"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "ton"))
	// assert.Equal(t, 0.002204622622, engineering.Convert(1, "g", "lton"))
	// assert.Equal(t, 11.2, engineering.Convert(0.8, "stone", "lbm"))
}
