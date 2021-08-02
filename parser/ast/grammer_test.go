package ast_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maniartech/x/parser/ast"
)

func TestGrammer(t *testing.T) {
	got, err := ast.ParseReader("", strings.NewReader("1 + 2.5"))
	if err != nil {
		t.Error(err)
	}

	fmt.Println("=", got)
}
