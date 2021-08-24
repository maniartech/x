package parser

import (
	"bytes"
	"fmt"
)

type ArrayNode struct {
	ASTNode

	Items []Node
}

func NewArrayNode(line int, column int, items []Node) Node {
	node := ArrayNode{
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

func (n ArrayNode) String() string {
	var b bytes.Buffer
	b.WriteString("[")
	for i, item := range n.Items {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(fmt.Sprintf("%v", item))
	}
	b.WriteString("]")

	return fmt.Sprintf("ArrayNode: %s", b.String())
}

// Eval parses the array node and returns the array node.
func (n ArrayNode) Eval(env map[string]interface{}) (interface{}, error) {

	results := make([]interface{}, len(n.Items))

	// Iterate each items and attach the evaluated results
	for i, item := range n.Items {
		// Evaluate the item
		evaluated, err := item.Eval(env)
		if err != nil {
			return nil, err
		}

		// Set the evaluated result
		results[i] = evaluated
	}
	return results, nil
}
