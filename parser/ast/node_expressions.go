package ast

import (
	"errors"
	"fmt"
)

type ExpressionNode struct {
	ASTNode
	Operator     string
	OperatorType OperatorType
	Left         Node
	Right        Node
}

func NewExpressionNode(line int, column int, operator string, left Node, right Node) *ExpressionNode {
	return &ExpressionNode{
		ASTNode: ASTNode{
			Type:     TypeFunc,
			Position: Position{Line: line, Column: column},
		},
		Operator: operator,
		Left:     left,
		Right:    right,
	}
}

func (n ExpressionNode) String() string {
	return fmt.Sprintf("{ Operator: %s, Left:%+v, Right:%+v }", n.Operator, n.Left, n.Right)
}

func (n ExpressionNode) Eval() (interface{}, error) {
	l, err := n.Left.Eval()
	if err != nil {
		return nil, err
	}

	r, err := n.Right.Eval()
	if err != nil {
		return nil, err
	}

	switch n.Operator {
	case "+":
		if v, ok := l.(Number); ok {
			return v + r.(Number), nil
		} else if v, ok := l.(string); ok {
			return v + r.(string), nil
		}
		return nil, errors.New(errInvalidType)
	case "-":
		if v, ok := l.(Number); ok {
			return v + r.(Number), nil
		}
		return nil, errors.New(errInvalidType)
		// case "*":
		// 	return l * r
		// case "/":
		// 	return l / r
		// case "%":
		// 	return l % r
		// case "=":
		// 	return l == r
		// case "<>":
		// 	return l != r
		// case ">":
		// 	return l > r
		// case "<":
		// 	return l < r
		// case ">=":
		// 	return l >= r
		// case "<=":
		// 	return l <= r
		// case "AND":
		// 	return l && r
		// case "OR":
		// 	return l || r
		// case "XOR":
		// 	return l != r
		// case "NOT":
		// 	return !r
	}
	return nil, errors.New(errInvalidOperator)
}
