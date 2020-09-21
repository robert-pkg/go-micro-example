package main

import (
	"time"

	pbHello "proto"

	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry/etcd"

	log "common/log"
)

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
