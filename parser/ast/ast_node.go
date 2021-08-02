package ast

type Position struct {
	Line   int
	Column int
	Offset int
}

type ASTNode struct {
	Type     NodeType
	Position Position
}

type Node interface {
	Eval() (interface{}, error)
}
