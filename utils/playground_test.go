package utils_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/maniartech/x/calc"
)

/*
 * This is a playground, many times while developing you need to test
 * some concept. Use this test to quickly run the ideas and concept
 * using VS Code's integrated test environment!
 */
func TestPlayground(t *testing.T) {
	f := 0.12
	g := 0.12
	for i := 0; i < 20; i++ {
		f = f + 0.01
		g = g - 0.01

		big.NewFloat(f)
		fmt.Printf(">>> %v\t%v\t%v\n", f, calc.Round(f, 15), calc.Round(g, 15))
	}
	println(calc.Round(0.0999999999999999, 15) * 2.0)

}
