package main

import (
	"fmt"
	"log"
	"micro-todolist/app/user/repositery/db/dao"
	"micro-todolist/app/user/service"
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

	//etcd 注册
	etcdReg := etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)),
	)

	//日志
	file, err := os.Create(config.UserPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	//new 一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()

	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())

	_ = microService.Run()
}
