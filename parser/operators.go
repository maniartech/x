package parser

import (
	"strings"
)

var ops = map[string]func(interface{}, interface{}) interface{}{
	"+": func(l, r interface{}) interface{} {
		return l.(float64) + r.(float64)
	},
	"-": func(l, r interface{}) interface{} {
		return l.(float64) - r.(float64)
	},
	"*": func(l, r interface{}) interface{} {
		return l.(float64) * r.(float64)
	},
	"/": func(l, r interface{}) interface{} {
		return l.(float64) / r.(float64)
	},
	"=": func(l, r interface{}) interface{} {
		return l == r
	},
	"<>": func(l, r interface{}) interface{} {
		return l != r
	},
	"<": func(l, r interface{}) interface{} {
		if v, ok := l.(float64); ok {
			return v < r.(float64)
		} else if v, ok := l.(string); ok {
			return strings.Compare(v, r.(string)) < 0
		}
		return false // TODO
	},
	"<=": func(l, r interface{}) interface{} {
		if v, ok := l.(float64); ok {
			return v <= r.(float64)
		} else if v, ok := l.(string); ok {
			return strings.Compare(v, r.(string)) <= 0
		}
		return false // TODO
	},
	">": func(l, r interface{}) interface{} {
		if v, ok := l.(float64); ok {
			return v > r.(float64)
		} else if v, ok := l.(string); ok {
			return strings.Compare(v, r.(string)) > 0
		}
		return false // TODO
	},
	">=": func(l, r interface{}) interface{} {
		if v, ok := l.(float64); ok {
			return v >= r.(float64)
		} else if v, ok := l.(string); ok {
			return strings.Compare(v, r.(string)) >= 0
		}
		return false // TODO
	},
}
