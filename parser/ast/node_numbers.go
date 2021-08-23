package ast

import (
	"fmt"
	"strconv"
)

// Number represents a number literal.
type Number float64

// String returns a string representation of the number.
func (n Number) String() string {
	return strconv.FormatFloat(float64(n), 'f', -1, 64)
}

// NumberNode represent a number literal.
type NumberNode struct {
	ASTNode

	// The value of the number.
	Value Number
}

// NewNumberNode creates a new NumberNode.
func NewNumberNode(line int, column int, value Number) Node {
	node := NumberNode{
		ASTNode: ASTNode{
			Type: TypeNumber,
			Position: Position{
				Line:   line,
				Column: column,
			},
		},
		Value: value,
	}
	return Node(node)
}

// String returns a string representation of the NumberNode.
func (n NumberNode) String() string {
	return fmt.Sprintf("NumberNode %s", n.Value)
}

// Eval evalues the NumberNode.
func (n NumberNode) Eval() (interface{}, error) {
	return n.Value, nil
}
