package registry

import (
	"fmt"
	"reflect"

	"github.com/maniartech/x/utils"
)

// Register registers a function with the function registry.
func Register(n string, fn interface{}) {
	// Ensure fn is a function
	if !utils.IsFunc(fn) {
		panic("The fn must be a function!")
	}

	// Ensure n is not already registered!
	if _, ok := registry[n]; ok {
		panic(fmt.Sprintf("The function %v is already registered!", n))
	}

	registry[n] = fn
}

// Unregister removes the function from the function registry.
func Unregister(n string) {

	// Ensure n is not already registered!
	if _, ok := registry[n]; !ok {
		panic(fmt.Sprintf("The function %v is not registered!", n))
	}

	delete(registry, n)
}

func AttachFunctions(env map[string]interface{}) {
	for k, v := range registry {
		env[k] = v
	}
}

func Get() map[string]interface{} {
	return registry
}

// Invoke calls the function with the given name.
// It first looks for the function key exists in the registry.
// If it does not, it looks for the function in the env.
// If it does not exist in either, it returns an error.
func Invoke(n string, args []interface{}, env map[string]interface{}) (interface{}, error) {
	fn, ok := registry[n]
	if !ok {
		fn, ok = env[n]
		if !ok {
			return nil, fmt.Errorf("the function %v does not exist", n)
		}
	}

	// Invoke function using reflection. Pass the args to the function.
	fnValue := reflect.ValueOf(fn)

	// Convert the args to the reflect.Value
	argsValue := make([]reflect.Value, len(args))
	for i, arg := range args {
		argsValue[i] = reflect.ValueOf(arg)
	}

	// Call the function
	returnVals := fnValue.Call(argsValue)

	result := returnVals[0].Interface()

	if len(returnVals) == 1 {
		return result, nil
	}

	err := returnVals[0].Interface().(error)
	return result, err
}
