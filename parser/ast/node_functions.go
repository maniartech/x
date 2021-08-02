package ast

import "fmt"

type FunctionNode struct {
	ASTNode
	Name string
	Args []Node
}

func NewFunctionNode(line int, column int, name string, args []Node) *FunctionNode {
	return &FunctionNode{
		ASTNode: ASTNode{
			Type:     TypeFunc,
			Position: Position{Line: line, Column: column},
		},
		Name: name,
		Args: args,
	}
}

func (n FunctionNode) String() string {
	return fmt.Sprintf("{ FunctionName: %v, Args:%+v }", n.Name, n.Args)
}
