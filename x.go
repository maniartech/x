package x

import (
	"github.com/antonmedv/expr"
	"github.com/maniartech/x/registry"
)

// export xeval
func Eval(e string, env Env) (interface{}, error) {

	if env == nil {
		env = make(Env)
	}
	registry.AttachFunctions(env)
	return expr.Eval(e, env)
}
