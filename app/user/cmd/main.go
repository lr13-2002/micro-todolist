package main

import (
	"fmt"
	"micro-todolist/app/user/repositery/db/dao"
	"micro-todolist/app/user/service"
	"micro-todolist/config"
	"micro-todolist/idl/pb"

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
