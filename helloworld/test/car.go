package car

import (
	"errors"
)

type Car struct {
	name  string
	price float32
}

func (c *Car) setName(name string) string {
	if name != "" {
		c.name = name
	}

	return c.name
}

func new(name string, price float32) (*Car, error) {
	if name == "" {
		return nil, errors.New("missing name")
	}

	return &Car{
		name:  name,
		price: price,
	}, nil
}
