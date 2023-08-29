package rpc

import (
	"context"
	"fmt"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/e"

	"github.com/gin-gonic/gin"
)

func UserLogin(ctx *gin.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)
	if err != nil || resp.Code != e.Success {
		resp.Code = e.Error
		return
	}
	return
}

// UserRegister 用户注册
func UserRegister(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	fmt.Printf("name : %v\n", req.UserName)
	resp, err = UserService.UserRegister(ctx, req)
	if err != nil {
		return
	}

	return
}
