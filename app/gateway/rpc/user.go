package rpc

import (
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
