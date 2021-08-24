package parser

import (
	"bytes"
	"fmt"
)

type ObjectItem struct {
	Key   string
	Value Node
}

type ObjectNode struct {
	ASTNode

	Items []ObjectItem
}

func NewObjectNode(line int, column int, items []ObjectItem) Node {
	node := ObjectNode{
		ASTNode: ASTNode{
			Type: TypeNull,
			Position: Position{
				Line:   line,
				Column: column,
			},
		},
		Items: items,
	}
	return Node(node)
}

func (n ObjectNode) String() string {
	var b bytes.Buffer
	b.WriteString("{")
	for i, item := range n.Items {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%s: %v", item.Key, item.Value))
	}
	b.WriteString("}")

	return fmt.Sprintf("ObjectNode: %s", b.String())
}

// Eval parses the array node and returns the array node.
func (n ObjectNode) Eval(env map[string]interface{}) (interface{}, error) {
	results := make(map[string]interface{})
	for _, item := range n.Items {
		result, err := item.Value.Eval(env)
		if err != nil {
			return nil, err
		}
		results[item.Key] = result
	}
	return results, nil
}
