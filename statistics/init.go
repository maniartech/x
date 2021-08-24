package statistics

import (
	"github.com/maniartech/x/registry"
)

func Initialize() {
	registry.Register("AVERAGE", Average)
	registry.Register("AVERAGEA", AverageA)
}
