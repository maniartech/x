package parser

type NullNode struct {
	ASTNode
}

func NewNullNode(line int, column int) *NullNode {
	return &NullNode{
		ASTNode: ASTNode{
			Type: TypeNull,
			Position: Position{
				Line:   line,
				Column: column,
			},
		},
	}
}

func (n NullNode) String() string {
	return "NullNode {null}"
}

func (n NullNode) Eval(map[string]interface{}) (interface{}, error) {
	return nil, nil
}
