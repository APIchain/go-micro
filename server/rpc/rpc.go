package rpc

import (
	"github.com/bottos-project/go-micro/server"
)

func NewServer(opts ...server.Option) server.Server {
	return server.NewServer(opts...)
}
