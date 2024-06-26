// Code generated by Kitex v0.9.1. DO NOT EDIT.

package konataservice

import (
	server "github.com/cloudwego/kitex/server"
	konata_client "github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler konata_client.KonataService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
