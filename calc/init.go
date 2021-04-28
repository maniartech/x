package calc

import "github.com/maniartech/x/registry"

func Initialize() {
	registry.Register("Ceiling", Ceiling)
	registry.Register("Round", Round)
	registry.Register("Floor", Floor)
	registry.Register("GCD", GCD)
}
