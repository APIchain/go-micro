package http

import (
	"github.com/bottos-project/go-micro/broker"
)

func NewBroker(opts ...broker.Option) broker.Broker {
	return broker.NewBroker(opts...)
}
