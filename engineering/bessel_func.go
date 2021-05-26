package engineering

import (
	"math"

	"github.com/maniartech/x/utils"
)

func BesselJ(x, n interface{}) float64 { return math.Jn(utils.ToInt(n), utils.ToFloat64(x)) }
func BesselY(x, n interface{}) float64 { return math.Yn(utils.ToInt(n), utils.ToFloat64(x)) }
