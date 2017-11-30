package pipeline

type (
	Context interface {
		Set(key string, val interface{}) error
		Get(key string) (interface{}, bool)
	}

	context struct {
		values map[string]interface{}
	}
)

func NewContext() Context {
	return &context{
		values: make(map[string]interface{}),
	}
}

func (c *context) Set(key string, val interface{}) error {
	c.values[key] = val
	return nil
}

func (c *context) Get(key string) (interface{}, bool) {
	v, ok := c.values[key]
	return v, ok
}
