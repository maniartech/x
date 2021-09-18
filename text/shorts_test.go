package text_test

import (
	"testing"

	"github.com/maniartech/x/text"
	"github.com/stretchr/testify/assert"
)

func TestShorts(t *testing.T) {
	//Char
	assert.Equal(t, "A", text.Char(65))
	assert.Equal(t, "a", text.Char(97))
	assert.Equal(t, " ", text.Char(32))
	assert.Equal(t, "😊", text.Char(128522))

	//Code
	assert.Equal(t, 32, text.Code(" "))
	assert.Equal(t, 97, text.Code("a"))
	assert.Equal(t, 65, text.Code("ASDF"))
	assert.Equal(t, 128522, text.Code("😊👌"))

	//Len
	assert.Equal(t, 6, text.Len(-65.23))
	assert.Equal(t, 11, text.Len("Hello World"))
	assert.Equal(t, 6, text.Len("こんにちは！"))
	assert.Equal(t, 27, text.Len("It's day 21! Good Morning😊"))

	//Upper
	assert.Equal(t, "-65.23", text.Upper(-65.23))
	assert.Equal(t, "HELLO WORLD", text.Upper("Hello World"))
	assert.Equal(t, "こんにちは！", text.Upper("こんにちは！"))
	assert.Equal(t, "IT'S DAY 21! GOOD MORNING😊", text.Upper("It's day 21! Good Morning😊"))

	//Lower
	assert.Equal(t, "-65.23", text.Lower(-65.23))
	assert.Equal(t, "hello world", text.Lower("Hello World"))
	assert.Equal(t, "こんにちは！", text.Lower("こんにちは！"))
	assert.Equal(t, "it's day 21! good morning😊", text.Lower("It's day 21! Good Morning😊"))

	//Proper
	assert.Equal(t, "-65.23", text.Proper(-65.23))
	assert.Equal(t, "Hello World", text.Proper("Hello wOrld"))
	assert.Equal(t, "こんにちは！", text.Proper("こんにちは！"))
	assert.Equal(t, "It's Day 21! Good Morning😊", text.Proper("It's day 21! Good moRniNg😊"))

	//Left
	assert.Equal(t, "-6", text.Left(-65.23, 2))
	assert.Equal(t, "Hello World", text.Left("Hello World", 20))
	assert.Equal(t, "こんにちは！", text.Left("こんにちは！", len("こんにちは！")))
	assert.Equal(t, "It's day 21! Good Mo", text.Left("It's day 21! Good Morning😊", 20))

	//Right
	assert.Equal(t, "23", text.Right(-65.23, 2))
	assert.Equal(t, "Hello World", text.Right("Hello World", 20))
	assert.Equal(t, "こんにちは！", text.Right("こんにちは！", len("こんにちは！")))
	assert.Equal(t, "y 21! Good Morning😊", text.Right("It's day 21! Good Morning😊", 20))

	//Fixed
	assert.Equal(t, "21.23", text.Fixed(21.2312, 2))
	assert.Equal(t, "200", text.Fixed(231.2312, -2))

	//Dollar
	assert.Equal(t, "$21.23", text.Dollar(21.2312, 2))
	assert.Equal(t, "$200", text.Dollar(231.2312, -2))

	//Substitute
	assert.Equal(t, "Good Evening", text.Substitute("Good Morning", "Morning", "Evening"))
	assert.Equal(t, "ニャー ニャー", text.Substitute("吠える 吠える", "吠える", "ニャー"))
	assert.Equal(t, "2021", text.Substitute(2019, 19, 21))
	assert.Equal(t, "Booooooo", text.Substitute("Baaaaaaa", "a", "o"))
	assert.Equal(t, "👋👋", text.Substitute("😊😊", "😊", "👋"))

	//Search
	assert.Equal(t, 8, text.Search("Sato", "Kazuma Sato", 5))
	assert.Equal(t, -1, text.Search("o", "Good Morning", 10))
	assert.Equal(t, 1, text.Search("-", -45.15, 0))
	assert.Equal(t, 7, text.Search(".", -12345.483, 3))

	//Replace
	assert.Equal(t, "Good Evening", text.Replace("Good Morning", 6, 9, "Evening"))
	assert.Equal(t, "21.", text.Replace("21:20", 3, 3, "."))

	//Rept
	assert.Equal(t, "😊😊😊😊😊", text.Rept("😊", 5))
	assert.Equal(t, "2020", text.Rept(20, 2))
	assert.Equal(t, "Baa Baa Baa ", text.Rept("Baa ", 3))

	//Exact
	assert.Equal(t, "true", text.Exact(20, "20"))
	assert.Equal(t, "false", text.Exact(01, "01"))
	assert.Equal(t, "false", text.Exact("go", "Go"))
	assert.Equal(t, "true", text.Exact(true, "true"))

}
