package main

import (
	"context"

	pbHello "proto"

	"github.com/micro/go-micro/v2/logger"
)

type GreeterHandler struct{}

// Hello ...
func (g *GreeterHandler) Hello(ctx context.Context, req *pbHello.HelloRequest, rsp *pbHello.HelloResponse) error {
	logger.Info("Hello.enter", "name", req.Name)

	rsp.Greeting = "Hello " + req.Name
	return nil
}
