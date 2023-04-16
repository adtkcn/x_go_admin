package admin

import (
	"x-gin-admin/admin/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	// 子级路由组
	admin := r.Group("/admin")
	{
		// var user = controller.User{}
		// admin.GET("/", user.SelectUser)
		// admin.GET("GetUsers", user.GetUsers)
		controller.RegisterUserRoutes(admin)
		controller.RegisterRoleRoutes(admin)

	}
}
