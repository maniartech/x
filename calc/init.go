package calc

import "github.com/maniartech/x/registry"

func Initialize() {
	registry.Register("Ceiling", Ceiling)
	registry.Register("Round", Round)
}
