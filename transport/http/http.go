package http

import (
	"github.com/bottos-project/go-micro/transport"
)

func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewTransport(opts...)
}
