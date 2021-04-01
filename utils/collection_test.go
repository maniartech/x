package utils_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maniartech/x/utils"
	"github.com/stretchr/testify/assert"
)

func TestForEach(t *testing.T) {
	arr := [10]interface{}{
		"A", "B", "C",
		"D", "E", "F",
		"G", "H",
	}

	subArr1 := [3]interface{}{"C1", "C2"}
	subArr2 := [3]interface{}{"F1", "F2", "F3"}

	arr[2] = subArr1    // Assign array
	arr[5] = subArr2[:] // Assign slice

	slist := []string{}

	count := utils.ForEach(func(i int, v interface{}) {
		slist = append(slist, fmt.Sprint(v))
	}, arr[:]...)

	assert.Equal(t, 14, count)
	assert.Equal(t,
		"A,B,C1,C2,<nil>,D,E,F1,F2,F3,G,H,<nil>,<nil>",
		strings.Join(slist, ","))
}
