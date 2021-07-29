package ast

// Defines the types of the nodes in the AST.

type NodeType int

const (
	Null      NodeType = iota // Null node
	Boolean                   // Boolean node
	Number                    // Number node
	String                    // String node
	Array                     // Array node
	Object                    // Object node
	Ident                     // Identifier node
	Keyword                   // Keyword node
	Operator                  // Operator node
	AssignOp                  // Assign operator node
	BinOp                     // Binary operator node
	CompOp                    // Comparison operator node
	LogicalOp                 // Logical operator node
	Func                      // Function node
	FuncArgs                  // Function arguments node
)

// Returns the string representation of the node type.
func (n NodeType) String() string {
	switch n {
	case Null:
		return "Null"
	case Boolean:
		return "Boolean"
	case Number:
		return "Number"
	case String:
		return "String"
	case Array:
		return "Array"
	case Object:
		return "Object"
	case Ident:
		return "Ident"
	case Keyword:
		return "Keyword"
	case Operator:
		return "Operator"
	case AssignOp:
		return "AssignOp"
	case BinOp:
		return "BinOp"
	case CompOp:
		return "CompOp"
	case LogicalOp:
		return "LogicalOp"
	case Func:
		return "Func"
	case FuncArgs:
		return "FuncArgs"
	}
	return "Unknown"
}
