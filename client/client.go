package main

import (
	"context"
	pbHello "proto"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"

	log "common/log"

	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	log.Init("./client.log")

	logger.Info("start")

	r := etcd.NewRegistry()

	service := grpc.NewService(
		service.Registry(r),
	)
	service.Init()

	// Create new greeter client
	greeter := pbHello.NewGreeterService("com.robert.api.greeter", service.Client())

	for i := 0; i < 10; i++ {

		ctx := context.Background()

		// Call the greeter
		rsp, err := greeter.Hello(ctx, &pbHello.HelloRequest{Name: "robert"})
		if err != nil {
			logger.Error(err)
			time.Sleep(time.Second)
			continue
		}

		logger.Info(rsp.Greeting)

		time.Sleep(time.Second)
	}

	logger.Info("exit ...")
}
