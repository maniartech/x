package parser

import (
	"fmt"

	"github.com/maniartech/x/registry"
)

type FunctionNode struct {
	ASTNode
	Name string
	Args []Node
}

func NewFunctionNode(line int, column int, name string, args []Node) Node {
	node := FunctionNode{
		ASTNode: ASTNode{
			Type:     TypeFunc,
			Position: Position{Line: line, Column: column},
		},
		Name: name,
		Args: args,
	}
	return Node(node)
}

func (n FunctionNode) String() string {
	return fmt.Sprintf("{ FunctionName: %v, Args:%+v }", n.Name, n.Args)
}

// Eval evaluates the function node and returns the result.
func (n FunctionNode) Eval(env map[string]interface{}) (interface{}, error) {
	args := make([]interface{}, 0, len(n.Args))
	for _, arg := range n.Args {
		arg, err := arg.Eval(env)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
	}
	return registry.Invoke(n.Name, args, env)
}
