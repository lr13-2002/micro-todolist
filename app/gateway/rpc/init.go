package rpc

import (
	"micro-todolist/app/gateway/wrappers"
	"micro-todolist/idl/pb"

	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
	TaskService pb.TaskService
)

func InitRPC() {
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	UserService = userService

	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())
	TaskService = taskService
}
