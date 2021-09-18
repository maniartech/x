package statistics

import "github.com/maniartech/x/utils"

// Count, counts the number of items passes in the func.
func Count(v ...interface{}) int {
	//Counting all the items
	c := utils.ForEach(func(_ int, x interface{}) {
	}, v...)
	return c
}

// CountA, counts the number of items passes in the func.
func CountA(v ...interface{}) int {
	//Counting all the items
	c := utils.ForEach(func(_ int, x interface{}) {
	}, v...)
	return c
}
