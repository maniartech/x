package statistics

import (
	"github.com/maniartech/x/registry"
)

func Initialize() {
	registry.Register("Average", Average)
	registry.Register("AverageA", AverageA)
}
