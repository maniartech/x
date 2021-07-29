package parser_test

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/maniartech/x/parser"
)

func TestParser(t *testing.T) {
	t1 := time.Now()
	got, err := parser.ParseReader("", strings.NewReader("AVG (\"aaa\" <= \"add\", 'C', SUM(2 * 10, 2) + 1, 30 <= 40, 2 + SUM(99) * 2, {\"a\": [true, 10 + 20] }, [10, 2, 3-2], 'testing')"))
	// got, err := parser.ParseReader("", strings.NewReader("SUM(true) + 10"))
	fmt.Printf("%+v\n", fmt.Sprintf("%v", time.Since(t1)))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("=", got)
}
