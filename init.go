package x

import (
	"github.com/maniartech/x/calc"
	"github.com/maniartech/x/currency"
	"github.com/maniartech/x/statistics"
)

// Represents the envionment passed to the expression
// evaluator.
type Env map[string]interface{}

// Initializes package x!
func init() {
	calc.Initialize()
	statistics.Initialize()
	currency.Initialize()
}
