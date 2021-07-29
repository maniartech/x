package x_test

import (
	"fmt"
	"testing"

	"github.com/maniartech/x"
	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {

	r, err := x.Exec("Average([t, t], t)", x.Env{"t": 10})
	assert.Equal(t, nil, err)
	assert.Equal(t, 10.0, r)

	// Normalization test!
	// The following expression resuts in floating point error
	// that returns 0.013000000000000001 instead of 0.013
	r, err = x.Exec("0.012 + 0.001", x.Env{}, false)

	assert.Equal(t, nil, err)
	assert.Equal(t, 0.013000000000000001, r)

	// The normalization should fix the floating point error!
	r, err = x.Exec("0.012 + 0.001", x.Env{}, true)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0.013, r)

	r, err = x.Exec("Digit2Word(Average(l))", x.Env{
		"l": []interface{}{15.0, 22, 30, 40, 50},
	})
	fmt.Println(r, err)
}
