package api

import (
	"x-gin-admin/admin/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	// 子级路由组
	api := r.Group("/api")
	{
		var user = controller.User{}
		api.GET("", user.SelectUser)
	}
}
