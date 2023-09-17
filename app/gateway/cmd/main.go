package main

import (
	"fmt"
	"log"
	"micro-todolist/app/gateway/router"
	"micro-todolist/app/gateway/rpc"
	"micro-todolist/config"
	"os"
	"time"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

func main() {
	config.Init()
	rpc.InitRPC()

	//etcd 注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	//日志
	file, err := os.Create(config.GateWayPath)
	log.Println(config.GateWayPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	//new 一个微服务实例
	webService := web.NewService(
		web.Name("httpService"),
		web.Address("localhost:4000"),
		web.Registry(etcdReg),
		web.Handler(router.NewRouter()),
		web.RegisterTTL(time.Second*30),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	_ = webService.Init()

	_ = webService.Run()
}
