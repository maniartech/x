package parser_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/parser"
	"github.com/maniartech/x/statistics"
)

func TestGrammer(t *testing.T) {
	calc.Initialize()
	statistics.Initialize()

	node, err := parser.ParseReader("", strings.NewReader("AVERAGE([10, 20] , 30, 40)"))
	if err != nil {
		t.Error(err)
	}

	startTime := time.Now()
	// for i := 0; i < 1000000; i++ {
	result, err := node.(parser.Node).Eval(map[string]interface{}{})
	// }

	endTime := time.Now()

	fmt.Println(">>>", endTime.Sub(startTime))

	if err != nil {
		t.Error(err)
	}

	fmt.Println("=", result)
}
