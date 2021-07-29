package registry

import (
	"fmt"

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
