package registry

var registry map[string]interface{}

func init() {
	registry = make(map[string]interface{})
}
