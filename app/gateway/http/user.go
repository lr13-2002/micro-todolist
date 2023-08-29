package http

import (
	"micro-todolist/app/gateway/rpc"
	"micro-todolist/idl/pb"
	"micro-todolist/pkg/ctl"
	"micro-todolist/pkg/jwt"
	"micro-todolist/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind"))
		return
	}
	userResp, err := rpc.UserRegister(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserRegister-RPC"))
		return
	}
	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, userResp))
}

func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-ShouldBind"))
	}
	userResp, err := rpc.UserLogin(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserRegisterHandler-UserLogin-RPC"))
	}

	token, err := jwt.GenerateToken(uint(userResp.UserDetail.Id))
	if err != nil {
		ctx.JSON(http.StatusOK, ctl.RespError(ctx, err, "UserLoginHandler-GenerateToken"))
	}

	res := &types.TokenData{
		User:  userResp,
		Token: token,
	}

	ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
}
