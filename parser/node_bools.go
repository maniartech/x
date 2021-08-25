package parser

import "fmt"

type BooleanNode struct {
	ASTNode
	Value bool
}

func NewBooleanNode(line int, column int, value bool) Node {
	node := BooleanNode{
		ASTNode: ASTNode{
			Type: TypeBoolean,
			Position: Position{
				Line:   line,
				Column: column,
			},
		},
		Value: value,
	}
	return Node(node)
}

func (n BooleanNode) String() string {
	return fmt.Sprintf("BooleanNode: %v", n.Value)
}

// Eval evaluates the BooleanNode
func (n BooleanNode) Eval(map[string]interface{}) (interface{}, error) {
	return n.Value, nil
}
