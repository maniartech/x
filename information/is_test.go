package information_test

import (
	"fmt"
	"testing"

	"github.com/maniartech/x/information"
	"github.com/maniartech/x/utils"
	"github.com/stretchr/testify/assert"
)

func TestIs(t *testing.T) {
	//IsLogical
	var testCases []interface{} = []interface{}{1, 32.45, true, false, "Foo", 'a', -1234, -152.573}
	var typeResults []interface{} = []interface{}{"int", "float64", "bool", "bool", "string", "int32", "int", "float64"}
	var isLogicalResults []bool = []bool{false, false, true, true, false, false, false, false}
	var isNonTextResults []bool = []bool{true, true, true, true, false, false, true, true}
	var isNumberResults []bool = []bool{true, true, false, false, false, false, true, true}
	var oddEvenTestCases []interface{} = []interface{}{123, -21.4, 1, 45, 12.35, -134.123, 412.341}
	var oddEvenResults []bool = []bool{true, true, true, true, false, false}
	var NResults []float64 = []float64{1.0, 32.45, 1, 0, 0, 0, -1234, -152.573}
	var length int = len(testCases)
	for i := 0; i < length; i++ {
		if !assert.Equal(t, isLogicalResults[i], information.IsLogical(testCases[i])) {
			printErr("isLogical", isLogicalResults[i], information.IsLogical(testCases[i]), testCases[i])
		}
		if !assert.Equal(t, isNonTextResults[i], information.IsNonText(testCases[i])) {
			printErr("isNonText", isNonTextResults[i], information.IsNonText(testCases[i]), testCases[i])
		}
		if !assert.Equal(t, !isNonTextResults[i], information.IsText(testCases[i])) {
			printErr("isText", !isNonTextResults[i], information.IsText(testCases[i]), testCases[i])
		}
		if !assert.Equal(t, isNumberResults[i], information.IsNumber(testCases[i])) {
			printErr("isNonText", isNumberResults[i], information.IsNumber(testCases[i]), testCases[i])
		}
		if !assert.Equal(t, NResults[i], information.N(testCases[i])) {
			printErr("isNonText", NResults[i], information.N(testCases[i]), testCases[i])
		}
		if !assert.Equal(t, typeResults[i], information.Type(testCases[i])) {
			printErr("isNonText", typeResults[i], information.Type(testCases[i]), testCases[i])
		}
	}
	// length2 := len(oddEvenTestCases)
	for i := 0; i < 6; i++ {
		if !assert.Equal(t, oddEvenResults[i], information.IsOdd(oddEvenTestCases[i])) {
			printErr("isOdd", oddEvenResults[i], information.IsOdd(oddEvenTestCases[i]), oddEvenTestCases[i])
		}
		if !assert.Equal(t, !oddEvenResults[i], information.IsEven(oddEvenTestCases[i])) {
			printErr("isEven", !oddEvenResults[i], information.IsEven(oddEvenTestCases[i]), oddEvenTestCases[i])
		}
	}
}
func printErr(function string, expected, got interface{}, value interface{}) {
	fmt.Printf("Expected %t got %t for value %s in %s Function!\n", expected, got, utils.ToString(value), function)
}
