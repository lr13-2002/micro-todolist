package rpc

import (
	"micro-todolist/idl/pb"

	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
	TaskService pb.TaskService
)

func InitRPC() {
	userMicroService := micro.NewService(micro.Name("userService.client"))
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	UserService = userService

	taskMicroService := micro.NewService(micro.Name("taskService.client"))
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())
	TaskService = taskService
}
