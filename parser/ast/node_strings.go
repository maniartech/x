package ast

import "fmt"

// StringNode is a node that represents a string literal.
type StringNode struct {
	ASTNode
	Value string
}

// NewStringNode returns a new StringNode.
func NewStringNode(line int, column int, value string) *StringNode {
	return &StringNode{
		ASTNode: ASTNode{
			Type: TypeString,
			Position: Position{
				Line:   line,
				Column: column,
			},
		},
		Value: value,
	}
}

// String returns a string representation of the StringNode.
func (n StringNode) String() string {
	return fmt.Sprint()
}
