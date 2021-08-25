package parser

import "fmt"

// StringNode is a node that represents a string literal.
type StringNode struct {
	ASTNode
	Value string
}

// NewStringNode returns a new StringNode.
func NewStringNode(line int, column int, value string) Node {
	node := StringNode{
		ASTNode: ASTNode{
			Type: TypeString,
			Position: Position{
				Line:   line,
				Column: column,
			},
		},
		Value: value,
	}
	return Node(node)
}

// String returns a string representation of the StringNode.
func (n StringNode) String() string {
	return fmt.Sprintf("StringNode: %s", n.Value)
}

// Eval evaluates the StringNode to string
func (n StringNode) Eval(env map[string]interface{}) (interface{}, error) {
	return n.Value, nil
}
