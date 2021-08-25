package parser_test

import (
	"testing"

	"github.com/maniartech/x/parser"
	"github.com/stretchr/testify/assert"
)

// TestDictFind tests the Find method in the Dict type.
// It recursively tests the Find method for each level of the Dict.
func TestDictFindCaseA(t *testing.T) {
	dict := getDict()

	// Test A
	result, ok := dict.Find("a")
	assert.True(t, ok)
	assert.EqualValues(t, parser.D{
		"aa": parser.D{
			"aaa": parser.D{},
			"aab": true,
			"aac": 123,
		},
	}, result)

	result, ok = dict.Find("a.aa")
	assert.True(t, ok)
	assert.EqualValues(t, parser.D{
		"aaa": parser.D{},
		"aab": true,
		"aac": 123,
	}, result)

	result, ok = dict.Find("a.aa.aaa")
	assert.True(t, ok)
	assert.EqualValues(t, parser.D{}, result)

	result, ok = dict.Find("a.aa.aab")
	assert.True(t, ok)
	assert.EqualValues(t, true, result)

	result, ok = dict.Find("a.aa.aac")
	assert.True(t, ok)
	assert.EqualValues(t, 123, result)
}

func TestDictFindCaseB(t *testing.T) {
	dict := getDict()

	// Test B
	result, ok := dict.Find("b")
	assert.True(t, ok)
	assert.EqualValues(t, map[string]interface{}{
		"bb": map[string]interface{}{
			"bbb": map[string]interface{}{},
			"bbc": true,
			"bbd": 123,
		},
	}, result)

	result, ok = dict.Find("b.bb")
	assert.True(t, ok)
	assert.EqualValues(t, map[string]interface{}{
		"bbb": map[string]interface{}{},
		"bbc": true,
		"bbd": 123,
	}, result)

	result, ok = dict.Find("b.bb.bbb")
	assert.True(t, ok)
	assert.EqualValues(t, map[string]interface{}{}, result)

	result, ok = dict.Find("b.bb.bbc")
	assert.True(t, ok)
	assert.EqualValues(t, true, result)

	result, ok = dict.Find("b.bb.bbd")
	assert.True(t, ok)
	assert.EqualValues(t, 123, result)
}

func TestDictFindCaseC(t *testing.T) {
	dict := getDict()

	// Test C
	result, ok := dict.Find("c")
	assert.True(t, ok)
	assert.EqualValues(t, []interface{}{
		map[string]interface{}{
			"ccc": true,
			"ccd": 123,
			"cce": []interface{}{1, 2, 3, 4, 5},
		}, true,
	}, result)

	result, ok = dict.Find("c.0")
	assert.True(t, ok)
	assert.EqualValues(t, map[string]interface{}{
		"ccc": true,
		"ccd": 123,
		"cce": []interface{}{1, 2, 3, 4, 5},
	}, result)

	result, ok = dict.Find("c.0.ccc")
	assert.True(t, ok)
	assert.EqualValues(t, true, result)

	result, ok = dict.Find("c.0.ccd")
	assert.True(t, ok)
	assert.EqualValues(t, 123, result)

	result, ok = dict.Find("c.0.cce")
	assert.True(t, ok)
	assert.EqualValues(t, []interface{}{1, 2, 3, 4, 5}, result)

	result, ok = dict.Find("c.0.cce.2")
	assert.True(t, ok)
	assert.EqualValues(t, 3, result)

	result, ok = dict.Find("c.1")
	assert.True(t, ok)
	assert.EqualValues(t, true, result)
}

func getDict() parser.D {
	return parser.D{
		"a": parser.D{
			"aa": parser.D{
				"aaa": parser.D{},
				"aab": true,
				"aac": 123,
			},
		},
		"b": map[string]interface{}{
			"bb": map[string]interface{}{
				"bbb": map[string]interface{}{},
				"bbc": true,
				"bbd": 123,
			},
		},
		"c": []interface{}{
			map[string]interface{}{
				"ccc": true,
				"ccd": 123,
				"cce": []interface{}{1, 2, 3, 4, 5},
			},
			true,
		},
	}
}
