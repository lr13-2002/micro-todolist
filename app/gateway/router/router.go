package router

import (
	"micro-todolist/app/gateway/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})

		v1.POST("/user/register", http.UserRegisterHandler)
		v1.POST("/user/login", http.UserLoginHandler)

	}
	return ginRouter
}
