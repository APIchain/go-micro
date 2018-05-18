package consul

import (
	"github.com/bottos-project/go-micro/registry"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}
