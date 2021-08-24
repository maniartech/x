package x

import (
	"io"
	"strings"

	"github.com/maniartech/x/parser"
)

// Evaluator evaluates the specified expression.
// It provides the methods like Compile, Run, Eval that
// can be used to evaluate the expression.
type Evaluator struct {
	node parser.Node
}

// Compiles the reader and returns the evaluator which can be
// used to evaluate the expression later.
func CompileReader(filename string, r io.Reader) (*Evaluator, error) {
	node, err := parser.ParseReader(filename, r)
	if err != nil {
		return nil, err
	}

	return &Evaluator{
		node: node.(parser.Node),
	}, nil
}

// Compile compiles the string expression which can be evaluated later.
func Compile(expression string) (*Evaluator, error) {
	node, err := parser.ParseReader("", strings.NewReader(expression))
	if err != nil {
		return nil, err
	}

	return &Evaluator{
		node: node.(parser.Node),
	}, nil
}

// Run evaluates the expression.
func (e *Evaluator) Run(env Env) (interface{}, error) {
	if e.node == nil {
		panic("expression not compiled")
	}

	return e.node.Eval(map[string]interface{}(env))
}

// Eval evaluates the expression. It compiles the expression if it is not
// compiled already.
func Eval2(expr string, env Env) (interface{}, error) {
	evaluator, err := Compile(expr)
	if err != nil {
		return nil, err
	}
	return evaluator.Run(env)
}
