package x

import (
	"github.com/antonmedv/expr"
	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/registry"
	"github.com/maniartech/x/statistics"
)

func Eval(e string) (interface{}, error) {

	env := registry.GetFunctions()
	return expr.Eval(e, env)
}

func init() {
	calc.Initialize()
	statistics.Initialize()
}
