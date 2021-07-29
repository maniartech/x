package parser

import (
	"errors"

	"github.com/maniartech/x/registry"
)

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func eval(first, rest interface{}) interface{} {
	l := first
	restSl := toIfaceSlice(rest)

	for _, v := range restSl {
		restExpr := toIfaceSlice(v)
		r := restExpr[3]
		op := ""

		if o, ok := restExpr[1].([]byte); ok {
			op = string(o)
		} else if o, ok := restExpr[1].(string); ok {
			op = o
		}

		if op != "" {
			l = ops[op](l, r)
		}
	}

	return l
}

// evalFn evalues the specified function expression and
// returns the value.
func EvalFn(funcName string, args ...interface{}) (interface{}, error) {
	fn, ok := registry.Get()[funcName].(func(args ...interface{}) (interface{}, error))
	if !ok {
		return nil, errors.New("function not found")
	}
	return fn(args...)
}
