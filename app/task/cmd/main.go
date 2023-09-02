package main

import (
	"context"
	"fmt"
	"micro-todolist/app/task/repository/db/dao"
	"micro-todolist/app/task/repository/mq"
	"micro-todolist/app/task/script"
	"micro-todolist/app/task/service"
	"micro-todolist/config"
	"micro-todolist/idl/pb"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()
	mq.InitRabbitMQ()
	loadingScript()
	//etcd 注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	//new 一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address(config.TaskServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()

	_ = pb.RegisterTaskServiceHandler(microService.Server(), service.GetTaskSrv())

	_ = microService.Run()
}

func loadingScript() {
	ctx := context.Background()
	go script.TaskCreateSync(ctx)
}
