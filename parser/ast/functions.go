package ast

import "fmt"

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func eval(line, col int, first, rest interface{}) interface{} {

	fmt.Printf(">> %+v %+v\n", first, rest)
	return nil

	// l := first
	// restSl := toIfaceSlice(rest)

	// for _, v := range restSl {
	// 	restExpr := toIfaceSlice(v)
	// 	if r, ok := restExpr[3].(Node); !ok {
	// 		panic("invalid-expression - assertion-failed!")
	// 		return nil
	// 	}
	// 	op := ""

	// 	if o, ok := restExpr[1].([]byte); ok {
	// 		op = string(o)
	// 	} else if o, ok := restExpr[1].(string); ok {
	// 		op = o
	// 	}

	// 	if op != "" {
	// 		fmt.Println("Operator: ", op)
	// 		l = NewExpressionNode(line, col, op, l, r)
	// 	}
	// }

	// return l
}

/*
func eval(line, col int, first, rest interface{}) Node {

	fmt.Printf(">> %+v %+v\n", first, rest)
	// return nil

	l := first
	restSl := toIfaceSlice(rest)

	for _, v := range restSl {
		restExpr := toIfaceSlice(v)
		if r, ok := restExpr[3].(Node); !ok {
			panic("invalid-expression - assertion-failed!")
			return nil
		}
		op := ""

		if o, ok := restExpr[1].([]byte); ok {
			op = string(o)
		} else if o, ok := restExpr[1].(string); ok {
			op = o
		}

		if op != "" {
			fmt.Println("Operator: ", op)
			l = NewExpressionNode(line, col, op, l, r)
		}
	}

	return l
}
*/

func getFunctionNode(line, col int, name string, args []Node) *FunctionNode {
	return NewFunctionNode(line, col, name, args)
}

func toByteSlice(val []interface{}) []byte {
	s := make([]byte, len(val))
	for i, v := range val {
		if b, ok := v.([]byte); ok {
			s[i] = b[0]
		} else if b, ok := v.(byte); ok {
			s[i] = b
		}
	}
	return s
}

func toString(val interface{}) string {
	if v, ok := val.([]interface{}); ok {
		bytes := toByteSlice(v)
		return string(bytes)
	}
	panic("not-implemented")
}
