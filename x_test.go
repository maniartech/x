package x_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/maniartech/x"
)

func TestEval(t *testing.T) {
	env := make(x.Env)
	env["ten"] = 10

	r, err := x.Eval("Average([ten, ten], ten)", env)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf(">>> %v\n", r)
}
