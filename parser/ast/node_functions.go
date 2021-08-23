package ast

import "fmt"

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
func (n FunctionNode) Eval() (interface{}, error) {
	panic(fmt.Sprintf("FunctionNode.Eval() not implemented for %v", n))
}
