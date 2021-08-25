package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/maniartech/x"
)

func main() {
	c := make(chan bool)
	js.Global().Set("goEval", js.FuncOf(xEval))
	<-c
}

// export xEval
func xEval(this js.Value, inputs []js.Value) interface{} {
	c := make(chan js.Value)
	go func() {
		expression := inputs[0].String()
		ctx := map[string]interface{}{}

		if len(inputs) > 1 {
			contextJSON := inputs[1].String()
			json.Unmarshal([]byte(contextJSON), &ctx)
		}

		result, err := x.Eval(expression, ctx)
		if err != nil {
			c <- js.ValueOf(fmt.Sprintf("xError:%s", fmt.Sprintf("%v", err)))
		}
		c <- js.ValueOf(result)
	}()
	defer func() {
		if r := recover(); r != nil {
			c <- js.ValueOf(fmt.Sprintf("xError:%s", fmt.Sprintf("%v", r)))
		}
	}()
	return <-c
}
