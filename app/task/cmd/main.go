package main

import (
	"context"
	"fmt"
	"log"
	"micro-todolist/app/task/repository/db/dao"
	"micro-todolist/app/task/repository/db/redis"
	"micro-todolist/app/task/repository/mq"
	"micro-todolist/app/task/script"
	"micro-todolist/app/task/service"
	"micro-todolist/config"
	"micro-todolist/idl/pb"
	"os"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()
	redis.InitRDB()
	mq.InitRabbitMQ()
	loadingScript()
	//etcd 注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	//日志
	file, err := os.Create(config.TaskPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

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
