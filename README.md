### micro example

micro example

### 启动etcd

```
切换到etcd所在目录
cd ~/myapp/golang/etcd/etcd-v3.3.18-darwin-amd64

./etcd
```

### 启动micro

```
切换到micro所在目录
cd ~/DVP/dvp_go/go/src/github.com/micro/micro/v2

查看服务列表
./micro --registry=etcd list services

启动micro
./micro --api_namespace=com.robert --registry=etcd api
```

### 通过curl调用接口

```
curl "http://localhost:8080/greeter/Greeter/Hello?name=John"

curl -H 'Content-Type: application/json' \
	-X POST \
	-d '{"name": "John"}' \
	http://localhost:8080/greeter/Greeter/Hello
```



