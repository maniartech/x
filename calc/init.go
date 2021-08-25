package calc

import "github.com/maniartech/x/registry"

func Initialize() {
	registry.Register("CEILING", Ceiling)
	registry.Register("ROUND", Round)
	registry.Register("FLOOR", Floor)
	registry.Register("GCD", GCD)
}
