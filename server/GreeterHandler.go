package main

import (
	"context"

	pbHello "proto"
)

type GreeterHandler struct{}

// Hello ...
func (g *GreeterHandler) Hello(ctx context.Context, req *pbHello.HelloRequest, rsp *pbHello.HelloResponse) error {

	rsp.Greeting = "Hello " + req.Name
	return nil
}
