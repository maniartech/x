package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maniartech/x/parser"
	// "github.com/maniartech/x/currency"
	// "github.com/maniartech/x/statistics"
	// "github.com/maniartech/x/statistics"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage: calculator 'EXPR'")
	}

	got, err := parser.ParseReader("", strings.NewReader(os.Args[1]))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=", got)
}
