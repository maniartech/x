package x

import (
	"github.com/antonmedv/expr"
	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/registry"
	"github.com/maniartech/x/statistics"
)

// Env represent the environment
type Env map[string]interface{}

func Eval(e string, env Env) (interface{}, error) {

	if env == nil {
		env = make(Env)
	}
	registry.AttachFunctions(env)
	return expr.Eval(e, env)
}

func init() {
	calc.Initialize()
	statistics.Initialize()
}
