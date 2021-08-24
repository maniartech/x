package parser

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
	Eval(map[string]interface{}) (interface{}, error)
}
