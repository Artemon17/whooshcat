package di

import "fmt"

type Factory func(c *Container) (any, error)

type Container struct {
	factories map[string]Factory
	instances map[string]any
}

func New() *Container {
	return &Container{
		factories: make(map[string]Factory),
		instances: make(map[string]any),
	}
}

func (c *Container) Register(name string, f Factory) {
	c.factories[name] = f
}

func (c *Container) Resolve(name string) (any, error) {
	if v, ok := c.instances[name]; ok {
		return v, nil
	}

	f, ok := c.factories[name]
	if !ok {
		return nil, fmt.Errorf("di: не зарегистрировано: %q", name)
	}

	v, err := f(c)
	if err != nil {
		return nil, fmt.Errorf("di: ошибка создания %q: %w", name, err)
	}

	c.instances[name] = v
	return v, nil
}

func (c *Container) MustResolve(name string) any {
	v, err := c.Resolve(name)
	if err != nil {
		panic(err)
	}
	return v
}
