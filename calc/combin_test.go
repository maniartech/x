package calc_test

import (
	"testing"

	"github.com/maniartech/x/calc"
	"github.com/stretchr/testify/assert"
)

func TestCombin(t *testing.T) {
	assert.Equal(t, 28, calc.Combin(8, 2))
}
