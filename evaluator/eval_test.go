package evaluator_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/maniartech/x/evaluator"
)

func TestEval(t *testing.T) {
	r, err := evaluator.Eval("AverageA(10 + (20 * 3)), 10, true)")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(">>> %v\n", r)
}
