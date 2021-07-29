package main

import (
	"fmt"

	"github.com/maniartech/x"
)

func main() {
	// c := make(chan bool)
	fmt.Println("Hello 21...")
	// js.Global().Set("xeval", js.FuncOf(Xeval))
	// println("...world")

	// <-c
}

// export xeval
func Xeval(expression string, context interface{}) (interface{}, error) {
	return x.Eval(expression, context.(map[string]interface{}))
}
