package main

import (
	"context"
	pbHello "proto"
	"time"

	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/server"

	log "common/log"
)

// 实现server.HandlerWrapper接口
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		requestID, _ := metadata.Get(ctx, "Requestid")

		logger.Infof("server invoke by grpc. requestID:%s service:%s method:%s", requestID, req.Service(), req.Endpoint())

		return fn(ctx, req, rsp)
	}
}

func main() {

	log.Init("./server.log")

	logger.Info("start")
	logger.Debug("start debug")

	r := etcd.NewRegistry()

	// 创建服务
	service := grpc.NewService(
		service.Name("com.robert.api.greeter"),
		service.Registry(r),
		service.RegisterTTL(time.Second*30),
		service.RegisterInterval(time.Second*10),
		service.WrapHandler(logWrapper),
	)

	// 初始化方法会解析命令行标识
	service.Init()

	// 注册处理器
	pbHello.RegisterGreeterHandler(service.Server(), &GreeterHandler{})

	// 运行服务
	if err := service.Run(); err != nil {
		logger.Fatal(err)
		return
	}

	logger.Info("exit server...")
}
