package main

import (
	"context"
	pbHello "proto"
	"time"

	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"

	log "common/log"

	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/registry/etcd"
	uuid "github.com/satori/go.uuid"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {

	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata)
	}

	newReqID := uuid.NewV4().String()
	md.Set("Requestid", newReqID)
	ctx = metadata.NewContext(ctx, md)

	err := l.Client.Call(ctx, req, rsp)

	if err == nil {
		logger.Infof("grpc call. service:%s method:%s, req:%v, rsp:%v", req.Service(), req.Endpoint(), req, rsp)
	} else {
		logger.Errorf("grpc call. service:%s method:%s", req.Service(), req.Endpoint())
	}

	return err
}

// 实现client.Wrapper，充当日志包装器
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	log.Init("./client.log")

	logger.Info("start")

	r := etcd.NewRegistry()

	service := grpc.NewService(
		service.Registry(r),
		service.WrapClient(logWrap),
	)
	service.Init()

	// Create new greeter client
	greeter := pbHello.NewGreeterService("com.robert.api.greeter", service.Client())

	for i := 0; i < 1; i++ {

		ctx := context.Background()

		// Call the greeter
		_, err := greeter.Hello(ctx, &pbHello.HelloRequest{Name: "robert"})
		if err != nil {
			logger.Error(err)
			time.Sleep(time.Second)
			continue
		}

		//logger.Info(rsp.Greeting)

		time.Sleep(time.Second)
	}

	logger.Info("exit ...")
}
