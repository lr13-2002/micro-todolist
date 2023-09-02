package router

import (
	"micro-todolist/app/gateway/http"
	"micro-todolist/app/gateway/middleware"

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

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task", http.CreateTaskHandler)
			authed.POST("update_task", http.UpdateTaskHandler) // task_id
			authed.POST("delete_task", http.DeleteTaskHandler) // task_id
			authed.GET("list_task", http.ListTaskHandler)
			authed.GET("get_task", http.GetTaskHandler) // task_id
		}
	}
	return ginRouter
}
