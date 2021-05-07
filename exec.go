package x

import (
	"github.com/antonmedv/expr"
	"github.com/maniartech/x/registry"
	"github.com/shopspring/decimal"
)

// Exec ececutes the specified expression.
func Exec(expression string, env Env, normalizeFloat ...bool) (interface{}, error) {

	// Initialize the environment
	if env == nil {
		env = make(Env)
	}
	registry.AttachFunctions(env)

	// Compile expression
	options := []expr.Option{
		expr.Env(env),
		expr.AllowUndefinedVariables(), // Allow to use undefined variables.
	}

	c, err := expr.Compile(expression, options...)
	if err != nil {
		return nil, err
	}

	// Execute expression
	result, err := expr.Run(c, env)
	if err != nil {
		return nil, err
	}

	// Return result. Generally float64 based calculataion causes
	// floating point errors.
	// => https://en.wikipedia.org/wiki/IEEE_754
	// => https://docs.oracle.com/cd/E19957-01/806-3568/ncg_goldberg.html
	//
	// Excel normalizes the result by rounding off the result to max 15 decimals!
	var fnormalize = true
	if len(normalizeFloat) == 1 {
		fnormalize = normalizeFloat[0]
	}
	if fnormalize {
		if f, ok := result.(float64); ok {
			r, _ := decimal.NewFromFloat(f).Round(15).Float64()
			return r, nil
		}
	}

	return result, nil
}
