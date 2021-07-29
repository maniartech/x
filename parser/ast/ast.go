package ast

type Position struct {
	Line   int
	Column int
	Offset int
}

type AST struct {
	Type        NodeType
	Children    []AST
	Value       interface{}
	Position    Position
	EndPosition Position
}
