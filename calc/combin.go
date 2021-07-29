package calc

import (
	"github.com/maniartech/x/core"
	"github.com/maniartech/x/utils"
)

func Combin(num, numChosen interface{}) int {
	n := utils.ToInt(num)
	k := utils.ToInt(numChosen)

	if n < 0 || k < 0 || n < k {
		panic(core.ErrInvalidInput)
	}
	return int(Divide(Fact(n), Fact(k)*Fact(n-k)))
}
