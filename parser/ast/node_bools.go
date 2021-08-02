package ast

import "fmt"

type BooleanNode struct {
	ASTNode
	Value bool
}

func NewBooleanNode(line int, column int, value bool) *BooleanNode {
	return &BooleanNode{
		ASTNode: ASTNode{
			Type: TypeBoolean,
			Position: Position{
				Line:   line,
				Column: column,
			},
		},
		Value: value,
	}
}

func (n BooleanNode) String() string {
	return fmt.Sprint()
}
