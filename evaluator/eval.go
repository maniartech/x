package evaluator

import "github.com/antonmedv/expr"

func Eval(e string) (interface{}, error) {
	return expr.Eval(e, nil)
}
