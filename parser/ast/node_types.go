package ast

// NodeType defines the types of the nodes in the AST.
type NodeType uint8

const (
	TypeNull       NodeType = iota // Null node
	TypeBoolean                    // Boolean node
	TypeNumber                     // Number node
	TypeString                     // String node
	TypeArray                      // Array node
	TypeObject                     // Object node
	TypeExpression                 // Expression node
	TypeFunc                       // Function node
	TypeIdent                      // Identifier node
)

// Returns the string representation of the node type.
func (n NodeType) String() string {
	switch n {

	// Null node
	case TypeNull:
		return "Null"

		// Boolean node
	case TypeBoolean:
		return "Boolean"

		// Number node
	case TypeNumber:
		return "Number"

		// String node
	case TypeString:
		return "String"

		// Array node
	case TypeArray:
		return "Array"

		// Object node
	case TypeObject:
		return "Object"

		// Identifier node
	case TypeIdent:
		return "Ident"

		// Expession Node
	case TypeExpression:
		return "Expression"

		// Function node
	case TypeFunc:
		return "Func"
	}
	// Unknown node
	return "Unknown"
}
