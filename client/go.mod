module github.com/zhanglihua2008/go-micro-example/client

go 1.13

require (
	common v0.0.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/satori/go.uuid v1.2.0
	proto v0.0.0
)

replace common => ../common

replace proto => ../proto
